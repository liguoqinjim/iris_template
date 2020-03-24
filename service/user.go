package service

import (
	"github.com/liguoqinjim/iris_template/datamodel"
	"github.com/liguoqinjim/iris_template/repository"
	"github.com/liguoqinjim/iris_template/web/param"
	"github.com/pkg/errors"
)

type userService struct {
	repoUser repository.UserRepository
}

var UserService = new(userService)

func init() {
	UserService.repoUser = repository.NewUserRepository()
}

func (s userService) Exist(username string) (bool, error) {
	return s.repoUser.Exist(username)
}

func (s *userService) Register(param *param.RegisterParam) (*datamodel.User, error) {
	user := &datamodel.User{Username: param.Username, Password: param.Password}
	user, err := s.repoUser.Insert(user)

	return user, err
}

func (s *userService) Login(param *param.LoginParam) (*datamodel.User, error) {
	if user, err := s.repoUser.Get(param.Username); err != nil {
		return nil, err
	} else {
		if user.Password != param.Password {
			return nil, errors.New("用户名或密码错误")
		} else {
			return user, nil
		}
	}
}
