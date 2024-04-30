

https://core.telegram.org/api/config#message-animated-emoji-max

Client configuration
The MTProto API has multiple configuration parameters that can be fetched with the appropriate methods.

MTProto configuration
config#cc1a241e flags:# default_p2p_contacts:flags.3?true preload_featured_stickers:flags.4?true revoke_pm_inbox:flags.6?true blocked_mode:flags.8?true force_try_ipv6:flags.14?true date:int expires:int test_mode:Bool this_dc:int dc_options:Vector<DcOption> dc_txt_domain_name:string chat_size_max:int megagroup_size_max:int forwarded_count_max:int online_update_period_ms:int offline_blur_timeout_ms:int offline_idle_timeout_ms:int online_cloud_timeout_ms:int notify_cloud_delay_ms:int notify_default_delay_ms:int push_chat_period_ms:int push_chat_limit:int edit_time_limit:int revoke_time_limit:int revoke_pm_time_limit:int rating_e_decay:int stickers_recent_limit:int channels_read_media_period:int tmp_sessions:flags.0?int call_receive_timeout_ms:int call_ring_timeout_ms:int call_connect_timeout_ms:int call_packet_timeout_ms:int me_url_prefix:string autoupdate_url_prefix:flags.7?string gif_search_username:flags.9?string venue_search_username:flags.10?string img_search_username:flags.11?string static_maps_provider:flags.12?string caption_length_max:int message_length_max:int webfile_dc_id:int suggested_lang_code:flags.2?string lang_pack_version:flags.2?int base_lang_pack_version:flags.2?int reactions_default:flags.15?Reaction autologin_token:flags.16?string = Config;
nearestDc#8e1a1775 country:string this_dc:int nearest_dc:int = NearestDc;

---functions---

help.getConfig#c4f9186b = Config;
help.getNearestDc#1fb33026 = NearestDc;
The huge config constructor contains lots of useful information, from chat and message size limitations, to privacy settings, online status refresh interval and timeout, VoIP configuration, default inline bot usernames for GIF, image and venue lookup, and lots of other global and user-specific information, check out the constructor page for more information.

Client configuration
jsonObjectValue#c0de1bd9 key:string value:JSONValue = JSONObjectValue;

jsonNull#3f6d7b68 = JSONValue;
jsonBool#c7345e6a value:Bool = JSONValue;
jsonNumber#2be0dfa4 value:double = JSONValue;
jsonString#b71e767a value:string = JSONValue;
jsonArray#f7444763 value:Vector<JSONValue> = JSONValue;
jsonObject#99c1d49d value:Vector<JSONObjectValue> = JSONValue;

help.appConfigNotModified#7cde641d = help.AppConfig;
help.appConfig#dd18782e hash:int config:JSONValue = help.AppConfig;

---functions---

help.getAppConfig#61e3f854 hash:int = help.AppConfig;
The help.getAppConfig method returns a JSON object containing rapidly evolving, client-specific configuration parameters.
While help.getConfig returns MTProto-specific configuration with information about server-side limitations and other MTProto-related information, help.getAppConfig returns configuration parameters useful for graphical Telegram clients.

When first invoking help.getAppConfig, pass 0 to hash; in future calls, use the hash contained in the previously returned help.appConfig; if the configuration hasn't changed, a help.appConfigNotModified will be returned instead of help.appConfig.

Example value of help.appConfig.config:
````
{
"test": 1,
"keep_alive_service": true,
"background_connection": true,
"emojies_animated_zoom": 0.625,
"emojies_send_dice": [
"\ud83c\udfb2",
"\ud83c\udfaf",
"\ud83c\udfc0",
"\u26bd",
"\u26bd\ufe0f",
"\ud83c\udfb0",
"\ud83c\udfb3"
],
"emojies_send_dice_success": {
"\ud83c\udfaf": {
"value": 6,
"frame_start": 62
},
"\ud83c\udfc0": {
"value": 5,
"frame_start": 110
},
"\u26bd": {
"value": 5,
"frame_start": 110
},
"\u26bd\ufe0f": {
"value": 5,
"frame_start": 110
},
"\ud83c\udfb0": {
"value": 64,
"frame_start": 110
},
"\ud83c\udfb3": {
"value": 6,
"frame_start": 110
}
},
"emojies_sounds": {
"\ud83c\udf83": {
"id": "4956223179606458539",
"access_hash": "-2107001400913062971",
"file_reference_base64": ""
},
"\u26b0": {
"id": "4956223179606458540",
"access_hash": "-1498869544183595185",
"file_reference_base64": ""
},
"\ud83e\udddf\u200d\u2642": {
"id": "4960929110848176331",
"access_hash": "3986395821757915468",
"file_reference_base64": ""
},
"\ud83e\udddf": {
"id": "4960929110848176332",
"access_hash": "-8929417974289765626",
"file_reference_base64": ""
},
"\ud83e\udddf\u200d\u2640": {
"id": "4960929110848176333",
"access_hash": "9161696144162881753",
"file_reference_base64": ""
},
"\ud83c\udf51": {
"id": "4963180910661861548",
"access_hash": "-7431729439735063448",
"file_reference_base64": ""
},
"\ud83c\udf8a": {
"id": "5094064004578410732",
"access_hash": "8518192996098758509",
"file_reference_base64": ""
},
"\ud83c\udf84": {
"id": "5094064004578410733",
"access_hash": "-4142643820629256996",
"file_reference_base64": ""
},
"\ud83e\uddbe": {
"id": "5094064004578410734",
"access_hash": "-8934384022571962340",
"file_reference_base64": ""
}
},
"gif_search_branding": "tenor",
"gif_search_emojies": [
"\ud83d\udc4d",
"\ud83d\ude18",
"\ud83d\ude0d",
"\ud83d\ude21",
"\ud83e\udd73",
"\ud83d\ude02",
"\ud83d\ude2e",
"\ud83d\ude44",
"\ud83d\ude0e",
"\ud83d\udc4e"
],
"stickers_emoji_suggest_only_api": false,
"stickers_emoji_cache_time": 86400,
"groupcall_video_participants_max": 1000,
"qr_login_camera": true,
"qr_login_code": "primary",
"dialog_filters_enabled": true,
"dialog_filters_tooltip": true,
"autoarchive_setting_available": false,
"pending_suggestions": [
"AUTOARCHIVE_POPULAR",
"VALIDATE_PASSWORD",
"VALIDATE_PHONE_NUMBER",
"NEWCOMER_TICKS"
],
"autologin_domains": [
"instantview.telegram.org",
"translations.telegram.org",
"contest.dev",
"contest.com",
"bugs.telegram.org",
"suggestions.telegram.org",
"themes.telegram.org",
"promote.telegram.org"
],
"url_auth_domains": [],
"round_video_encoding": {
"diameter": 384,
"video_bitrate": 1000,
"audio_bitrate": 64,
"max_size": 12582912
},
"chat_read_mark_expire_period": 604800,
"chat_read_mark_size_threshold": 100,
"reactions_uniq_max": 11,
"ringtone_duration_max": 5,
"ringtone_size_max": 307200,
"ringtone_saved_count_max": 100,
"premium_purchase_blocked": false,
"channels_limit_default": 500,
"channels_limit_premium": 1000,
"saved_gifs_limit_default": 200,
"saved_gifs_limit_premium": 400,
"stickers_faved_limit_default": 5,
"stickers_faved_limit_premium": 10,
"dialog_filters_limit_default": 10,
"dialog_filters_limit_premium": 20,
"dialog_filters_chats_limit_default": 100,
"dialog_filters_chats_limit_premium": 200,
"dialogs_pinned_limit_default": 5,
"dialogs_pinned_limit_premium": 10,
"dialogs_folder_pinned_limit_default": 100,
"dialogs_folder_pinned_limit_premium": 200,
"channels_public_limit_default": 10,
"channels_public_limit_premium": 20,
"caption_length_limit_default": 1024,
"caption_length_limit_premium": 2048,
"upload_max_fileparts_default": 4000,
"upload_max_fileparts_premium": 8000,
"about_length_limit_default": 70,
"about_length_limit_premium": 140,
"stickers_premium_by_emoji_num": 0,
"stickers_normal_by_emoji_per_premium_num": 3,
"message_animated_emoji_max": 100,
"reactions_user_max_default": 1,
"reactions_user_max_premium": 3,
"reactions_in_chat_max": 100,
"default_emoji_statuses_stickerset_id": "773947703670341676",
"premium_promo_order": [
"double_limits",
"voice_to_text",
"faster_download",
"translations",
"animated_emoji",
"more_upload",
"emoji_status",
"profile_badge",
"advanced_chat_management",
"no_ads",
"app_icons",
"infinite_reactions",
"animated_userpics",
"premium_stickers"
],
"premium_bot_username": "PremiumBot",
"whitelisted_domains": [
"telegram.dog",
"telegram.me",
"telegram.org",
"t.me",
"telesco.pe",
"fragment.com"
],
"topics_pinned_limit": 5,
"telegram_antispam_user_id": "5434988373",
"telegram_antispam_group_size_min": 200,
"fragment_prefixes": [
"888"
],
"hidden_members_group_size_min": 100,
"premium_gift_attach_menu_icon": true,
"premium_gift_text_field_icon": false,
"chatlist_invites_limit_default": 3,
"chatlist_invites_limit_premium": 100,
"chatlists_joined_limit_default": 2,
"chatlists_joined_limit_premium": 20,
"chatlist_update_period": 300,
"small_queue_max_active_operations_count": 5,
"large_queue_max_active_operations_count": 2
}
````
Typical fields included in the resulting JSON object are:

emojies_animated_zoom
Animated emojis and animated dice should be scaled by this factor before being shown to the user (float)

keep_alive_service
Whether app clients should start a keepalive service to keep the app running and fetch updates even when the app is closed (boolean)

background_connection
Whether app clients should start a background TCP connection for MTProto update fetching (boolean)

emojies_send_dice
A list of supported animated dice stickers (array of strings).

emojies_send_dice_success
For animated dice emojis other than the basic 🎲, indicates the winning dice value and the final frame of the animated sticker, at which to show the fireworks 🎆 (object with emoji keys and object values, containing value and frame_start float values)

emojies_sounds
A map of soundbites to be played when the user clicks on the specified animated emoji; the file reference field should be base64-decoded before downloading the file (map of file IDs (inputDocument.id), with emoji string keys)

gif_search_branding
Specifies the name of the service providing GIF search through gif_search_username (string)

gif_search_emojies
Specifies a list of emojis that should be suggested as search term in a bar above the GIF search box (array of string emojis)

stickers_emoji_suggest_only_api
Specifies that the app should not display local sticker suggestions » for emojis at all and just use the result of messages.getStickers (bool)

stickers_emoji_cache_time
Specifies the validity period of the local cache of messages.getStickers, also relevant when generating the pagination hash when invoking the method. (integer)

qr_login_camera
Whether the Settings->Devices menu should show an option to scan a QR login code (boolean)

qr_login_code
Whether the login screen should show a QR code login option, possibly as default login method (string, "disabled", "primary" or "secondary")

dialog_filters_enabled
Whether clients should show an option for managing dialog filters AKA folders (boolean)

dialog_filters_tooltip
Whether clients should actively show a tooltip, inviting the user to configure dialog filters AKA folders; typically this happens when the chat list is long enough to start getting cluttered. (boolean)

autoarchive_setting_available
Whether clients can invoke account.setGlobalPrivacySettings with globalPrivacySettings.archive_and_mute_new_noncontact_peers = boolTrue, to automatically archive and mute new incoming chats from non-contacts. (boolean)

pending_suggestions
Contains a list of suggestions that should be actively shown as a tooltip to the user. (Array of strings, possible values shown in the suggestions section ».

topics_pinned_limit
Maximum number of topics that can be pinned in a single forum. (integer)

telegram_antispam_user_id
The ID of the official native antispam bot, that will automatically delete spam messages if enabled as specified in the native antispam documentation ».
When fetching the admin list of a supergroup using channels.getParticipants, if native antispam functionality in the specified supergroup, the bot should be manually added to the admin list displayed to the user. (numeric string that represents a Telegram user/bot ID, should be casted to an int64)

telegram_antispam_group_size_min
Minimum number of group members required to enable native antispam functionality. (integer)

fragment_prefixes
List of phone number prefixes for anonymous Fragment phone numbers. (array of strings).

hidden_members_group_size_min
Minimum number of participants required to hide the participants list of a supergroup using channels.toggleParticipantsHidden. (integer)

url_auth_domains
A list of domains that support automatic login with manual user confirmation, click here for more info on URL authorization ». (array of strings)

autologin_domains
A list of Telegram domains that support automatic login with no user confirmation, click here for more info on URL authorization ». (array of strings)

whitelisted_domains
A list of Telegram domains that can always be opened without additional user confirmation, when clicking on in-app links where the URL is not fully displayed (i.e. messageEntityTextUrl entities). (array of strings)

Note that when opening named bot web app links for the first time, confirmation should still be requested from the user, even if the domain of the containing deep link is whitelisted (i.e. t.me/<bot_username>/<short_name>?startapp=<start_parameter>, where t.me is whitelisted).

Confirmation should always be asked, even if we already opened the named bot web app before, if the link is not visible (i.e. messageEntityTextUrl text links, inline buttons etc.).

round_video_encoding
Contains a set of recommended codec parameters for round videos. (object, as described in the example)

chat_read_mark_size_threshold
Per-user read receipts, fetchable using messages.getMessageReadParticipants, will be available in groups with an amount of participants less or equal to chat_read_mark_size_threshold. (integer)

chat_read_mark_expire_period
To protect user privacy, read receipts are only stored for chat_read_mark_expire_period seconds after the message was sent. (integer)

groupcall_video_participants_max
Maximum number of participants in a group call (livestreams allow ∞ participants) (integer)

reactions_uniq_max
Maximum number of unique reactions for any given message: for example, if there are 2000 👍 and 1000 custom emoji 😁 reactions and reactions_uniq_max = 2, you can't add a 👎 reaction, because that would raise the number of unique reactions to 3 > 2. (integer)

reactions_in_chat_max
Maximum number of reactions that can be marked as allowed in a chat using chatReactionsSome. (integer)

reactions_user_max_default
Maximum number of reactions that can be added to a single message by a non-Premium user. (integer)

reactions_user_max_premium
Maximum number of reactions that can be added to a single message by a Premium user. (integer)

default_emoji_statuses_stickerset_id
Default emoji status stickerset ID. (integer)
Note that the stickerset can be fetched using inputStickerSetEmojiDefaultStatuses.

ringtone_duration_max
The maximum duration in seconds of uploadable notification sounds » (integer)

ringtone_size_max
The maximum post-conversion size in bytes of uploadable notification sounds »

ringtone_saved_count_max
The maximum number of saveable notification sounds »

message_animated_emoji_max
The maximum number of custom emojis that may be present in a message. (integer)

stickers_premium_by_emoji_num
Defines how many Premium stickers to show in the sticker suggestion popup when entering an emoji into the text field, see the sticker docs for more info (integer, defaults to 0)

stickers_normal_by_emoji_per_premium_num
For Premium users, used to define the suggested sticker list, see the sticker docs for more info (integer, defaults to 2)

premium_purchase_blocked
The user can't purchase Telegram Premium. The app must also hide all Premium features, including stars for other users, et cetera. (boolean)

channels_limit_default
The maximum number of channels and supergroups a non-Premium user may join (integer)

channels_limit_premium
The maximum number of channels and supergroups a Premium user may join (integer)

saved_gifs_limit_default
The maximum number of GIFs a non-Premium user may save (integer)

saved_gifs_limit_premium
The maximum number of GIFs a Premium user may save (integer)

stickers_faved_limit_default
The maximum number of stickers a non-Premium user may add to Favorites » (integer)

stickers_faved_limit_premium
The maximum number of stickers a Premium user may add to Favorites » (integer)

dialog_filters_limit_default
The maximum number of folders a non-Premium user may create (integer)

dialog_filters_limit_premium
The maximum number of folders a Premium user may create (integer)

dialog_filters_chats_limit_default
The maximum number of chats a non-Premium user may add to a folder (integer)

dialog_filters_chats_limit_premium
The maximum number of chats a Premium user may add to a folder (integer)

dialogs_pinned_limit_default
The maximum number of chats a non-Premium user may pin (integer)

dialogs_pinned_limit_premium
The maximum number of chats a Premium user may pin (integer)

dialogs_folder_pinned_limit_default
The maximum number of chats a non-Premium user may pin in a folder (integer)

dialogs_folder_pinned_limit_premium
The maximum number of chats a Premium user may pin in a folder (integer)

channels_public_limit_default
The maximum number of public channels or supergroups a non-Premium user may create (integer)

channels_public_limit_premium
The maximum number of public channels or supergroups a Premium user may create (integer)

caption_length_limit_default
The maximum UTF-8 length of media captions sendable by non-Premium users (integer)

caption_length_limit_premium
The maximum UTF-8 length of media captions sendable by non-Premium users (integer)

upload_max_fileparts_default
The maximum number of file parts uploadable by non-Premium users (integer, the maximum file size can be extrapolated by multiplying this value by 524288, the biggest possible chunk size)

upload_max_fileparts_premium
The maximum number of file parts uploadable by Premium users (integer, the maximum file size can be extrapolated by multiplying this value by 524288, the biggest possible chunk size)

about_length_limit_default
The maximum UTF-8 length of bios of non-Premium users (integer)

about_length_limit_premium
The maximum UTF-8 length of bios of Premium users (integer)

premium_promo_order
Array of string identifiers, indicating the order of Telegram Premium features in the Telegram Premium promotion popup, see here for the possible values »

premium_bot_username
Contains the username of the official Telegram Premium bot that may be used to buy a Telegram Premium subscription, see here for detailed instructions » (string)

premium_invoice_slug
Contains an invoice slug that may be used to buy a Telegram Premium subscription, see here for detailed instructions » (string)

premium_gift_attach_menu_icon
Whether a gift icon should be shown in the attachment menu in private chats with users, offering the current user to gift a Telegram Premium subscription to the other user in the chat. (boolean)

premium_gift_text_field_icon
Whether a gift icon should be shown in the text bar in private chats with users (ie like the / icon in chats with bots), offering the current user to gift a Telegram Premium subscription to the other user in the chat. Can only be true if premium_gift_attach_menu_icon is also true. (boolean)

chatlist_update_period
Users that import a folder using a chat folder deep link » should retrieve additions made to the folder by invoking chatlists.getChatlistUpdates at most every chatlist_update_period seconds. (integer)

chatlist_invites_limit_default
Maximum number of per-folder chat folder deep links » that can be created by non-Premium users. (integer)

chatlist_invites_limit_premium
Maximum number of per-folder chat folder deep links » that can be created by Premium users. (integer)

chatlists_joined_limit_default
Maximum number of shareable folders non-Premium users may have. (integer)

chatlists_joined_limit_premium
Maximum number of shareable folders Premium users may have. (integer)

small_queue_max_active_operations_count
A soft limit, specifying the maximum number of files that should be downloaded in parallel from the same DC, for files smaller than 20MB. (integer)

large_queue_max_active_operations_count
A soft limit, specifying the maximum number of files that should be downloaded in parallel from the same DC, for files bigger than 20MB. (integer)

Suggestions
The API can return a set of useful suggestions for users of graphical clients.

Basic suggestions
jsonObjectValue#c0de1bd9 key:string value:JSONValue = JSONObjectValue;

jsonNull#3f6d7b68 = JSONValue;
jsonBool#c7345e6a value:Bool = JSONValue;
jsonNumber#2be0dfa4 value:double = JSONValue;
jsonString#b71e767a value:string = JSONValue;
jsonArray#f7444763 value:Vector<JSONValue> = JSONValue;
jsonObject#99c1d49d value:Vector<JSONObjectValue> = JSONValue;

---functions---

help.getAppConfig#61e3f854 hash:int = help.AppConfig;
The help.getAppConfig method returns a JSON object containing rapidly evolving, client-specific configuration parameters.
A full list of these parameters can be seen in the Client Configuration section », but we're mostly interested in the pending_suggestions and autoarchive_setting_available fields of the returned JSON object:

autoarchive_setting_available - Whether clients can invoke account.setGlobalPrivacySettings with globalPrivacySettings.archive_and_mute_new_noncontact_peers = boolTrue, to automatically archive and mute new incoming chats from non-contacts. (boolean)
pending_suggestions - Contains a list of suggestions that should be actively shown as a tooltip to the user. Array of strings, possible values shown below:
"AUTOARCHIVE_POPULAR" - Users should invoke account.setGlobalPrivacySettings with globalPrivacySettings.archive_and_mute_new_noncontact_peers = boolTrue, to automatically archive and mute new incoming chats from non-contacts.
"VALIDATE_PASSWORD" - Users should make sure they still remember their 2-step verification password.
"VALIDATE_PHONE_NUMBER" - Users should check whether their authorization phone number is correct and change the phone number if it is inaccessible.
"NEWCOMER_TICKS" - Show the user a hint about the meaning of one and two ticks on sent messages.
Channel suggestions
messages.chatFull#e5d7d19c full_chat:ChatFull chats:Vector<Chat> users:Vector<User> = messages.ChatFull;

channelFull#f2355507 flags:# can_view_participants:flags.3?true can_set_username:flags.6?true can_set_stickers:flags.7?true hidden_prehistory:flags.10?true can_set_location:flags.16?true has_scheduled:flags.19?true can_view_stats:flags.20?true blocked:flags.22?true flags2:# can_delete_channel:flags2.0?true antispam:flags2.1?true participants_hidden:flags2.2?true translations_disabled:flags2.3?true id:long about:string participants_count:flags.0?int admins_count:flags.1?int kicked_count:flags.2?int banned_count:flags.2?int online_count:flags.13?int read_inbox_max_id:int read_outbox_max_id:int unread_count:int chat_photo:Photo notify_settings:PeerNotifySettings exported_invite:flags.23?ExportedChatInvite bot_info:Vector<BotInfo> migrated_from_chat_id:flags.4?long migrated_from_max_id:flags.4?int pinned_msg_id:flags.5?int stickerset:flags.8?StickerSet available_min_id:flags.9?int folder_id:flags.11?int linked_chat_id:flags.14?long location:flags.15?ChannelLocation slowmode_seconds:flags.17?int slowmode_next_send_date:flags.18?int stats_dc:flags.12?int pts:int call:flags.21?InputGroupCall ttl_period:flags.24?int pending_suggestions:flags.25?Vector<string> groupcall_default_join_as:flags.26?Peer theme_emoticon:flags.27?string requests_pending:flags.28?int recent_requesters:flags.28?Vector<long> default_send_as:flags.29?Peer available_reactions:flags.30?ChatReactions = ChatFull;

---functions---

channels.getFullChannel#8736a09 channel:InputChannel = messages.ChatFull;
Some channel/supergroup-related suggestions can also be contained in the pending_suggestions field of the channelFull constructor, returned by channels.getFullChannel.
Here's a list of possible suggestions:

CONVERT_GIGAGROUP - The supergroup has many participants: the admin should call channels.convertToGigagroup to convert it to a gigagroup.
Dismissing suggestions
boolFalse#bc799737 = Bool;
boolTrue#997275b5 = Bool;

---functions---

help.dismissSuggestion#f50dbaa1 peer:InputPeer suggestion:string = Bool;
help.dismissSuggestion can be used to dismiss a suggestion.
Pass inputPeerEmpty to peer for basic suggestions and the channel/supergroup's peer for channel suggestions.

App-specific configuration
help.appUpdate#ccbbce30 flags:# can_not_skip:flags.0?true id:int version:string text:string entities:Vector<MessageEntity> document:flags.1?Document url:flags.2?string sticker:flags.3?Document = help.AppUpdate;
help.noAppUpdate#c45a6536 = help.AppUpdate;

updates#74ae4240 updates:Vector<Update> users:Vector<User> chats:Vector<Chat> date:int seq:int = Updates;
updateServiceNotification#ebe46819 flags:# popup:flags.0?true inbox_date:flags.1?int type:string message:string media:MessageMedia entities:Vector<MessageEntity> = Update;

help.inviteText#18cb9f78 message:string = help.InviteText;

---functions---

help.getAppUpdate#522d5a7d source:string = help.AppUpdate;
help.getAppChangelog#9010ef6f prev_app_version:string = Updates;

help.getInviteText#4d392343 = help.InviteText;
help.getAppUpdate - Get info about an application update, can contain the updated application binary in the attached document
help.getAppChangelog - Get a list of service messages with app-specific changelogs
help.getInviteText - Returns a localized invitation message that can be sent via SMS to contacts that haven't signed up to Telegram yet
Terms of service
help.termsOfServiceUpdateEmpty#e3309f7f expires:int = help.TermsOfServiceUpdate;
help.termsOfServiceUpdate#28ecf961 expires:int terms_of_service:help.TermsOfService = help.TermsOfServiceUpdate;

help.termsOfService#780a0310 flags:# popup:flags.0?true id:DataJSON text:string entities:Vector<MessageEntity> min_age_confirm:flags.1?int = help.TermsOfService;

auth.authorizationSignUpRequired#44747e9a flags:# terms_of_service:flags.0?help.TermsOfService = auth.Authorization;

---functions---

help.getTermsOfServiceUpdate#2ca51fd1 = help.TermsOfServiceUpdate;
help.acceptTermsOfService#ee72f79a id:DataJSON = Bool;

auth.signIn#8d52a951 flags:# phone_number:string phone_code_hash:string phone_code:flags.0?string email_verification:flags.1?EmailVerification = auth.Authorization;

account.deleteAccount#a2c0cf74 flags:# reason:string password:flags.0?InputCheckPasswordSRP = Bool;
These methods can be used for managing consent to Telegram's Terms Of Service.

Typically, before a user signs up by invoking auth.signUp, apps should show a pop-up (if the popup flag of the help.termsOfService method is set), asking the user to accept Telegram's terms of service; in case of denial, the user is to be returned to the initial page of the login flow.

When signing up for the first time, the help.termsOfService is to be obtained from the auth.authorizationSignUpRequired constructor returned by the auth.signIn.

After signing up, or when logging in as an existing user, apps are supposed to call help.getTermsOfServiceUpdate to check for any updates to the Terms of Service; this call should be repeated after expires seconds have elapsed.
If an update to the Terms Of Service is available, clients are supposed to show a consent popup; if accepted, clients should call help.acceptTermsOfService, providing the termsOfService id JSON object; in case of denial, clients are to delete the account using account.deleteAccount, providing Decline ToS update as deletion reason.

Example implementation: android (signup), android (after login)

Telegram support info
user#8f97c628 flags:# self:flags.10?true contact:flags.11?true mutual_contact:flags.12?true deleted:flags.13?true bot:flags.14?true bot_chat_history:flags.15?true bot_nochats:flags.16?true verified:flags.17?true restricted:flags.18?true min:flags.20?true bot_inline_geo:flags.21?true support:flags.23?true scam:flags.24?true apply_min_photo:flags.25?true fake:flags.26?true bot_attach_menu:flags.27?true premium:flags.28?true attach_menu_enabled:flags.29?true flags2:# bot_can_edit:flags2.1?true id:long access_hash:flags.0?long first_name:flags.1?string last_name:flags.2?string username:flags.3?string phone:flags.4?string photo:flags.5?UserProfilePhoto status:flags.6?UserStatus bot_info_version:flags.14?int restriction_reason:flags.18?Vector<RestrictionReason> bot_inline_placeholder:flags.19?string lang_code:flags.22?string emoji_status:flags.30?EmojiStatus usernames:flags2.0?Vector<Username> = User;

help.support#17c6b5f6 phone_number:string user:User = help.Support;
help.supportName#8c05f1c9 name:string = help.SupportName;

---functions---

help.getSupport#9cdf08cd = help.Support;
help.getSupportName#d360e72c = help.SupportName;
These methods can be used for fetching info about Telegram's support user, that users can use to get support and ask questions about the app.

help.getSupport - Will return the user object that can be used for contacting support.
help.getSupportName - Will return a localized name for the support chat.
Country information and login phone patterns
help.countryCode#4203c5ef flags:# country_code:string prefixes:flags.0?Vector<string> patterns:flags.1?Vector<string> = help.CountryCode;

help.country#c3878e23 flags:# hidden:flags.0?true iso2:string default_name:string name:flags.1?string country_codes:Vector<help.CountryCode> = help.Country;

help.countriesListNotModified#93cc1f32 = help.CountriesList;
help.countriesList#87d0759e countries:Vector<help.Country> hash:int = help.CountriesList;

---functions---
help.getCountriesList#735787a8 lang_code:string hash:int = help.CountriesList;
help.getCountriesList can be used to fetch a list of localized names for all available countries and phone code patterns for logging in.

The phone code pattern should be used when showing the login screen, or when changing phone number: for example, a pattern value of XXX XXX XXX with country_code +39 indicates that the phone field for login should accept a spaced pattern like +39 123 456 789.
Also, the beginning of the national part of the phone number (123 456 789) should match one of the prefixes, if any were returned.

Additionally, the fragment_prefixes client configuration parameter contains a list of phone number prefixes for anonymous Fragment phone numbers.

