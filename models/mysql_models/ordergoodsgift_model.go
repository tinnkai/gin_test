package mysql_models

import (
	"time"
)

type OrderGoodsGift struct {
	Id             int       `orm:"column(id);auto"`
	OrderId        uint      `orm:"column(order_id)"`
	GoodsId        int       `orm:"column(goods_id);null"`
	GiftGoodsId    int       `orm:"column(gift_goods_id);null"`
	Num            int       `orm:"column(num);null"`
	Amount         float64   `orm:"column(amount);null;digits(10);decimals(2)"`
	DeliveryStatus string    `orm:"column(delivery_status);size(50);null"`
	UpdateTime     time.Time `orm:"column(update_time);type(datetime);null"`
	CreateTime     time.Time `orm:"column(create_time);type(datetime);null"`
}

// save order goods gift
func SaveOrderGoodsGift(o []OrderGoodsGift) error {
	for _, v := range o {
		err := db.Create(&v).Error
		if err != nil {
			return err
		}
	}
	return nil
}
