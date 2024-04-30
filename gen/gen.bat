@echo off

set CUR_DIR=%~dp0..

set PATH=%CUR_DIR%/tools/protoc/win32/;%PATH%

set SRC_DIR=%CUR_DIR%/schema
set MTPROTO_DST_DIR=../mtproto
set RPC_DST_DIR=../mtproto

set GOGOPROTO_PATH=%CUR_DIR%/github.com/gogo/protobuf/;%CUR_DIR%/github.com/gogo/protobuf/protobuf/

REM protoc -I=%SRC_DIR% --proto_path=%GOGOPROTO_PATH% ^
REM      --gogo_out=plugins=grpc,Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,:%MTPROTO_DST_DIR% ^
REM      %SRC_DIR%/schema.*.proto

REM protoc -I=%SRC_DIR% --proto_path=%GOGOPROTO_PATH% ^
REM      --gogo_out=plugins=grpc,Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,:%RPC_DST_DIR% ^
REM      %SRC_DIR%/rpc.*.proto

rem python auto.py

protoc -I=%SRC_DIR% --proto_path=%GOGOPROTO_PATH% ^
     --gogo_out=plugins=grpc,Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,:%MTPROTO_DST_DIR% ^
     %SRC_DIR%/schema.*.proto

protoc -I=%SRC_DIR% --proto_path=%GOGOPROTO_PATH% ^
     --gogo_out=plugins=grpc,Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,:%RPC_DST_DIR% ^
     %SRC_DIR%/rpc.*.proto

gofmt -w %MTPROTO_DST_DIR%/*.go
gofmt -w %RPC_DST_DIR%/*.go