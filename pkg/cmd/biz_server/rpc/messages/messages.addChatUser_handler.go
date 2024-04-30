/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.addChatUser_handler.go
 * @Desc :
 *
 */

package rpc

import (
	"context"
	"github.com/gogo/protobuf/types"
	"novachat_engine/mtproto"
	chatService "novachat_engine/pkg/cmd/chat/rpc_client"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
	"novachat_engine/service/compat"
	"novachat_engine/service/constants"
	"time"
)

/*
 Send: { core_message
  msg_id: 6990261265084088676 [LONG],
  seq_no: 165 [INT],
  bytes: 28 [INT],
  body: { messages_addChatUser
    chat_id: 522114107 [INT],
    user_id: { inputUser
      user_id: 1358002840 [INT],
      access_hash: 8378484006885420883 [LONG],
    },
    fwd_limit: 100 [INT],
  },
}
Recv: { core_message
  msg_id: 6990261266965615617 [LONG],
  seq_no: 159 [INT],
  bytes: 576 [INT],
  body: { rpc_result
    req_msg_id: 6990261265084088676 [LONG],
    result: [GZIPPED] { updates
      updates: [ vector<0x0>
        { updateMessageID
          id: 202 [INT],
          random_id: 8634338456267912039 [LONG],
        },
        { updateChatParticipants
          participants: { chatParticipants
            chat_id: 522114107 [INT],
            participants: [ vector<0x0>
              { chatParticipant
                user_id: 1358002840 [INT],
                inviter_id: 1648887153 [INT],
                date: 1627547030 [INT],
              },
              { chatParticipant
                user_id: 1060156807 [INT],
                inviter_id: 1648887153 [INT],
                date: 1627546959 [INT],
              },
              { chatParticipant
                user_id: 1352087200 [INT],
                inviter_id: 1648887153 [INT],
                date: 1627546959 [INT],
              },
              { chatParticipantCreator
                user_id: 1648887153 [INT],
              },
            ],
            version: 2 [INT],
          },
        },
        { updateNewMessage
          message: { messageService
            flags: 258 [INT],
            out: YES [ BY BIT 1 IN FIELD flags ],
            mentioned: [ SKIPPED BY BIT 4 IN FIELD flags ],
            media_unread: [ SKIPPED BY BIT 5 IN FIELD flags ],
            silent: [ SKIPPED BY BIT 13 IN FIELD flags ],
            post: [ SKIPPED BY BIT 14 IN FIELD flags ],
            legacy: [ SKIPPED BY BIT 19 IN FIELD flags ],
            id: 202 [INT],
            from_id: { peerUser
              user_id: 1648887153 [INT],
            },
            peer_id: { peerChat
              chat_id: 522114107 [INT],
            },
            reply_to: [ SKIPPED BY BIT 3 IN FIELD flags ],
            date: 1627547030 [INT],
            action: { messageActionChatAddUser
              users: [ vector<0xffffffffa8509bda>
                1358002840 [INT],
              ],
            },
          },
          pts: 395 [INT],
          pts_count: 1 [INT],
        },
      ],
      users: [ vector<0x0>
        { user
          flags: 33560699 [INT],
          self: [ SKIPPED BY BIT 10 IN FIELD flags ],
          contact: YES [ BY BIT 11 IN FIELD flags ],
          mutual_contact: YES [ BY BIT 12 IN FIELD flags ],
          deleted: [ SKIPPED BY BIT 13 IN FIELD flags ],
          bot: [ SKIPPED BY BIT 14 IN FIELD flags ],
          bot_chat_history: [ SKIPPED BY BIT 15 IN FIELD flags ],
          bot_nochats: [ SKIPPED BY BIT 16 IN FIELD flags ],
          verified: [ SKIPPED BY BIT 17 IN FIELD flags ],
          restricted: [ SKIPPED BY BIT 18 IN FIELD flags ],
          min: [ SKIPPED BY BIT 20 IN FIELD flags ],
          bot_inline_geo: [ SKIPPED BY BIT 21 IN FIELD flags ],
          support: [ SKIPPED BY BIT 23 IN FIELD flags ],
          scam: [ SKIPPED BY BIT 24 IN FIELD flags ],
          apply_min_photo: YES [ BY BIT 25 IN FIELD flags ],
          id: 1358002840 [INT],
          access_hash: 8378484006885420883 [LONG],
          first_name: "alacom" [STRING],
          last_name: [ SKIPPED BY BIT 2 IN FIELD flags ],
          username: "alacom_baggins" [STRING],
          phone: "8618675553281" [STRING],
          photo: { userProfilePhoto
            flags: 0 [INT],
            has_video: [ SKIPPED BY BIT 0 IN FIELD flags ],
            photo_id: 6075578944020655064 [LONG],
            photo_small: { fileLocationToBeDeprecated
              volume_id: 500077400791 [LONG],
              local_id: 106001 [INT],
            },
            photo_big: { fileLocationToBeDeprecated
              volume_id: 500077400791 [LONG],
              local_id: 106003 [INT],
            },
            dc_id: 5 [INT],
          },
          status: { userStatusOffline
            was_online: 1627456422 [INT],
          },
          bot_info_version: [ SKIPPED BY BIT 14 IN FIELD flags ],
          restriction_reason: [ SKIPPED BY BIT 18 IN FIELD flags ],
          bot_inline_placeholder: [ SKIPPED BY BIT 19 IN FIELD flags ],
          lang_code: [ SKIPPED BY BIT 22 IN FIELD flags ],
        },
        { user
          flags: 33561723 [INT],
          self: YES [ BY BIT 10 IN FIELD flags ],
          contact: YES [ BY BIT 11 IN FIELD flags ],
          mutual_contact: YES [ BY BIT 12 IN FIELD flags ],
          deleted: [ SKIPPED BY BIT 13 IN FIELD flags ],
          bot: [ SKIPPED BY BIT 14 IN FIELD flags ],
          bot_chat_history: [ SKIPPED BY BIT 15 IN FIELD flags ],
          bot_nochats: [ SKIPPED BY BIT 16 IN FIELD flags ],
          verified: [ SKIPPED BY BIT 17 IN FIELD flags ],
          restricted: [ SKIPPED BY BIT 18 IN FIELD flags ],
          min: [ SKIPPED BY BIT 20 IN FIELD flags ],
          bot_inline_geo: [ SKIPPED BY BIT 21 IN FIELD flags ],
          support: [ SKIPPED BY BIT 23 IN FIELD flags ],
          scam: [ SKIPPED BY BIT 24 IN FIELD flags ],
          apply_min_photo: YES [ BY BIT 25 IN FIELD flags ],
          id: 1648887153 [INT],
          access_hash: 1113455867854165891 [LONG],
          first_name: "D" [STRING],
          last_name: [ SKIPPED BY BIT 2 IN FIELD flags ],
          username: "b" [STRING],
          phone: "86" [STRING],
          photo: { userProfilePhoto
            flags: 0 [INT],
            has_video: [ SKIPPED BY BIT 0 IN FIELD flags ],
            photo_id: 6316686472303979484 [LONG],
            photo_small: { fileLocationToBeDeprecated
              volume_id: 500077000982 [LONG],
              local_id: 248125 [INT],
            },
            photo_big: { fileLocationToBeDeprecated
              volume_id: 500077000982 [LONG],
              local_id: 248127 [INT],
            },
            dc_id: 5 [INT],
          },
          status: { userStatusOnline
            expires: 1627547310 [INT],
          },
          bot_info_version: [ SKIPPED BY BIT 14 IN FIELD flags ],
          restriction_reason: [ SKIPPED BY BIT 18 IN FIELD flags ],
          bot_inline_placeholder: [ SKIPPED BY BIT 19 IN FIELD flags ],
          lang_code: [ SKIPPED BY BIT 22 IN FIELD flags ],
        },
        { user
          flags: 33560701 [INT],
          self: [ SKIPPED BY BIT 10 IN FIELD flags ],
          contact: YES [ BY BIT 11 IN FIELD flags ],
          mutual_contact: YES [ BY BIT 12 IN FIELD flags ],
          deleted: [ SKIPPED BY BIT 13 IN FIELD flags ],
          bot: [ SKIPPED BY BIT 14 IN FIELD flags ],
          bot_chat_history: [ SKIPPED BY BIT 15 IN FIELD flags ],
          bot_nochats: [ SKIPPED BY BIT 16 IN FIELD flags ],
          verified: [ SKIPPED BY BIT 17 IN FIELD flags ],
          restricted: [ SKIPPED BY BIT 18 IN FIELD flags ],
          min: [ SKIPPED BY BIT 20 IN FIELD flags ],
          bot_inline_geo: [ SKIPPED BY BIT 21 IN FIELD flags ],
          support: [ SKIPPED BY BIT 23 IN FIELD flags ],
          scam: [ SKIPPED BY BIT 24 IN FIELD flags ],
          apply_min_photo: YES [ BY BIT 25 IN FIELD flags ],
          id: 1060156807 [INT],
          access_hash: 7685283431854918581 [LONG],
          first_name: [ SKIPPED BY BIT 1 IN FIELD flags ],
          last_name: "z" [STRING],
          username: "q" [STRING],
          phone: "861" [STRING],
          photo: { userProfilePhoto
            flags: 0 [INT],
            has_video: [ SKIPPED BY BIT 0 IN FIELD flags ],
            photo_id: 4553338815153022889 [LONG],
            photo_small: { fileLocationToBeDeprecated
              volume_id: 500137900129 [LONG],
              local_id: 97042 [INT],
            },
            photo_big: { fileLocationToBeDeprecated
              volume_id: 500137900129 [LONG],
              local_id: 97044 [INT],
            },
            dc_id: 5 [INT],
          },
          status: { userStatusOffline
            was_online: 1627459039 [INT],
          },
          bot_info_version: [ SKIPPED BY BIT 14 IN FIELD flags ],
          restriction_reason: [ SKIPPED BY BIT 18 IN FIELD flags ],
          bot_inline_placeholder: [ SKIPPED BY BIT 19 IN FIELD flags ],
          lang_code: [ SKIPPED BY BIT 22 IN FIELD flags ],
        },
        { user
          flags: 33560699 [INT],
          self: [ SKIPPED BY BIT 10 IN FIELD flags ],
          contact: YES [ BY BIT 11 IN FIELD flags ],
          mutual_contact: YES [ BY BIT 12 IN FIELD flags ],
          deleted: [ SKIPPED BY BIT 13 IN FIELD flags ],
          bot: [ SKIPPED BY BIT 14 IN FIELD flags ],
          bot_chat_history: [ SKIPPED BY BIT 15 IN FIELD flags ],
          bot_nochats: [ SKIPPED BY BIT 16 IN FIELD flags ],
          verified: [ SKIPPED BY BIT 17 IN FIELD flags ],
          restricted: [ SKIPPED BY BIT 18 IN FIELD flags ],
          min: [ SKIPPED BY BIT 20 IN FIELD flags ],
          bot_inline_geo: [ SKIPPED BY BIT 21 IN FIELD flags ],
          support: [ SKIPPED BY BIT 23 IN FIELD flags ],
          scam: [ SKIPPED BY BIT 24 IN FIELD flags ],
          apply_min_photo: YES [ BY BIT 25 IN FIELD flags ],
          id: 1352087200 [INT],
          access_hash: 15240523245817097092 [LONG],
          first_name: "f" [STRING],
          last_name: [ SKIPPED BY BIT 2 IN FIELD flags ],
          username: "g" [STRING],
          phone: "862" [STRING],
          photo: { userProfilePhoto
            flags: 0 [INT],
            has_video: [ SKIPPED BY BIT 0 IN FIELD flags ],
            photo_id: 6319051410146176196 [LONG],
            photo_small: { fileLocationToBeDeprecated
              volume_id: 500172700916 [LONG],
              local_id: 120531 [INT],
            },
            photo_big: { fileLocationToBeDeprecated
              volume_id: 500172700916 [LONG],
              local_id: 120533 [INT],
            },
            dc_id: 5 [INT],
          },
          status: { userStatusOffline
            was_online: 1626835841 [INT],
          },
          bot_info_version: [ SKIPPED BY BIT 14 IN FIELD flags ],
          restriction_reason: [ SKIPPED BY BIT 18 IN FIELD flags ],
          bot_inline_placeholder: [ SKIPPED BY BIT 19 IN FIELD flags ],
          lang_code: [ SKIPPED BY BIT 22 IN FIELD flags ],
        },
      ],
      chats: [ vector<0x0>
        { chat
          flags: 262145 [INT],
          creator: YES [ BY BIT 0 IN FIELD flags ],
          kicked: [ SKIPPED BY BIT 1 IN FIELD flags ],
          left: [ SKIPPED BY BIT 2 IN FIELD flags ],
          deactivated: [ SKIPPED BY BIT 5 IN FIELD flags ],
          call_active: [ SKIPPED BY BIT 23 IN FIELD flags ],
          call_not_empty: [ SKIPPED BY BIT 24 IN FIELD flags ],
          id: 522114107 [INT],
          title: "dddd test" [STRING],
          photo: { chatPhoto
            flags: 0 [INT],
            has_video: [ SKIPPED BY BIT 0 IN FIELD flags ],
            photo_small: { fileLocationToBeDeprecated
              volume_id: 500158500201 [LONG],
              local_id: 229219 [INT],
            },
            photo_big: { fileLocationToBeDeprecated
              volume_id: 500158500201 [LONG],
              local_id: 229221 [INT],
            },
            dc_id: 5 [INT],
          },
          participants_count: 4 [INT],
          date: 1627546959 [INT],
          version: 2 [INT],
          migrated_to: [ SKIPPED BY BIT 6 IN FIELD flags ],
          admin_rights: [ SKIPPED BY BIT 14 IN FIELD flags ],
          default_banned_rights: { chatBannedRights
            flags: 0 [INT],
            view_messages: [ SKIPPED BY BIT 0 IN FIELD flags ],
            send_messages: [ SKIPPED BY BIT 1 IN FIELD flags ],
            send_media: [ SKIPPED BY BIT 2 IN FIELD flags ],
            send_stickers: [ SKIPPED BY BIT 3 IN FIELD flags ],
            send_gifs: [ SKIPPED BY BIT 4 IN FIELD flags ],
            send_games: [ SKIPPED BY BIT 5 IN FIELD flags ],
            send_inline: [ SKIPPED BY BIT 6 IN FIELD flags ],
            embed_links: [ SKIPPED BY BIT 7 IN FIELD flags ],
            send_polls: [ SKIPPED BY BIT 8 IN FIELD flags ],
            change_info: [ SKIPPED BY BIT 10 IN FIELD flags ],
            invite_users: [ SKIPPED BY BIT 15 IN FIELD flags ],
            pin_messages: [ SKIPPED BY BIT 17 IN FIELD flags ],
            until_date: 2147483647 [INT],
          },
        },
      ],
      date: 1627547029 [INT],
      seq: 0 [INT],
    },
  },
}
*/

//  messages.addChatUser#f9a0aa09 chat_id:int user_id:InputUser fwd_limit:int = Updates;
//
func (s *MessagesServiceImpl) MessagesAddChatUser(ctx context.Context, request *mtproto.TLMessagesAddChatUser) (*mtproto.Updates, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Debugf("MessagesAddChatUser %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	addUserResp, err := chatService.GetChatClientByKeyId(constants.PeerTypeFromChatIDType32(request.ChatId).ToInt()).ReqAddUser(
		ctx,
		&chatService.AddUser{
			UserId:    md.UserId,
			ChatId:    constants.PeerTypeFromChatIDType32(request.ChatId).ToInt(),
			PeerId:    constants.PeerTypeFromUserIDType32(request.UserId.UserId).ToInt(),
			AuthKeyId: md.AuthKeyId,
			FwdLimit:  request.FwdLimit,
			Date:      int32(time.Now().Unix()),
		})

	if err != nil {
		log.Errorf("MessagesAddChatUser %v, request: %v error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
		return nil, err
	}

	var updates mtproto.Updates
	types.UnmarshalAny(addUserResp, &updates)
	log.Infof("MessagesAddChatUser %v, request: %v ok!!!!!!!!!!!!!!!!!", metadata.RpcMetaDataDebug(md), request)
	return compat.UpdatesCompat(&updates, md.Layer), nil
}
