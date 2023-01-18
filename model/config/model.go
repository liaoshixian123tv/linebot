package config

// ServerSettings 伺服器設定
type ServerSettings struct {
	RunMode  string
	HttpPort string
}

// DatabaseSettings 資料庫設定
type DatabaseSettings struct {
	Username string
	Password string
	Host     string
	Port     string
	DBName   string
}

// LinebotSettings Linebot設定
type LinebotSettings struct {
	ChannelScrect      string
	ChannelAccessToken string
}
