@echo OFF
setlocal EnableDelayedExpansion

set CUR_DIR=%~dp0

set ROOT_DIR=%CUR_DIR%../../../../

set PATH=%ROOT_DIR%/tools/protoc/win32;%PATH%

set SRC_DIR=%CUR_DIR%

set PB_DST_DIR=.

set GOGOPROTO_PATH=%ROOT_DIR%;%ROOT_DIR%/github.com/gogo/protobuf/;%ROOT_DIR%/github.com/gogo/protobuf/protobuf/;%ROOT_DIR%/schema;%SRC_DIR%

protoc -I=%SRC_DIR% --proto_path=%GOGOPROTO_PATH%;./ --gogo_out=plugins=grpc,Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,Mschema.tl.core_types.proto=novachat_engine/mtproto,Mschema.tl.sync_service.proto=novachat_engine/mtproto,Mschema.tl.sync.proto=novachat_engine/mtproto,Mschema.tl.sync.tl.proto=novachat_engine/mtproto,Mschema.tl.sync.proto=novachat_engine/mtproto,Mschema.tl.transport.proto=novachat_engine/mtproto,Mschema.data.proto=novachat_engine/mtproto:%PB_DST_DIR% %SRC_DIR%/*.proto

for %%i in (%PB_DST_DIR%/*.pb.go) do (
	echo %%i
	gofmt -w %%i
)
