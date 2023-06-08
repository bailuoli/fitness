package conf

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

func InitConfig() {
	dir, err := os.Getwd()
	if err != nil {
		log.Panic("读取根目录失败")
	}
	viper.SetConfigName("settings")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(dir + "/conf")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Panic("读取配置文件失败")
		} else {
			log.Println("读取配置文件成功！")
		}
	}
}
