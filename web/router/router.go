package router

import (
	"errors"
	"github.com/iris-contrib/swagger/v12"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/_examples/mvc/login/web/controllers"
	"github.com/kataras/iris/v12/mvc"
	"github.com/liguoqinjim/iris_template/config"
	"github.com/liguoqinjim/iris_template/logger"
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
		if config.Conf.JwtFlag {
			app.Router.Use(iris.NewConditionalHandler(func(ctx iris.Context) bool {

				if strings.HasSuffix(ctx.Path(), "login") {
					return false
				}
				return true
			}))
		}

		//app.HandleError(core.HandleError)

		app.Party("/login").Handle(new(controllers.UserController))
	})

	//rootParty := app.Party("/", middleware.Cors(), middleware.RequestId, middleware.LoggerHandler) //这里可以加上cors的middleware
	//{
	//
	//	v1 := rootParty.Party(apiPrefix)
	//	app := mvc.New(v1)
	//	app.HandleError(core.HandleError)
	//
	//	if config.Conf.JwtFlag {
	//		v1.Use(iris.NewConditionalHandler(func(ctx iris.Context) bool {
	//
	//			if strings.HasSuffix(ctx.Path(), "login") {
	//				return false
	//			}
	//			return true
	//		}))
	//	}
	//
	//	{
	//
	//		mvc.New(v1.Party("/login")).Handle(new(controllers.LoginController))
	//		mvc.New(v1.Party("/callback")).Handle(new(controllers.CallbackController))
	//
	//		//todo jwt
	//
	//		mvc.New(v1.Party("/user")).Handle(new(controllers.UserController))
	//		mvc.New(v1.Party("/auth")).Handle(new(controllers.AuthController))
	//	}
	//
	//}

	//todo
	//app.Get("/ws", websocket.Handler(ws.Ws))

	//swagger
	swag(app)

	app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
		// error code handlers are not sharing the same middleware as other routes, so we have
		// to call them inside their body.
		logger.Errorf(ctx.Path())
		core.Response(ctx, nil, errors.New(ctx.Path()+" path not found"))
	})

}

func swag(app *iris.Application) {
	//swaggerUrl := "http://localhost:18080/swagger/doc2.json"
	//url := swagger.URL(swaggerUrl)
	//_ = url
	//app.Get("/swagger/{any:path}", swagger.WrapHandler(swaggerFiles.Handler, url))
	//logger.Log("swag flag:", config.Conf.SwaggerFlag)

	//可关闭
	app.Get("/swagger/{any:path}", swagger.DisablingWrapHandler(swaggerFiles.Handler, config.Conf.SwaggerFlag))
}
