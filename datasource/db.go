package datasource

import (
	"fmt"
	"github.com/liguoqinjim/iris_template/config"
	"github.com/liguoqinjim/iris_template/logger"
	"github.com/liguoqinjim/iris_template/model"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"time"
)

func initDBMysql() {
	//mysql链接样例：username:password@tcp(host:port)/dbname?charset=utf8&parseTime=True&loc=Local&tls=skip-verify&autocommit=true
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		config.Config.Mysql.User,
		config.Config.Mysql.Password,
		config.Config.Mysql.Host,
		config.Config.Mysql.Port,
		config.Config.Mysql.DBName)
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

	logger.Infof("db mysql connect success")
}

func initDBPostgres() {
	//postgresql链接样例：host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		config.Config.Postgres.Host,
		config.Config.Postgres.User,
		config.Config.Postgres.Password,
		config.Config.Postgres.DBName,
		config.Config.Postgres.Port)
	logger.Infof("DB postgres connectInfo:%s", dsn)

	gormCf := &gorm.Config{
		AllowGlobalUpdate: false,
	}
	if config.Config.Debug {
		gormCf.Logger = gormlogger.Default.LogMode(gormlogger.Info)
	}

	var err error
	DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), gormCf)
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

	logger.Infof("db postgres connect success")
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
