# [iris_template](https://github.com/liguoqinjim/iris_template)

---
[![Build Status](https://travis-ci.org/liguoqinjim/iris_template.svg?branch=master)](https://travis-ci.org/liguoqinjim/iris_template)
[![codecov](https://codecov.io/gh/liguoqinjim/iris_template/branch/master/graph/badge.svg)](https://codecov.io/gh/liguoqinjim/iris_template)
---

## TODO
 - [x] go mod
 - [x] 登录，注册，得到用户信息功能
 - [ ] 密码加密存储
 - [ ] iris mvc controller
 - [ ] error handler
 - [ ] 错误处理
 - [x] api的单元测试 base_test.go
 - [x] gorm
 - [ ] sql脚本
 - [ ] redis
 - [x] swagger
 - [ ] jwt
 - [ ] cors
 - [ ] validator
 - [ ] travis
 - [ ] middleware
 - [ ] logger
 - [x] config /viper
 - [ ] 启动脚本
 - [ ] 打包脚本
 - [ ] Dockerfile
 - [ ] config_test.go
 - [ ] Makefile
 - [ ] 返回静态文件
 - [ ] websocket
 - [ ] BeforeActivation
 - [ ] casbin
 - [ ] 测试模式，可配置
 - [ ] /根路径的处理
 - [x] codecov覆盖率
 - [x] 重命名脚本
 - [x] 新建模板脚本
 - [ ] scp脚本
 - [ ] github badge
 - [ ] consts整理
 - [ ] 定时任务
 - [ ] recover middleware
 - [ ] 测试用的配置文件(比如travis.yml的mysql密码要为空，但是本地测试的数据库密码为123456)
 - [ ] server项目的错误处理移过来
 - [ ] server项目的错误返回
 - [ ] server项目的导出excel
 - [ ] sentry
 - [x] 升级到gorm v2
 - [ ] 在iris_template下再创建一个目录，把代码都移进去。感觉现在做的几个都是多个服务器分开的。默认就放在server下。但是这个再考虑下，其实也可以直接在路径下创建就可以了。
 - [ ] yapi
 - [x] datamodel->model

## swagger使用
 - 安装
 - `swag init`
 - 访问地址：`http://localhost:30100/swagger/index.html`
 
## NOTICE
 - `package consts`用了复数是因为const是golang关键字

## 脚本使用
 - 