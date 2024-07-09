#!/bin/bash

CUR_DIR=$(dirname $(realpath $0))
ROOT_DIR=$CUR_DIR/..

PLAT=`uname`
MAC="Darwin"
if [[ $PLAT =~ $MAC ]]; then
  PATH=$ROOT_DIR/tools/protoc/mac/:$PATH
else
  PATH=$ROOT_DIR/tools/protoc/linux/:$PATH
fi


cd $ROOT_DIR/schema

chmod 755 gen.sh
#./gen.sh

if [ $? -ne 0 ]; then
    echo "gen mtproto error ..........."
    exit
fi

MOD_LIST=(
    $ROOT_DIR/pkg/rpc/metadata
    $ROOT_DIR/pkg/rpc/rpc
    $ROOT_DIR/pkg/cmd/auth/rpc_client
    $ROOT_DIR/pkg/cmd/chat/rpc_client
    $ROOT_DIR/pkg/cmd/msg/rpc_client
    $ROOT_DIR/pkg/cmd/relay/rpc_client
    $ROOT_DIR/pkg/cmd/session/rpc_client
    $ROOT_DIR/pkg/cmd/sfs/rpc_client
    $ROOT_DIR/pkg/cmd/sync/rpc_client
    $ROOT_DIR/service/data
    )

for line in ${MOD_LIST[@]}
do
    echo $line
    cd $line
    chmod 755 gen.sh
    ./gen.sh
done

# SRC_DIR=$ROOT_DIR/schema
# MTPROTO_DST_DIR=../mtproto

# GOGOPROTO_PATH=$ROOT_DIR/github.com/gogo/protobuf/:$ROOT_DIR/github.com/gogo/protobuf/protobuf/

# echo "Start gen mtproto ..........."

# protoc -I=$SRC_DIR --proto_path=$GOGOPROTO_PATH:./ \
    # --gogo_out=plugins=grpc,Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,:$MTPROTO_DST_DIR \
    # $SRC_DIR/*.proto

# echo "End gen mtproto ..........."

# echo "Start format mtproto ..........."
# gofmt -w $MTPROTO_DST_DIR/*.go
# echo "End format mtproto ..........."

#
#cd $CUR_DIR/rpc
#RET=`./gen.sh`
#echo "$RET"
#
#cd $CUR_DIR/gen
#
#cd $CUR_DIR/service/statistics/
#RET=`./gen.sh`
#echo "$RET"
#cd ../
#
#echo "Start format dbproxy dataobject ..........."
#cd $CUR_DIR/server/dbproxy/server/core/dataobject/schema/
#RET=`./gen.sh`
#echo "$RET"
#echo "End format dbproxy dataobject ..........."
#cd ../../../../../
#
cd $CUR_DIR
