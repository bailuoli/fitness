package dao

//场地表

type Area struct {
	ID        uint   `gorm:"primarykey"`
	AreaId    string `gorm:"type:varchar(200)" json:"area_id"`    //场地id
	AreaName  string `gorm:"type:varchar(200)" json:"area_name"`  // 场地名称
	AreaLocal string `gorm:"type:varchar(200)" json:"area_local"` //场地位置
	AreaDesc  string `gorm:"type:varchar(200)" json:"area_desc"`  // 场地描述
	State     string `gorm:"type:varchar(200)" json:"state"`      //场地状态 0== 空闲   1== 已经被预约 2==使用中
	AreaFee   string `gorm:"type:varchar(200)" json:"area_fee"`   //场地费
	Url       string `gorm:"type:varchar(200)" json:"url"`
	Type      string `gorm:"type:varchar(200)" json:"type"`
}
