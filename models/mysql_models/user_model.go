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

// 实例化
var UserRepository = newUserRepository()

func newUserRepository() *userRepository {
	return &userRepository{}
}

// 生日礼包获取数据结构体
type userRepository struct {
}

func (t *User) TableName() string {
	return "gin_user"
}

// 根据手机号码查询用户信息
func (u *userRepository) GetUserInfoByPhone(phone int64, field string) (User, error) {
	user := User{}
	err := db.Select(field).Where("phone = ? AND status = ?", phone, "ENABELD").First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return user, err
	}

	return user, nil
}

// Get User Info By Id
func (u *userRepository) GetUserInfoById(id int64, field string) (User, error) {
	user := User{}
	err := db.Select(field).Where("id = ? AND status = ?", id, "ENABELD").First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return user, err
	}

	return user, nil
}
