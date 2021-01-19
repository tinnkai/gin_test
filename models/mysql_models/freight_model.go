package mysql_models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Freight struct {
	Id         int       `gorm:"column(id);auto"`
	CityId     int       `gorm:"column(city_id);null"`
	Amount     float64   `gorm:"column(amount);null;digits(5);decimals(2)"`
	UpdateTime time.Time `gorm:"column(update_time);type(datetime);null"`
	CreateTime time.Time `gorm:"column(create_time);type(datetime);null"`
}

type OrderFreight struct {
	CityId int     `gorm:"column(city_id);null"`
	Amount float64 `gorm:"column(amount);null;digits(5);decimals(2)"`
}

func (t *Freight) TableName() string {
	return "gin_freight"
}

// 根据城市id查询运费
func GetFreightInfoByCityId(cityId int) (OrderFreight, error) {
	var freight Freight
	var orderFreight OrderFreight
	err := db.Table(freight.TableName()).Where("city_id = ?", cityId).Scan(&orderFreight).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return orderFreight, err
	}
	return orderFreight, nil
}
