package datasource

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var (
	//mysql
	DB *gorm.DB
	//redis
	RedisClient *redis.Client
)

func init() {
	//initDBMysql()
	initDBPostgres()
	initRedis()
}
