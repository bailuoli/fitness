package errorHandle

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func AppErrorHandleMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 捕获panic错误
		defer func() {
			if err := recover(); err != nil {
				// 打印出错信息
				logrus.Infof("Panic caught: %v\n", err)
				// 返回响应
				c.JSON(500, gin.H{
					"error": "内部错误",
				})
			}
		}()
		// 执行后续的中间件和处理器
		c.Next()
	}
}
