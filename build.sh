#!/bin/bash

set -e

SOURCE_FILE_NAME=main
TARGET_FILE_NAME=iris_template
RELEASE_PATH=releases/

mkdir -p releases
rm -fr ${TARGET_FILE_NAME}*
rm -rf ${RELEASE_PATH}*

build(){
    echo "build:"$GOOS $GOARCH

    tname=${TARGET_FILE_NAME}_${GOOS}_${GOARCH}${EXT}

    env GOOS=$GOOS GOARCH=$GOARCH \
    go build -o ${tname} \
    -v *.go

    chmod +x ${tname}
    mv ${tname} ${TARGET_FILE_NAME}${EXT}

    if [ ${GOOS} == "windows" ];then
        #zip ${tname}.zip ${TARGET_FILE_NAME}${EXT} config.ini ../public/
        filename=${tname}.zip
        zip $filename ${TARGET_FILE_NAME}${EXT} app.toml
        mv ${filename} ${RELEASE_PATH}${filename}
    else
        filename=${tname}.tar.gz
        #tar --exclude=*.gz  --exclude=*.zip  --exclude=*.git --exclude=log -czvf ${filename} ${TARGET_FILE_NAME}${EXT} app.toml *.sh -C ./ .
        tar -czvf ${filename} ${TARGET_FILE_NAME}${EXT} app.toml
        mv ${filename} ${RELEASE_PATH}${filename}
    fi

    # 只上传压缩之后的包
    # mv ${TARGET_FILE_NAME}${EXT} ${RELEASE_PATH}${tname}
    rm ${TARGET_FILE_NAME}${EXT}
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