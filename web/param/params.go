package params

type RegisterParam struct {
	Username string `json:"username" validate:"required" comment:"username"`
	Password string `json:"password" validate:"required" comment:"password"`
}

type LoginParam struct {
	Username string `json:"username" validate:"required" comment:"用户名"`
	Password string `json:"password" validate:"required" comment:"密码"`
}
