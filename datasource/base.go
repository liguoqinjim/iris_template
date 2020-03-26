package datasource

import (
	"github.com/go-redis/redis/v7"
	"github.com/jinzhu/gorm"
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
