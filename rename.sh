#!/bin/bash

echo 'new module name:'$1

# -e 测试
#origin="github.com/liguoqinjim/iris_template"
#sed -e 's,'"$origin"','"$1"',g' "web/controller/user.go"

origin="github.com/liguoqinjim/iris_template"
# sed -i '' 's,'"$a"','"$1"',g' "web/controller/user.go"

# 修改所有go文件
find . -name '*.go' -print0 | xargs -0 sed -i '' 's,'"$origin"','"$1"',g'
# 修改go.mod
sed -i '' 's,'"$origin"','"$1"',g' go.mod

pwd=$(PWD)
arr=(${pwd//\// })
len=${#arr[@]}-1
# 当前文件夹名
dir=${arr[$len]}
# 替换replace的路径
sed -i '' 's,'iris_template','$dir',g' go.mod

