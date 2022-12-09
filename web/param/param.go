package param

type RegisterParam struct {
	Username string `json:"username" validate:"required" comment:"username"` //用户名
	Password string `json:"password" validate:"required" comment:"password"` //密码
}

type LoginParam struct {
	Username string `json:"username" validate:"required" comment:"用户名"` //用户名
	Password string `json:"password" validate:"required" comment:"密码"`  //密码
}

type QueryParam struct {
	Username string `json:"username" url:"username" validate:"required" comment:"username"` //用户名
	Page     int    `json:"page" url:"page" validate:"gte=1"`                               //页码
	PageSize int    `json:"pageSize" url:"page_size" validate:"gte=1,lte=100"`              //数量
}
