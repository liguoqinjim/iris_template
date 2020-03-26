package datasource

import (
	"github.com/go-redis/redis/v7"
	"github.com/liguoqinjim/iris_template/config"
	"log"
)

func initRedis() {
	Client = redis.NewClient(&redis.Options{
		Addr:     config.Config.Redis.Addr,
		Password: config.Config.Redis.Password,
		DB:       config.Config.Redis.DB,
	})

	if err := Client.Ping().Err(); err != nil {
		log.Fatalf("redis ping error:%v", err)
	}
}
