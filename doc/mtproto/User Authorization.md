# [User Authorization](https://core.telegram.org/api/auth)

Authorization is associated with a client’s encryption key identifier: **auth_key_id**. No additional parameters need to be passed into methods following authorization.

To log in as a [bot](https://core.telegram.org/bots), follow [these instructions »](https://core.telegram.org/api/bots).

### Sending a verification code

Example implementations: [telegram for android](https://github.com/DrKLO/Telegram/blob/master/TMessagesProj/src/main/java/org/telegram/ui/LoginActivity.java), [tdlib](https://github.com/tdlib/td/tree/master/td/telegram/SendCodeHelper.cpp).

To show a nicely formatted and validated phone number field, the [help.countriesList](https://core.telegram.org/constructor/help.countriesList) constructor can be obtained using the [help.getCountriesList](https://core.telegram.org/method/help.getCountriesList) method.
The [help.countriesList](https://core.telegram.org/constructor/help.countriesList) config is then used as described [here »](https://core.telegram.org/api/config#country-information-and-login-phone-patterns).

Authorization requires that a text message containing an authorization code first be sent to the user’s phone.
This may be done using the [auth.sendCode](https://core.telegram.org/method/auth.sendCode) method. The system will automatically choose how to send the authorization code; there are four possible ways the code can arrive:

- [Telegram code](https://core.telegram.org/constructor/auth.sentCodeTypeApp)
- [SMS code](https://core.telegram.org/constructor/auth.sentCodeTypeSms)
- [Phone call](https://core.telegram.org/constructor/auth.sentCodeTypeCall): a synthesized voice will tell the user which verification code to input
- [Flash phone call](https://core.telegram.org/constructor/auth.sentCodeTypeFlashCall): the code will be sent via a flash phone call, that will be closed immediately. In the last case, the phone code will then be the phone number itself, just make sure that the phone number matches the specified pattern (see [auth.sentCodeTypeFlashCall](https://core.telegram.org/constructor/auth.sentCodeTypeFlashCall)).

The [auth.sendCode](https://core.telegram.org/method/auth.sendCode) method also has parameters for enabling/disabling use of flash calls, and allows passing an SMS token that will be included in the sent SMS. For example, the latter is required in newer versions of android, to use the [android SMS receiver APIs](https://developers.google.com/identity/sms-retriever/overview).

The returned [auth.SentCode](https://core.telegram.org/type/auth.SentCode) object will contain multiple parameters:

| **flags**           | [#](https://core.telegram.org/type/%23)                      | Flags, see [TL conditional fields](https://core.telegram.org/mtproto/TL-combinators#conditional-fields) |
| ------------------- | ------------------------------------------------------------ | ------------------------------------------------------------ |
| **type**            | [auth.SentCodeType](https://core.telegram.org/type/auth.SentCodeType) | Phone code type                                              |
| **phone_code_hash** | [string](https://core.telegram.org/type/string)              | Phone code hash, to be stored and later re-used with [auth.signIn](https://core.telegram.org/method/auth.signIn) |
| **next_type**       | [flags](https://core.telegram.org/mtproto/TL-combinators#conditional-fields).1?[auth.CodeType](https://core.telegram.org/type/auth.CodeType) | Phone code type that will be sent next, if the phone code is not received within `timeout` seconds: to send it use [auth.resendCode](https://core.telegram.org/method/auth.resendCode) |
| **timeout**         | [flags](https://core.telegram.org/mtproto/TL-combinators#conditional-fields).2?[int](https://core.telegram.org/type/int) | Timeout for reception of the phone code                      |

If the message takes too long (`timeout` seconds) to arrive at the phone, the [auth.resendCode](https://core.telegram.org/method/auth.resendCode) method may be invoked to resend a code of type `next_type`. If the same happens again, you can use [auth.resendCode](https://core.telegram.org/method/auth.resendCode) with the `next_type` returned by the previous call to [auth.resendCode](https://core.telegram.org/method/auth.resendCode). To cancel the verification code use [auth.cancelCode](https://core.telegram.org/method/auth.cancelCode).

### Sign in/sign up

When user enters verification code, the [auth.signIn](https://core.telegram.org/method/auth.signIn) method must be used to validate it and possibly sign user in.

If the code was entered correctly, but the method returns [auth.authorizationSignUpRequired](https://core.telegram.org/constructor/auth.authorizationSignUpRequired), it means that account with this phone number doesn't exist yet: user needs to provide basic information, accept terms of service and then the new user registration method ([auth.signUp](https://core.telegram.org/method/auth.signUp)) must be invoked.

### 2FA

When trying to sign in using [auth.signIn](https://core.telegram.org/method/auth.signIn), an [error 400 SESSION_PASSWORD_NEEDED](https://core.telegram.org/method/auth.signIn#possible-errors) may be returned, if the user has two-factor authentication enabled. In this case, instructions for [SRP 2FA authentication](https://core.telegram.org/api/srp) must be followed.

To set up two-factor authorization on an already authorized account, follow the [SRP 2FA authentication docs](https://core.telegram.org/api/srp).

#### Test Phone Numbers

Each phone number is limited to only a certain amount of logins per day (e.g. 5, but this is subject to change) after which the API will return a FLOOD error until the next day. This might not be enough for testing the implementation of User Authorization flows in client applications.

There are several reserved phone number prefixes for testing that your application handles redirects between DCs, sign up, sign in and 2FA flows correctly. These numbers are only available on **Test DCs** (their IP addresses for TCP transport are availble in [API development tools](https://my.telegram.org/apps) panel after [api_id was obtained](https://core.telegram.org/api/obtaining_api_id#obtaining-api-id), [URI format](https://core.telegram.org/mtproto/transports#uri-format) for HTTPS/Websocket transport).

If you wish to emulate an application of a user associated with DC number X, it is sufficient to specify the phone number as `99966XYYYY`, where YYYY are random numbers, when registering the user. A user like this would always get XXXXX as the login confirmation code (the DC number, repeated five times). Note that the value of X must be in the range of 1-3 because there are only 3 Test DCs. When the flood limit is reached for any particular test number, just choose another number (changing the YYYY random part).

Do not store any important or private information in the messages of such test accounts; anyone can make use of the simplified authorization mechanism – and we periodically wipe all information stored there.

Proceed with User Authorization flows in **Production DCs** only after you make sure everything works correctly on **Test DCs** first to avoid reaching flood limits.

> To help you with working on production DCs, logins with the same phone number with which the `api_id` was registered have more generous flood limits.

### We are authorized

As a result of authorization, the client key, **auth_key_id**, becomes associated with the user, and each subsequent API call with this key will be executed with that user’s identity. The authorization method itself returns the relevant user. It is best to immediately store the User ID locally in a binding with the key.

Only a small portion of the API methods are available to **unauthorized** users:

- [auth.sendCode](https://core.telegram.org/method/auth.sendCode)
- [auth.resendCode](https://core.telegram.org/method/auth.resendCode)
- [account.getPassword](https://core.telegram.org/method/account.getPassword)
- [auth.checkPassword](https://core.telegram.org/method/auth.checkPassword)
- [auth.checkPhone](https://core.telegram.org/method/auth.checkPhone)
- [auth.signUp](https://core.telegram.org/method/auth.signUp)
- [auth.signIn](https://core.telegram.org/method/auth.signIn)
- [auth.importAuthorization](https://core.telegram.org/method/auth.importAuthorization)
- [help.getConfig](https://core.telegram.org/method/help.getConfig)
- [help.getNearestDc](https://core.telegram.org/method/help.getNearestDc)
- [help.getAppUpdate](https://core.telegram.org/method/help.getAppUpdate)
- [help.getCdnConfig](https://core.telegram.org/method/help.getCdnConfig)
- [langpack.getLangPack](https://core.telegram.org/method/langpack.getLangPack)
- [langpack.getStrings](https://core.telegram.org/method/langpack.getStrings)
- [langpack.getDifference](https://core.telegram.org/method/langpack.getDifference)
- [langpack.getLanguages](https://core.telegram.org/method/langpack.getLanguages)
- [langpack.getLanguage](https://core.telegram.org/method/langpack.getLanguage)

Other methods will result in an error: [**401 UNAUTHORIZED**](https://core.telegram.org/api/errors#401-unauthorized).