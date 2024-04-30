#!/bin/bash



MOD_LIST=(
gate
sync
auth
dbproxy
biz_server
session
msg
sfs
npns
chatgptbot
)

for e in ${MOD_LIST[@]}
do
echo "check $e"

ps -ef |grep ./$e |grep -v "grep"

done
