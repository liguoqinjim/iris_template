package core

import (
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/kataras/iris/v12"
	"github.com/liguoqinjim/iris_template/config"
	"github.com/liguoqinjim/iris_template/consts"
	"github.com/liguoqinjim/iris_template/logger"
	"github.com/liguoqinjim/iris_template/web/viewmodel"
)

func responseError(ctx iris.Context, err error) {
	//ctx.StopExecution()
	//ctx.JSON(viewmodel.Response{
	//	Code: common.SuccessCode,
	//	Msg:  err.Error(),
	//})

	//ctx.StopExecution()

	switch err.(type) {
	case consts.E:
		if config.Config.Debug {
			ctx.JSON(viewmodel.ResponseDebug{
				Code:  err.(consts.E).Code,
				Msg:   err.(consts.E).Msg,
				Data:  nil,
				Debug: err.(consts.E).Internal,
			})
		} else {
			ctx.JSON(viewmodel.Response{
				Code: err.(consts.E).Code,
				Msg:  err.(consts.E).Msg,
				Data: nil,
			})
		}
	default:
		ctx.JSON(viewmodel.Response{
			Code: consts.ErrInternal.Code,
			Msg:  err.Error(),
			Data: nil,
		})
	}
}

func Response(ctx iris.Context, response interface{}, err error) {
	if err != nil {
		responseError(ctx, err)
	} else {
		ctx.JSON(viewmodel.Response{
			Code: consts.SuccessCode,
			Msg:  consts.SuccessMsg,
			Data: response,
		})
	}
}

func ResponseExcel(ctx iris.Context, response interface{}, err error) {
	if err != nil {
		responseError(ctx, err)
	} else {

		xlsx := response.(*excelize.File)

		ctx.Header("Content-Type", "application/octet-stream")
		ctx.Header("Content-Disposition", "attachment; filename="+"Workbook.xlsx")
		ctx.Header("Content-Transfer-Encoding", "binary")
		ctx.Header("Expires", "0")

		xlsx.Write(ctx.ResponseWriter())
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
