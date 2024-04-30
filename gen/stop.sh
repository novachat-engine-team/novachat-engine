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

cmd=$1
echo $cmd

for e in ${MOD_LIST[@]}
do
	if [ $cmd == $e ]; then
		echo "kill $e"
		kill -9 $(pidof $e)
		break
	elif [ $cmd == 'all' ]; then
		echo "kill $e"
		kill -9 $(pidof $e)
	fi
done

