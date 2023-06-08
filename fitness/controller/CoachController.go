package controller

import (
	"fitness/dao"
	"fitness/service"
	"fitness/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 创建教练
func CreateCoach(c *gin.Context) {
	var coach dao.Coach
	if err := c.ShouldBind(&coach); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "bind json error",
		})
		return
	}
	if len(coach.Mobile) != 11 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 1003,
			"msg":  "手机号必须是11位",
		})
		return
	}
	if len(coach.PassWord) > 11 && len(coach.PassWord) < 6 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 1004,
			"msg":  "密码必须在6-11位之间",
		})
		return
	}
	//if coach.CertainPassword != coach.PassWord {
	//	c.JSON(http.StatusBadRequest, gin.H{
	//		"code": 1008,
	//		"msg":  "两次密码不一样",
	//	})
	//	return
	//}

	coach.CoachId = utils.RandInt(12)
	//判断用户是否存在
	if exists := service.CoachIsMobileExists(coach.Mobile); exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 1001,
			"msg":  "该用户已存在",
		})
		return
	}

	//密码加密
	hashpwd := utils.GetHashPwd(coach.PassWord)
	hashcheckPassWord := utils.GetHashPwd(coach.CertainPassword)
	coach.PassWord = string(hashpwd)
	coach.CertainPassword = string(hashcheckPassWord)
	if err := service.CreateCoach(&coach); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "service error",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg":      "create user success",
			"code":     200,
			"user_id:": coach.Mobile,
		})
		return
	}
}

// 删除教练
func DeleteCoachById(c *gin.Context) {
	var coach dao.Coach
	if err := c.ShouldBindJSON(&coach); err != nil {
		c.JSON(http.StatusNoContent, gin.H{
			"msg": "bind json error",
		})
		return
	}
	if err := service.DeleteCoachById(strconv.FormatUint(uint64(coach.ID), 10)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "delete coach error",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "delete coach success",
			"code": 200,
		})
		return
	}
}
func DeleteCoachByMobile(c *gin.Context) {
	var coach dao.Coach
	if err := c.BindJSON(&coach); err != nil {
		c.JSON(http.StatusNoContent, gin.H{
			"msg": "bind json error",
		})
		return
	}
	if err := service.DeleteCoachByMobile(coach.Mobile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "delete coach error",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "delete coach success",
			"code": 200,
		})
		return
	}
}

// UpdateCoach 编辑教练信息
func UpdateCoach(c *gin.Context) {
	var coachs dao.Coach
	if err := c.ShouldBind(&coachs); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "bind json error",
		})
		return
	}

	log.Println("111111111111111111+", coachs.Mobile)
	tx := service.UpdateCoach(coachs.Mobile, &coachs)
	if tx.RowsAffected > 0 {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "更新教练成功",
			"code": 200,
		})
		return
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "更新教练失败",
			"code": 1006,
		})
		return
	}
}

// 查询教练
func GetCoachById(c *gin.Context) {
	var coach dao.Coach
	if err := c.BindJSON(&coach); err != nil {
		c.JSON(http.StatusNoContent, gin.H{
			"msg": "bind json error",
		})
		return
	}
	todoList, err := service.GetCoachById(strconv.FormatUint(uint64(coach.ID), 10))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "service error",
		})
		return
	} else if todoList.Mobile == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "The coach does not exist",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  "get coach success",
		"code": 200,
		"data": todoList,
	})
}
func GetCoachByMobile(c *gin.Context) {
	var coach dao.Coach
	if err := c.ShouldBindJSON(&coach); err != nil {
		c.JSON(http.StatusNoContent, gin.H{
			"msg": "bind json error",
		})
		return
	}
	println("教练手机号：", coach.Mobile)
	if todoList, err := service.GetCoachByMobile(coach.Mobile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "service error",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "get coach success",
			"code": 200,
			"data": todoList,
		})
		return
	}
}
func GetAllCoach(c *gin.Context) {
	pagenum, _ := strconv.Atoi(c.Query("pagenum"))
	pagesize, _ := strconv.Atoi(c.Query("pagesize"))
	total, user := service.GetAllUser(pagenum, pagesize)
	if total == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 1010,
			"msg":  "未查询到数据",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "查询成功",
			"data": user,
		})
	}
}

// 列表
func CoachInfo(c *gin.Context) {
	//var users dao.User
	//if err := c.ShouldBindJSON(&users); err != nil {
	//	c.JSON(400, gin.H{
	//		"code": 400,
	//	})
	//}
	total, coachs := service.GetAllCoachinfo()
	c.JSON(http.StatusOK, gin.H{
		"code":  200,
		"data":  coachs,
		"total": total,
	})
}

// 登录
func CoachLogin(c *gin.Context) {
	var coachs dao.Coach
	if err := c.BindJSON(&coachs); err != nil {
		c.JSON(http.StatusNoContent, gin.H{
			"msg": "bind json error",
		})
		return
	}
	if exists := service.CoachIsMobileExists(coachs.Mobile); !exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 1002,
			"msg":  "该用户不存在",
		})
		return
	}

	coach := service.CoachLogin(coachs.Mobile, coachs.PassWord)
	if utils.ComparePwd(coach.PassWord, coachs.PassWord) {
		token, _ := utils.GenerateToken(coachs.Mobile, coachs.PassWord)
		c.JSON(http.StatusOK, gin.H{
			"code":  200,
			"msg":   "登录成功!!!",
			"token": token,
		})
		return
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 1005,
			"msg":  "密码错误",
		})
		return
	}
}
