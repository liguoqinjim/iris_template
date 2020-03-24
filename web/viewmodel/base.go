package viewmodel

import (
	"github.com/liguoqinjim/iris_template/consts"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type QueryContent struct {
	PageNum  int         `json:"page_num"`
	PageSize int         `json:"page_size"`
	Total    int         `json:"total"`
	Result   interface{} `json:"result"`
}

var (
	ResponseSuccess = &Response{Code: consts.SuccessCode, Msg: consts.SuccessMessage}
	//ResponseParamError  = &Response{Code: consts.ErrorCodeParamError, Msg: consts.ErrorCodeParamErrorMessage}
	//ResponseDBError     = &Response{Code: consts.ErrorCodeDBError, Msg: consts.ErrorCodeDBErrorMessage}
	//ResponseSystemError = &Response{Code: consts.ErrorCodeSystemError, Msg: consts.ErrorCodeSystemErrorMessage}
	//
	//ResponseRequestError = &Response{Code: consts.ErrorCodeRequestParse, Msg: consts.ErrorCodeRequestParseMessage}
	//
	//ResponseLoginError    = &Response{Code: consts.ErrorCodeLogin, Msg: consts.ErrorCodeLoginMessage}
	//ResponseNotNeedModify = &Response{Code: consts.ErrorCodeNotNeedModify, Msg: consts.ErrorCodeNotNeedModifyMessage}
	//
	//ResponseSealPermissionError = &Response{Code: consts.ErrorCodeSealPermission, Msg: consts.ErrorCodeSealPermissionMessage}
	//
	//ResponseNoRecord = &Response{Code: consts.ErrorCodeNoRecord, Msg: consts.ErrorCodeNoRecordMessage}
)
