package main

import (
	"github.com/kataras/iris/v12"
	"github.com/liguoqinjim/iris_template/config"
	"github.com/liguoqinjim/iris_template/logger"
	"github.com/liguoqinjim/iris_template/web/router"
)

func NewApp() *iris.Application {
	app := iris.New()
	app.Logger().SetLevel(config.Conf.IrisLoggerLevel)
	app.Logger().Install(logger.Log)

	routers.API(app)

	iris.RegisterOnInterrupt(func() {
		//todo 在这里处理关闭的操作，比如数据库关闭
		logger.Infof("interrupt...")
	})

	return app
}

func NewTestApp() *iris.Application {
	app := iris.New()
	routers.API(app)

	iris.RegisterOnInterrupt(func() {
		//todo 在这里处理关闭的操作，比如数据库关闭
		logger.Infof("interrupt...")
	})

	return app
}
