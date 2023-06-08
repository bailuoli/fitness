package service

import (
	"errors"
	"fitness/dao"
)

// IsExistsArea 判断场地是否存在
func IsExistsArea(area_id string) bool {
	var area dao.Area
	dao.DB.Where("area_id = ?", area_id).First(&area)
	return area.ID != 0
}

// CreateArea 增加场地
func CreateArea(area *dao.Area) (err error) {
	if err := dao.DB.Create(&area).Error; err != nil {
		return err
	}
	return nil
}

// DeleteAreaById 删除场地根据id
func DeleteArea(area_id string) (err error) {
	if err := dao.DB.Unscoped().Where("area_id = ?", area_id).Delete(&dao.Area{}).Error; err != nil {
		return errors.New("del area dao error")
	}
	return nil
}

// UpdateArea 修改场地信息
func UpdateArea(area *dao.Area) error {
	err := dao.DB.Where("area_id = ?", area.AreaId).Updates(&area).Error
	if err != nil {
		return errors.New("更新场地信息失败")
	}
	return nil
}

// UpdateAreaState 修改场地状态
func UpdateAreaState(AreaId string, state string) {
	dao.DB.Model(&dao.Area{}).Where("area_id = ?", AreaId).Update("state", state)
}

// GetAllAreas 查询场地列表
func GetAreaList() (total int, area []dao.Area) {
	// 计算偏移量

	// 查询所有的user
	result := dao.DB.Find(&area)
	// 查不到数据时
	if result.RowsAffected == 0 {
		return 0, nil
	}
	// 获取user总数
	total = len(area)
	// 查询数据
	result.Find(&area)
	return total, area
}

// GetAreaById 根据id查询场地
func GetAreaById(id string) (area *dao.Area, err error) {
	if err = dao.DB.Where("id = ?", id).First(&area).Error; err != nil {
		return nil, err
	}
	return
}
func GetAreaByAreaId(areaid string) (area *dao.Area, err error) {
	if err = dao.DB.Where("area_id = ?", areaid).First(&area).Error; err != nil {
		return
	}
	return
}

func GetArea(Type string) (areas []dao.Area, err error) {
	dao.DB.Raw("select * from areas where type = ?", Type).Scan(&areas)
	return areas, err
}
