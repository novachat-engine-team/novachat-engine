@echo off

set C_DIR=%CD%
set CUR_DIR=%~dp0..
set BASE_DIR=%~dp0

REM  set GOBIN=%BASE_DIR%install/go
set PATH=%CUR_DIR%/tools/protoc/win32/;%PATH%

for /f "delims=" = %%P in ('dir /ad/b %CUR_DIR%\cmd') do (
	echo start build %%P
	cd %CUR_DIR%\server\%%P
	go build -o %BASE_DIR%
	echo end %%P
)

cd %C_DIR%
