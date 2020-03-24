package consts

import "github.com/pkg/errors"

var (
	ErrorDB           = errors.New("db error")
	ErrorParam        = errors.New("请求解析错误")
	ErrorValidate     = errors.New("参数格式错误")
	ErrorDataNotFound = errors.New("data not found")

	ErrorUserAlreadyExist = errors.New("用户已存在")
)

// 也可以查看https://github.com/EDDYCJY/go-gin-example​里面的处理
