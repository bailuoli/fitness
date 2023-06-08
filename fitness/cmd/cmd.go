package cmd

import (
	"fitness/conf"
	"fitness/dao"
	"fitness/log/logrus"
	"fitness/routes"
	"fmt"
	logr "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func Start() {
	logrus.InitLog()
	//初始化配置文件
	conf.InitConfig()
	//初始化数据库
	dao.InitDB()
	//初始化router
	r := routes.SetRouter()
	if port := viper.GetString("server.port"); port != "" {
		logr.Panic(r.Run(":" + port))
	} else {
		logr.Panic(r.Run())
	}
}

func Clean() {
	fmt.Println("=====clean=====")
}
