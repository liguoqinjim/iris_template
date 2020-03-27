package main

import (
	"github.com/kataras/iris/v12"
	"testing"
)

// 成功登录
func TestUserLoginSuccess(t *testing.T) {
	username := "admin"
	password := "123456"
	oj := map[string]string{
		"username": username,
		"password": password,
	}

	login(t, oj, iris.StatusOK, "success")
}

// 输入错误的登录密码
func TestUserLoginWithErrorPwd(t *testing.T) {
	oj := map[string]string{
		"username": "admin",
		"password": "1234567",
	}
	login(t, oj, iris.StatusOK, "用户名或密码错误")
}

// 不输入密码
func TestUserLoginWithNoPwd(t *testing.T) {
	oj := map[string]string{
		"username": "username",
		"password": "",
	}

	login(t, oj, iris.StatusOK, "参数格式错误")
}
