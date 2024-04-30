#!/bin/bash


BASE_DIR=`pwd`
echo $BASE_DIR
CUR_DIR=$(dirname $(realpath $0))/..
echo $CUR_DIR
RET=

MOD_LIST=$(ls $CUR_DIR/cmd)

cmd=$1
if [ ! -n "$cmd" ]; then
  cmd="all"
fi

PLAT=`uname`
MAC="Darwin"
if [[ $PLAT =~ $MAC ]]; then
  PATH=$ROOT_DIR/tools/protoc/mac/:$PATH
else
  PATH=$ROOT_DIR/tools/protoc/linux/:$PATH
fi

#PROJECTNAME=codegen
#
#.PHONY: linux
#linux:
#	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build $(RACE) -o codegen .
#
#
#
#.PHONY: win
#win:
#	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build $(RACE) -o codegen.exe .
#
#
#.PHONY: mac
#mac:
#	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build $(RACE) -o codegen .
#
#.PHONY: aix
#aix:
#	CGO_ENABLED=0 GOOS=aix GOARCH=ppc64 go build $(RACE) -o codegen .
#

GOBIN=$CUR_DIR/gen/out

source go_build_env.sh

echo "Begin Build..............."
for e in ${MOD_LIST[@]}
do
    if [ "all" == $cmd ]; then
    echo $e
    cd $CUR_DIR/cmd/$e/
    echo "start build $e"
    RET=`CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "$LDFlags" -o $GOBIN . 2>&1`
    if [ $? == 0 ]; then
      echo end $e
      cd ../
    else
      echo "result: $RET "
      cd ../
      exit
    fi
  elif [ $cmd == $e ]; then
    echo $e
    cd $CUR_DIR/cmd/$e/
    echo "start build $e"
    RET=`CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "$LDFlags" -o $GOBIN . 2>&1`
    if [ $? == 0 ]; then
      echo end $e
      cd ../
    else
      echo "result: $RET "
      cd ../
      exit
    fi
  fi

done
echo "End Build ..............."

if [ "all" == $cmd ]; then
  cd $GOBIN

  echo ""
  echo "Begin Md5 ..............."
  md5sum  ${MOD_LIST[*]} > bin.md5
  echo "End Md5 ..............."
  echo ""

  echo "Begin tar ..............."
  #tar -zcvf bin.tar.gz ${MOD_LIST[*]} bin.md5
  echo "End tar ..............."
fi

cd $BASE_DIR

