package linebotrouter

import (
	"linebot/global"
	"linebot/model"
	"linebot/service/linebotservice"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

// AddLineBotRouter 新增line bot可用的router
func AddLineBotRouter(router *gin.Engine) {
	api := router.Group("")
	{
		api.GET("getmessage", getMessage)
		api.POST("receivemessage", receiveMessage)
	}
}

/*
Header: startTime 、 endTime ， 倘若startTime 、 endTime 都為空則取全部 否則取時間費圍內的聊天紀錄含頭尾
*/
// getMessage 取得聊天紀錄
func getMessage(c *gin.Context) {
	var res Message
	var statusCode int
	var arr []model.History
	startTime, endTime := c.GetHeader("startTime"), c.GetHeader("endTime")
	if startTime == "" && endTime == "" {
		tmpArr, err := linebotservice.GetAllChat()
		statusCode, res.Message = errorHandle(err)
		if err != nil {
			c.JSON(statusCode, res)
			return
		}
		arr = tmpArr
	} else {
		tmpArr, err := linebotservice.GetByTimeRange(startTime, endTime)
		statusCode, res.Message = errorHandle(err)
		if err != nil {
			c.JSON(statusCode, res)
			return
		}
		arr = tmpArr
	}
	c.JSON(statusCode, arr)
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
