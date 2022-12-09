package repository

import (
	"errors"
	"github.com/liguoqinjim/iris_template/consts"
	"github.com/liguoqinjim/iris_template/model"
	"github.com/liguoqinjim/iris_template/util"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserRepo interface {
	Insert(*model.User) (*model.User, error)
	Get(username string) (*model.User, error)
	Exist(username string) (bool, error)
	Query(username string, page, pageSize int) (int, []*model.User, error)
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

func (r *userRepo) Query(username string, page, pageSize int) (int, []*model.User, error) {
	/*
		select * from t_user where username like '%admin%';
	*/

	var total int64
	queryTotal := r.db.Table("t_user").
		Where("username like ?", util.LikeParam(username))

	if err := queryTotal.Count(&total).Error; err != nil {
		r.logger.Errorf("user Query total error:%v", err)
		return 0, nil, consts.ErrDB
	}
	if total == 0 {
		return 0, nil, nil
	}

	offset, limit := util.GetPageQueryParams(page, pageSize)
	query := r.db.Table("t_user").
		Where("username like ?", util.LikeParam(username)).
		Offset(offset).
		Limit(limit).
		Order("id")

	var us []*model.User
	if err := query.Find(&us).Error; err != nil {
		r.logger.Errorf("user Query find error:%v", err)
		return 0, nil, consts.ErrDB
	} else {
		return int(total), us, nil
	}
}
