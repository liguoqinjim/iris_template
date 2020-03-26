package datasource

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/liguoqinjim/iris_template/config"
	"github.com/liguoqinjim/iris_template/datamodel"
	"github.com/liguoqinjim/iris_template/logger"
	"time"
)

func initDB() {
	connectInfo := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		config.Config.Database.User,
		config.Config.Database.Password,
		config.Config.Database.Host,
		config.Config.Database.Port,
		config.Config.Database.DBName)
	logger.Infof("DB connectInfo:%s", connectInfo)

	var err error
	DB, err = gorm.Open("mysql", connectInfo)
	if err != nil {
		logger.Fatalf("open DB error:%v", err)
	}

	DB.DB().SetConnMaxLifetime(time.Second * 60)
	DB.DB().SetMaxIdleConns(5)
	DB.DB().SetMaxOpenConns(10)

	if config.Config.TestMode {
		DB.LogMode(true)
	}

	DB.BlockGlobalUpdate(true)
}

func InitTestDB() {
	DB.AutoMigrate(
		&datamodel.User{})

	//测试数据
	DB.Create(&datamodel.User{
		Id:       1,
		Username: "admin",
		Password: "123456",
	})
}

func ResetTestDB() {
	DB.DropTable(
		&datamodel.User{})
}
