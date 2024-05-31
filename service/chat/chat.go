package chat

import (
	"novachat_engine/mtproto"
	chatService "novachat_engine/pkg/cmd/chat/rpc_client"
	"novachat_engine/service/app/dc"
	"novachat_engine/service/banned_right"
	"novachat_engine/service/compat"
	"novachat_engine/service/constants"
	data_chat "novachat_engine/service/data/chat"
)

func ToChat(userId int64, dataChat *chatService.Chat, layer int32) *mtproto.Chat {
	chat := mtproto.NewTLChannel(&mtproto.Chat{
		Creator:               dataChat.ChatData.Creator == userId,
		Kicked:                false,
		Megagroup:             true,
		Title:                 dataChat.ChatData.Title,
		Left:                  false,
		Deactivated:           false,
		CallActive:            false, //TODO:(Coderxw)
		CallNotEmpty:          false, //TODO:(Coderxw)
		Id:                    constants.PeerTypeFromChatIDType(dataChat.ChatData.ChatId).ToInt32(),
		Photo:                 nil,
		ParticipantsCount:     dataChat.Count,
		Date:                  dataChat.ChatData.Date,
		Version:               0, //TODO:(Coderxw)
		MigratedTo:            nil,
		AdminRights4DF3083493: nil,
		DefaultBannedRights:   nil,
	}).To_Chat()

	for _, vv := range dataChat.ParticipantList {
		if vv.UserId == userId {
			if vv.State != data_chat.ParticipantState_ParticipantStateNormal {
				switch vv.State {
				case data_chat.ParticipantState_ParticipantStateLeft:
					chat.Left = true
				case data_chat.ParticipantState_ParticipantStateDelete:
					chat.Kicked = true
				default:
					chat.Deactivated = true
				}
			}
			break
		}
	}

	chat.DefaultBannedRights = banned_right.RightsToChatBannedRight(dataChat.ChatData.BannedRights, dataChat.ChatData.RightsUtilDate)

	if dataChat.ChatData.Photo == nil || dataChat.ChatData.Photo.Photo == nil {
		chat.Photo = mtproto.NewTLChatPhotoEmpty(nil).To_ChatPhoto()
	} else {
		chat.Photo = mtproto.NewTLChatPhoto(&mtproto.ChatPhoto{
			PhotoSmall: nil,
			PhotoBig:   nil,
			DcId:       dc.DefaultDc,
			HasVideo:   false,
		}).To_ChatPhoto()

		chat.GetPhoto().PhotoSmall = compat.NewFileLocationByLayer(
			dataChat.ChatData.Photo.Photo.VolumeId,
			dataChat.ChatData.Photo.Photo.LocalId,
			layer,
			dc.DefaultDc)
		if dataChat.ChatData.Photo.Video != nil {
			chat.GetPhoto().PhotoBig = compat.NewFileLocationByLayer(
				dataChat.ChatData.Photo.Video.VolumeId,
				dataChat.ChatData.Photo.Video.LocalId,
				layer,
				dc.DefaultDc)
		}

		chat.GetPhoto().HasVideo = dataChat.ChatData.Photo.Video != nil
	}
	return chat
}

func ToChatParticipant(participant *data_chat.ChatParticipant, ownerId, userId int64) *mtproto.ChannelParticipant {
	if participant == nil || participant.State == data_chat.ParticipantState_ParticipantStateLeft {
		if participant == nil {
			mtproto.NewTLChannelParticipantLeft(&mtproto.ChannelParticipant{
				UserId: constants.PeerTypeFromUserIDType(userId).ToInt32(),
			}).To_ChannelParticipant()
		}
		return mtproto.NewTLChannelParticipantLeft(&mtproto.ChannelParticipant{
			UserId: constants.PeerTypeFromUserIDType(participant.UserId).ToInt32(),
		}).To_ChannelParticipant()
	} else if participant.State == data_chat.ParticipantState_ParticipantStateBan {
		return mtproto.NewTLChannelParticipantBanned(&mtproto.ChannelParticipant{
			UserId:                 constants.PeerTypeFromUserIDType(participant.UserId).ToInt32(),
			Left:                   false,
			KickedBy:               constants.PeerTypeFromUserIDType(participant.InviteId).ToInt32(),
			Date:                   participant.Date,
			BannedRights1C0FACAF93: banned_right.RightsToChatBannedRight(participant.Rights, 0),
		}).To_ChannelParticipant()
	} else if participant.State == data_chat.ParticipantState_ParticipantStateDelete {
		return mtproto.NewTLChannelParticipantKicked(&mtproto.ChannelParticipant{
			UserId:   constants.PeerTypeFromUserIDType(participant.UserId).ToInt32(),
			KickedBy: constants.PeerTypeFromUserIDType(participant.InviteId).ToInt32(),
		}).To_ChannelParticipant()
	} else {
		if participant.UserId == ownerId {
			return mtproto.NewTLChannelParticipantCreator(&mtproto.ChannelParticipant{
				UserId:                constants.PeerTypeFromUserIDType(participant.UserId).ToInt32(),
				AdminRights5DAA6E2393: banned_right.RightsToChatAdminBannedRight(participant.AdminRights),
			}).To_ChannelParticipant()
		} else if participant.Admin {
			return mtproto.NewTLChannelParticipantAdmin(&mtproto.ChannelParticipant{
				UserId:                constants.PeerTypeFromUserIDType(participant.UserId).ToInt32(),
				CanEdit:               true,
				InviterId:             constants.PeerTypeFromUserIDType(participant.InviteId).ToInt32(),
				PromotedBy:            constants.PeerTypeFromUserIDType(participant.PromotedBy).ToInt32(),
				Self:                  participant.UserId == userId,
				AdminRights5DAA6E2393: banned_right.RightsToChatAdminBannedRight(participant.AdminRights),
				Rank:                  participant.Rank,
			}).To_ChannelParticipant()
		} else {
			if participant.UserId == userId {
				return mtproto.NewTLChannelParticipantSelf(&mtproto.ChannelParticipant{
					UserId:    constants.PeerTypeFromUserIDType(participant.UserId).ToInt32(),
					InviterId: constants.PeerTypeFromUserIDType(participant.InviteId).ToInt32(),
					Date:      participant.Date,
				}).To_ChannelParticipant()
			} else {
				return mtproto.NewTLChannelParticipant(&mtproto.ChannelParticipant{
					UserId: constants.PeerTypeFromUserIDType(participant.UserId).ToInt32(),
					Date:   participant.Date,
				}).To_ChannelParticipant()
			}
		}
	}
}
