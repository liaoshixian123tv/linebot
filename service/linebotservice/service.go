package linebotservice

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"linebot/global"
	"linebot/model"
	"time"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

// ReceiveMessage 接收訊息
func ReceiveMessage(events []*linebot.Event) (err error) {
	var details []model.History
	for _, event := range events {
		userID := event.Source.UserID
		timestamp := event.Timestamp.UnixNano() / int64(time.Millisecond)
		if event.Type == linebot.EventTypeMessage {
			var detail model.History
			switch message := event.Message.(type) {
			case *linebot.TextMessage: // 一般文字
				if detail, err = textHandle(message, timestamp, userID); err != nil {
					return err
				}
			case *linebot.ImageMessage: // 圖片檔案
				if detail, err = imageHandle(message, timestamp, userID); err != nil {
					return err
				}
			case *linebot.VideoMessage: // 影片檔案
				if detail, err = videoHandle(message, timestamp, userID); err != nil {
					return err
				}
			case *linebot.AudioMessage: // 音頻檔案
				if detail, err = aduioHandle(message, timestamp, userID); err != nil {
					return err
				}
			case *linebot.LocationMessage: // 位置檔案
				if detail, err = locationHandle(message, timestamp, userID); err != nil {
					return err
				}
			case *linebot.StickerMessage: // 貼圖檔案
				if detail, err = stickerHandle(message, timestamp, userID); err != nil {
					return err
				}
			}
			details = append(details, detail)
		}
	}
	fmt.Printf("Detail %+v \n", details)
	return
}

// getInfoByContent 根據content取得內容  (圖片，音頻，影片)
func getInfoByContent(messageID string, messageType model.MessageType) (byteArr []byte, err error) {
	if messageType == model.LocationType || messageType == model.TextType || messageType == model.StickerType {
		return byteArr, errors.New("not accept type")
	}
	content, err := global.LineBotClient.GetMessageContent(messageID).Do()
	if err != nil {
		return
	}
	defer content.Content.Close()
	byteArr, err = ioutil.ReadAll(content.Content)
	if err != nil {
		return
	}
	return
}

// textHandle 文字信息處理
func textHandle(message *linebot.TextMessage, timestamp int64, userID string) (detail model.History, err error) {
	byteArr, err := json.Marshal(message.Text)
	if err != nil {
		return detail, err
	}
	detail = model.History{
		Content:     byteArr,
		MessageID:   message.ID,
		Timestamp:   timestamp,
		UserID:      userID,
		MessageType: model.TextType,
	}
	fmt.Println(string(byteArr))
	return
}

// imageHandle 圖片信息處理
func imageHandle(message *linebot.ImageMessage, timestamp int64, userID string) (detail model.History, err error) {
	byteArr, err := getInfoByContent(message.ID, model.ImageType)
	if err != nil {
		return detail, err
	}
	detail = model.History{
		Content:     byteArr,
		MessageID:   message.ID,
		Timestamp:   timestamp,
		UserID:      userID,
		MessageType: model.ImageType,
	}
	return
}

// videoHandle 影片信息處理
func videoHandle(message *linebot.VideoMessage, timestamp int64, userID string) (detail model.History, err error) {
	byteArr, err := getInfoByContent(message.ID, model.VideoType)
	if err != nil {
		return detail, err
	}
	detail = model.History{
		Content:     byteArr,
		MessageID:   message.ID,
		Timestamp:   timestamp,
		UserID:      userID,
		MessageType: model.VideoType,
	}
	return
}

// aduioHandle 音頻信息處理
func aduioHandle(message *linebot.AudioMessage, timestamp int64, userID string) (detail model.History, err error) {
	byteArr, err := getInfoByContent(message.ID, model.AudioType)
	if err != nil {
		return detail, err
	}
	detail = model.History{
		Content:     byteArr,
		MessageID:   message.ID,
		Timestamp:   timestamp,
		UserID:      userID,
		MessageType: model.AudioType,
	}
	return
}

// locationHandle 地點信息處理
func locationHandle(message *linebot.LocationMessage, timestamp int64, userID string) (detail model.History, err error) {
	var location model.Location
	locationBytes, err := message.MarshalJSON()
	if err != nil {
		return detail, err
	}
	err = json.Unmarshal(locationBytes, &location)
	if err != nil {
		return detail, err
	}
	detail = model.History{
		Location:    location,
		MessageID:   message.ID,
		Timestamp:   timestamp,
		UserID:      userID,
		MessageType: model.LocationType,
	}
	return
}

// stickerHandle 貼圖信息處理
func stickerHandle(message *linebot.StickerMessage, timestamp int64, userID string) (detail model.History, err error) {
	var sticker model.Sticker
	stickerBytes, err := message.MarshalJSON()
	if err != nil {
		return detail, err
	}
	err = json.Unmarshal(stickerBytes, &sticker)
	if err != nil {
		return detail, err
	}
	detail = model.History{
		Sticker:     sticker,
		MessageID:   message.ID,
		Timestamp:   timestamp,
		UserID:      userID,
		MessageType: model.StickerType,
	}
	return
}
