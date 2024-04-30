- [API](https://core.telegram.org/api)
- [Scheduled messages](https://core.telegram.org/api/scheduled-messages)

# Scheduled messages

Telegram allows scheduling messages.

```
message#58ae39c9 flags:# out:flags.1?true mentioned:flags.4?true media_unread:flags.5?true silent:flags.13?true post:flags.14?true from_scheduled:flags.18?true legacy:flags.19?true edit_hide:flags.21?true pinned:flags.24?true id:int from_id:flags.8?Peer peer_id:Peer fwd_from:flags.2?MessageFwdHeader via_bot_id:flags.11?int reply_to:flags.3?MessageReplyHeader date:int message:string media:flags.9?MessageMedia reply_markup:flags.6?ReplyMarkup entities:flags.7?Vector<MessageEntity> views:flags.10?int forwards:flags.10?int replies:flags.23?MessageReplies edit_date:flags.15?int post_author:flags.16?string grouped_id:flags.17?long restriction_reason:flags.22?Vector<RestrictionReason> = Message;

updateNewScheduledMessage#39a51dfb message:Message = Update;
updateDeleteScheduledMessages#90866cee peer:Peer messages:Vector<int> = Update;

---functions---

messages.sendMessage#520c3870 flags:# no_webpage:flags.1?true silent:flags.5?true background:flags.6?true clear_draft:flags.7?true peer:InputPeer reply_to_msg_id:flags.0?int message:string random_id:long reply_markup:flags.2?ReplyMarkup entities:flags.3?Vector<MessageEntity> schedule_date:flags.10?int = Updates;

messages.sendMedia#3491eba9 flags:# silent:flags.5?true background:flags.6?true clear_draft:flags.7?true peer:InputPeer reply_to_msg_id:flags.0?int media:InputMedia message:string random_id:long reply_markup:flags.2?ReplyMarkup entities:flags.3?Vector<MessageEntity> schedule_date:flags.10?int = Updates;
```

To schedule a message, simply provide a future unixtime in the `schedule_date` flag of [messages.sendMessage](https://core.telegram.org/method/messages.sendMessage) or [messages.sendMedia](https://core.telegram.org/method/messages.sendMedia).

The specified message or media will be added to a server-side schedule queue for the current chat, and will be automatically sent at the specified time.
The method call generates the following updates:

- Immediately, an [updateNewScheduledMessage](https://core.telegram.org/constructor/updateNewScheduledMessage), with ID equal to the ID of the message in the schedule queue for the current chat (each PM, chat, supergroup and channel has its own schedule queue and ID sequence).
- At `schedule_date`, an [updateNewMessage](https://core.telegram.org/constructor/updateNewMessage) or [updateNewChannelMessage](https://core.telegram.org/constructor/updateNewChannelMessage) with the `from_scheduled` flag set, indicating to the sender that the specified scheduled message was sent.
- At `schedule_date`, an [updateDeleteScheduledMessages](https://core.telegram.org/constructor/updateDeleteScheduledMessages), indicating that the message was flushed from the schedule queue.

If the `schedule_date` is less than 10 seconds in the future, the message will be sent immediately, generating a normal [updateNewMessage](https://core.telegram.org/constructor/updateNewMessage)/[updateNewChannelMessage](https://core.telegram.org/constructor/updateNewChannelMessage) .

### Manipulating the schedule queue

```
updateNewScheduledMessage#39a51dfb message:Message = Update;
updateDeleteScheduledMessages#90866cee peer:Peer messages:Vector<int> = Update;

---functions---

messages.getScheduledHistory#e2c2685b peer:InputPeer hash:int = messages.Messages;
messages.getScheduledMessages#bdbb0464 peer:InputPeer id:Vector<int> = messages.Messages;
messages.sendScheduledMessages#bd38850a peer:InputPeer id:Vector<int> = Updates;
messages.deleteScheduledMessages#59ae2b16 peer:InputPeer id:Vector<int> = Updates;

messages.editMessage#48f71778 flags:# no_webpage:flags.1?true peer:InputPeer id:int message:flags.11?string media:flags.14?InputMedia reply_markup:flags.2?ReplyMarkup entities:flags.3?Vector<MessageEntity> schedule_date:flags.15?int = Updates;
```

Clients can manually edit the schedule queue of a certain chat, providing the scheduled message ID obtained from [updateNewScheduledMessage](https://core.telegram.org/constructor/updateNewScheduledMessage).

- [messages.getScheduledHistory](https://core.telegram.org/method/messages.getScheduledHistory) obtains all messages in the schedule queue for the specified chat
- [messages.getScheduledMessages](https://core.telegram.org/method/messages.getScheduledMessages) obtains information about specific messages in the schedule queue for the specified chat
- [messages.sendScheduledMessages](https://core.telegram.org/method/messages.sendScheduledMessages) flushes messages from the schedule queue, sending them immediately
- [messages.deleteScheduledMessages](https://core.telegram.org/method/messages.deleteScheduledMessages) deletes messages from the schedule queue, without sending them
- [messages.editMessage](https://core.telegram.org/method/messages.editMessage) can be used to modify the scheduled date of a specific message in a schedule queue.

Modifying scheduled messages will generate an [updateNewScheduledMessage](https://core.telegram.org/constructor/updateNewScheduledMessage) with the same ID, and updated information.
Deleting scheduled messages will generate an [updateDeleteScheduledMessages](https://core.telegram.org/constructor/updateDeleteScheduledMessages).