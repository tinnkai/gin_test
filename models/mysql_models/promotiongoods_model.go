package mysql_models

import (
	"time"
)

type PromotionGoods struct {
	Id          int       `gorm:"column(id);auto"`
	PromotionId int       `gorm:"null"`
	GoodsId     int       `gorm:"column(goods_id);null"`
	Amount      float64   `gorm:"column(amount);null;digits(10);decimals(2)"`
	UpdateTime  time.Time `gorm:"column(update_time);type(datetime);null"`
	CreateTime  time.Time `gorm:"column(create_time);type(datetime);null"`
}

func (t *PromotionGoods) TableName() string {
	return "gin_promotion_goods"
}
