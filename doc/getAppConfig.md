- [API](https://core.telegram.org/api)
- 

- [Client configuration](https://core.telegram.org/api/config)

# Client configuration

The MTProto API has multiple configuration parameters that can be fetched with the appropriate methods.

## MTProto configuration

```
config#330b4067 flags:# phonecalls_enabled:flags.1?true default_p2p_contacts:flags.3?true preload_featured_stickers:flags.4?true ignore_phone_entities:flags.5?true revoke_pm_inbox:flags.6?true blocked_mode:flags.8?true pfs_enabled:flags.13?true date:int expires:int test_mode:Bool this_dc:int dc_options:Vector<DcOption> dc_txt_domain_name:string chat_size_max:int megagroup_size_max:int forwarded_count_max:int online_update_period_ms:int offline_blur_timeout_ms:int offline_idle_timeout_ms:int online_cloud_timeout_ms:int notify_cloud_delay_ms:int notify_default_delay_ms:int push_chat_period_ms:int push_chat_limit:int saved_gifs_limit:int edit_time_limit:int revoke_time_limit:int revoke_pm_time_limit:int rating_e_decay:int stickers_recent_limit:int stickers_faved_limit:int channels_read_media_period:int tmp_sessions:flags.0?int pinned_dialogs_count_max:int pinned_infolder_count_max:int call_receive_timeout_ms:int call_ring_timeout_ms:int call_connect_timeout_ms:int call_packet_timeout_ms:int me_url_prefix:string autoupdate_url_prefix:flags.7?string gif_search_username:flags.9?string venue_search_username:flags.10?string img_search_username:flags.11?string static_maps_provider:flags.12?string caption_length_max:int message_length_max:int webfile_dc_id:int suggested_lang_code:flags.2?string lang_pack_version:flags.2?int base_lang_pack_version:flags.2?int = Config;

nearestDc#8e1a1775 country:string this_dc:int nearest_dc:int = NearestDc;

---functions---

help.getConfig#c4f9186b = Config;
help.getNearestDc#1fb33026 = NearestDc;
```

The huge [config](https://core.telegram.org/constructor/config) constructor contains lots of useful information, from chat and message size limitations, to privacy settings, online status refresh interval and timeout, VoIP configuration, default inline bot usernames for GIF, image and venue lookup, and lots of other global and user-specific information, check out the [constructor page](https://core.telegram.org/constructor/config) for more information.

## Client configuration

```
jsonObjectValue#c0de1bd9 key:string value:JSONValue = JSONObjectValue;

jsonNull#3f6d7b68 = JSONValue;
jsonBool#c7345e6a value:Bool = JSONValue;
jsonNumber#2be0dfa4 value:double = JSONValue;
jsonString#b71e767a value:string = JSONValue;
jsonArray#f7444763 value:Vector<JSONValue> = JSONValue;
jsonObject#99c1d49d value:Vector<JSONObjectValue> = JSONValue;

---functions---

help.getAppConfig#98914110 = JSONValue;
```

The [help.getAppConfig](https://core.telegram.org/method/help.getAppConfig) method returns a JSON object containing rapidly evolving, client-specific configuration parameters.
While [help.getConfig](https://core.telegram.org/method/help.getConfig) returns MTProto-specific configuration with information about server-side limitations and other MTProto-related information, [help.getAppConfig](https://core.telegram.org/method/help.getAppConfig) returns configuration parameters useful for graphical Telegram clients.

Typical fields included in the resulting JSON object are:

- `emojies_animated_zoom` - [Animated emojis](https://core.telegram.org/api/animated-emojis) and [animated dice](https://core.telegram.org/api/dice) should be scaled by this factor before being shown to the user (float)
- `keep_alive_service` - Whether app clients should start a keepalive service to keep the app running and fetch updates even when the app is closed (boolean)
- `background_connection` - Whether app clients should start a background TCP connection for MTProto update fetching (boolean)
- `emojies_send_dice` - A list of supported [animated dice](https://core.telegram.org/api/dice) stickers (array of strings).
- `emojies_send_dice_success` - For [animated dice](https://core.telegram.org/api/dice) emojis other than the basic ![🎲](https://telegram.org/img/emoji/40/F09F8EB2.png), indicates the winning dice value and the final frame of the animated sticker, at which to show the fireworks ![🎆](https://telegram.org/img/emoji/40/F09F8E86.png) (object with emoji keys and object values, containing `value` and `frame_start` float values)
- `emojies_sounds` - A map of soundbites to be played when the user clicks on the specified [animated emoji](https://core.telegram.org/api/animated-emojis); the [file reference field](https://core.telegram.org/api/file_reference) should be base64-decoded before [downloading the file](https://core.telegram.org/api/files) (map of [file IDs](https://core.telegram.org/api/files), with emoji string keys)
- `gif_search_branding` - Specifies the name of the service providing GIF search through [gif_search_username](https://core.telegram.org/api/config#mtproto-configuration) (string)
- `gif_search_emojies` - Specifies a list of emojies that should be suggested as search term in a bar above the GIF search box (array of string emojis)
- `qr_login_camera` - Whether the Settings->Devices menu should show an option to scan a [QR login code](https://core.telegram.org/api/qr-login) (boolean)
- `qr_login_code` - Whether the login screen should show a [QR code login option](https://core.telegram.org/api/qr-login), possibly as default login method (string, "disabled", "primary" or "secondary")
- `dialog_filters_enabled` - Whether clients should show an option for managing [dialog filters AKA folders](https://core.telegram.org/api/folders) (boolean)
- `dialog_filters_tooltip` - Whether clients should actively show a tooltip, inviting the user to configure [dialog filters AKA folders](https://core.telegram.org/api/folders); typically this happens when the chat list is long enough to start getting cluttered. (boolean)

Example value:

```
{
    "test": 1,
    "emojies_animated_zoom": 0.625,
    "keep_alive_service": true,
    "background_connection": true,
    "emojies_send_dice": [
        "\ud83c\udfb2",
        "\ud83c\udfaf",
        "\ud83c\udfc0",
        "\u26bd",
        "\u26bd\ufe0f",
        "\ud83c\udfb0"
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
        }
    },
    "emojies_sounds": {
        "\ud83c\udf83": {
            "id": "4956223179606458539",
            "access_hash": "-2107001400913062971",
            "file_reference_base64": "AF-4ApC7ukC0UWEPZN0TeSJURe7T"
        },
        "\u26b0": {
            "id": "4956223179606458540",
            "access_hash": "-1498869544183595185",
            "file_reference_base64": "AF-4ApCLKMGt96WCvLm58kbqZHd3"
        },
        "\ud83e\udddf\u200d\u2642": {
            "id": "4960929110848176331",
            "access_hash": "3986395821757915468",
            "file_reference_base64": "AF-4ApAedNln3IMEHH-SUQuH8L9g"
        },
        "\ud83e\udddf": {
            "id": "4960929110848176332",
            "access_hash": "-8929417974289765626",
            "file_reference_base64": "AF-4ApArGURtGa2KVC-Yovh1kQoW"
        },
        "\ud83e\udddf\u200d\u2640": {
            "id": "4960929110848176333",
            "access_hash": "9161696144162881753",
            "file_reference_base64": "AF-4ApD-eOqXvTBmcszAEkzQN615"
        },
        "\ud83c\udf51": {
            "id": "4963180910661861548",
            "access_hash": "-7431729439735063448",
            "file_reference_base64": "AF-4ApBimvRxhcXg-iQ5Gw4Eelit"
        },
        "\u2764": {
            "id": "4978826754966683841",
            "access_hash": "3926211553285686901",
            "file_reference_base64": "AF-4ApDBkyjgN2Tk9zJvXPhfJXPA"
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
    "qr_login_camera": true,
    "qr_login_code": "secondary",
    "dialog_filters_enabled": true,
    "dialog_filters_tooltip": false
}
```

## App-specific configuration

```
help.appUpdate#1da7158f flags:# can_not_skip:flags.0?true id:int version:string text:string entities:Vector<MessageEntity> document:flags.1?Document url:flags.2?string = help.AppUpdate;
help.noAppUpdate#c45a6536 = help.AppUpdate;

updates#74ae4240 updates:Vector<Update> users:Vector<User> chats:Vector<Chat> date:int seq:int = Updates;
updateServiceNotification#ebe46819 flags:# popup:flags.0?true inbox_date:flags.1?int type:string message:string media:MessageMedia entities:Vector<MessageEntity> = Update;

help.inviteText#18cb9f78 message:string = help.InviteText;

---functions---

help.getAppUpdate#522d5a7d source:string = help.AppUpdate;
help.getAppChangelog#9010ef6f prev_app_version:string = Updates;

help.getInviteText#4d392343 = help.InviteText;
```

- [help.getAppUpdate](https://core.telegram.org/method/help.getAppUpdate) - Get info about an application update, can contain the updated application binary in the attached document
- [help.getAppChangelog](https://core.telegram.org/method/help.getAppChangelog) - Get a list of service messages with app-specific changelogs
- [help.getInviteText](https://core.telegram.org/method/help.getInviteText) - Returns a localized invitation message that can be sent via SMS to contacts that haven't signed up to Telegram yet

## Terms of service

```
help.termsOfServiceUpdateEmpty#e3309f7f expires:int = help.TermsOfServiceUpdate;
help.termsOfServiceUpdate#28ecf961 expires:int terms_of_service:help.TermsOfService = help.TermsOfServiceUpdate;

help.termsOfService#780a0310 flags:# popup:flags.0?true id:DataJSON text:string entities:Vector<MessageEntity> min_age_confirm:flags.1?int = help.TermsOfService;

auth.authorizationSignUpRequired#44747e9a flags:# terms_of_service:flags.0?help.TermsOfService = auth.Authorization;

---functions---

help.getTermsOfServiceUpdate#2ca51fd1 = help.TermsOfServiceUpdate;
help.acceptTermsOfService#ee72f79a id:DataJSON = Bool;

auth.signIn#bcd51581 phone_number:string phone_code_hash:string phone_code:string = auth.Authorization;

account.deleteAccount#418d4e0b reason:string = Bool;
```

These methods can be used for managing consent to Telegram's [Terms Of Service](https://telegram.org/tos).

Typically, before a user [signs up](https://core.telegram.org/api/auth#sign-insign-up) by invoking [auth.signUp](https://core.telegram.org/method/auth.signUp), apps should show a pop-up (if the `popup` flag of the [help.termsOfService](https://core.telegram.org/constructor/help.termsOfService) method is set), asking the user to accept Telegram's terms of service; in case of denial, the user is to be returned to the initial page of the login flow.

When signing up for the first time, the [help.termsOfService](https://core.telegram.org/constructor/help.termsOfService) is to be obtained from the [auth.authorizationSignUpRequired](https://core.telegram.org/constructor/auth.authorizationSignUpRequired) constructor returned by the [auth.signIn](https://core.telegram.org/method/auth.signIn).

After signing up, or when logging in as an existing user, apps are supposed to call [help.getTermsOfServiceUpdate](https://core.telegram.org/method/help.getTermsOfServiceUpdate) to check for any updates to the Terms of Service; this call should be repeated after `expires` seconds have elapsed.
If an update to the Terms Of Service is available, clients are supposed to show a consent popup; if accepted, clients should call [help.acceptTermsOfService](https://core.telegram.org/method/help.acceptTermsOfService), providing the [termsOfService `id` JSON object](https://core.telegram.org/constructor/help.termsOfService); in case of denial, clients are to delete the account using [account.deleteAccount](https://core.telegram.org/method/account.deleteAccount), providing `Decline ToS update` as deletion `reason`.

Example implementation: [android (signup)](https://github.com/DrKLO/Telegram/blob/dbf81a34affcd1c24d78e1403347ea8b3a186154/TMessagesProj/src/main/java/org/telegram/ui/LoginActivity.java#L3510), [android (after login)](https://github.com/DrKLO/Telegram/blob/ed9e38da5b3b6ca80a7cb719a000d310d07497b0/TMessagesProj/src/main/java/org/telegram/ui/Components/TermsOfServiceView.java)

## Telegram support info

```
user#938458c1 flags:# self:flags.10?true contact:flags.11?true mutual_contact:flags.12?true deleted:flags.13?true bot:flags.14?true bot_chat_history:flags.15?true bot_nochats:flags.16?true verified:flags.17?true restricted:flags.18?true min:flags.20?true bot_inline_geo:flags.21?true support:flags.23?true scam:flags.24?true apply_min_photo:flags.25?true id:int access_hash:flags.0?long first_name:flags.1?string last_name:flags.2?string username:flags.3?string phone:flags.4?string photo:flags.5?UserProfilePhoto status:flags.6?UserStatus bot_info_version:flags.14?int restriction_reason:flags.18?Vector<RestrictionReason> bot_inline_placeholder:flags.19?string lang_code:flags.22?string = User;

help.support#17c6b5f6 phone_number:string user:User = help.Support;
help.supportName#8c05f1c9 name:string = help.SupportName;

---functions---

help.getSupport#9cdf08cd = help.Support;
help.getSupportName#d360e72c = help.SupportName;
```

These methods can be used for fetching info about Telegram's support user, that users can use to get support and ask questions about the app.

- [help.getSupport](https://core.telegram.org/method/help.getSupport) - Will return the [user](https://core.telegram.org/constructor/user) object that can be used for contacting support.
- [help.getSupportName](https://core.telegram.org/method/help.getSupportName) - Will return a localized name for the support chat.

## Country information and login phone patterns

```
help.countryCode#4203c5ef flags:# country_code:string prefixes:flags.0?Vector<string> patterns:flags.1?Vector<string> = help.CountryCode;

help.country#c3878e23 flags:# hidden:flags.0?true iso2:string default_name:string name:flags.1?string country_codes:Vector<help.CountryCode> = help.Country;

help.countriesListNotModified#93cc1f32 = help.CountriesList;
help.countriesList#87d0759e countries:Vector<help.Country> hash:int = help.CountriesList;

---functions---
help.getCountriesList#735787a8 lang_code:string hash:int = help.CountriesList;
```

[help.getCountriesList](https://core.telegram.org/method/help.getCountriesList) can be used to fetch a list of localized names for all available countries and phone code patterns for logging in.

The phone code pattern should be used when showing the [login](https://core.telegram.org/api/auth) screen, or when changing phone number: for example, a pattern value of `XXX XXX XXX` with `country_code` `+39` indicates that the phone field for login should accept a spaced pattern like `+39 123 456 789`.
Also, the beginning of the national part of the phone number (`123 456 789`) should with match one of the `prefixes`, if any were returned.