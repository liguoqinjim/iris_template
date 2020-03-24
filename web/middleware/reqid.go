package middleware

import (
	"github.com/kataras/iris/v12"
	"github.com/satori/go.uuid"
	"github.com/liguoqinjim/iris_template/logger"
)

func RequestId(ctx iris.Context) {
	logger.Debugw("requestId middleware ing...")

	requestId := ctx.Request().Header.Get("req-id")

	if requestId == "" {
		u4 := uuid.NewV4()
		requestId = u4.String()
	}

	ctx.Values().Set("req-id", requestId)

	ctx.Header("req-id", requestId)
	ctx.Next()
}
