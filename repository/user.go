package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/liguoqinjim/iris_template/consts"
	"github.com/liguoqinjim/iris_template/datamodel"
	"github.com/liguoqinjim/iris_template/datasource"
	"github.com/liguoqinjim/iris_template/logger"
)

type UserRepository interface {
	Insert(*datamodel.User) (*datamodel.User, error)
	Get(username string) (*datamodel.User, error)
	Exist(username string) (bool, error)
}

type userRepository struct{}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) Insert(user *datamodel.User) (*datamodel.User, error) {
	err := datasource.DB.Create(user).Error
	logger.Debugf("user1:%v", user)

	return user, err
}

func (r *userRepository) Get(username string) (*datamodel.User, error) {
	user := new(datamodel.User)
	err := datasource.DB.Where("username = ?", username).Take(user).Error
	if err != nil && gorm.IsRecordNotFoundError(err) {
		err = consts.ErrorDataNotFound
	}

	return user, err
}

func (r *userRepository) Exist(username string) (bool, error) {
	user := new(datamodel.User)

	if err := datasource.DB.Select("id").Model(&datamodel.User{}).Where("username = ?", username).Take(user).Error; err != nil && !gorm.IsRecordNotFoundError(err) {
		return false, err
	} else {
		if user.Id != 0 {
			return true, nil
		} else {
			return false, nil
		}
	}
}
