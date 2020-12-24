package mysql_models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type OrderGoods struct {
	Id             int       `gorm:"column(id);auto"`
	OrderId        uint      `gorm:"column(order_id)"`
	UserId         int64     `gorm:"column(user_id);null"`
	GoodsId        int       `gorm:"column(goods_id);null"`
	GoodsType      string    `gorm:"column(goods_type);size(50);null"`
	GoodsNum       int       `gorm:"column(goods_num);null"`
	GoodsAmount    float64   `gorm:"column(goods_amount);null;digits(10);decimals(2)"`
	DeliveryStatus string    `gorm:"column(delivery_status);size(50);null"`
	UpdateTime     time.Time `gorm:"column(update_time);type(datetime);null"`
	CreateTime     time.Time `gorm:"column(create_time);type(datetime);null"`
}

type OrderGoodsModel struct {
}

// 获取商品对应的订单数量
func (o *OrderGoods) GetOrderGoodsAll(goodsId []int, userId int64, field string) ([]OrderGoods, error) {
	var orderGoods []OrderGoods
	err := db.Select(field).Where("user_id = ? AND goods_id in (?) ", userId, goodsId).Find(&orderGoods).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return orderGoods, err
	}

	return orderGoods, nil
}

// save order goods
func (this *OrderGoodsModel) SaveOrderGoods(orderGoodsList []OrderGoods) error {

	for _, v := range orderGoodsList {
		err := db.Create(&v).Error
		if err != nil {
			return err
		}
	}
	return nil
}
