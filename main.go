package main

import (
	"context"
	"fmt"
	"linebot/global"
	"linebot/init/linebotinit"
	"linebot/init/mongodbinit"
	"linebot/router/linebotrouter"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	linebotinit.LineBotInit()
	mongodbinit.MongoDBInit()
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(corsMiddleware())
	linebotrouter.AddLineBotRouter(router)

	router.NoRoute(func(c *gin.Context) {
		var res pageNotFound
		res.Response = "PAGE_NOT_FOUND"
		c.JSON(http.StatusNotFound, res)
	})

	defer func() {
		if err := global.MongoDBClient.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	fmt.Println("running on: 9090")
	router.Run(":9090")
}

// corsMiddleware 允許cors請求
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 核心處理方式
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST")
		c.Set("content-type", "application/json")

		c.Next()
	}
}

// pageNotFound 錯誤router 的 response
type pageNotFound struct {
	Response string `json:"response"`
}
