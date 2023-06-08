package utils

//比较预约时间大小
import (
	"time"

	"github.com/spf13/viper"
)

func ComPareTime(timeStr string) bool {
	TimeLocation, _ := time.LoadLocation("Asia/Shanghai")
	now, _ := time.ParseInLocation(viper.GetString("timeLayout"), time.Now().Format(viper.GetString("timeLayout")), TimeLocation)
	a, _ := time.ParseInLocation(viper.GetString("timeLayout"), timeStr, time.Local)
	return now.After(a)
}

func ComPareTimeWith(str1, str2 string) bool {
	TimeLocation, _ := time.LoadLocation("Asia/Shanghai")
	strres1, _ := time.ParseInLocation(viper.GetString("timeLayout"), str1, TimeLocation)
	strres2, _ := time.ParseInLocation(viper.GetString("timeLayout"), str2, TimeLocation)
	return strres1.After(strres2)
}
