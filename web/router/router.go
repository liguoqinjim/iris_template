package router

import (
	"errors"
	"github.com/iris-contrib/swagger/v12"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/liguoqinjim/iris_template/config"
	"github.com/liguoqinjim/iris_template/logger"
	"github.com/liguoqinjim/iris_template/web/controller"
	"github.com/liguoqinjim/iris_template/web/core"
	"github.com/liguoqinjim/iris_template/web/middleware"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"strings"

	// swagger middleware for Iris
	// swagger embed files
	_ "github.com/liguoqinjim/iris_template/docs"
)

const (
	apiPrefix = "/api/v1"
)

func API(app *iris.Application) {

	mvc.Configure(app.Party(apiPrefix), func(app *mvc.Application) {
		app.Router.Use(middleware.Cors(), middleware.RequestId, middleware.LoggerHandler)

		if config.Config.JwtFlag {
			app.Router.Use(iris.NewConditionalHandler(func(ctx iris.Context) bool {
				if strings.HasSuffix(ctx.Path(), "login") {
					return false
				}
				return true
			}, middleware.Jwt()))
		}

		//总的错误处理
		app.HandleError(core.HandleError)

		app.Party("/user").Handle(new(controller.UserController))
		app.Party("/ping").Handle(new(controller.PingController))
	})

	//swagger
	swag(app)

	app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
		logger.Errorf(ctx.Path())
		core.Response(ctx, nil, errors.New(ctx.Path()+" path not found"))
	})

}

func swag(app *iris.Application) {
	//swaggerUrl := "http://localhost:18080/swagger/doc2.json"
	//url := swagger.URL(swaggerUrl)
	//_ = url
	//app.Get("/swagger/{any:path}", swagger.WrapHandler(swaggerFiles.Handler, url))
	//logger.Log("swag flag:", config.Config.SwaggerFlag)

	//可关闭
	app.Get("/swagger/{any:path}", swagger.DisablingWrapHandler(swaggerFiles.Handler, config.Config.SwaggerFlag))
}
