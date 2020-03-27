#!/bin/bash

set -e

SOURCE_FILE_NAME=main
TARGET_FILE_NAME=iris_template

rm -fr ${TARGET_FILE_NAME}*

# todo 修改配置文件名称

build(){
    echo $GOOS $GOARCH

    tname=${TARGET_FILE_NAME}_${GOOS}_${GOARCH}${EXT}

    env GOOS=$GOOS GOARCH=$GOARCH \
    go build -o ${tname} \
    -v *.go

    chmod +x ${tname}
    mv ${tname} ${TARGET_FILE_NAME}${EXT}

    if [ ${GOOS} == "windows" ];then
        #zip ${tname}.zip ${TARGET_FILE_NAME}${EXT} config.ini ../public/
        zip ${tname}.zip ${TARGET_FILE_NAME}${EXT} app.toml
    else
        tar --exclude=*.gz  --exclude=*.zip  --exclude=*.git -czvf ${tname}.tar.gz ${TARGET_FILE_NAME}${EXT} app.toml *.sh -C ./ .
    fi
    mv ${TARGET_FILE_NAME}${EXT} ${tname}
}

CGO_ENABLED=0
#mac os 64
GOOS=darwin
GOARCH=amd64
build

#linux 64
GOOS=linux
GOARCH=amd64
build

#windows
#64
GOOS=windows
GOARCH=amd64
build

#32
GOARCH=386
build

ls -al