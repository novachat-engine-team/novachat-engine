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

GOGOPROTO_PATH=$ROOT_DIR:$ROOT_DIR/github.com/gogo/protobuf/:$ROOT_DIR/github.com/gogo/protobuf/protobuf/:$ROOT_DIR/service/data/fs:$ROOT_DIR/service/data/restriction:$SRC_DIR

protoc -I=$SRC_DIR --proto_path=$GOGOPROTO_PATH:./ \
    --gogo_out=plugins=grpc,\
Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,\
Mphoto.proto=novachat_engine/service/data/fs,\
Mvideo.proto=novachat_engine/service/data/fs,\
Mphoto_profile.proto=novachat_engine/service/data/fs,\
Mgeopoint.proto=novachat_engine/service/data/fs,\
Mdocument.proto=novachat_engine/service/data/fs,\
Mrestriction_reason.proto=novachat_engine/service/data/restriction:$PB_DST_DIR \
$SRC_DIR/*.proto

gofmt -w $PB_DST_DIR/*.go
