package controller

import (
	"fitness/dao"
	"fitness/service"
	"fitness/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func CreateUser(c *gin.Context) {

	var user dao.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusNoContent, gin.H{
			"msg": "bind json error",
		})
		return
	}
	//if len(user.Mobile) != 11 {
	//	c.JSON(http.StatusBadRequest, gin.H{
	//		"code": 1003,
	//		"msg":  "手机号必须是11位",
	//	})
	//	return
	//}
	//if len(user.PassWord) > 11 && len(user.PassWord) < 6 {
	//	c.JSON(http.StatusBadRequest, gin.H{
	//		"code": 1004,
	//		"msg":  "密码的长度在6-11位之间",
	//	})
	//	return
	//}
	if user.Type == "" {
		user.Type = "普通用户"
	}
	//if user.CertainPassword != user.PassWord {
	//	c.JSON(http.StatusBadRequest, gin.H{
	//		"code": 1008,
	//		"msg":  "两次的密码不相同",
	//	})
	//	return
	//}

	user.UserId = utils.RandInt(12)
	//判断用户是否存在
	if exists := service.IsMobileExists(user.Mobile); exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 1001,
			"msg":  "该用户已存在",
		})
		return
	}

	//密码加密
	hashpwd := utils.GetHashPwd(user.PassWord)
	hashcheckPassWord := utils.GetHashPwd(user.CertainPassword)
	user.PassWord = hashpwd
	user.CertainPassword = hashcheckPassWord
	if err := service.CreateUser(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "service error",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg":      "注册成功",
			"code":     200,
			"user_id:": user.Mobile,
		})
		return
	}
}

func RegisterUser(c *gin.Context) {

	var user dao.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusNoContent, gin.H{
			"msg": "bind json error",
		})
		return
	}
	if len(user.Mobile) != 11 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 1003,
			"msg":  "手机号必须是11位",
		})
		return
	}
	if len(user.PassWord) > 11 || len(user.PassWord) < 6 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 1004,
			"msg":  "密码的长度在6-11位之间",
		})
		return
	}
	if user.Type == "" {
		user.Type = "普通用户"
	}
	if user.CertainPassword != user.PassWord {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 1008,
			"msg":  "两次的密码不相同",
		})
		return
	}

	user.UserId = utils.RandInt(12)
	//判断用户是否存在
	if exists := service.IsMobileExists(user.Mobile); exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 1001,
			"msg":  "该用户已存在",
		})
		return
	}

	//密码加密
	hashpwd := utils.GetHashPwd(user.PassWord)
	user.PassWord = hashpwd
	hashcertainPass := utils.GetHashPwd(user.CertainPassword)
	user.CertainPassword = hashcertainPass
	if err := service.CreateUser(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "service error",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg":     "注册成功",
			"code":    200,
			"mobile:": user.Mobile,
		})
		return
	}
}

func GetUser(c *gin.Context) {
	var users dao.User
	if err := c.BindJSON(&users); err != nil {
		c.JSON(400, gin.H{
			"msg": "bind json error",
		})
		return
	}
	if service.IsMobileExists(users.Mobile) {
		user := service.GetUser(users.Mobile)
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "查询成功",
			"data": user,
		})
		return
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": "1002",
			"msg":  "用户不存在",
		})
		return
	}
}

func UserInfo(c *gin.Context) {
	//var users dao.User
	//if err := c.ShouldBindJSON(&users); err != nil {
	//	c.JSON(400, gin.H{
	//		"code": 400,
	//	})
	//}
	total, users := service.GetAllUserinfo()
	c.JSON(http.StatusOK, gin.H{
		"code":  200,
		"data":  users,
		"total": total,
	})
}

func DeleteUserByMobile(c *gin.Context) {
	var users dao.User
	if err := c.ShouldBind(&users); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "bind json error",
		})
		return
	}
	log.Println(users.Mobile)
	if service.IsMobileExists(users.Mobile) {
		if err := service.DeleteUserByMobile(users.Mobile); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 1007,
				"msg":  "delete user error",
			})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"msg":  "delete user success",
				"code": 200,
			})
			return
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 1002,
			"msg":  "用户不存在",
		})
		return
	}

}

func Update(c *gin.Context) {
	var users dao.User

	if err := c.ShouldBind(&users); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "bind json error",
		})
		return
	}
	if service.IsMobileExists(users.Mobile) {
		tx := service.Update(users.Mobile, &users)
		if tx.RowsAffected > 0 {
			c.JSON(http.StatusOK, gin.H{
				"msg":  "update user success",
				"code": 200,
			})
			return
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg":  "update user error",
				"code": 1006,
			})
			return
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 1002,
			"msg":  "用户不存在",
		})
	}

}

func UserLogin(c *gin.Context) {
	var users dao.User
	if err := c.BindJSON(&users); err != nil {
		c.JSON(http.StatusNoContent, gin.H{
			"msg": "bind json error",
		})
		return
	}
	if exists := service.IsMobileExists(users.Mobile); !exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 1002,
			"msg":  "该用户不存在",
		})
		return
	}

	user := service.UserLogin(users.Mobile)
	if utils.ComparePwd(user.PassWord, users.PassWord) {
		//if users.PassWord == user.PassWord {
		token, _ := utils.GenerateToken(users.Mobile, users.PassWord)
		c.JSON(http.StatusOK, gin.H{
			"code":  200,
			"msg":   "登录成功!!!",
			"token": token,
			"data":  user,
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

func UserLogOut(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "退出登录成功~",
	})
}

func User(c *gin.Context) {

	mobileAny := c.MustGet("mobile")
	mobile := mobileAny.(string)
	user := service.GetUser(mobile)
	c.JSON(200, gin.H{
		"data": user,
	})
	return

}

func UpdatePassword(c *gin.Context) {
	// 从请求中获取原密码、新密码、确认密码
	var req struct {
		OldPassword     string `json:"old_password"`
		NewPassword     string `json:"pass_word"`
		ConfirmPassword string `json:"certain_password"`
	}

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 验证新密码和确认密码是否一致
	if req.NewPassword != req.ConfirmPassword {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "俩次密码不相同"})
		return
	}

	// 获取当前登录用户的信息
	currentUserID, _ := c.Get("mobile")
	currentUserIDStr := fmt.Sprintf("%v", currentUserID)

	// 从数据库中获取当前登录用户的信息
	var user dao.User
	if err := dao.DB.Where("mobile = ?", currentUserIDStr).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	// 验证原密码是否正确
	if !utils.ComparePwd(user.PassWord, req.OldPassword) {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "旧密码错误"})
		return
	}

	// 更新密码
	hashedPassword := utils.GetHashPwd(req.NewPassword)
	user.PassWord = hashedPassword
	if err := dao.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "修改密码失败"})
		return
	}

	// 返回成功响应
	c.JSON(http.StatusOK, gin.H{"msg": "修改密码成功"})
}

func UploadUserAvatar(c *gin.Context) {

	token, err := c.Request.Cookie("token")
	if err != nil {
		// 处理错误
		c.JSON(401, gin.H{
			"message": "未携带有效的身份认证信息，请登录后重试",
		})
		c.Abort()
		return
	}

	log.Printf("token: %v", token.Value)

	mobileAny, exists := c.Get("mobile")
	if !exists {
		c.JSON(401, gin.H{
			"code": 401,
			"msg":  "未授权访问",
		})
		return
	}
	mobile := mobileAny.(string)

	file, err := c.FormFile("file")
	if err != nil {
		// 处理错误
		c.JSON(500, gin.H{
			"code": 500,
			"msg":  "上传文件错误",
		})
		log.Println("UploadUserAvatar error: ", err)
		return
	}
	if file == nil {
		// 处理错误，上传的文件为空
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "上传文件为空",
		})
		log.Println("UploadUserAvatar error: upload file is empty")
		return
	}

	err = utils.SaveUploadedFile(file, "./static/picture/"+file.Filename)
	if err != nil {
		// 处理错误
		c.JSON(500, gin.H{
			"code": 500,
			"msg":  "保存上传文件错误",
		})
		log.Println("UploadUserAvatar error: ", err)
		return
	}
	code, url := utils.UploadToQiNiu(file)
	if code == 0 {
		service.UploadUserAvatar(mobile, url)
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "更新成功",
			"url":  url,
		})
		return
	} else {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "更新头像失败",
		})
		log.Println("UploadUserAvatar error: upload to qiniu failed")
		return
	}
}
func EditUserAvatar(c *gin.Context) {

	var users dao.User
	if err2 := c.ShouldBind(&users); err2 != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "bind json error",
		})
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		// 处理错误
		c.JSON(500, gin.H{
			"code": 500,
			"msg":  "上传文件错误",
		})
		log.Println("UploadUserAvatar error: ", err)
		return
	}
	if file == nil {
		// 处理错误，上传的文件为空
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "上传文件为空",
		})
		log.Println("UploadUserAvatar error: upload file is empty")
		return
	}

	err = utils.SaveUploadedFile(file, "./static/picture/"+file.Filename)
	if err != nil {
		// 处理错误
		c.JSON(500, gin.H{
			"code": 500,
			"msg":  "保存上传文件错误",
		})
		log.Println("UploadUserAvatar error: ", err)
		return
	}
	code, url := utils.UploadToQiNiu(file)
	if code == 0 {
		service.UploadUserAvatar(users.Mobile, url)
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "更新成功",
			"url":  url,
		})
		return
	} else {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "更新头像失败",
		})
		log.Println("UploadUserAvatar error: upload to qiniu failed")
		return
	}
}
