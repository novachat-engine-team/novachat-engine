package constants

import (
	"fmt"
	"novachat_engine/mtproto"
)

// MTToChannelParticipantFilterType
//  + TL_channelParticipantsRecent
//  + TL_channelParticipantsAdmins
//  + TL_channelParticipantsKicked
//  + TL_channelParticipantsBots
//  + TL_channelParticipantsBanned
//  + TL_channelParticipantsSearch
//  + TL_channelParticipantsContacts
//  + TL_channelParticipantsMentions
func MTToChannelParticipantFilterType(filter *mtproto.ChannelParticipantsFilter) ChannelParticipantsFilterType {
	if filter == nil {
		return ChannelParticipantsFilterNone
	}

	switch filter.ClassName {
	case mtproto.ClassChannelParticipantsContacts:
		return ChannelParticipantsFilterContacts
	case mtproto.ClassChannelParticipantsRecent:
		return ChannelParticipantsFilterRecent
	case mtproto.ClassChannelParticipantsAdmins:
		return ChannelParticipantsFilterAdmin
	case mtproto.ClassChannelParticipantsKicked:
		return ChannelParticipantsFilterKicked
	case mtproto.ClassChannelParticipantsBots:
		return ChannelParticipantsFilterBots
	case mtproto.ClassChannelParticipantsBanned:
		return ChannelParticipantsFilterBanned
	case mtproto.ClassChannelParticipantsSearch:
		return ChannelParticipantsFilterSearch
	case mtproto.ClassChannelParticipantsMentions:
		return ChannelParticipantsFilterMentions
	default:
		panic(fmt.Errorf("filter:%s error", filter.ClassName))
	}
}
