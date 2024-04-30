#!/usr/bin/env bash

MONGO_PATH=""
DEST_HOST="192.168.1.80:27017"
DST_DIR="out"

MOD_LIST=(
	access_hash
	chat_updates
	chats
	contacts
	conversations
	file_info
	message_boxes
	message_chat_boxes
	user_status
	use_updates
)

if [ ! -f "$MONGO_PATH/bin/mongorestore" ]; then
  echo "$MONGO_PATH/bin/mongorestore not found"
  exit
fi

if [ ! -d "$DST_DIR" ]; then
  echo "$DST_DIR not found data dir"
  exit
fi

echo "Begin mongorestore gen ..............."
for e in ${MOD_LIST[@]}
do
  echo "mongorestore $e ..."

  $MONGO_PATH/bin/mongorestore  -h $DEST_HOST -d $e  $DST_DIR/$e/$e/
done
echo "End mongorestore gen  ..............."
