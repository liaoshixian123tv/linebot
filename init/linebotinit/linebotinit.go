package linebotinit

import (
	"linebot/global"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

// LineBotInit 初始化line bot 機器人
func LineBotInit() (err error) {
	global.LineBotClient, err = linebot.New(global.LinebotSettings.ChannelScrect, global.LinebotSettings.ChannelAccessToken)
	if err != nil {
		panic(err)
	}
	return
}
