package param

type RegisterParam struct {
	Username string `json:"username" validate:"required" comment:"username"` //用户名
	Password string `json:"password" validate:"required" comment:"password"` //密码
}

type LoginParam struct {
	Username string `json:"username" validate:"required" comment:"用户名"` //用户名
	Password string `json:"password" validate:"required" comment:"密码"`  //密码
}
