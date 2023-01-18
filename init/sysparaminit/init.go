package sysparaminit

import (
	"linebot/global"

	"github.com/spf13/viper"
)

type Setting struct {
	vp *viper.Viper
}

const (
	serverSection   string = "Server"
	databasSectione string = "Database"
	linebotSection  string = "Linebot"
)

// NewSetting 新增設定
func NewSetting() (err error) {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.AddConfigPath("config/")
	vp.SetConfigType("yaml")
	if err = vp.ReadInConfig(); err != nil {
		return err
	}
	var tmpSetting = &Setting{vp}

	if err = tmpSetting.readSection(serverSection, &global.ServerSetting); err != nil {
		return err
	}
	if err = tmpSetting.readSection(databasSectione, &global.DatabaseSettings); err != nil {
		return err
	}
	if err = tmpSetting.readSection(linebotSection, &global.LinebotSettings); err != nil {
		return err
	}
	return nil
}

// readSection 讀取設定值(section)
func (s *Setting) readSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	return nil
}
