@echo OFF
setlocal EnableDelayedExpansion

set CUR_DIR=%~dp0../

set PATH=%CUR_DIR%/tools/protoc/win32/;%PATH%

set SRC_DIR=%CUR_DIR%/schema
set MTPROTO_DST_DIR=../mtproto

set GOGOPROTO_PATH=%CUR_DIR%/github.com/gogo/protobuf/;%CUR_DIR%/github.com/gogo/protobuf/protobuf/

echo "Start gen mtproto ..........."

protoc -I=%SRC_DIR% --proto_path=%GOGOPROTO_PATH%;./ --gogo_out=plugins=grpc,Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,:%MTPROTO_DST_DIR% %SRC_DIR%/*.proto

echo "End gen mtproto ..........."

echo "Start format mtproto ..........."
gofmt -w %MTPROTO_DST_DIR%/*.go
echo "End format mtproto ..........."

cd ..