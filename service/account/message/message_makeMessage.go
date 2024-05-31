/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time :
 * @File :
 */

package message

import (
	"novachat_engine/mtproto"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/mention"
	"novachat_engine/service/common/errors"
	"novachat_engine/service/constants"
	"novachat_engine/service/core/account/account"
	"novachat_engine/service/data/users"
	"novachat_engine/service/input"
	"novachat_engine/service/webpage"
	"strings"
	"time"
)

func (c *Core) makeMessage(userId int64, inputPeer *input.InputPeer, request *mtproto.TLMessagesSendMessage, dataUsers *data_users.Users) (*mtproto.Message, error) {
	message := mtproto.NewTLMessage(&mtproto.Message{
		Out:               true,
		Silent:            request.Silent,
		FromId286FA604119: input.MakeInputPeerValue(userId, constants.PeerTypeUser).ToPeer(),
		PeerId:            inputPeer.ToPeer(),
		Date:              int32(time.Now().Unix()),
		Message:           request.Message,
		Media:             mtproto.NewTLMessageMediaEmpty(nil).To_MessageMedia(),
		Entities:          make([]*mtproto.MessageEntity, 0, len(request.Entities)),
		FromId90DDDC1171:  constants.PeerTypeFromUserIDType(userId).ToInt32(),
		ToId:              inputPeer.ToPeer(),
		//ReplyTo: 			  request.ReplyMarkup
		ReplyMarkup: request.ReplyMarkup,
	})

	//TODO:Coder
	// is channel not chat/migrate chat
	// message.SetPost(true)
	// message.SetViews()  views counts

	for _, ent := range request.Entities {
		if ent.ClassName == mtproto.ClassMessageEntityMentionName {
			message.SetMentioned(true)
			entityMentionName := mtproto.NewTLMessageEntityMentionName(nil)
			entityMentionName.SetOffset(ent.Offset)
			entityMentionName.SetLength(ent.Length)
			entityMentionName.SetUserId352DCA5871(ent.UserId352DCA5871)
			message.Data2.Entities = append(message.Data2.Entities, entityMentionName.To_MessageEntity())
		} else {
			message.Data2.Entities = append(message.Data2.Entities, ent)
		}
	}

	//idxList := _split(request.Message)
	posList, ok := ParseUrl(request.Message)
	if ok == true {

		for _, v := range posList {
			entityUrl := mtproto.NewTLMessageEntityTextUrl(nil)
			entityUrl.SetOffset(v.Beg)
			entityUrl.SetLength(v.End)
			entityUrl.SetUrl(v.Text)

			message.Data2.Entities = append(message.Data2.Entities, entityUrl.To_MessageEntity())
		}

		canEmbedLink := true
		if canEmbedLink == true {
			media := mtproto.NewTLMessageMediaWebPage(nil)
			media.SetWebpage(webpage.GetWebPagePreview(request.Message))
			message.SetMedia(media.To_MessageMedia())
		}
	}

	var err error
	nameList, PosList, ok := ParseMention(request.Message)
	if ok == true {
		var userCacheList []string
		var userInfoList []*data_users.Users
		var u *data_users.Users
		userCacheList, err = c.accountCore.UsernameUserList(nameList)
		if err != nil {
			log.Warnf("UsernameUserList nameList:%v error:%s", nameList, err.Error())
			userInfoList, err = c.userCore.UsernameList(nameList)
			if err != nil {
				err = errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_INTERNAL)
				log.Errorf("SendMessage ReplyMessage error:%s", err.Error())
				return nil, err
			}
		} else {
			userInfoList = make([]*data_users.Users, 0, len(userCacheList))
			for _, v := range userCacheList {
				u, _ = account.CacheInfo2User(v)
				userInfoList = append(userInfoList, u)
			}
		}
		u = nil
		for idx, v := range PosList {

			for _, u1 := range userInfoList {
				if strings.Compare(u.GetUsername(), nameList[idx]) == 0 {
					u = u1
					break
				}
			}
			if u != nil {
				mention := mtproto.NewTLMessageEntityMention(nil)
				mention.SetOffset(v.Beg)
				mention.SetLength(v.End - v.Beg)
				mention.Data2.UserId352DCA5871 = constants.PeerTypeFromUserIDType(u.Id).ToInt32()
				mention.Data2.UserId208E68C971 = mtproto.NewTLInputUser(&mtproto.InputUser{
					UserId: constants.PeerTypeFromUserIDType(u.Id).ToInt32(),
				}).To_InputUser()

				message.Data2.Entities = append(message.Data2.Entities, mention.To_MessageEntity())
			}
		}
	}

	tags := mention.GetTags('#', request.GetMessage())
	for _, tag := range tags {
		hashtag := mtproto.NewTLMessageEntityHashtag(&mtproto.MessageEntity{
			Offset: int32(tag.Index),
			Length: int32(len(tag.Tag)),
		})
		message.Data2.Entities = append(message.Data2.Entities, hashtag.To_MessageEntity())
	}

	if inputPeer.GetPeerType() == constants.PeerTypeUser && (inputPeer.GetPeerId() == constants.BotFatherId || dataUsers != nil && dataUsers.Bot == 1) {
		index := strings.Index(request.GetMessage(), "/")
		commands := strings.Split(request.GetMessage(), " ")
		if len(commands[0]) > 1 && commands[0][0] == '/' {
			entity := mtproto.NewTLMessageEntityBotCommand(&mtproto.MessageEntity{
				Offset: int32(index) - 1,
				Length: int32(len(commands[0])),
			})
			message.Data2.Entities = append(message.Data2.Entities, entity.To_MessageEntity())
		}
	}

	return message.To_Message(), nil
}
