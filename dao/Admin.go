package dao

// 管理员表
type Admin struct {
	ID             uint   `gorm:"primarykey"`
	UserName       string `gorm:"type:varchar(200)" json:"user_name"`
	PassWord       string `gorm:"type:varchar(200)" json:"pass_word"`
	InvitationCode string `gorm:"type:varchar(200)" json:"invitation_code"`
}
