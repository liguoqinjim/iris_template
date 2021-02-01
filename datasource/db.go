package datasource

import (
	"fmt"
	"github.com/liguoqinjim/iris_template/config"
	"github.com/liguoqinjim/iris_template/datamodel"
	"github.com/liguoqinjim/iris_template/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"time"
)

func initDB() {
	//mysql链接样例：username:password@tcp(host:port)/dbname?charset=utf8&parseTime=True&loc=Local&tls=skip-verify&autocommit=true
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		config.Config.Database.User,
		config.Config.Database.Password,
		config.Config.Database.Host,
		config.Config.Database.Port,
		config.Config.Database.DBName)
	logger.Infof("DB connectInfo:%s", dsn)

	gormCf := &gorm.Config{
		AllowGlobalUpdate: false,
	}
	if config.Config.Debug {
		gormCf.Logger = gormlogger.Default.LogMode(gormlogger.Info)
	}

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), gormCf)
	if err != nil {
		logger.Fatalf("open DB error:%v", err)
	}

	sqlDB, err := DB.DB()
	if err != nil {
		logger.Fatalf("DB.DB() error:%v", err)
	}

	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetConnMaxLifetime(time.Second * 60)
}

func InitTestDB() {
	DB.AutoMigrate(&datamodel.User{})

	//测试数据
	DB.Create(&datamodel.User{
		Id:       1,
		Username: "admin",
		Password: "123456",
	})
}

func ResetTestDB() {
	DB.Migrator().DropTable(&datamodel.User{})
}
