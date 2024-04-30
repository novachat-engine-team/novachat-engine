#!/bin/bash

CUR_DIR=$(cd "$(dirname "$0")";pwd)
ROOT_DIR=$CUR_DIR/../

PLAT=`uname`
MAC="Darwin"
if [[ $PLAT =~ $MAC ]]; then
  PATH=$ROOT_DIR/tools/protoc/mac/:$PATH
else
  PATH=$ROOT_DIR/tools/protoc/linux/:$PATH
fi

SRC_DIR=$ROOT_DIR/schema
MTPROTO_DST_DIR=../mtproto

GOGOPROTO_PATH=$ROOT_DIR:$ROOT_DIR/github.com/gogo/protobuf/:$ROOT_DIR/github.com/gogo/protobuf/protobuf/

echo "Start gen mtproto ..........."

protoc -I=$SRC_DIR --proto_path=$GOGOPROTO_PATH:./ \
    --gogo_out=plugins=grpc,Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,:$MTPROTO_DST_DIR \
    $SRC_DIR/*.proto

echo "End gen mtproto ..........."

echo "Start format mtproto ..........."
gofmt -w $MTPROTO_DST_DIR/*.go
echo "End format mtproto ..........."

cd ..
