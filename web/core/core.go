package core

import (
	"github.com/kataras/iris/v12"
	"github.com/liguoqinjim/iris_template/logger"
	"github.com/liguoqinjim/iris_template/web/viewmodel"
)

func responseError(ctx iris.Context, err error) {
	ctx.StopExecution()
	ctx.JSON(viewmodel.Response{
		Code: 0,
		Msg:  err.Error(),
	})
}

func Response(ctx iris.Context, response interface{}, err error) {
	if err != nil {
		responseError(ctx, err)
	} else {
		ctx.JSON(viewmodel.Response{
			Code: 0,
			Msg:  "success",
			Data: response,
		})
	}
}

func GetReqID(ctx iris.Context) string {
	requestId := ctx.Values().GetString("req-id")

	return requestId
}

func HandleError(ctx iris.Context, err error) {
	logger.Debugf("core handler error:%v", err)

	ctx.JSON(viewmodel.Response{
		Code: 0,
		Msg:  err.Error(),
		Data: nil,
	})
}
