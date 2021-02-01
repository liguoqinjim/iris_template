package repository

import (
	"errors"
	"github.com/liguoqinjim/iris_template/consts"
	"github.com/liguoqinjim/iris_template/datasource"
	"github.com/liguoqinjim/iris_template/logger"
	"github.com/liguoqinjim/iris_template/model"
	"gorm.io/gorm"
)

type UserRepo interface {
	Insert(*model.User) (*model.User, error)
	Get(username string) (*model.User, error)
	Exist(username string) (bool, error)
}

type userRepo struct{}

func NewUserRepo() UserRepo {
	return &userRepo{}
}

func (r *userRepo) Insert(user *model.User) (*model.User, error) {
	if err := datasource.DB.Create(user).Error; err != nil {
		logger.Errorf("user Insert error:%v", err)
		return nil, consts.ErrDB
	} else {
		return user, nil
	}
}

func (r *userRepo) Get(username string) (*model.User, error) {
	user := new(model.User)
	if err := datasource.DB.Where("username = ?", username).Take(user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, consts.ErrUserNotFound
		} else {
			logger.Errorf("user Get error:%v", err)
			return nil, consts.ErrDB
		}
	} else {
		return user, nil
	}
}

func (r *userRepo) Exist(username string) (bool, error) {
	var count int64
	if err := datasource.DB.Model(&model.User{}).Where("username = ?", username).Count(&count).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logger.Errorf("user Exist error:%v", err)
		return false, consts.ErrDB
	} else {
		if count != 0 {
			return true, nil
		} else {
			return false, nil
		}
	}
}
