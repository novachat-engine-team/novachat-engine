@echo off

set CUR_DIR=%~dp0
set ROOT_DIR=%CUR_DIR%../../../../

set PATH=%ROOT_DIR%/tools/protoc/win32/;%PATH%

set SRC_DIR=%CUR_DIR%

set PB_DST_DIR=.

set GOGOPROTO_PATH=%ROOT_DIR%;%ROOT_DIR%/github.com/gogo/protobuf/;%ROOT_DIR%/github.com/gogo/protobuf/protobuf/

protoc -I=%SRC_DIR% --proto_path=%GOGOPROTO_PATH%;./ --gogo_out=plugins=grpc,Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,Mpkg/rpc/rpc/rpc.proto=novachat_engine/pkg/rpc/rpc:%PB_DST_DIR% %SRC_DIR%/*.proto

for %%i in (%PB_DST_DIR%/*.pb.go) do (
	echo %%i
	gofmt -w %%i
)
