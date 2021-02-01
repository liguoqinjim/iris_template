package viewmodel

import (
	"github.com/jinzhu/copier"
	"github.com/liguoqinjim/iris_template/model"
	"github.com/liguoqinjim/iris_template/logger"
	"time"
)

type User struct {
	Id         int       `json:"id"`
	Username   string    `json:"username"`
	Phone      string    `json:"phone"`
	CreateTime time.Time `json:"create_time"`
	Token      string    `json:"token"`
}

func NewUser(user *model.User, token string) *User {
	u := new(User)
	err := copier.Copy(u, user)
	u.Token = token
	if err != nil {
		logger.Errorf("copy error:%v", err)
	}
	logger.Debugf("user:%v", u)
	return u
}
