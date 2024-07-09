#!/bin/bash

CUR_DIR=$(cd "$(dirname "$0")";pwd)
ROOT_DIR=$CUR_DIR/../../../../

PLAT=`uname`
MAC="Darwin"
if [[ $PLAT =~ $MAC ]]; then
  PATH=$ROOT_DIR/tools/protoc/mac/:$PATH
else
  PATH=$ROOT_DIR/tools/protoc/linux/:$PATH
fi

SRC_DIR=$CUR_DIR

PB_DST_DIR=.

if [ ! -d "$PB_DST_DIR" ]; then
    mkdir $PB_DST_DIR
fi

GOGOPROTO_PATH=$ROOT_DIR:$ROOT_DIR/github.com/gogo/protobuf/:$ROOT_DIR/github.com/gogo/protobuf/protobuf/:$ROOT_DIR/service/data/fs:$ROOT_DIR/service/data/restriction:$ROOT_DIR/service/data/chat:$ROOT_DIR/service/data/setting:$ROOT_DIR/schema:$ROOT_DIR/pkg/cmd/msg/rpc_client:$ROOT_DIR/service/data/fs:$SRC_DIR

protoc -I=$SRC_DIR --proto_path=$GOGOPROTO_PATH:./ \
    --gogo_out=plugins=grpc,\
Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,\
Mphoto.proto=novachat_engine/service/data/fs,\
Mgeopoint.proto=novachat_engine/service/data/fs,\
Mdocument.proto=novachat_engine/service/data/fs,\
Mservice/data/chat/chat.proto=novachat_engine/service/data/chat,\
Mphoto_profile.proto=novachat_engine/service/data/fs,\
Mservice/data/chat/chat_participant.proto=novachat_engine/service/data/chat,\
Mschema.msg.proto=novachat_engine/pkg/cmd/msg/rpc_client,\
Mschema.tl.sync.proto=novachat_engine/mtproto,\
Mservice/data/setting/notify_setting.proto=novachat_engine/service/data/setting,\
Mrestriction_reason.proto=novachat_engine/service/data/restriction:$PB_DST_DIR \
$SRC_DIR/*.proto

gofmt -w $PB_DST_DIR/*.go
