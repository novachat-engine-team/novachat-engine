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

SRC_DIR=$CUR_DIR

PB_DST_DIR=.

if [ ! -d "$PB_DST_DIR" ]; then
    mkdir $PB_DST_DIR
fi

#GOGOPROTO_PATH=$ROOT_DIR/github.com/gogo/protobuf/:$ROOT_DIR/github.com/gogo/protobuf/protobuf/:$ROOT_DIR/schema:$ROOT_DIR/rpc/schema/session:$SRC_DIR
GOGOPROTO_PATH=$ROOT_DIR:$ROOT_DIR/github.com/gogo/protobuf/:$ROOT_DIR/github.com/gogo/protobuf/protobuf/

protoc -I=$SRC_DIR --proto_path=$GOGOPROTO_PATH:./:$ROOT_DIR \
    --gogo_out=plugins=grpc,Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,Mpkg/rpc/metadata/meta_data.proto=novachat_engine/pkg/rpc/metadata:$PB_DST_DIR \
    $SRC_DIR/*.proto

gofmt -w $PB_DST_DIR/*.go
