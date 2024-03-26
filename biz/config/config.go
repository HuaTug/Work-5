package config

import (
	"Hertz_refactored/biz/pkg/logging"
	"github.com/spf13/viper"
)

var ConfigInfo config

func Init() {
	viper.SetConfigType("yaml")
	viper.SetConfigName("config.yaml")
	viper.AddConfigPath("./biz/config") //这个路径为调用Init的路径为始点进行查找起点
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			logging.Fatal("config file not found: ", err)
		} else {
			logging.Fatal("config error :", err)
		}
	}
	if err := viper.Unmarshal(&ConfigInfo); err != nil {
		logging.Fatal("config decode error: ", err)
	}
}
