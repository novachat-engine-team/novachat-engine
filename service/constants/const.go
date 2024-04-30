/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/4/21 11:35
 * @File : const.go
 */

package constants

type UserIDType int64
type ChannelIDType int64
type ChatIDType int64

const (
	EmojiStickerSetName = "AnimatedEmojies"
)

const (
	RpcDefaultTimeout = 3
)

const (
	SystemFolderType   = int32(0)
	ArchivedFolderType = int32(1)
	TimeToLiveAccount  = int32(180)
)

const (
	CurrentDhConfigVersion = 3
)

type AuthKeyStatus int32

const (
	AuthKeyStatusDC   AuthKeyStatus = 0
	AuthKeyStatusTemp AuthKeyStatus = 1
)

func (m AuthKeyStatus) ToInt32() int32 {
	return int32(m)
}

const (
	ServerSaltIncorrect = 48
	MsgIdTooLow         = 16
	MsgIdTooHigh        = 17
	MsgIdMod4           = 18
	MsgIdCollision      = 19
	MsgIdTooOld         = 20
	SeqNoTooLow         = 3
	SeqNoTooHigh        = 33
	SeqNoNotEven        = 34
	SeqNoNotOdd         = 35
	InvalidContainer    = 64
)

type DevicePlatformType string

const (
	DevicePlatformTypeUnknown DevicePlatformType = ""
	DevicePlatformTypePC      DevicePlatformType = "tdesktop" //pc
	DevicePlatformTypeAndroid DevicePlatformType = "android"
	DevicePlatformTypeIOS     DevicePlatformType = "ios"
	DevicePlatformTypeWeb     DevicePlatformType = "web"
)
const DCOptionAppStoreReviews string = "review"

type PeerType int32

const (
	PeerTypeEmpty              PeerType = 0
	PeerTypeSelf               PeerType = 1
	PeerTypeChat               PeerType = 2
	PeerTypeUser               PeerType = 3
	PeerTypeChannel            PeerType = 4
	PeerTypeUserFromMessage    PeerType = 5 // unused
	PeerTypeChannelFromMessage PeerType = 6 // unused
)

type PeerNotifyType int32

const (
	PeerNotifyInputNotifyPeer       PeerNotifyType = 0 //
	PeerNotifyInputNotifyUsers      PeerNotifyType = 1 // user
	PeerNotifyInputNotifyChats      PeerNotifyType = 2 // chat
	PeerNotifyInputNotifyBroadcasts PeerNotifyType = 3 // channel
	PeerNotifyInputNotifyAll        PeerNotifyType = 4 // layer: 71 old
	PeerGlobalSetting               PeerNotifyType = 5 //
)

type InputUserType int32

const (
	InputUserTypeEmpty           InputUserType = 0
	InputUserTypeSelf            InputUserType = 1
	InputUserTypeUser            InputUserType = 2
	InputUserTypeUserFromMessage InputUserType = 3
)

const (
	GlobalNotifySettingTag int32 = 0
)

/*
1 - APNS (device token for apple push)
2 - FCM (firebase token for google firebase)
3 - MPNS (channel URI for microsoft push)
4 - Deprecated: Simple push (endpoint for firefox's simple push API)
5 - Ubuntu phone (token for ubuntu push)
6 - Blackberry (token for blackberry push)
7 - MTProto separate session
8 - WNS (windows push)
9 - APNS VoIP (token for apple push VoIP)
10 - Web push (web push, see below)
11 - MPNS VoIP (token for microsoft push VoIP)
12 - Tizen (token for tizen push)
For 10 web push, the token must be a JSON-encoded object with the following keys:

endpoint: Absolute URL exposed by the push service where the application server can send push messages
keys: P-256 elliptic curve Diffie-Hellman parameters in the following object
p256dh: Base64url-encoded P-256 elliptic curve Diffie-Hellman public key
auth: Base64url-encoded authentication secret


Notification encryption:
{
  "data": {
    "loc_key": "CHAT_MESSAGE_CONTACT",
    "loc_args": ["John Doe", "My magical group", "Contact Exchange"],
    "user_id": 14124122,
    "custom": {
      "chat_id": 241233,
      "msg_id": 123
    },
    "sound": "sound1.mp3",
  }
}
*/

type RegisterDeviceType int32

const (
	RegisterDeviceTypeDefault     RegisterDeviceType = 0
	RegisterDeviceTypeAPNS        RegisterDeviceType = 1
	RegisterDeviceTypeFCM         RegisterDeviceType = 2
	RegisterDeviceTypeMPNS        RegisterDeviceType = 3
	RegisterDeviceTypeSimplePush  RegisterDeviceType = 4
	RegisterDeviceTypeUbuntuPhone RegisterDeviceType = 5
	RegisterDeviceTypeBlackberry  RegisterDeviceType = 6
	RegisterDeviceTypeMTPROTO     RegisterDeviceType = 7
	RegisterDeviceTypeWNS         RegisterDeviceType = 8
	RegisterDeviceTypeAPNSVOIP    RegisterDeviceType = 9
	RegisterDeviceTypeWebPush     RegisterDeviceType = 10
	RegisterDeviceTypeMPNSVOIP    RegisterDeviceType = 11
	RegisterDeviceTypeTIZEN       RegisterDeviceType = 12

	RegisterDeviceTypeHuawei RegisterDeviceType = 130
	RegisterDeviceTypeXiaomi RegisterDeviceType = 131
	RegisterDeviceTypeVivo   RegisterDeviceType = 132
	RegisterDeviceTypeOppo   RegisterDeviceType = 133
	RegisterDeviceTypeMeizu  RegisterDeviceType = 134
)

type MessageBoxesType int32

const (
	MessageBoxesNone MessageBoxesType = 0
	MessageBoxesIn   MessageBoxesType = 1
	MessageBoxesOut  MessageBoxesType = 2
)

type UpdateType int32

const (
	UpdateTypeNone                              UpdateType = 0
	UpdateTypeUpdateNewMessage                  UpdateType = 1
	UpdateTypeUpdateMessageID                   UpdateType = 2
	UpdateTypeUpdateDeleteMessages              UpdateType = 3
	UpdateTypeUpdateUserTyping                  UpdateType = 4
	UpdateTypeUpdateChatUserTyping              UpdateType = 5
	UpdateTypeUpdateChatParticipants            UpdateType = 6
	UpdateTypeUpdateUserStatus                  UpdateType = 7
	UpdateTypeUpdateUserName                    UpdateType = 8
	UpdateTypeUpdateUserPhoto                   UpdateType = 9
	UpdateTypeUpdateNewEncryptedMessage         UpdateType = 10
	UpdateTypeUpdateEncryptedChatTyping         UpdateType = 11
	UpdateTypeUpdateEncryption                  UpdateType = 12
	UpdateTypeUpdateEncryptedMessagesRead       UpdateType = 13
	UpdateTypeUpdateChatParticipantAdd          UpdateType = 14
	UpdateTypeUpdateChatParticipantDelete       UpdateType = 15
	UpdateTypeUpdateDcOptions                   UpdateType = 16
	UpdateTypeUpdateNotifySettings              UpdateType = 17
	UpdateTypeUpdateServiceNotification         UpdateType = 18
	UpdateTypeUpdatePrivacy                     UpdateType = 19
	UpdateTypeUpdateUserPhone                   UpdateType = 20
	UpdateTypeUpdateReadHistoryInbox            UpdateType = 21
	UpdateTypeUpdateReadHistoryOutbox           UpdateType = 22
	UpdateTypeUpdateWebPage                     UpdateType = 23
	UpdateTypeUpdateReadMessagesContents        UpdateType = 24
	UpdateTypeUpdateChannelTooLong              UpdateType = 25
	UpdateTypeUpdateChannel                     UpdateType = 26
	UpdateTypeUpdateNewChannelMessage           UpdateType = 27
	UpdateTypeUpdateReadChannelInbox            UpdateType = 28
	UpdateTypeUpdateDeleteChannelMessages       UpdateType = 29
	UpdateTypeUpdateChannelMessageViews         UpdateType = 30
	UpdateTypeUpdateChatParticipantAdmin        UpdateType = 31
	UpdateTypeUpdateNewStickerSet               UpdateType = 32
	UpdateTypeUpdateStickerSetsOrder            UpdateType = 33
	UpdateTypeUpdateStickerSets                 UpdateType = 34
	UpdateTypeUpdateSavedGifs                   UpdateType = 35
	UpdateTypeUpdateBotInlineQuery              UpdateType = 36
	UpdateTypeUpdateBotInlineSend               UpdateType = 37
	UpdateTypeUpdateEditChannelMessage          UpdateType = 38
	UpdateTypeUpdateBotCallbackQuery            UpdateType = 39
	UpdateTypeUpdateEditMessage                 UpdateType = 40
	UpdateTypeUpdateInlineBotCallbackQuery      UpdateType = 41
	UpdateTypeUpdateReadChannelOutbox           UpdateType = 42
	UpdateTypeUpdateDraftMessage                UpdateType = 43
	UpdateTypeUpdateReadFeaturedStickers        UpdateType = 44
	UpdateTypeUpdateRecentStickers              UpdateType = 45
	UpdateTypeUpdateConfig                      UpdateType = 46
	UpdateTypeUpdatePtsChanged                  UpdateType = 47
	UpdateTypeUpdateChannelWebPage              UpdateType = 48
	UpdateTypeUpdateDialogPinned                UpdateType = 49
	UpdateTypeUpdatePinnedDialogs               UpdateType = 50
	UpdateTypeUpdateBotWebhookJSON              UpdateType = 51
	UpdateTypeUpdateBotWebhookJSONQuery         UpdateType = 52
	UpdateTypeUpdateBotShippingQuery            UpdateType = 53
	UpdateTypeUpdateBotPrecheckoutQuery         UpdateType = 54
	UpdateTypeUpdatePhoneCall                   UpdateType = 55
	UpdateTypeUpdateLangPackTooLong             UpdateType = 56
	UpdateTypeUpdateLangPack                    UpdateType = 57
	UpdateTypeUpdateFavedStickers               UpdateType = 58
	UpdateTypeUpdateChannelReadMessagesContents UpdateType = 59
	UpdateTypeUpdateContactsReset               UpdateType = 60
	UpdateTypeUpdateChannelAvailableMessages    UpdateType = 61
	UpdateTypeUpdateDialogUnreadMark            UpdateType = 62
	UpdateTypeUpdateMessagePoll                 UpdateType = 63
	UpdateTypeUpdateChatDefaultBannedRights     UpdateType = 64
	UpdateTypeUpdateFolderPeers                 UpdateType = 65
	UpdateTypeUpdatePeerSettings                UpdateType = 66
	UpdateTypeUpdatePeerLocated                 UpdateType = 67
	UpdateTypeUpdateNewScheduledMessage         UpdateType = 68
	UpdateTypeUpdateDeleteScheduledMessages     UpdateType = 69
	UpdateTypeUpdateTheme                       UpdateType = 70
	UpdateTypeUpdateGeoLiveViewed               UpdateType = 71
	UpdateTypeUpdateLoginToken                  UpdateType = 72
	UpdateTypeUpdateMessagePollVote             UpdateType = 73
	UpdateTypeUpdateDialogFilter                UpdateType = 74
	UpdateTypeUpdateDialogFilterOrder           UpdateType = 75
	UpdateTypeUpdateDialogFilters               UpdateType = 76
	UpdateTypeUpdatePhoneCallSignalingData      UpdateType = 77
	UpdateTypeUpdateChannelParticipant          UpdateType = 78
	UpdateTypeUpdateChannelMessageForwards      UpdateType = 79
	UpdateTypeUpdateReadChannelDiscussionInbox  UpdateType = 80
	UpdateTypeUpdateReadChannelDiscussionOutbox UpdateType = 81
	UpdateTypeUpdatePeerBlocked                 UpdateType = 82
	UpdateTypeUpdateChannelUserTyping           UpdateType = 83
	UpdateTypeUpdatePinnedMessages              UpdateType = 84
	UpdateTypeUpdatePinnedChannelMessages       UpdateType = 85
	UpdateTypeUpdateChat                        UpdateType = 86
	UpdateTypeUpdateGroupCallParticipants       UpdateType = 87
	UpdateTypeUpdateGroupCall                   UpdateType = 88
	UpdateTypeUpdateContactRegistered           UpdateType = 89
	UpdateTypeUpdateContactLink                 UpdateType = 90
	UpdateTypeUpdateUserBlocked                 UpdateType = 91
	UpdateTypeUpdateChatAdmins                  UpdateType = 92
	UpdateTypeUpdateChannelPinnedMessage        UpdateType = 93
	UpdateTypeUpdateRequestContacts             UpdateType = 94

	UpdateTypeUpdates                UpdateType = 100
	UpdateTypeUpdatesTooLong         UpdateType = 102
	UpdateTypeUpdateShortMessage     UpdateType = 103
	UpdateTypeUpdateShortChatMessage UpdateType = 104
	UpdateTypeUpdateShort            UpdateType = 105
	UpdateTypeUpdatesCombined        UpdateType = 106
	UpdateTypeUpdateShortSentMessage UpdateType = 107
)

type UpdateGetDifferenceType int32

const (
	UpdateGetDifferenceNone    UpdateGetDifferenceType = 0
	UpdateGetDifferenceCommon  UpdateGetDifferenceType = 1
	UpdateGetDifferenceChat    UpdateGetDifferenceType = 2
	UpdateGetDifferenceSecret  UpdateGetDifferenceType = 4
	UpdateGetDifferenceChannel UpdateGetDifferenceType = 8
)

type MessageType int32

const (
	MessageTypeEmpty          MessageType = 0
	MessageTypeMessage        MessageType = 1
	MessageTypeMessageService MessageType = 2
)

//
//type DialogType int32
//const (
//	DialogTypeUser DialogType = 0
//	DialogTypeSecret DialogType = 1
//	DialogTypeChat DialogType = 2
//	DialogTypeChannel DialogType = 3
//	DialogTypeFold DialogType = 4
//)

type AccessHashType string

const (
	AccessHashTypeUnknown AccessHashType = "unknown"
	//AccessHashTypeContact AccessHashType = "contacts",
	AccessHashTypeInputUser    AccessHashType = "inputUser"    // int32
	AccessHashTypeUser         AccessHashType = "users"        // int32
	AccessHashTypeInputChannel AccessHashType = "inputChannel" // int32
	AccessHashTypeChannel      AccessHashType = "channel"      // int32
	AccessHashTypeInputChat    AccessHashType = "inputChat"    // int32
	AccessHashTypeChat         AccessHashType = "chat"         // int32
	AccessHashTypeSecretChat   AccessHashType = "secretChat"   // int32
	AccessHashTypePeer         AccessHashType = "peer"         // int32

	AccessHashTypeWebLocation AccessHashType = "webLocation" // url:string
	AccessHashTypeWebDocument AccessHashType = "webDocument" // url:string

	AccessHashTypeGeoPoint AccessHashType = "geoPoint" // none  long lat

	AccessHashTypeInputGame               AccessHashType = "inputGame" // int64
	AccessHashTypeGame                    AccessHashType = "game"
	AccessHashTypeInputBotInlineMessageID AccessHashType = "inputBotInlineMessageID"
	AccessHashTypeInputPhoto              AccessHashType = "inputPhoto" //int64
	AccessHashTypePhoto                   AccessHashType = "photo"
	AccessHashTypeInputWallPaper          AccessHashType = "inputWallPaper" //int64
	AccessHashTypeWallPaper               AccessHashType = "wallPaper"
	AccessHashTypeGroupCall               AccessHashType = "groupCall" //int64
	AccessHashTypeInputPhoneCall          AccessHashType = "inputPhoneCall"
	AccessHashTypePhoneCall               AccessHashType = "phoneCall"       //int64
	AccessHashTypeEncryptedFile           AccessHashType = "encryptedFile"   // int64
	AccessHashTypeInputStickerSet         AccessHashType = "inputStickerSet" //int64
	AccessHashTypeStickerSet              AccessHashType = "stickerSet"
	AccessHashTypeFile                    AccessHashType = "file"
	AccessHashTypeFileLocation            AccessHashType = "fileLocation" // int64
	AccessHashTypeDocument                AccessHashType = "document"     // int64
	AccessHashTypeTheme                   AccessHashType = "theme"        // int64
)

type InputMessagesFilterType int32

const (
	InputMessagesFilterUnknown             InputMessagesFilterType = 0
	InputMessagesFilterEmpty               InputMessagesFilterType = 1
	InputMessagesFilterPhotos              InputMessagesFilterType = 2
	InputMessagesFilterVideo               InputMessagesFilterType = 3
	InputMessagesFilterPhotoVideo          InputMessagesFilterType = 4
	InputMessagesFilterDocument            InputMessagesFilterType = 5
	InputMessagesFilterUrl                 InputMessagesFilterType = 6
	InputMessagesFilterGif                 InputMessagesFilterType = 7
	InputMessagesFilterVoice               InputMessagesFilterType = 8
	InputMessagesFilterMusic               InputMessagesFilterType = 9
	InputMessagesFilterChatPhotos          InputMessagesFilterType = 10
	InputMessagesFilterPhoneCalls          InputMessagesFilterType = 11
	InputMessagesFilterRoundVoice          InputMessagesFilterType = 12
	InputMessagesFilterRoundVideo          InputMessagesFilterType = 13
	InputMessagesFilterMyMentions          InputMessagesFilterType = 14
	InputMessagesFilterGeo                 InputMessagesFilterType = 15
	InputMessagesFilterContacts            InputMessagesFilterType = 16
	InputMessagesFilterPinned              InputMessagesFilterType = 17
	InputMessagesFilterPhotoVideoDocuments InputMessagesFilterType = 18
)

type LoadHistoryType int32

const (
	LoadHistoryTypeBackward           LoadHistoryType = 0
	LoadHistoryTypeForward            LoadHistoryType = 1
	LoadHistoryTypeFirstUnRead        LoadHistoryType = 2
	LoadHistoryTypeFirstAroundMessage LoadHistoryType = 3
	LoadHistoryTypeFirstAroundDate    LoadHistoryType = 4
	LoadHistoryTypeMax                LoadHistoryType = 5
)

type PhoneCallStatus int32

const (
	PhoneCallStatusNoop    PhoneCallStatus = 0
	PhoneCallStatusRequest PhoneCallStatus = 1
	PhoneCallStatusWaiting PhoneCallStatus = 2
	PhoneCallStatusCalling PhoneCallStatus = 3
)

type ChannelParticipantsType int32

//  + TL_channelParticipantsRecent
//  + TL_channelParticipantsAdmins
//  + TL_channelParticipantsKicked
//  + TL_channelParticipantsBots
//  + TL_channelParticipantsBanned
//  + TL_channelParticipantsSearch
//  + TL_channelParticipantsContacts
//  + TL_channelParticipantsMentions
const (
	ChannelParticipantsTypeContacts ChannelParticipantsType = 0
	ChannelParticipantsTypeRecent   ChannelParticipantsType = 1
	ChannelParticipantsTypeAdmin    ChannelParticipantsType = 2
	ChannelParticipantsTypeKicked   ChannelParticipantsType = 3
	ChannelParticipantsTypeBots     ChannelParticipantsType = 4
	ChannelParticipantsTypeBanned   ChannelParticipantsType = 5
	ChannelParticipantsTypeSearch   ChannelParticipantsType = 6
	ChannelParticipantsTypeMentions ChannelParticipantsType = 7
)

///////////////////////////////////////////////////////////////////////////////
// MessageAction <--
//  + TL_messageActionEmpty
//  + TL_messageActionChatCreate
//  + TL_messageActionChatEditTitle
//  + TL_messageActionChatEditPhoto
//  + TL_messageActionChatDeletePhoto
//  + TL_messageActionChatAddUser
//  + TL_messageActionChatDeleteUser
//  + TL_messageActionChatJoinedByLink
//  + TL_messageActionChannelCreate
//  + TL_messageActionChatMigrateTo
//  + TL_messageActionChannelMigrateFrom
//  + TL_messageActionPinMessage
//  + TL_messageActionHistoryClear
//  + TL_messageActionGameScore
//  + TL_messageActionPaymentSentMe
//  + TL_messageActionPaymentSent
//  + TL_messageActionScreenshotTaken
//  + TL_messageActionCustomAction
//  + TL_messageActionBotAllowed
//  + TL_messageActionSecureValuesSentMe
//  + TL_messageActionSecureValuesSent
//  + TL_messageActionContactSignUp
//  + TL_messageActionPhoneCall
//  + TL_messageActionGeoProximityReached
//  + TL_messageActionGroupCall
//  + TL_messageActionInviteToGroupCall
//
type MessageAction int32

const (
	MessageActionEmpty               MessageAction = 0
	MessageActionChatCreate          MessageAction = 1
	MessageActionChatEditTitle       MessageAction = 2
	MessageActionChatEditPhoto       MessageAction = 3
	MessageActionChatDeletePhoto     MessageAction = 4
	MessageActionChatAddUser         MessageAction = 5
	MessageActionChatDeleteUser      MessageAction = 6
	MessageActionChatJoinedByLink    MessageAction = 7
	MessageActionChannelCreate       MessageAction = 8
	MessageActionChatMigrateTo       MessageAction = 9
	MessageActionChannelMigrateFrom  MessageAction = 10
	MessageActionPinMessage          MessageAction = 11
	MessageActionHistoryClear        MessageAction = 12
	MessageActionGameScore           MessageAction = 13
	MessageActionPaymentSentMe       MessageAction = 14
	MessageActionPaymentSent         MessageAction = 15
	MessageActionScreenshotTaken     MessageAction = 16
	MessageActionCustomAction        MessageAction = 17
	MessageActionBotAllowed          MessageAction = 18
	MessageActionSecureValuesSentMe  MessageAction = 19
	MessageActionSecureValuesSent    MessageAction = 20
	MessageActionContactSignUp       MessageAction = 21
	MessageActionPhoneCall           MessageAction = 22
	MessageActionGeoProximityReached MessageAction = 23
	MessageActionGroupCall           MessageAction = 24
	MessageActionInviteToGroupCall   MessageAction = 25
)

///////////////////////////////////////////////////////////////////////////////
// MessageMedia <--
//  + TL_messageMediaEmpty
//  + TL_messageMediaPhoto
//  + TL_messageMediaGeo
//  + TL_messageMediaContact
//  + TL_messageMediaUnsupported
//  + TL_messageMediaDocument
//  + TL_messageMediaWebPage
//  + TL_messageMediaVenue
//  + TL_messageMediaGame
//  + TL_messageMediaInvoice
//  + TL_messageMediaGeoLive
//  + TL_messageMediaPoll
//  + TL_messageMediaDice
type MessageMedia int32

const (
	MessageMediaEmpty       MessageMedia = 0
	MessageMediaPhoto       MessageMedia = 1
	MessageMediaGeo         MessageMedia = 2
	MessageMediaContact     MessageMedia = 3
	MessageMediaUnsupported MessageMedia = 4
	MessageMediaDocument    MessageMedia = 5
	MessageMediaWebPage     MessageMedia = 6
	MessageMediaVenue       MessageMedia = 7
	MessageMediaGame        MessageMedia = 8
	MessageMediaInvoice     MessageMedia = 9
	MessageMediaGeoLive     MessageMedia = 10
	MessageMediaPoll        MessageMedia = 11
	MessageMediaDice        MessageMedia = 12
)

///////////////////////////////////////////////////////////////////////////////
// InputMessage <--
//  + TL_inputMessageID
//  + TL_inputMessageReplyTo
//  + TL_inputMessagePinned
//  + TL_inputMessageCallbackQuery
type InputMessageType int32

const (
	InputMessageID              InputMessageType = 0
	InputMessageReplyTo         InputMessageType = 1
	InputMessagePinned          InputMessageType = 2
	InputMessageIDCallbackQuery InputMessageType = 3
)

func (i InputMessageType) ToInt32() int32 {
	return int32(i)
}
