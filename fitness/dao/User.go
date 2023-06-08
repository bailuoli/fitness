package dao

import (
	"fitness/utils"
	"gorm.io/gorm"
)

// 用户表
type User struct {
	gorm.Model
	UserId          string `gorm:"type:varchar(200)" json:"user_id"`   //userid
	NickName        string `gorm:"type:varchar(200)" json:"nick_name"` //用户名
	PassWord        string `gorm:"type:varchar(200)" json:"pass_word" validate:"required,max=11,min=6"`
	CertainPassword string `gorm:"type:varchar(200)" json:"certain_password" validate:"required,max=11,min=6"`
	Gender          string `gorm:"type:varchar(200)" json:"gender"`
	Mobile          string `gorm:"type:varchar(200)" json:"mobile"`
	WeChat          string `gorm:"type:varchar(200)" json:"we_chat"`
	SportType       string `gorm:"type:varchar(200)" json:"sport_type"`
	Type            string `gorm:"type:varchar(200)" json:"type"`
	Avatar          string `gorm:"type:varchar(200)" json:"avatar"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {

	if u.Type == "" {
		u.Type = "普通用户"
	}
	u.UserId = utils.RandInt(12)
	return
}
