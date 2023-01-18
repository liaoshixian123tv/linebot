package global

import (
	"linebot/model/config"

	"github.com/line/line-bot-sdk-go/v7/linebot"
	"go.mongodb.org/mongo-driver/mongo"
)

var LineBotClient *linebot.Client
var MongoDBClient *mongo.Client

var ServerSetting *config.ServerSettings
var DatabaseSettings *config.DatabaseSettings
var LinebotSettings *config.LinebotSettings

func init() {
	LineBotClient = new(linebot.Client)
	MongoDBClient = new(mongo.Client)

	ServerSetting = new(config.ServerSettings)
	DatabaseSettings = new(config.DatabaseSettings)
	LinebotSettings = new(config.LinebotSettings)
}
