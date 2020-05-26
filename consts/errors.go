package consts

var (
	SuccessCode = 200
	SuccessMsg  = "success"
)

var (
	//基础错误
	ErrInternal     = newError(-1, "内部错误") //内部错误，未通过自己定义的errorCode
	ErrDB           = newError(10001, "数据库错误")
	ErrParam        = newError(10002, "参数解析错误")
	ErrValidate     = newError(10003, "参数格式错误")
	ErrDataNotFound = newError(10004, "data not found")
	ErrSystem       = newError(10005, "系统内部错误")
	ErrRedis        = newError(10006, "redis error")
	ErrSms          = newError(10007, "短信SDK错误")
	ErrNoNeedUpdate = newError(10008, "数据无需更新")
	ErrNoAuth       = newError(10009, "无操作权限")
	ErrPhone        = newError(10010, "手机号不正确")

	//基础错误 jwt token
	ErrJwtTokenExpired = newError(10101, "jwt token过期")
	ErrJwtTokenUserId  = newError(10102, "jwt token中未发现userId")

	//
	ErrUserNotFound = newError(11008, "用户不存在")
)

type E struct {
	Msg      string
	Code     int
	Internal string //内部错误
}

func (e E) Error() string {
	return e.Msg
}

func newError(code int, msg string) E {
	return E{Code: code, Msg: msg}
}

func (e E) Clone(internal string) E {
	ee := newError(e.Code, e.Msg)
	ee.Internal = internal
	return ee
}
