package datamodel

import (
	"time"
)

type Auth struct {
	Id         int       `gorm:"column:id;primary_key;AUTO_INCREMENT;not null" json:"id"`
	AuthCode   string    `gorm:"column:auth_code;type:varchar(16);size:16;not null" json:"auth_code"`
	UserId     int       `gorm:"column:user_id;type:int(11);not null" json:"-"`
	AuthType   int       `gorm:"column:auth_type;type:tinyint(4);not null" json:"auth_type"`
	BuyType    int       `gorm:"column:buy_type;type:tinyint(4);not null" json:"buy_type"`
	BuyTime    time.Time `gorm:"column:buy_time;type:datetime;not null;default:CURRENT_TIMESTAMP" json:"buy_time"`
	ExpireTime time.Time `gorm:"column:expire_time;type:datetime;not null;default:0" json:"expire_time"`
}

func (Auth) TableName() string {
	return "t_auth"
}

type User struct {
	Id         int       `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	Username   string    `gorm:"column:username;type:varchar(32);size:32;not null" json:"username"`
	Password   string    `gorm:"column:password;type:varchar(64);size:64;not null" json:"password"`
	Phone      string    `gorm:"column:phone;type:char(11);size:11" json:"phone"`
	UserId     int       `gorm:"AUTO_INCREMENT"`
	CreateTime time.Time `gorm:"column:create_time;type:datetime;not null;default:CURRENT_TIMESTAMP" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time;type:datetime" json:"update_time"`
}

func (User) TableName() string {
	return "t_user"
}
