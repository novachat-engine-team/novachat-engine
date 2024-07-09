#!/usr/bin/env bash

MONGO_PATH=""
HOST="192.168.95.132:27017"
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

if [ ! -f "$MONGO_PATH/bin/mongodump" ]; then
  echo "$MONGO_PATH/bin/mongodump not found"
  exit
fi

if [ ! -d "$DST_DIR" ]; then
  mkdir $DST_DIR
fi


echo "Begin mongodump gen ..............."
for e in ${MOD_LIST[@]}
do
  echo "mongodump $e ..."
  $MONGO_PATH/bin/mongodump  -h $HOST -d $e -o $DST_DIR/$e
done
echo "End mongodump gen  ..............."