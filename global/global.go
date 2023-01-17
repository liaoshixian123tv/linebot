package global

import (
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"go.mongodb.org/mongo-driver/mongo"
)

var ChannelScrect string
var ChannelAccessToken string

var LineBotClient *linebot.Client
var MongoDBClient *mongo.Client

func init() {
	ChannelScrect = "9b2c1d8418b748fa85dd1b7eb6c67f9c"
	ChannelAccessToken = "0nQO8GdQRPIpERKz4UUCqRu7PqRop+SOmovUVTPcActC1cdXbE8t70rQIvfOzUbVrYYSs64BcaWaneYZoRbXbhMTf6cdKbrGAXn3gYMHM2oTosUGXN4IfFJncTurX7pT9ObXZoqqyLONKv0XsvbOBwdB04t89/1O/w1cDnyilFU="
	// lineID := "Ua4ca938776eddf76aa17a8e000f62e97"
	LineBotClient = new(linebot.Client)
	MongoDBClient = new(mongo.Client)

}
