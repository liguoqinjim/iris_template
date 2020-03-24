#!/bin/bash

echo 'new module name:'$1


#find . -name '*.go' -print0 | xargs -0 sed -i "" "s/form/forms/g"

# -e 测试
#origin="github.com/liguoqinjim/iris_template"
#sed -e 's,'"$origin"','"$1"',g' "web/controller/user.go"

# 修改所有go文件
find . -name '*.go' -print0 | xargs -0 sed -i '' 's,'"$origin"','"$1"',g'