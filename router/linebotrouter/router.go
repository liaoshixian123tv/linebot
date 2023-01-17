package linebotrouter

import (
	"linebot/global"
	"linebot/service/linebotservice"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

// AddLineBotRouter 新增line bot可用的router
func AddLineBotRouter(router *gin.Engine) {
	api := router.Group("")
	{
		api.POST("callback", receiveMessage)
	}
}

// receiveMessage 接收聊天資訊
func receiveMessage(c *gin.Context) {
	var res Message
	var statusCode int
	events, err := global.LineBotClient.ParseRequest(c.Request)
	statusCode, res.Message = errorHandle(err)
	if err != nil {
		c.JSON(statusCode, res)
		return
	}

	err = linebotservice.ReceiveMessage(events)
	statusCode, res.Message = errorHandle(err)
	if err != nil {
		c.JSON(statusCode, res)
		return
	}

	c.JSON(statusCode, res)
}

// errorHandle 錯誤分類
func errorHandle(err error) (int, string) {
	switch err {
	case nil:
		return http.StatusOK, "success"
	case linebot.ErrInvalidSignature:
		return http.StatusBadRequest, err.Error()
	default:
		return http.StatusInternalServerError, err.Error()
	}
}

// Message API回傳格式
type Message struct {
	Message string `json: "message"`
}
