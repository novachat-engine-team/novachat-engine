/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.editChatPhoto_handler.go
 * @Desc :
 *
 */

package rpc

import (
	"context"
	"fmt"
	"github.com/gogo/protobuf/types"
	"novachat_engine/mtproto"
	chatService "novachat_engine/pkg/cmd/chat/rpc_client"
	sfsService "novachat_engine/pkg/cmd/sfs/rpc_client"
	"novachat_engine/pkg/cmd/sfs/utils"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
	"novachat_engine/service/common/errors"
	"novachat_engine/service/compat"
	"novachat_engine/service/constants"
	data_fs "novachat_engine/service/data/fs"
	"time"
)

/*
Send: { core_message
  msg_id: 6990261149009436260 [LONG],
  seq_no: 154 [INT],
  bytes: 136 [INT],
  body: { msg_container
    messages: [ vector<0xffffffffffffffff>
      { core_message
        msg_id: 6990261149009136468 [LONG],
        seq_no: 153 [INT],
        bytes: 76 [INT],
        body: { messages_editChatPhoto
          chat_id: 522114107 [INT],
          photo: { inputChatUploadedPhoto
            flags: 1 [INT],
            file: { inputFile
              id: 13941031204537930410 [LONG],
              parts: 1 [INT],
              name: ".jpg" [STRING],
              md5_checksum: "75c2dc5e46c5f03436eb23872d4cd30b" [STRING],
            },
            video: [ SKIPPED BY BIT 1 IN FIELD flags ],
            video_start_ts: [ SKIPPED BY BIT 2 IN FIELD flags ],
          },
        },
      },
    ],
  },
}

Recv: { core_message
  msg_id: 6990261149043490817 [LONG],
  seq_no: 147 [INT],
  bytes: 620 [INT],
  body: { rpc_result
    req_msg_id: 6990261149009136468 [LONG],
    result: [GZIPPED] { updates
      updates: [ vector<0x0>
        { updateMessageID
          id: 201 [INT],
          random_id: 3220297111532950832 [LONG],
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
            id: 201 [INT],
            from_id: { peerUser
              user_id: 1648887153 [INT],
            },
            peer_id: { peerChat
              chat_id: 522114107 [INT],
            },
            reply_to: [ SKIPPED BY BIT 3 IN FIELD flags ],
            date: 1627547002 [INT],
            action: { messageActionChatEditPhoto
              photo: { photo
                flags: 0 [INT],
                has_stickers: [ SKIPPED BY BIT 0 IN FIELD flags ],
                id: 6059637081979858453 [LONG],
                access_hash: 8211330806298609294 [LONG],
                file_reference: 03 00 00 00 C9 61 02 65 7A C5 E6 BB C7 DE EB 89 49 D5 66 59 CA 4D 24 EA 38 [25 BYTES],
                date: 1627547002 [INT],
                sizes: [ vector<0x0>
                  { photoSize
                    type: "a" [STRING],
                    location: { fileLocationToBeDeprecated
                      volume_id: 500158500201 [LONG],
                      local_id: 229219 [INT],
                    },
                    w: 160 [INT],
                    h: 160 [INT],
                    size: 3686 [INT],
                  },
                  { photoSize
                    type: "b" [STRING],
                    location: { fileLocationToBeDeprecated
                      volume_id: 500158500201 [LONG],
                      local_id: 229220 [INT],
                    },
                    w: 320 [INT],
                    h: 320 [INT],
                    size: 6315 [INT],
                  },
                  { photoSize
                    type: "c" [STRING],
                    location: { fileLocationToBeDeprecated
                      volume_id: 500158500201 [LONG],
                      local_id: 229221 [INT],
                    },
                    w: 640 [INT],
                    h: 640 [INT],
                    size: 17922 [INT],
                  },
                  { photoStrippedSize
                    type: "i" [STRING],
                    bytes: 01 28 28 C5 A2 A5 58 25 68 1A 60 84 C6 A7 05 AA... [207 BYTES],
                  },
                ],
                video_sizes: [ SKIPPED BY BIT 1 IN FIELD flags ],
                dc_id: 5 [INT],
              },
            },
          },
          pts: 394 [INT],
          pts_count: 1 [INT],
        },
      ],
      users: [ vector<0x0>
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
            expires: 1627547301 [INT],
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
      date: 1627547001 [INT],
      seq: 0 [INT],
    },
  },
}
*/

//  messages.editChatPhoto#ca4c79d8 chat_id:int photo:InputChatPhoto = Updates;
//
func (s *MessagesServiceImpl) MessagesEditChatPhoto(ctx context.Context, request *mtproto.TLMessagesEditChatPhoto) (*mtproto.Updates, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Debugf("MessagesEditChatPhoto %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	var err error
	//var action *mtproto.MessageAction
	//var photoId int64
	var photoInfo *sfsService.PhotoInfo
	dataChatPhoto := &data_fs.PhotoProfile{}

	chatPhoto := request.GetPhoto()
	switch chatPhoto.ClassName {
	case mtproto.ClassInputChatPhotoEmpty:
		//action = mtproto.NewTLMessageActionChatDeletePhoto(nil).To_MessageAction()
	case mtproto.ClassInputChatUploadedPhoto:
		var documentInfo *sfsService.DocumentInfo
		if request.Photo.Video != nil {
			documentInfo, err = sfsService.GetSfsClient(fmt.Sprintf("%d", chatPhoto.GetVideo().Id)).
				ReqUploadedDocument(ctx,
					&sfsService.UploadedDocument{
						AuthKeyId: md.AuthKeyId,
						FileId:    chatPhoto.GetVideo().Id,
						FileName:  chatPhoto.GetVideo().Name,
						Parts:     chatPhoto.GetVideo().Parts,
						Thumb: &sfsService.UploadedDocument_Thumb{
							FileId:   chatPhoto.GetFile().Id,
							FileName: chatPhoto.GetFile().Name,
							Parts:    chatPhoto.GetFile().Parts,
							Sticker:  nil,
						},
						VideoStartTs: chatPhoto.GetVideoStartTs(),
						Video:        true,
					})
		} else {
			photoInfo, err = sfsService.GetSfsClient(fmt.Sprintf("%d", chatPhoto.GetFile().Id)).
				ReqUploadPhoto(ctx,
					&sfsService.UploadPhoto{
						AuthKeyId: md.AuthKeyId,
						FileId:    chatPhoto.GetFile().Id,
						FileName:  chatPhoto.GetFile().Name,
						Parts:     chatPhoto.GetFile().Parts,
						Stickers:  nil,
					})
		}
		if err != nil {
			log.Errorf("MessagesEditChatPhoto error: %v", err)
			return nil, err
		}
		if documentInfo == nil {
			dataChatPhoto.Photo = utils.ToDataPhoto(photoInfo).Detail[0]
		} else {
			dataChatPhoto.Photo = utils.ToDataPhoto(documentInfo.Thumbs[0]).Detail[0]
			dataChatPhoto.Video = utils.DocumentToVideo(documentInfo).Detail[0]
		}
	case mtproto.ClassInputChatPhoto:
		photoInfo, err = sfsService.GetSfsClient(fmt.Sprintf("%d", chatPhoto.GetId().Id)).
			ReqGetPhoto(context.Background(),
				&sfsService.GetPhoto{
					PhotoId: chatPhoto.GetId().Id,
				})
		if err != nil {
			log.Errorf("MessagesEditChatPhoto InputChatPhoto error: %v", err)
			return nil, err
		}
		dataChatPhoto.Photo = utils.ToDataPhoto(photoInfo).Detail[0]
	}

	chatId := constants.PeerTypeFromChatIDType32(request.ChatId).ToInt()
	if chatId == 0 {
		err := errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_PEER_ID_INVALID)
		log.Errorf("MessagesEditChatPhoto %v, request: %v error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
		return nil, err
	}

	editPhotoResp, err := chatService.GetChatClientByKeyId(chatId).
		ReqEditPhoto(ctx, &chatService.EditPhoto{
			ChatId:    chatId,
			Photo:     dataChatPhoto,
			UserId:    md.UserId,
			AuthKeyId: md.AuthKeyId,
			Date:      int32(time.Now().Unix()),
		})
	if err != nil {
		log.Errorf("MessagesEditChatTitle %v, request: %v ReqEditTitle error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
		return nil, err
	}
	var updates mtproto.Updates
	types.UnmarshalAny(editPhotoResp, &updates)
	return compat.UpdatesCompat(&updates, md.Layer), nil
}
