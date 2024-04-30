@echo off

set CUR_DIR=%~dp0
set ROOT_DIR=%CUR_DIR%../../../..

set PATH=%ROOT_DIR%/tools/protoc/win32/;%PATH%

set SRC_DIR=%CUR_DIR%

set PB_DST_DIR=.

GOGOPROTO_PATH=%ROOT_DIR%;%ROOT_DIR%/github.com/gogo/protobuf/;%ROOT_DIR%/github.com/gogo/protobuf/protobuf/;%ROOT_DIR%/service/data/fs;%ROOT_DIR%/service/data/restriction;%ROOT_DIR%/service/data/chat;%ROOT_DIR%/schema;%SRC_DIR%

protoc -I=%SRC_DIR% --proto_path=%GOGOPROTO_PATH%;./ --gogo_out=plugins=grpc,Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,Mphoto.proto=novachat_engine/service/data/fs,Mgeopoint.proto=novachat_engine/service/data/fs,Mdocument.proto=novachat_engine/service/data/fs,Mrestriction_reason.proto=novachat_engine/service/data/restriction:%PB_DST_DIR% %SRC_DIR%/*.proto
for %%i in (%PB_DST_DIR%/*.pb.go) do (
	echo %%i
	gofmt -w %%i
)
