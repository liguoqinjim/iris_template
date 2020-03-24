package middleware

import (
	"github.com/kataras/iris/v12"
	"github.com/liguoqinjim/iris_template/logger"
	"github.com/liguoqinjim/iris_template/web/core"
	"net/http/httputil"
	"time"
)

func LoggerHandler(ctx iris.Context) {
	start := time.Now().UTC()
	path := ctx.Request().URL.Path

	//todo
	//跳过一些path

	//跳过健康检测请求
	if path == "/swggg/health" || path == "/sd/ram" {
		return
	}

	ip := ctx.RemoteAddr()
	dumpReq, _ := httputil.DumpRequest(ctx.Request(), true)
	if dumpReq != nil {
		logger.Debugw("Request start", "requestId", core.GetReqID(ctx), "description", string(dumpReq))
	}

	ctx.Record()
	ctx.Next()

	end := time.Now().UTC()
	latency := end.Sub(start).String()

	//要使用ctx.Recorder()，需要先调用ctx.Record()
	logger.Infow("Request end", "requestId", core.GetReqID(ctx), "latency", latency, "ip", ip, "path", path, "body", ctx.Recorder().Body())
}
