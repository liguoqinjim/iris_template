#!/bin/bash

# 上传服务器
host=myhost

#打包
target=iris_template
GOOS=linux GOARCH=amd64 go build -o $target

#上传可执行文件
scp $target root@$host:/root/Workspace/iris_template
scp app_server.toml root@$host:/root/Workspace/iris_template/app.toml

echo '上传完成'