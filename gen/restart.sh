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
		echo "kill $e"
		kill -9 $(pidof $e)
		sleep 1
		echo "start $e"
		nohup ./$e $e.yaml 1>>/dev/null 2>>$LOG_PATH/$e.err &
		break
	elif [ $cmd == 'all' ]; then
		echo "kill $e"
		kill -9 $(pidof $e)
		sleep 1
		echo "start $e"
		nohup ./$e $e.yaml 1>>/dev/null 2>>$LOG_PATH/$e.err &

	fi

done


