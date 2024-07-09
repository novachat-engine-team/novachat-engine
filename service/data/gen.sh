#!/bin/bash

MOD_LIST=(
  chat
  fs
  messages/id/message_id_pts
  messages/message
  restriction
  setting
  stickerset
  update
)

BASE_DIR=`pwd`
CUR_DIR=$(cd "$(dirname "$0")";pwd)
RET=

echo "Begin data gen ..............."
for e in ${MOD_LIST[@]}
do
	cd $CUR_DIR/$e/
	echo "---- start gen $e ----"
	RET=`./gen.sh`
	if [ $? == 0 ]; then
		echo "---- end $e ----"
		cd $CUR_DIR
	else
		echo "**** $e result: $RET ****"
		cd $CUR_DIR
		exit
	fi

done
echo "End data gen ..............."
