package main

import (
	"github.com/gavv/httpexpect"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/httptest"
	"github.com/liguoqinjim/iris_template/datasource"
	"os"
	"strings"
	"testing"
)

var (
	app *iris.Application
)

func TestMain(m *testing.M) {
	app = newTestApp()
	datasource.InitTestDB()

	exitCode := m.Run()

	//清空数据库
	datasource.ResetTestDB()

	os.Exit(exitCode)
}

func login(t *testing.T, object interface{}, statusCode int, values ...interface{}) (e *httpexpect.Expect) {
	e = httptest.New(t, app)

	e.POST("/api/v1/user/login").WithJSON(object).Expect().Status(statusCode).
		JSON().Object().Values().Contains(values...)

	return e
}

func post(t *testing.T, url string, object interface{}, statusCode int, values []interface{}) (e *httpexpect.Expect) {
	if !strings.HasPrefix(url, "/") {
		url = "/" + url
	}

	e = httptest.New(t, app)

	e.POST(url).WithJSON(object).Expect().Status(statusCode).
		JSON().Object().Values().Contains(values...)

	return e
}
