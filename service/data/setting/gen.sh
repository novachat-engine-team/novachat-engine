#!/bin/bash

CUR_DIR=$(cd "$(dirname "$0")";pwd)
ROOT_DIR=$CUR_DIR/../../../

PLAT=`uname`
MAC="Darwin"
if [[ $PLAT =~ $MAC ]]; then
  PATH=$ROOT_DIR/tools/protoc/mac/:$PATH
else
  PATH=$ROOT_DIR/tools/protoc/linux/:$PATH
fi

SRC_DIR=$CUR_DIR/
PB_DST_DIR=.

GOGOPROTO_PATH=$ROOT_DIR:$ROOT_DIR/github.com/gogo/protobuf/:$ROOT_DIR/github.com/gogo/protobuf/protobuf/:$ROOT_DIR/schema:$SRC_DIR:$ROOT_DIR/service/data/messages/message:$ROOT_DIR/service/data/fs

protoc -I=$SRC_DIR --proto_path=$GOGOPROTO_PATH:./ \
    --gogo_out=plugins=grpc,Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,Mschema.tl.core_types.proto=novachat_engine/mtproto,Mschema.tl.sync_service.proto=novachat_engine/mtproto,Mschema.tl.sync.proto=novachat_engine/mtproto,Mschema.tl.sync.tl.proto=novachat_engine/mtproto,Mschema.tl.sync.proto=novachat_engine/mtproto,Mschema.tl.transport.proto=novachat_engine/mtproto,Mschema.data.proto=novachat_engine/mtproto,Mmessage.proto=novachat_engine/service/data/messages/message,Mmedia.proto=novachat_engine/service/data/messages/message:$PB_DST_DIR \
    $SRC_DIR/*.proto

gofmt -w $PB_DST_DIR/*.go
