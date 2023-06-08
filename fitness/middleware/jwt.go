package middleware

import (
	"fitness/utils"
	"github.com/gin-gonic/gin"
	"log"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Request.Cookie("token")
		if err != nil {
			c.JSON(500, gin.H{
				"status": -1,
				"msg":    "请求未携带token，无权限访问",
			})
			c.Abort()
			return
		}

		log.Print("get token: ", token)

		// parseToken 解析token包含的信息
		claims, err := utils.ParseToken(token.Value)
		if err != nil {
			if err == utils.TokenExpired {
				c.JSON(500, gin.H{
					"status": -1,
					"msg":    "授权已过期",
				})
				c.Abort()
				return
			}
			c.JSON(500, gin.H{
				"status": -1,
				"msg":    "token无效",
			})
			c.Abort()
			return
		}

		if claims.Mobile == "" {
			c.JSON(500, gin.H{
				"status": -1,
				"msg":    "请求未携带token，无权限访问",
			})
			c.Abort()
			return
		}

		// 继续交由下一个路由处理,并将解析出的信息传递下去
		c.Set("mobile", claims.Mobile)
	}
}
