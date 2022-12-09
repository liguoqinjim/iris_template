package controller

import (
	"github.com/kataras/iris/v12"
	"github.com/liguoqinjim/iris_template/logger"
	"github.com/liguoqinjim/iris_template/web/core"
)

type PingController struct{}

// Get @Summary ping
// @Description ping
// @Tags ping
// @Router /ping [get]
func (c *PingController) Get(ctx iris.Context) error {
	core.Response(ctx, "pong", nil)

	return nil
}

// HandleError 每个controller各自的error handler可以覆盖总的error handler
func (c *PingController) HandleError(ctx iris.Context, err error) {
	logger.Errorf("ping controller handler error:%v", err)

	core.Response(ctx, nil, err)
}
