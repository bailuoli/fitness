package dao

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func InitDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("datasource.username"),
		viper.GetString("datasource.password"),
		viper.GetString("datasource.host"),
		viper.GetString("datasource.port"),
		viper.GetString("datasource.database"))
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic("failed to connect database")
	}

	err = DB.AutoMigrate(&User{}, &Admin{}, &Coach{}, &Area{}, &Order{})
	if err != nil {
		log.Panic("failed to auto migrate database")
	}
}

// GetDb 获取Db操作数据库

func GetDB() *gorm.DB {
	return DB
}
