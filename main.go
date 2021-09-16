package main

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/liguoqinjim/iris_template/config"
	"github.com/liguoqinjim/iris_template/logger"
)

// @title Iris Template API
// @version 1.0
// @description This is a sample iris server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host http://localhost:18080
// @BasePath /api/v1
func main() {
	app := NewApp()

	if err := app.Run(
		iris.Addr(fmt.Sprintf(":%d", config.Config.Web.Port)),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	); err != nil {
		logger.Fatalf("app.Run error:%v", err)
	}
}
