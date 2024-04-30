#!/bin/bash

export GOTRACEBACK=crash
ulimit -c unlimited

LOG_PATH=./log

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

cmd=$1

if [ ! "$cmd" ]; then
	echo input empty
	exit
fi

for e in ${MOD_LIST[@]}
do
	if [ $cmd == $e ]; then
		echo "start $e"
		nohup ./$e 1>>/dev/null 2>>$LOG_PATH/$e.err &
	elif [ $cmd == 'all' ]; then
		echo "start $e"
		nohup ./$e 1>>/dev/null 2>>$LOG_PATH/$e.err &
	fi
done

