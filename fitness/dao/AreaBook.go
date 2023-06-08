package dao

import (
	"gorm.io/gorm"
)

// 场地预约表

type AreaBook struct {
	gorm.Model
	BookingId         string `gorm:"type:varchar(200)" json:"booking_id"` //订单id
	BookingUsername   string `gorm:"type:varchar(200)" json:"booking_username"`
	BookingUsermobile string `gorm:"type:varchar(200)" json:"booking_usermobile"`
	BookingContext    string `gorm:"type:varchar(200)" json:"booking_context"`
	BookingStatus     string `gorm:"type:varchar(200)" json:"booking_status"` // 0==预约中 1==预约结束
	StartTime         string `gorm:"type:varchar(200)" json:"start_time"`
	EndTime           string `gorm:"type:varchar(200)" json:"end_time"`
}

type AreaBookRes struct {
	OrderId string `gorm:"type:varchar(200)" json:"order_id"`
	Area
	AreaBook
}
