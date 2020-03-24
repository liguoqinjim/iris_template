package consts

const (
	//token有效期
	TokenPeriodValidity = 3600 * 24 * 7
)

var (
	AESKey = []byte("app001_seal_1209")
)

const (
	OperatorInit      = iota + 1 //未使用
	OperatorSeal                 //施封完成
	OperatorCheck                //巡检中
	OperatorOpen                 //拆封
	OperatorException            //异常
)

const (
	CtxKeyTokenUser = "u"
)

//注册状态
const (
	RegisterStatusInit = iota //未处理
	RegisterStatusPass        //通过
	RegisterStatusDeny        //拒绝
)

//账号是否可用
const (
	UserEnable  = 1
	UserDisable = 0
)

const (
	CompanyParentFlag = -1 //总公司的标记
)

const (
	TokenTypeAppUser = 1 //app用户的token
	TokenTypeAdmin   = 2 //admin的token
)

const (
	PageNum  = 1
	PageSize = 10
)

const (
	ExceptionType01 = iota + 1 //封条外部有损坏
	ExceptionType02            //封条丢失
	ExceptionType03            //封条读出的封条号和外部封条号不一致
	ExceptionType04            //封条无法读出
	ExceptionType05            //其他原因
)

//异常是否确认
const (
	ExceptionConfirmed    = 1  //已确认为异常
	ExceptionNotConfirmed = -1 //没有确认为异常
)

//版本类型
const (
	VersionTypeApp    = iota + 1 //app版本号
	VersionTypeServer            //服务器版本号
)

const (
	ReportImageEmpty = "" //上传的时候没有照片的时候用这个值
)
