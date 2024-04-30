/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.createChat_handler.go
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
	"novachat_engine/service/app/help"
	"novachat_engine/service/common/errors"
	"novachat_engine/service/constants"
	"strings"
	"time"
)

/*
Send: { core_message
  msg_id: 6990260960371023972 [LONG],
  seq_no: 132 [INT],
  bytes: 132 [INT],
  body: { msg_container
    messages: [ vector<0xffffffffffffffff>
      { core_message
        msg_id: 6990260960370715588 [LONG],
        seq_no: 131 [INT],
        bytes: 56 [INT],
        body: { messages_createChat
          users: [ vector<0x0>
            { inputUser
              user_id: 1352087200 [INT],
              access_hash: 15240523245817097092 [LONG],
            },
            { inputUser
              user_id: 1060156807 [INT],
              access_hash: 7685283431854918581 [LONG],
            },
          ],
          title: "dddd test" [STRING],
        },
      },
    ],
  },
}

Recv: { core_message
  msg_id: 6990260963101946881 [LONG],
  seq_no: 127 [INT],
  bytes: 476 [INT],
  body: { rpc_result
    req_msg_id: 6990260960370715588 [LONG],
    result: [GZIPPED] { updates
      updates: [ vector<0x0>
        { updateMessageID
          id: 200 [INT],
          random_id: 1371323332285119233 [LONG],
        },
        { updateChatParticipants
          participants: { chatParticipants
            chat_id: 522114107 [INT],
            participants: [ vector<0x0>
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
            version: 1 [INT],
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
            id: 200 [INT],
            from_id: { peerUser
              user_id: 1648887153 [INT],
            },
            peer_id: { peerChat
              chat_id: 522114107 [INT],
            },
            reply_to: [ SKIPPED BY BIT 3 IN FIELD flags ],
            date: 1627546959 [INT],
            action: { messageActionChatCreate
              title: "dddd test" [STRING],
              users: [ vector<0xffffffffa8509bda>
                1648887153 [INT],
                1352087200 [INT],
                1060156807 [INT],
              ],
            },
          },
          pts: 393 [INT],
          pts_count: 1 [INT],
        },
      ],
      users: [ vector<0x0>
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
            expires: 1627547147 [INT],
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
          photo: { chatPhotoEmpty },
          participants_count: 3 [INT],
          date: 1627546959 [INT],
          version: 1 [INT],
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
      date: 1627546958 [INT],
      seq: 0 [INT],
    },
  },
}


Recv: { core_message
  msg_id: 6990261488012345345 [LONG],
  seq_no: 171 [INT],
  bytes: 452 [INT],
  body: [GZIPPED] { updates
    updates: [ vector<0x0>
      { updateChannel
        channel_id: 1538335065 [INT],
      },
      { updateNotifySettings
        peer: { notifyPeer
          peer: { peerChannel
            channel_id: 1538335065 [INT],
          },
        },
        notify_settings: { peerNotifySettings
          flags: 0 [INT],
          show_previews: [ SKIPPED BY BIT 0 IN FIELD flags ],
          silent: [ SKIPPED BY BIT 1 IN FIELD flags ],
          mute_until: [ SKIPPED BY BIT 2 IN FIELD flags ],
          sound: [ SKIPPED BY BIT 3 IN FIELD flags ],
        },
      },
      { updateReadChannelInbox
        flags: 0 [INT],
        folder_id: [ SKIPPED BY BIT 0 IN FIELD flags ],
        channel_id: 1538335065 [INT],
        max_id: 1 [INT],
        still_unread_count: 0 [INT],
        pts: 2 [INT],
      },
      { updateNewChannelMessage
        message: { messageService
          flags: 2 [INT],
          out: YES [ BY BIT 1 IN FIELD flags ],
          mentioned: [ SKIPPED BY BIT 4 IN FIELD flags ],
          media_unread: [ SKIPPED BY BIT 5 IN FIELD flags ],
          silent: [ SKIPPED BY BIT 13 IN FIELD flags ],
          post: [ SKIPPED BY BIT 14 IN FIELD flags ],
          legacy: [ SKIPPED BY BIT 19 IN FIELD flags ],
          id: 1 [INT],
          from_id: [ SKIPPED BY BIT 8 IN FIELD flags ],
          peer_id: { peerChannel
            channel_id: 1538335065 [INT],
          },
          reply_to: [ SKIPPED BY BIT 3 IN FIELD flags ],
          date: 1627547081 [INT],
          action: { messageActionChannelMigrateFrom
            title: "dddd test" [STRING],
            chat_id: 522114107 [INT],
          },
        },
        pts: 2 [INT],
        pts_count: 1 [INT],
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
          id: 204 [INT],
          from_id: { peerUser
            user_id: 1648887153 [INT],
          },
          peer_id: { peerChat
            chat_id: 522114107 [INT],
          },
          reply_to: [ SKIPPED BY BIT 3 IN FIELD flags ],
          date: 1627547081 [INT],
          action: { messageActionChatMigrateTo
            channel_id: 1538335065 [INT],
          },
        },
        pts: 398 [INT],
        pts_count: 1 [INT],
      },
      { updateReadHistoryOutbox
        peer: { peerChat
          chat_id: 522114107 [INT],
        },
        max_id: 204 [INT],
        pts: 399 [INT],
        pts_count: 1 [INT],
      },
    ],
    users: [ vector<0x0>
      { user
        flags: 33636395 [INT],
        self: [ SKIPPED BY BIT 10 IN FIELD flags ],
        contact: [ SKIPPED BY BIT 11 IN FIELD flags ],
        mutual_contact: [ SKIPPED BY BIT 12 IN FIELD flags ],
        deleted: [ SKIPPED BY BIT 13 IN FIELD flags ],
        bot: YES [ BY BIT 14 IN FIELD flags ],
        bot_chat_history: [ SKIPPED BY BIT 15 IN FIELD flags ],
        bot_nochats: YES [ BY BIT 16 IN FIELD flags ],
        verified: [ SKIPPED BY BIT 17 IN FIELD flags ],
        restricted: [ SKIPPED BY BIT 18 IN FIELD flags ],
        min: [ SKIPPED BY BIT 20 IN FIELD flags ],
        bot_inline_geo: [ SKIPPED BY BIT 21 IN FIELD flags ],
        support: [ SKIPPED BY BIT 23 IN FIELD flags ],
        scam: [ SKIPPED BY BIT 24 IN FIELD flags ],
        apply_min_photo: YES [ BY BIT 25 IN FIELD flags ],
        id: 1087968824 [INT],
        access_hash: 14916718766862931948 [LONG],
        first_name: "Group" [STRING],
        last_name: [ SKIPPED BY BIT 2 IN FIELD flags ],
        username: "GroupAnonymousBot" [STRING],
        phone: [ SKIPPED BY BIT 4 IN FIELD flags ],
        photo: { userProfilePhoto
          flags: 0 [INT],
          has_video: [ SKIPPED BY BIT 0 IN FIELD flags ],
          photo_id: 5159307831025969322 [LONG],
          photo_small: { fileLocationToBeDeprecated
            volume_id: 806529792 [LONG],
            local_id: 188482 [INT],
          },
          photo_big: { fileLocationToBeDeprecated
            volume_id: 806529792 [LONG],
            local_id: 188484 [INT],
          },
          dc_id: 1 [INT],
        },
        status: [ SKIPPED BY BIT 6 IN FIELD flags ],
        bot_info_version: 3 [INT],
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
          expires: 1627547336 [INT],
        },
        bot_info_version: [ SKIPPED BY BIT 14 IN FIELD flags ],
        restriction_reason: [ SKIPPED BY BIT 18 IN FIELD flags ],
        bot_inline_placeholder: [ SKIPPED BY BIT 19 IN FIELD flags ],
        lang_code: [ SKIPPED BY BIT 22 IN FIELD flags ],
      },
    ],
    chats: [ vector<0x0>
      { channel
        flags: 286977 [INT],
        creator: YES [ BY BIT 0 IN FIELD flags ],
        left: [ SKIPPED BY BIT 2 IN FIELD flags ],
        broadcast: [ SKIPPED BY BIT 5 IN FIELD flags ],
        verified: [ SKIPPED BY BIT 7 IN FIELD flags ],
        megagroup: YES [ BY BIT 8 IN FIELD flags ],
        restricted: [ SKIPPED BY BIT 9 IN FIELD flags ],
        signatures: [ SKIPPED BY BIT 11 IN FIELD flags ],
        min: [ SKIPPED BY BIT 12 IN FIELD flags ],
        scam: [ SKIPPED BY BIT 19 IN FIELD flags ],
        has_link: [ SKIPPED BY BIT 20 IN FIELD flags ],
        has_geo: [ SKIPPED BY BIT 21 IN FIELD flags ],
        slowmode_enabled: [ SKIPPED BY BIT 22 IN FIELD flags ],
        call_active: [ SKIPPED BY BIT 23 IN FIELD flags ],
        call_not_empty: [ SKIPPED BY BIT 24 IN FIELD flags ],
        id: 1538335065 [INT],
        access_hash: 473190190564556043 [LONG],
        title: "dddd test" [STRING],
        username: [ SKIPPED BY BIT 6 IN FIELD flags ],
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
        date: 1627547081 [INT],
        version: 0 [INT],
        restriction_reason: [ SKIPPED BY BIT 9 IN FIELD flags ],
        admin_rights: { chatAdminRights
          flags: 6847 [INT],
          change_info: YES [ BY BIT 0 IN FIELD flags ],
          post_messages: YES [ BY BIT 1 IN FIELD flags ],
          edit_messages: YES [ BY BIT 2 IN FIELD flags ],
          delete_messages: YES [ BY BIT 3 IN FIELD flags ],
          ban_users: YES [ BY BIT 4 IN FIELD flags ],
          invite_users: YES [ BY BIT 5 IN FIELD flags ],
          pin_messages: YES [ BY BIT 7 IN FIELD flags ],
          add_admins: YES [ BY BIT 9 IN FIELD flags ],
          anonymous: [ SKIPPED BY BIT 10 IN FIELD flags ],
          manage_call: YES [ BY BIT 11 IN FIELD flags ],
        },
        banned_rights: [ SKIPPED BY BIT 15 IN FIELD flags ],
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
        participants_count: [ SKIPPED BY BIT 17 IN FIELD flags ],
      },
      { chat
        flags: 97 [INT],
        creator: YES [ BY BIT 0 IN FIELD flags ],
        kicked: [ SKIPPED BY BIT 1 IN FIELD flags ],
        left: [ SKIPPED BY BIT 2 IN FIELD flags ],
        deactivated: YES [ BY BIT 5 IN FIELD flags ],
        call_active: [ SKIPPED BY BIT 23 IN FIELD flags ],
        call_not_empty: [ SKIPPED BY BIT 24 IN FIELD flags ],
        id: 522114107 [INT],
        title: "dddd test" [STRING],
        photo: { chatPhotoEmpty },
        participants_count: 0 [INT],
        date: 1627546959 [INT],
        version: 3 [INT],
        migrated_to: { inputChannel
          channel_id: 1538335065 [INT],
          access_hash: 473190190564556043 [LONG],
        },
        admin_rights: [ SKIPPED BY BIT 14 IN FIELD flags ],
        default_banned_rights: [ SKIPPED BY BIT 18 IN FIELD flags ],
      },
    ],
    date: 1627547080 [INT],
    seq: 7 [INT],
  },
}
*/

//  messages.createChat#9cb126e users:Vector<InputUser> title:string = Updates;
//
func (s *MessagesServiceImpl) MessagesCreateChat(ctx context.Context, request *mtproto.TLMessagesCreateChat) (*mtproto.Updates, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Debugf("MessagesCreateChat %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	tlUser, err := s.accountUsersCore.GetUser(md.UserId, md.UserId)
	if err != nil {
		err = errors.NewRpcErrorWithRpcErrorCodeString(mtproto.RpcErrorCode_INTERNAL, err.Error())
		log.Errorf("MessagesCreateChat %v, request: %v GetByUserId error:%s", md, request, err.Error())
		return nil, err
	}
	if tlUser.Restricted || tlUser.Deleted {
		err = errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_UNAUTHORIZED_AUTH_KEY_UNREGISTERED)
		return nil, err
	}

	title := strings.TrimSpace(request.Title)
	if len(title) == 0 {
		err = errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_CHAT_TITLE_EMPTY)
		log.Errorf("MessagesCreateChat %v, request: %v error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
		return nil, err
	}

	userList := make([]int64, 0, len(request.Users))
	for _, v := range request.Users {
		userList = append(userList, constants.PeerTypeFromUserIDType32(v.UserId).ToInt())
	}

	createChatResp, err := chatService.GetChatClientByKeyId(md.UserId).ReqCreateChat(context.Background(),
		&chatService.CreateChat{
			UserId:      md.UserId,
			AuthKeyId:   md.AuthKeyId,
			PeerId:      userList,
			Title:       title,
			ChatSizeMax: help.GetConfig().ChatSizeMax,
			Date:        int32(time.Now().Unix()),
		})
	if err != nil {
		log.Errorf("MessagesCreateChat %v, request: %v CreateChat error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
		return nil, err
	}

	var updates mtproto.Updates
	types.UnmarshalAny(createChatResp, &updates)
	log.Infof("MessagesCreateChat %v, request: %v OK!!!!!!!!!!!!!", metadata.RpcMetaDataDebug(md), request)
	return &updates, nil
}
