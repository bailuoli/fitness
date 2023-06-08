package dao

import "gorm.io/gorm"

// 教练表

type Coach struct {
	gorm.Model
	CoachId         string `gorm:"type:varchar(200)" json:"coach_id"`
	CoachName       string `gorm:"type:varchar(200)" json:"coach_name"`
	PassWord        string `gorm:"type:varchar(200)" json:"pass_word"`
	CertainPassword string `gorm:"type:varchar(200)" json:"certain_password"`
	Url             string `gorm:"type:varchar(200)" json:"url"` //头像
	WeChat          string `gorm:"type:varchar(200)" json:"we_chat"`
	Gender          string `gorm:"type:varchar(200)" json:"gender"` //性别
	Mobile          string `gorm:"type:varchar(200)" json:"mobile"`
	Type            string `gorm:"type:varchar(200)" json:"type"` //执教类型
	CoachAge        string `gorm:"type:varchar(200)" json:"coach_age"`
	State           string `gorm:"type:varchar(200)" json:"state"`      // 0==空闲 1==培训中
	CoachFee        string `gorm:"type:varchar(200)" json:"coach_fee"`  //教练费用
	CoachDesc       string `gorm:"type:varchar(200)" json:"coach_desc"` //教练简介
}
