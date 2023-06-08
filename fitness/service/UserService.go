package service

import (
	"errors"
	"fitness/dao"
	"gorm.io/gorm"
)

// 判断用户是否存在
func IsMobileExists(mobile string) bool {
	tx := dao.DB.Where("mobile = ?", mobile).First(&dao.User{})
	if tx.RowsAffected != 0 {
		return true
	}
	return false
}

// 新建用户
func CreateUser(user *dao.User) (err error) {
	if err = dao.DB.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

// 查询用户列表
func GetAllUser(pageNum int, pageSize int) (total int, users []dao.User) {
	// 计算偏移量
	offset := (pageNum - 1) * pageSize
	// 查询所有的user
	result := dao.DB.Offset(offset).Limit(pageSize).Find(&users)
	// 查不到数据时
	if result.RowsAffected == 0 {
		return 0, nil
	}
	// 获取user总数
	total = len(users)
	// 查询数据
	result.Offset(offset).Limit(pageSize).Find(&users)
	return total, users
}

// 删除用户
func DeleteUserById(id string) (err error) {
	if err = dao.DB.Unscoped().Where("id = ?", id).Delete(&dao.User{}).Error; err != nil {
		return errors.New("删除用户失败")
	}
	return nil
}
func DeleteUserByMobile(mobile string) (err error) {
	if err := dao.DB.Unscoped().Where("mobile = ?", mobile).Delete(&dao.User{}).Error; err != nil {
		return errors.New("删除用户失败")
	}
	return
}
func DeleteUserBy(user_id string) (err error) {
	if err := dao.DB.Raw("DELETE FROM users where user_id = ?", user_id).Error; err != nil {
		return err
	}
	return nil
}

// 查询单个用户
func GetUser(mobile string) (user dao.User) {
	dao.DB.Where("mobile = ?", mobile).Find(&user)
	return user
}

func GetAllUserinfo() (total int, users []dao.User) {
	// 计算偏移量

	// 查询所有的user
	result := dao.DB.Find(&users)
	// 查不到数据时
	if result.RowsAffected == 0 {
		return 0, nil
	}
	// 获取user总数
	total = len(users)
	// 查询数据
	result.Find(&users)
	return total, users
}

// 更新用户信息
func UpdateUser(id string, user *dao.User) (tx *gorm.DB) {
	tx = dao.DB.Where("id = ?", id).Updates(&user)
	return tx
}
func Update(mobile string, user *dao.User) (tx *gorm.DB) {
	tx = dao.DB.Model(user).Where("mobile = ?", mobile).Updates(user)
	return tx
}

// 用户登录
func UserLogin(mobile string) (user *dao.User) {
	dao.DB.Where("mobile = ?", mobile).First(&user)
	return user
}

// 查询密码
func UserPassWord(mobile string) (password string) {
	var user dao.User
	dao.DB.Where("mobile = ?", mobile).First(&user)
	return user.PassWord
}

// 修改密码
func ChangePassword(mobile string, data *dao.User) *gorm.DB {
	tx := dao.DB.Select("pass_word").Where("mobile = ?", mobile).Updates(&data)
	return tx
}

// 更新头像
func UploadUserAvatar(mobile string, avatar string) {
	dao.DB.Model(&dao.User{}).Where("mobile = ?", mobile).Update("avatar", avatar)
}
