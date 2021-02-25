package datasource

import (
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

var (
	//db
	DB *gorm.DB
	//redis
	RedisClient *redis.Client
	//mongodb
	MongoClient *mongo.Client
)

func init() {
	//initDBMysql()
	//initDBPostgres()
	//initRedis()
	//initMongo()
}
