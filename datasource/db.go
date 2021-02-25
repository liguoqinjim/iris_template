package datasource

import (
	"fmt"
	"github.com/liguoqinjim/iris_template/config"
	"github.com/liguoqinjim/iris_template/logger"
	"github.com/liguoqinjim/iris_template/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"time"
)

func initDBMysql() {
	//mysql链接样例：username:password@tcp(host:port)/dbname?charset=utf8&parseTime=True&loc=Local&tls=skip-verify&autocommit=true
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		config.Config.Database.User,
		config.Config.Database.Password,
		config.Config.Database.Host,
		config.Config.Database.Port,
		config.Config.Database.DBName)
	logger.Infof("DB mysql connectInfo:%s", dsn)

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

func initDBPostgres() {
	//postgresql链接样例：host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		config.Config.Database.Host,
		config.Config.Database.User,
		config.Config.Database.Password,
		config.Config.Database.DBName,
		config.Config.Database.Port)
	logger.Infof("DB postgres connectInfo:%s", dsn)

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
	DB.AutoMigrate(&model.User{})

	//测试数据
	DB.Create(&model.User{
		Id:       1,
		Username: "admin",
		Password: "123456",
	})
}

func ResetTestDB() {
	DB.Migrator().DropTable(&model.User{})
}
