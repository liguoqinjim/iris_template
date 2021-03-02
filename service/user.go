package service

import (
	"github.com/iris-contrib/middleware/jwt"
	"github.com/liguoqinjim/iris_template/config"
	"github.com/liguoqinjim/iris_template/consts"
	"github.com/liguoqinjim/iris_template/datasource"
	"github.com/liguoqinjim/iris_template/logger"
	"github.com/liguoqinjim/iris_template/model"
	"github.com/liguoqinjim/iris_template/repository"
	"github.com/liguoqinjim/iris_template/web/param"
	"time"
)

type userService struct {
	repoUser repository.UserRepo
}

var UserService = &userService{
	repoUser: repository.NewUserRepo(datasource.DB, logger.Log.Get()),
}

func (s *userService) Register(p *param.RegisterParam) (interface{}, error) {
	exist, err := s.repoUser.Exist(p.Username)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, consts.ErrUserRegistered
	}

	user := &model.User{
		Username:   p.Username,
		Password:   p.Password,
		CreateTime: time.Now(),
	}

	return s.repoUser.Insert(user)
}

func (s *userService) Login(p *param.LoginParam) (interface{}, error) {
	user, err := s.repoUser.Get(p.Username)
	if err != nil {
		return nil, err
	}
	if user.Password != p.Password {
		return nil, consts.ErrUserPassword
	}

	//jwt
	token := jwt.NewTokenWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Hour).Unix(),
		"iat": time.Now().Unix(),
		"u":   user.Id,
	})
	if t, err := token.SignedString([]byte(config.Config.Secret.Jwt)); err != nil {
		logger.Errorf("jwt token", "err", err)
		return nil, consts.ErrSystem
	} else {
		user.Token = t
		//todo 需要更新到mysql或者redis
	}

	return user, nil
}
