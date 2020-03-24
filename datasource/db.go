package datasource

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"li.com/weixin/server/config"
	"li.com/weixin/server/datamodels"
	"li.com/weixin/server/logger"
	"time"
)

var (
	DB *gorm.DB
)

func init() {
	initDB()
}

func initDB() {
	connectInfo := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		config.Conf.Database.User,
		config.Conf.Database.Password,
		config.Conf.Database.Host,
		config.Conf.Database.Port,
		config.Conf.Database.DBName)
	logger.Infof("DB connectInfo:%s", connectInfo)

	var err error
	DB, err = gorm.Open("mysql", connectInfo)
	if err != nil {
		logger.Errorf("open DB error:%v", err)
	}

	DB.DB().SetConnMaxLifetime(time.Second * 60)
	DB.DB().SetMaxIdleConns(5)
	DB.DB().SetMaxOpenConns(10)

	if config.Conf.Debug {
		DB.LogMode(true)
	}

	DB.BlockGlobalUpdate(true)
}

func InitTestDB() {
	DB.AutoMigrate(
		&datamodels.User{}, &datamodels.Auth{})

	//测试数据
	DB.Create(&datamodels.User{
		Id:       1,
		Username: "admin",
		Password: "123456",
	})
}

func ResetTestDB() {
	DB.DropTable(
		&datamodels.User{}, &datamodels.Auth{})
}
