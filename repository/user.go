package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/liguoqinjim/iris_template/consts"
	"github.com/liguoqinjim/iris_template/datamodel"
	"github.com/liguoqinjim/iris_template/datasource"
	"github.com/liguoqinjim/iris_template/logger"
)

type UserRepo interface {
	Insert(*datamodel.User) (*datamodel.User, error)
	Get(username string) (*datamodel.User, error)
	Exist(username string) (bool, error)
}

type userRepo struct{}

func NewUserRepo() UserRepo {
	return &userRepo{}
}

func (r *userRepo) Insert(user *datamodel.User) (*datamodel.User, error) {
	if err := datasource.DB.Create(user).Error; err != nil {
		logger.Errorf("user Insert error:%v", err)
		return nil, consts.ErrDB
	} else {
		return user, nil
	}
}

func (r *userRepo) Get(username string) (*datamodel.User, error) {
	user := new(datamodel.User)
	if err := datasource.DB.Where("username = ?", username).Take(user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
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
	var count int
	if err := datasource.DB.Model(&datamodel.User{}).Where("username = ?", username).Count(&count).Error; err != nil && !gorm.IsRecordNotFoundError(err) {
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
