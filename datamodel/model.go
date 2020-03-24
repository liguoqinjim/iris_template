package datamodel

import (
	"time"
)

type User struct {
	Id         int       `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	Username   string    `gorm:"column:username;type:varchar(32);size:32;not null" json:"username"`
	Password   string    `gorm:"column:password;type:varchar(64);size:64;not null" json:"-"`
	Phone      string    `gorm:"column:phone;type:char(11);size:11" json:"phone"`
	CreateTime time.Time `gorm:"column:create_time;type:datetime;not null;default:CURRENT_TIMESTAMP" json:"-"`
	UpdateTime time.Time `gorm:"column:update_time;type:datetime" json:"-"`
	Token      string    `gorm:"-" json:"token"`
}

func (User) TableName() string {
	return "t_user"
}
