package service

import (
	"errors"
	"fitness/dao"
	"gorm.io/gorm"
)

// 修改订单
func UpdateOrder(order *dao.Order) (tx *gorm.DB) {
	tx = dao.DB.Model(&dao.Order{}).Where("booking_id = ?", order.BookingId).Updates(order)
	return tx
}

// 获取订单列表
func GetOrderList() (total int, order []dao.Order) {
	result := dao.DB.Find(&order)
	// 查不到数据时
	if result.RowsAffected == 0 {
		return 0, nil
	}
	// 获取user总数
	total = len(order)
	// 查询数据
	result.Find(&order)
	return total, order
}

func CheckOrder(order *dao.Order) (ok bool, orders []dao.Order) {
	var count int64
	dao.DB.Model(&order).
		Where("area_id = ? AND ((start_time <= ? AND end_time >= ?) OR (start_time <= ? AND end_time >= ?))",
			order.AreaId, order.StartTime, order.EndTime, order.StartTime, order.EndTime).Count(&count).Scan(&orders)
	if count > 0 {
		for i := 0; i < len(orders); i++ {
			if orders[i].BookingStatus != "预约结束" {
				return false, orders
			}
			return true, nil
		}
	}
	return true, nil
}

// 获取单个订单
func GetOrder(bookingId string) (order dao.Order) {
	dao.DB.Where("booking_id = ?", bookingId).First(&order)
	return
}
func GetOrderByAreaId(areaId string) (order []dao.Order) {
	dao.DB.Where("area_id = ?", areaId).Find(&order)
	return
}

func GetOrderByMobile(mobile string) (total int, order []dao.Order) {
	result := dao.DB.Where("booking_usermobile = ?", mobile).Find(&order)
	// 查不到数据时
	if result.RowsAffected == 0 {
		return 0, nil
	}
	// 获取user总数
	total = len(order)
	// 查询数据
	result.Find(&order)
	return total, order
}

// 根据userid查询
func GetOrderByUserId(UserId string) (total int, order []dao.Order) {
	result := dao.DB.Where("user_id = ?", UserId).Find(&order)
	// 查不到数据时
	if result.RowsAffected == 0 {
		return 0, nil
	}
	// 获取user总数
	total = len(order)
	// 查询数据
	result.Find(&order)
	return total, order
}

// 更新订单状态
func UpdateOrderStatus(BookingId string, state string) {
	dao.DB.Model(&dao.Order{}).Where("booking_id = ?", BookingId).Update("booking_status", state)
}
func UpdateOrderState(BookingId string, state string) {
	dao.DB.Model(&dao.Order{}).Where("booking_id = ?", BookingId).Update("state", state)
}

// 添加订单
// 添加预约
func AddOrder(order *dao.Order) (err error) {
	dao.DB.Create(&order)
	return nil
}

// 删除订单
func DeleteOrder(bookingId string) error {
	if err := dao.DB.Unscoped().Where("booking_id = ?", bookingId).Delete(&dao.Order{}).Error; err != nil {
		return errors.New("删除订单失败")
	}
	return nil
}
