package service

import (
	"fitness/dao"
)

// CheckArea 查看场地是否预约
func CheckArea(areaid string) *dao.Area {
	dao.DB.Where("area_id = ?", areaid).First(&dao.Area{})
	return &dao.Area{}
}

// CheckAreaBookExist 查看预约是否存在
func CheckAreaBookExist(id uint) bool {
	if err := dao.DB.Where("id = ?", id).First(&dao.AreaBook{}).Error; err != nil {
		return false
	}
	return true

}

// AddAreaBooking 添加预约
func AddAreaBooking(areabookres *dao.AreaBookRes) (err error) {
	dao.DB.Create(&dao.AreaBook{
		BookingId:         areabookres.BookingId,
		BookingUsername:   areabookres.BookingUsername,
		BookingUsermobile: areabookres.BookingUsermobile,
		BookingContext:    areabookres.BookingContext,
		BookingStatus:     areabookres.BookingStatus,
		StartTime:         areabookres.StartTime,
		EndTime:           areabookres.EndTime,
	})
	dao.DB.Create(&areabookres)
	return nil
}

// UpdateAreaBookState 修改场地预约状态
func UpdateAreaBookState(BookingId string, state int) {
	dao.DB.Model(&dao.AreaBook{}).Where("booking_id = ?", BookingId).Update("booking_status", state)
}

func UpdateAreaBookResBookState(BookingId string, state int) {
	dao.DB.Model(&dao.AreaBookRes{}).Where("booking_id = ?", BookingId).Update("booking_status", state)
}
func UpdateAreaBookResState(BookingId string, state int) {
	dao.DB.Model(&dao.AreaBookRes{}).Where("booking_id = ?", BookingId).Update("state", state)
}

// GetAllAreaBooks 查询场地预约列表
func GetAreaBookList() (total int, areabooks []dao.AreaBook) {
	// 计算偏移量

	// 查询所有的user
	result := dao.DB.Find(&areabooks)
	// 查不到数据时
	if result.RowsAffected == 0 {
		return 0, nil
	}
	// 获取user总数
	total = len(areabooks)
	// 查询数据
	result.Find(&areabooks)
	return total, areabooks
}

// GetAreaBookById 根据id查询场地预约
func GetAreaBookById(id string) (areabook *dao.AreaBook, err error) {
	if err = dao.DB.Where("id = ?", id).First(&areabook).Error; err != nil {
		return nil, err
	}
	return
}

func GetAreaBookByBookId(id string) (areabook *dao.AreaBook, err error) {
	if err = dao.DB.Where("booking_id = ?", id).First(&areabook).Error; err != nil {
		return nil, err
	}
	return
}

// DeleteAreaBookById 根据id删除预约
func DeleteAreaBookById(id string) bool {
	if err := dao.DB.Unscoped().Where("id = ?", id).Delete(&dao.AreaBook{}).Error; err != nil {
		return false
	}
	return true
}
