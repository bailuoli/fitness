package dao

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	BookingId         string `gorm:"type:varchar(200)" json:"booking_id"` //订单号
	BookingUsername   string `gorm:"type:varchar(200)" json:"booking_username"`
	BookingUsermobile string `gorm:"type:varchar(200)" json:"booking_usermobile"`
	BookingContext    string `gorm:"type:varchar(200)" json:"booking_context"`
	BookingStatus     string `gorm:"type:varchar(200)" json:"booking_status"` // 0==预约中 1==预约结束
	StartTime         string `gorm:"type:varchar(200)" json:"start_time"`
	EndTime           string `gorm:"type:varchar(200)" json:"end_time"`
	UserId            string `gorm:"type:varchar(200)" json:"user_id"`
	Area
}
