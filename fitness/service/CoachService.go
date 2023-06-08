package service

import (
	"errors"
	"fitness/dao"

	"gorm.io/gorm"
)

// 判断教练是否存在

func CoachIsMobileExists(mobile string) bool {
	var coach dao.Coach
	dao.DB.Where("mobile = ?", mobile).First(&coach)
	return coach.ID != 0
}

// 新建教练

func CreateCoach(coach *dao.Coach) (err error) {
	if err = dao.DB.Create(&coach).Error; err != nil {
		return errors.New("创建教练失败")
	}
	return nil
}

// 删除教练

func DeleteCoachById(id string) (err error) {
	if err := dao.DB.Unscoped().Where("id = ?", id).Delete(&dao.Coach{}).Error; err != nil {
		return errors.New("删除教练失败")
	}
	return
}

func DeleteCoachByMobile(mobile string) (err error) {
	if err := dao.DB.Unscoped().Where("mobile = ?", mobile).Delete(&dao.Coach{}).Error; err != nil {
		return errors.New("删除教练失败")
	}
	return
}

// 查询教练

func GetCoachById(id string) (coach *dao.Coach, err error) {
	if err = dao.DB.Where("id = ?", id).First(&coach).Error; err != nil {
		return nil, errors.New("查询此教练失败")
	}
	return
}
func GetCoachByMobile(mobile string) (coach *dao.Coach, err error) {
	if err = dao.DB.Where("mobile = ?", mobile).Find(&coach).Error; err != nil {
		return nil, errors.New("查询此教练失败")
	}
	return
}

// 查询教练列表
func GetAllCoachinfo() (total int, coachs []dao.Coach) {
	// 计算偏移量

	// 查询所有的user
	result := dao.DB.Find(&coachs)
	// 查不到数据时
	if result.RowsAffected == 0 {
		return 0, nil
	}
	// 获取user总数
	total = len(coachs)
	// 查询数据
	result.Find(&coachs)
	return total, coachs
}

func GetAllCoach(pageNum int, pageSize int) (total int, coach []dao.Coach) {
	// 计算偏移量
	offset := (pageNum - 1) * pageSize
	// 查询所有的user
	result := dao.DB.Offset(offset).Limit(pageSize).Find(&coach)
	// 查不到数据时
	if result.RowsAffected == 0 {
		return 0, nil
	}
	// 获取user总数
	total = len(coach)
	// 查询数据
	result.Offset(offset).Limit(pageSize).Find(&coach)
	return total, coach
}

// 更新教练信息
func UpdateCoach(mobile string, coach *dao.Coach) (tx *gorm.DB) {
	tx = dao.DB.Model(coach).Where("mobile = ?", mobile).Updates(coach)
	return tx
}

func CoachLogin(mobile string, password string) (coach dao.Coach) {
	dao.DB.Where("mobile = ?", mobile).First(&coach)
	return coach
}
