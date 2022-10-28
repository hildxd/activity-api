package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	C = new(Config)
)

// load config file
func MustLoad() {
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath("./configs")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("读取配置文件失败: %v", err)
	}
	err := viper.Unmarshal(&C)
	if err != nil {
		log.Fatal("viper 解析配置文件失败")
	}
}

type Config struct {
	Port  int
	Name  string
	MySql string
}
