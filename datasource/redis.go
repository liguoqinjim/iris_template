package datasource

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/liguoqinjim/iris_template/config"
	"log"
)

func initRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     config.Config.Redis.Addr,
		Password: config.Config.Redis.Password,
		DB:       config.Config.Redis.DB,
	})

	if err := RedisClient.Ping(context.TODO()).Err(); err != nil {
		log.Fatalf("redis ping error:%v", err)
	}
}
