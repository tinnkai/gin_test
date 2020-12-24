package mysql_models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	Id         int64     `gorm:"primary_key" json:"id"`
	Username   string    `gorm:"column(username);size(50);json:"username"`
	Password   string    `gorm:"column(password);size(50);json:"-"`
	Phone      int64     `gorm:"column(phone);json:"phone"`
	Group      int       `gorm:"column(group);json:"group"`
	Status     string    `gorm:"column(status);size(10);json:"status"`
	UpdateTime time.Time `gorm:"column(update_time);type(datetime);null"`
	CreateTime time.Time `gorm:"column(create_time);type(datetime);null"`
}

// 根据手机号码查询用户信息
func (u *User) GetUserInfoByPhone(field string) error {
	err := db.Select(field).Where("phone = ? AND status = ?", u.Phone, "ENABELD").First(u).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	return nil
}

// Get User Info By Id
func (u *User) GetUserInfoById(field string) error {
	var user User
	err := db.Select(field).Where("id = ? AND status = ?", u.Id, "ENABELD").First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	return nil
}
