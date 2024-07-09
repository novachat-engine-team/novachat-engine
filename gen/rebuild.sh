#!/bin/bash


GEN=./gen.sh
BUILD=./build.sh

echo start gen
RET=`${GEN}`
if [ $? == 0 ]; then
	echo end gen
else
	exit
fi

echo start build
RET=`${BUILD}`
if [ $? == 0 ]; then
	echo end build
else
	exit
fi
