#!/bin/bash

# 判断参数个数是否正确
if [ $# != 2 ]
then
    echo "usage: new.sh /new/path new.com/project/module"
    exit
fi

# 拷贝至新目录
# 忽略.idea和.git目录
rsync -av --exclude .idea --exclude .git --exclude logs --exclude releases --exclude app.toml --exclude .gitignore --exclude new.sh --exclude test.sh --exclude .travis.yml . $1

# 重命名包
cd $1
bash rename.sh $2
rm rename.sh

# 重命名app_template.toml
mv app_template.toml app.toml

# 重命名template.gitignore
mv template.gitignore .gitignore
