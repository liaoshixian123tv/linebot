package model

const (
	TextType     MessageType = iota + 1 // TextType 文字
	ImageType                           // ImageType 圖片
	AudioType                           // AudioType 音頻
	VideoType                           // VideoType 影片
	LocationType                        // LocationType 地點位置
	StickerType                         // StickerType 貼圖
)

// History 聊天記錄
type History struct {
	Content     []byte      `json:"content,omitempty"` // Content 內容
	MessageID   string      `json:"messageId"`         // MessageID 信息ID
	UserID      string      `json:"userId"`            // UserID 使用者ID
	Timestamp   int64       `json:"timestamp"`         // Timestamp 發送時間
	Sticker     Sticker     `json:"sticker,omitempty"` // Sticker 貼圖
	Location    Location    `json:"location,omitempty"`
	MessageType MessageType `json:"messageType"`          // MessageType 信息種類
	ContentStr  string      `json:"contentStr,omitempty"` // Content 內容
}

// MessageType 訊息種類
type MessageType int

// Sticker 貼圖
type Sticker struct {
	PackageID           string   `json:"packageId"`
	StickerID           string   `json:"stickerId"`
	StickerResourceType string   `json:"stickerResourceType"`
	Keywords            []string `json:"keywords"`
}

type Location struct {
	Address   string  `json:"address"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
