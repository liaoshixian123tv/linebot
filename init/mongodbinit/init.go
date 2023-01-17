package mongodbinit

import (
	"context"
	"fmt"
	"linebot/global"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// MongoDBInit MongoDB初始化
func MongoDBInit() (err error) {
	url := "mongodb://localhost:27017"
	global.MongoDBClient, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(url))
	if err != nil {
		panic(err)
	}

	if err := global.MongoDBClient.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("MongoDB is ready")

	return
}
