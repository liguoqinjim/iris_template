package viewmodel

import (
	"github.com/liguoqinjim/iris_template/consts"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type ResponseDebug struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data"`
	Debug string      `json:"debug"`
}

//分页结果
type QueryContent struct {
	Page   int         `json:"page"`   //当前页
	Total  int         `json:"total"`  //总数
	Result interface{} `json:"result"` //结果
}

var (
	ResponseSuccess = &Response{Code: consts.SuccessCode, Msg: consts.SuccessMsg}
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
