#!/bin/bash


BASE_DIR=`pwd`
echo $BASE_DIR
CUR_DIR=$(dirname $(realpath $0))/..
echo $CUR_DIR
RET=

OUT=$CUR_DIR/gen/out
if [ ! -d "$OUT" ]; then
    mkdir $OUT
fi

if [ ! -d "$OUT/log" ]; then
    mkdir $OUT/log
fi


MOD_LIST=$(ls $CUR_DIR/cmd)

echo "Begin package yaml ............."

for e in ${MOD_LIST[@]}
do
    cp $CUR_DIR/pkg/cmd/$e/conf/conf.yaml $OUT/$e.yaml
done
echo "End package yaml ............."

echo "Begin package config.list ............."

for line in `cat $CUR_DIR/gen/config.list`
do
    echo $line
    cp $CUR_DIR/$line $OUT/
done

echo "Begin package config.list ............."

cd $BASE_DIR
