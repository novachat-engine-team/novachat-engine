

# [Channels](https://core.telegram.org/api/channel)

### Channels, chats & supergroups

[Channels](https://telegram.org/tour/channels) are a tool for broadcasting your messages to large audiences. They can have an unlimited number of subscribers, they can be public with a permanent URL and each post in a channel has its own view counter. Technically, they are represented by [channel](https://core.telegram.org/constructor/channel) constructors.

[Supergroups](https://telegram.org/tour/groups) are a powerful tool for building communities and can support up to 200,000 members each. Technically, supergroups are actually channels: they are represented by [channel](https://core.telegram.org/constructor/channel) constructors, with the `megagroup` flag set to true.

Channels and supergroup can be created using the [channels.createChannel](https://core.telegram.org/method/channels.createChannel) method, by setting the appropriate `broadcast` or `megagroup` flags. Supergroups can also be assigned a `geo_point` to become [geochats](https://telegram.org/blog/contacts-local-groups).

In previous versions of telegram, only normal groups (represented by [chat](https://core.telegram.org/constructor/chat) constructors) could be created using [messages.createChat](https://core.telegram.org/method/messages.createChat): these groups have fewer features, and can only have 200 members at max.

### Migration

To upgrade a legacy group to a supergroup, [messages.migrateChat](https://core.telegram.org/method/messages.migrateChat) can be used. The `chats` field of the result will have two objects:

- A [chat](https://core.telegram.org/constructor/chat) constructor with a `migrated_to` field, indicating the address of the new supergroup
- The new [channel](https://core.telegram.org/constructor/channel) megagroup constructor

When [getting full info](https://core.telegram.org/method/channels.getFullChannel) about the migrated channel, the [channelFull](https://core.telegram.org/constructor/channelFull) object will have `migrated_from_chat_id` and `migrated_from_max_id` fields indicating the original ID of the chat, and the message ID in the original chat at which the group was migrated.

All users of the chat will receive an [updateNewMessage](https://core.telegram.org/constructor/updateNewMessage) from the old chat with a [messageService](https://core.telegram.org/constructor/messageService) containing a [messageActionChatMigrateTo](https://core.telegram.org/constructor/messageActionChatMigrateTo) constructor.

All new messages have to be sent to the new supergroup.

When working with migrated groups clients need to handle loading of the message history (as well as search results et cetera) from both the legacy group and the new supergroup. This is done by merging the two messages lists (requested with different [Peer](https://core.telegram.org/type/Peer) values) client side.

### Rights

Channels, legacy groups and supergroups allow setting [granular permissions](https://telegram.org/blog/permissions-groups-undo) both for admins and specific users; channels, supergroups and legacy groups also allow setting global granular permissions for users.

For more info on how to set and modify rights, see [here »](https://core.telegram.org/api/rights).

### Pinned messages

Telegram allows pinning multiple messages on top in a chat, group, supergroup or channel.

See [here »](https://core.telegram.org/api/pin) for more info on pinning and unpinning messages.

### Discussion

Groups can be associated to a channel as a [discussion group](https://telegram.org/blog/privacy-discussions-web-bots), to allow users to discuss about posts.

For more info on how to set a discussion group in channel, see [here »](https://core.telegram.org/api/discussion)

### Recent actions

Both supergroups and channels offer a so-called [admin log](https://telegram.org/blog/admin-revolution), a log of recent relevant supergroup and channel actions, like the modification of group/channel settings or information on behalf of an admin, user kicks and bans, and more.

See [here »](https://core.telegram.org/api/recent-actions) for more info.