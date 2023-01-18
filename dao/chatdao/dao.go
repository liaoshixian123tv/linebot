package chatdao

import (
	"context"
	"linebot/global"
	"linebot/global/tbname"
	"linebot/model"

	"go.mongodb.org/mongo-driver/bson"
)

// InsertMany 新增複數筆資料
func InsertOne(doc interface{}) (err error) {
	collection := global.MongoDBClient.Database(global.DatabaseSettings.DBName).Collection(tbname.ChatHistoryCollection())
	_, err = collection.InsertOne(context.TODO(), doc)
	return
}

// InsertMany 新增複數筆資料
func InsertMany(docs []interface{}) (err error) {
	collection := global.MongoDBClient.Database(global.DatabaseSettings.DBName).Collection(tbname.ChatHistoryCollection())
	_, err = collection.InsertMany(context.TODO(), docs)
	return
}

// GetAll 取得全部
func GetAll() (docs []model.History, err error) {
	collection := global.MongoDBClient.Database("linebot").Collection(tbname.ChatHistoryCollection())
	cur, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return
	}
	err = cur.All(context.TODO(), &docs)
	if err != nil {
		return
	}
	for idx, v := range docs {
		if v.MessageType == model.TextType {
			docs[idx].ContentStr = string(v.Content)
		}
	}
	return
}

// GetAll 取得全部
func GetByTimeRange(st, et int64) (docs []model.History, err error) {
	collection := global.MongoDBClient.Database("linebot").Collection(tbname.ChatHistoryCollection())
	filter := bson.M{"timestamp": bson.M{"$gte": st, "$lte": st}}
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return
	}
	err = cur.All(context.TODO(), &docs)
	if err != nil {
		return
	}
	for idx, v := range docs {
		if v.MessageType == model.TextType {
			docs[idx].ContentStr = string(v.Content)
		}
	}
	return
}
