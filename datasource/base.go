package datasource

import (
	"github.com/go-redis/redis/v7"
	"gorm.io/gorm"
)

var (
	//mysql
	DB *gorm.DB
	//redis
	Client *redis.Client
)

func init() {
	initDB()
	initRedis()
}
