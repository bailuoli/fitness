package controller

import (
	"fitness/dao"
	"fitness/service"
	"fitness/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(c *gin.Context) {

	var admin dao.Admin
	if err := c.BindJSON(&admin); err != nil {
		c.JSON(http.StatusNoContent, gin.H{
			"msg": "bind json error",
		})
		return
	}
	randInt := "storm_818"
	if admin.InvitationCode == randInt {
		if err := service.Register(&admin); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 400,
				"msg":  "注册失败",
			})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": 200,
				"msg":  "注册成功",
				"data": admin,
			})
			return
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "邀请码错误",
		})
		return
	}
}

func Login(c *gin.Context) {
	var admin dao.Admin
	if err := c.BindJSON(&admin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "bind json error",
		})
		return
	}
	admins := service.Login(admin.UserName)
	if admins.PassWord == admin.PassWord {
		token, _ := utils.GenerateToken(admin.UserName, admin.PassWord)
		c.JSON(http.StatusOK, gin.H{
			"code":  200,
			"msg":   "登陆成功",
			"token": token,
			"data":  admins,
		})
		return
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "登陆失败",
		})
		return
	}
}

func Logout(c *gin.Context) {

	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "退出成功",
	})
}
