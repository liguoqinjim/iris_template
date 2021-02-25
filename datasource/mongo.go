package datasource

import (
	"context"
	"fmt"
	"github.com/liguoqinjim/iris_template/config"
	"github.com/liguoqinjim/iris_template/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

func initMongo() {
	//mongodb://root@localhost:27017/?authSource=admin
	uri := fmt.Sprintf(
		"mongodb://%s:%s@%s:%d/?authSource=%s",
		config.Config.Mongo.User,
		config.Config.Mongo.Password,
		config.Config.Mongo.Host,
		config.Config.Mongo.Port,
		config.Config.Mongo.AuthSource)
	logger.Infof("mongo uri=%s", uri)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	var err error
	defer cancel()
	MongoClient, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		logger.Fatalf("mongo connect error:%v", err)
	}

	// Ping the primary
	if err := MongoClient.Ping(ctx, readpref.Primary()); err != nil {
		logger.Fatalf("mongo ping error:%v", err)
	} else {
		logger.Infof("mongo ping success")
	}
}
