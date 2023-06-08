package controller

import (
	"context"
	"fitness/utils"
	"github.com/gin-gonic/gin"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
)

func UpLoad(c *gin.Context) {
	var AccessKey = viper.GetString("AccessKey")
	var SecretKey = viper.GetString("SecretKey")
	var Bucket = viper.GetString("Bucket")
	var ImgUrl = viper.GetString("QiniuServer")
	var code = 200
	file, err := c.FormFile("file")
	if err != nil {
		// 处理错误
		logrus.Error(err)
		panic(err)
		return
	}
	if file == nil {
		// 处理错误，上传的文件为空
		logrus.Error("上传的文件为空")
		return
	}

	err = utils.SaveUploadedFile(file, "./static/picture/"+file.Filename)
	if err != nil {
		logrus.Panic("下载图片失败")
		return
	}

	src, err := file.Open()
	if err != nil {
		return
	}

	putPlicy := storage.PutPolicy{
		Scope: Bucket,
	}
	mac := qbox.NewMac(AccessKey, SecretKey)

	// 获取上传凭证
	upToken := putPlicy.UploadToken(mac)

	// 配置参数
	cfg := storage.Config{
		Zone:          &storage.ZoneHuadongZheJiang2,
		UseCdnDomains: false,
		UseHTTPS:      false, // 非https
	}
	formUploader := storage.NewFormUploader(&cfg)

	ret := storage.PutRet{}        // 上传后返回的结果
	putExtra := storage.PutExtra{} // 额外参数

	// 上传 自定义key，可以指定上传目录及文件名和后缀，
	key := "img/" + file.Filename // 上传路径，如果当前目录中已存在相同文件，则返回上传失败错误
	err = formUploader.Put(context.Background(), &ret, upToken, key, src, file.Size, &putExtra)
	if err != nil {
		code = 501
		return
	}

	url := ImgUrl + ret.Key // 返回上传后的文件访问路径
	logrus.Info(ret)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"url":  url,
	})

}
