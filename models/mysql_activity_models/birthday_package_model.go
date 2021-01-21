package mysql_activity_models

import (
	"gin_test/pkg/utils"
	"time"
)

// 生日礼包模型结构体
type BirthdayPackage struct {
	Id         int                  `gorm:"column(id);auto" json:"id"`
	Title      string               `gorm:"column(title);size(100);null"  json:"title"`
	StartTime  utils.CustomDatetime `gorm:"column(start_time);type(datetime);null" json:"start_time"`
	EndTime    utils.CustomDatetime `gorm:"column(end_time);type(datetime);null" json:"end_time"`
	UpdateTime time.Time            `gorm:"column(update_time);type(datetime);null"`
	CreateTime time.Time            `gorm:"column(create_time);type(datetime);null"`
}

// 实例化
var BirthdayPackageRepository = newBirthdayPackageRepository()

func newBirthdayPackageRepository() *birthdayPackageRepository {
	return &birthdayPackageRepository{}
}

// 生日礼包获取数据结构体
type birthdayPackageRepository struct {
}

func (t *BirthdayPackage) TableName() string {
	return "birthday_package"
}

// 获取符合时间的活动
func (t *birthdayPackageRepository) GetInfoByTime() (BirthdayPackage, error) {
	nowDateTime := utils.NowDateTime()
	birthdayPackageModel := BirthdayPackage{}
	err := db.Where("start_time <= ? AND end_time >= ?", nowDateTime, nowDateTime).
		Table(birthdayPackageModel.TableName()).
		First(&birthdayPackageModel).Error

	return birthdayPackageModel, err
}

// 不限定where条件,返回一条数据
func (this *birthdayPackageRepository) GetOneByWhere(where ...interface{}) (BirthdayPackage, error) {
	ret := BirthdayPackage{}
	if err := db.Take(&ret, where...).Error; err != nil {
		return ret, err
	}
	return ret, nil
}
