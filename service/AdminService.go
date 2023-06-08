package service

import "fitness/dao"

func Register(admin *dao.Admin) (err error) {
	if err = dao.DB.Create(&admin).Error; err != nil {
		return err
	}
	return nil
}

func Login(user_name string) (admin *dao.Admin) {
	dao.DB.Where("user_name = ?", user_name).First(&admin)
	return admin
}
