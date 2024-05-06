package rpc

import (
	"context"
	"errors"
	"github.com/gogo/protobuf/types"
	"novachat_engine/mtproto"
	chatService "novachat_engine/pkg/cmd/chat/rpc_client"
	"novachat_engine/pkg/cmd/chat/service"
	syncService "novachat_engine/pkg/cmd/sync/rpc_client"
	"novachat_engine/pkg/log"
	errorsService "novachat_engine/service/common/errors"
	"novachat_engine/service/constants"
	data_chat "novachat_engine/service/data/chat"
)

func (impl *Impl) ReqDeleteChat(ctx context.Context, request *chatService.DeleteChat) (*types.Any, error) {
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

	chatData := chat.GetChatInfo().ChatData
	if chatData.Creator != request.UserId {
		participant := chat.GetChatInfo().GetParticipants(request.UserId)
		if participant == nil {
			return nil, errorsService.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_CHANNEL_PRIVATE)
		}
		if !participant.Admin {
			return nil, errorsService.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_CHAT_ADMIN_REQUIRED)
		}
	}

	_, _, err = chat.EditDeleted()
	if err != nil {
		log.Errorf("ReqDeleteChat request:%+v", request, err.Error())
		return nil, err
	}

	updates := mtproto.NewTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{
			mtproto.NewTLUpdateChannel(&mtproto.Update{
				ChannelId: constants.PeerTypeFromChatIDType(request.ChatId).ToInt32(),
			}).To_Update(),
		},
	}).To_Updates()
	updateDataList := make([]*syncService.UpdateData, 0, chat.GetChatInfo().Count)
	chat.GetChatInfo().Iteration(func(participant *data_chat.ChatParticipant) {
		if request.UserId == participant.UserId {
			updateDataList = append(updateDataList, &syncService.UpdateData{
				UserId:          participant.UserId,
				IgnoreAuthKeyId: request.AuthKeyId,
				Updates:         updates,
			})
		} else {
			updateDataList = append(updateDataList, &syncService.UpdateData{
				UserId:  participant.UserId,
				Updates: updates,
			})
		}
	})

	_, err = syncService.GetSyncClientById(request.UserId).ReqSyncUpdates(ctx, &syncService.SyncUpdates{
		UpdateDataList: updateDataList,
		Push:           false,
		PeerType:       constants.PeerTypeUser.ToInt32(),
	})
	if err != nil {
		log.Warnf("ReqDeleteChat ReqSyncUpdates:%+v", request, err.Error())
	}

	return types.MarshalAny(updates)
}
