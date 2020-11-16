package main

import (
	"github.com/liguoqinjim/iris_template/consts"
	"net/http"
	"testing"
)

//注册
func TestUserRegister(t *testing.T) {
	url := "/api/v1/user/register"

	t.Run("注册成功", func(t *testing.T) {
		obj := map[string]interface{}{
			"username": "admin2",
			"password": "1234567",
		}

		post(t, url, obj, http.StatusOK, []interface{}{0, "success"})
	})

	// todo 这个功能改到登录成功之后，使用jwt token获取资源
	//t.Run("注册成功的账号的使用授权码", func(t *testing.T) {
	//	userId := 2
	//	e := httptest.New(t, app)
	//
	//	things := e.GET("/api/v1/auth/all/{x}", "2").Expect().Status(http.StatusOK).JSON()
	//
	//	things.Schema(auth_schema)
	//
	//	things.Path("$.code").Number().Equal(0)
	//	things.Path("$.msg").String().Equal("success")
	//
	//	things.Path("$.data.auths[*].auth_code").Array().Contains(fmt.Sprintf("test_%09d", userId))
	//})

	t.Run("注册参数不正确", func(t *testing.T) {
		obj1 := map[string]interface{}{
			"username": "admin",
		}
		obj2 := map[string]interface{}{
			"password": "123456",
		}

		post(t, url, obj1, http.StatusOK, []interface{}{consts.ErrValidate.Error()})
		post(t, url, obj2, http.StatusOK, []interface{}{consts.ErrValidate.Error()})
	})

	t.Run("注册用户名已存在", func(t *testing.T) {
		obj := map[string]interface{}{
			"username": "admin",
			"password": "123456",
		}

		post(t, url, obj, http.StatusOK, []interface{}{consts.ErrUserNotFound.Error()})
	})
}

var auth_schema = `{
    "$schema": "http://json-schema.org/draft-07/schema",
    "$id": "http://example.com/example.json",
    "type": "object",
    "title": "The Root Schema",
    "description": "The root schema comprises the entire JSON document.",
    "required": [
        "code",
        "msg",
        "data"
    ],
    "properties": {
        "code": {
            "$id": "#/properties/code",
            "type": "integer",
            "title": "The Code Schema",
            "description": "An explanation about the purpose of this instance.",
            "default": 0
        },
        "msg": {
            "$id": "#/properties/msg",
            "type": "string",
            "title": "The Msg Schema",
            "description": "An explanation about the purpose of this instance.",
            "default": ""
        },
        "data": {
            "$id": "#/properties/data",
            "type": "object",
            "title": "The Data Schema",
            "description": "An explanation about the purpose of this instance.",
            "default": {},
            "required": [
                "auths"
            ],
            "properties": {
                "auths": {
                    "$id": "#/properties/data/properties/auths",
                    "type": "array",
                    "title": "The Auths Schema",
                    "description": "An explanation about the purpose of this instance.",
                    "default": [],
                    "items": {
                        "$id": "#/properties/data/properties/auths/items",
                        "type": "object",
                        "title": "The Items Schema",
                        "description": "An explanation about the purpose of this instance.",
                        "default": {},
                        "required": [
                            "id",
                            "auth_code",
                            "auth_type",
                            "buy_type",
                            "buy_time",
                            "expire_time"
                        ],
                        "properties": {
                            "id": {
                                "$id": "#/properties/data/properties/auths/items/properties/id",
                                "type": "integer",
                                "title": "The Id Schema",
                                "description": "An explanation about the purpose of this instance.",
                                "default": 0
                            },
                            "auth_code": {
                                "$id": "#/properties/data/properties/auths/items/properties/auth_code",
                                "type": "string",
                                "title": "The Auth_code Schema",
                                "description": "An explanation about the purpose of this instance.",
                                "default": ""
                            },
                            "auth_type": {
                                "$id": "#/properties/data/properties/auths/items/properties/auth_type",
                                "type": "integer",
                                "title": "The Auth_type Schema",
                                "description": "An explanation about the purpose of this instance.",
                                "default": 0
                            },
                            "buy_type": {
                                "$id": "#/properties/data/properties/auths/items/properties/buy_type",
                                "type": "integer",
                                "title": "The Buy_type Schema",
                                "description": "An explanation about the purpose of this instance.",
                                "default": 0
                            },
                            "buy_time": {
                                "$id": "#/properties/data/properties/auths/items/properties/buy_time",
                                "type": "string",
                                "title": "The Buy_time Schema",
                                "description": "An explanation about the purpose of this instance.",
                                "default": ""
                            },
                            "expire_time": {
                                "$id": "#/properties/data/properties/auths/items/properties/expire_time",
                                "type": "string",
                                "title": "The Expire_time Schema",
                                "description": "An explanation about the purpose of this instance.",
                                "default": ""
                            }
                        }
                    }
                }
            }
        }
    }
}
`
