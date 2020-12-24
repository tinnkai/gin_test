package mysql_models

import (
	"time"
)

type Address struct {
	Id         int       `gorm:"column(id);auto"`
	UserId     int64     `gorm:"column(user_id);null"`
	AreasId    int       `gorm:"column(areas_id);null"`
	Detail     string    `gorm:"column(detail);size(255);null"`
	UpdateTime time.Time `gorm:"column(update_time);type(datetime);null"`
	CreateTime time.Time `gorm:"column(create_time);type(datetime);null"`
}

func (t *Address) TableName() string {
	return "gin_address"
}

// 根据Id查询收货地址获取省市区信息
func GetAddressInfoById(id int, userId int64) (*Address, error) {
	var address Address
	err := db.Where("id= ? AND user_id = ?", id, userId).
		First(&address).Error

	return &address, err
}
