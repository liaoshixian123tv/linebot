package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

var ChannelScrect string
var channelAccessToken string

func init() {
	ChannelScrect = "9b2c1d8418b748fa85dd1b7eb6c67f9c"
	channelAccessToken = "0nQO8GdQRPIpERKz4UUCqRu7PqRop+SOmovUVTPcActC1cdXbE8t70rQIvfOzUbVrYYSs64BcaWaneYZoRbXbhMTf6cdKbrGAXn3gYMHM2oTosUGXN4IfFJncTurX7pT9ObXZoqqyLONKv0XsvbOBwdB04t89/1O/w1cDnyilFU="
	// lineID := "Ua4ca938776eddf76aa17a8e000f62e97"
}

func main() {

	bot, err := linebot.New(ChannelScrect, channelAccessToken)
	if err != nil {
		fmt.Println(err)
		return
	}
	http.HandleFunc("/callback", func(w http.ResponseWriter, req *http.Request) {
		events, err := bot.ParseRequest(req)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				w.WriteHeader(400)
			} else {
				w.WriteHeader(500)
			}
			return
		}
		for _, event := range events {
			var detail History
			tmp, err := event.MarshalJSON()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Printf("event: %s \n", string(tmp))

			if event.Type == linebot.EventTypeMessage {

				switch message := event.Message.(type) {

				case *linebot.TextMessage: // 一般文字
					detail.Content = message.Text
					detail.MessageID = message.ID
					detail.Timestamp = event.Timestamp.UnixMilli()
					detail.UserID = event.Source.UserID
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("Leo曰: "+message.Text)).Do(); err != nil {
						fmt.Println(err)
					}
				case *linebot.ImageMessage: // 圖片檔案

					by, err := message.MarshalJSON()
					if err != nil {
						fmt.Println(err)
					}
					fmt.Println("Image: ", string(by))
					content, err := bot.GetMessageContent(message.ID).Do()
					if err != nil {
						fmt.Println(err)
					}
					defer content.Content.Close()
					byts, err := ioutil.ReadAll(content.Content)
					if err != nil {
						fmt.Println(err)
					}
					if err = ioutil.WriteFile("/Users/shixianliao/workspace/src/linebot/abc.png", byts, 0); err != nil {
						fmt.Println(err)
					}

				case *linebot.VideoMessage: // 影片檔案
					by, err := message.MarshalJSON()
					if err != nil {
						fmt.Println(err)
					}
					fmt.Println("VideoMessage ", string(by))
					content, err := bot.GetMessageContent(message.ID).Do()
					if err != nil {
						fmt.Println(err)
					}
					defer content.Content.Close()
					byts, err := ioutil.ReadAll(content.Content)
					if err != nil {
						fmt.Println(err)
					}
					if err = ioutil.WriteFile("/Users/shixianliao/workspace/src/linebot/abc.mp4", byts, 0); err != nil {
						fmt.Println(err)
					}
				case *linebot.AudioMessage: // 音頻檔案
					by, err := message.MarshalJSON()
					if err != nil {
						fmt.Println(err)
					}
					fmt.Println("AudioMessage ", string(by))
					content, err := bot.GetMessageContent(message.ID).Do()
					if err != nil {
						fmt.Println(err)
					}
					defer content.Content.Close()
					byts, err := ioutil.ReadAll(content.Content)
					if err != nil {
						fmt.Println(err)
					}
					if err = ioutil.WriteFile("/Users/shixianliao/workspace/src/linebot/abcsound.mp4", byts, 0); err != nil {
						fmt.Println(err)
					}
				case *linebot.LocationMessage: // 位置檔案
					by, err := message.MarshalJSON()
					if err != nil {
						fmt.Println(err)
					}
					fmt.Println("LocationMessage ", string(by))

				case *linebot.StickerMessage: // 貼圖檔案
					by, err := message.MarshalJSON()
					if err != nil {
						fmt.Println(err)
					}
					fmt.Println("StickerMessage ", string(by))
					_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewStickerMessage("1257552", "10443983")).Do()
					if err != nil {
						fmt.Println(err)
					}
				}
			}
		}
	})

	// This is just sample code.
	// For actual use, you must support HTTPS by using `ListenAndServeTLS`, a reverse proxy or something else.
	fmt.Println("running on: 9090")
	if err := http.ListenAndServe(":9090", nil); err != nil {
		log.Fatal(err)
	}

}

// History 聊天記錄
type History struct {
	MessageID string  `json:"messageId"`         // MessageID 信息ID
	UserID    string  `json:"userId"`            // UserID 使用者ID
	Timestamp int64   `json:"timestamp"`         // Timestamp 發送時間
	Content   string  `json:"content,omitempty"` // Content 內容
	Sticker   Sticker `json:"sticker,omitempty"` // Sticker 貼圖
}

// Sticker 貼圖
type Sticker struct {
	PackageID string
	StickerID string
}
