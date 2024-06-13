/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time :
 * @File :
 */

package rpc

import (
	"context"
	"errors"
	"fmt"
	"github.com/gogo/protobuf/types"
	"novachat_engine/mtproto"
	chatService "novachat_engine/pkg/cmd/chat/rpc_client"
	"novachat_engine/pkg/cmd/chat/service"
	errorsService "novachat_engine/service/common/errors"
	"novachat_engine/service/constants"
	data_chat "novachat_engine/service/data/chat"
	"regexp"
	"sort"
)

func (impl *Impl) ReqGetParticipants(ctx context.Context, request *chatService.GetParticipants) (*types.Any, error) {
	chat, err := impl.chatManager.GetChat(request.ChatId)
	if err != nil {
		if errors.Is(err, service.DefaultErrChatIdInvalid) {
			return nil, errorsService.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_CHAT_ID_INVALID)
		} else {
			return nil, err
		}
	}

	if chat.Invalid() || chat.Deleted() {
		return nil, errorsService.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_CHAT_INVALID)
	}

	chatData := chatService.Chat{
		ChatData:        chat.GetChatInfo().ChatData,
		ParticipantList: []*data_chat.ChatParticipant{},
		Count:           chat.GetChatInfo().Count,
	}
	if request.UserId != 0 {
		participant := chat.GetChatInfo().GetParticipants(request.UserId)
		if participant != nil {
			chatData.ParticipantList = append(chatData.ParticipantList, participant)
		}
		return types.MarshalAny(&chatData)
	}
	if request.Filter == constants.ChannelParticipantsFilterContacts.ToInt32() ||
		request.Filter == constants.ChannelParticipantsFilterBots.ToInt32() ||
		request.Filter == constants.ChannelParticipantsFilterMentions.ToInt32() ||
		request.Filter == constants.ChannelParticipantsFilterSearch.ToInt32() && len(request.Q) == 0 {
		return types.MarshalAny(&chatData)
	}

	var r *regexp.Regexp
	if request.Filter == constants.ChannelParticipantsFilterSearch.ToInt32() {
		r, err = regexp.Compile(fmt.Sprintf("?%s?", request.Q))
		if err != nil {
			return types.MarshalAny(&chatData)
		}
		_ = r
	}

	chatParticipantList := make([]*data_chat.ChatParticipant, 0, chat.GetChatInfo().Count)
	chat.GetChatInfo().Iteration(func(participant *data_chat.ChatParticipant) {
		switch constants.ChannelParticipantsFilterTypeFromInt32(request.Filter) {
		case constants.ChannelParticipantsFilterAdmin:
			if participant.Admin && participant.State == data_chat.ParticipantState_ParticipantStateNormal {
				chatParticipantList = append(chatParticipantList, participant)
			}
		case constants.ChannelParticipantsFilterKicked:
			if participant.State == data_chat.ParticipantState_ParticipantStateDelete {
				chatParticipantList = append(chatParticipantList, participant)
			}
		case constants.ChannelParticipantsFilterBanned:
			if participant.State == data_chat.ParticipantState_ParticipantStateBan {
				chatParticipantList = append(chatParticipantList, participant)
			}
		case constants.ChannelParticipantsFilterSearch:
			fallthrough
		case constants.ChannelParticipantsFilterRecent:
			fallthrough
		default:
			if participant.State == data_chat.ParticipantState_ParticipantStateNormal {
				chatParticipantList = append(chatParticipantList, participant)
			}
		}
	})

	sort.SliceStable(chatParticipantList, func(i, j int) bool {
		return chatParticipantList[i].UserId < chatParticipantList[j].UserId
	})

	if int32(len(chatParticipantList)) > request.Offset {
		if request.Offset+request.Limit < int32(len(chatParticipantList)) {
			chatData.ParticipantList = chatParticipantList[request.Offset:request.Limit]
		} else {
			chatData.ParticipantList = chatParticipantList[request.Offset:]
		}
	}

	return types.MarshalAny(&chatData)
}
