package repository

import (
	"errors"
	"github.com/liguoqinjim/iris_template/consts"
	"github.com/liguoqinjim/iris_template/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserRepo interface {
	Insert(*model.User) (*model.User, error)
	Get(username string) (*model.User, error)
	Exist(username string) (bool, error)
}

type userRepo struct {
	db     *gorm.DB
	logger *zap.SugaredLogger
}

func NewUserRepo(db *gorm.DB, logger *zap.SugaredLogger) UserRepo {
	return &userRepo{
		db:     db,
		logger: logger,
	}
}

func (r *userRepo) Insert(user *model.User) (*model.User, error) {
	if err := r.db.Create(user).Error; err != nil {
		r.logger.Errorf("user Insert error:%v", err)
		return nil, consts.ErrDB
	} else {
		return user, nil
	}
}

func (r *userRepo) Get(username string) (*model.User, error) {
	user := new(model.User)
	if err := r.db.Where("username = ?", username).Take(user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, consts.ErrUserNotFound
		} else {
			r.logger.Errorf("user Get error:%v", err)
			return nil, consts.ErrDB
		}
	} else {
		return user, nil
	}
}

func (r *userRepo) Exist(username string) (bool, error) {
	var count int64
	if err := r.db.Model(&model.User{}).Where("username = ?", username).Count(&count).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		r.logger.Errorf("user Exist error:%v", err)
		return false, consts.ErrDB
	} else {
		if count != 0 {
			return true, nil
		} else {
			return false, nil
		}
	}
}
