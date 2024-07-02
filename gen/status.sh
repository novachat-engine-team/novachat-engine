#!/bin/bash



MOD_LIST=(
auth
biz_server
gate
chat
session
msg
sync
sfs
relay
)

for e in ${MOD_LIST[@]}
do
echo "check $e"

ps -ef |grep ./$e |grep -v "grep"

done
