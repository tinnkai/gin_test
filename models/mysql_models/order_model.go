package mysql_models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Order struct {
	Id                uint      `gorm:"column(id);auto"`
	OrderNo           string    `gorm:"column(order_no);size(100)"`
	UserId            int64     `gorm:"column(user_id);null"`
	OrderStatus       string    `gorm:"column(order_status);size(50);null"`
	PayStatus         string    `gorm:"column(pay_status);size(50);null"`
	AftermarketStatus string    `gorm:"column(aftermarket_status);size(50);null"`
	TotalAmount       float64   `gorm:"column(total_amount);null;digits(10);decimals(2)"`
	RealAmount        float64   `gorm:"column(real_amount);null;digits(10);decimals(2)"`
	PayAmount         float64   `gorm:"column(pay_amount);null;digits(10);decimals(2)"`
	PromotionAmount   float64   `gorm:"column(promotion_amount);null;digits(10);decimals(2)"`
	PayTime           time.Time `gorm:"column(pay_time);type(datetime);null"`
	UpdateTime        time.Time `gorm:"column(update_time);type(datetime);null"`
	CreateTime        time.Time `gorm:"column(create_time);type(datetime);null"`
}

type GoodsBuyNum struct {
	GoodsId  int
	GoodsNum int
}

func (t *Order) TableName() string {
	return "gin_order"
}

// 获取商品对应的订单数量
func (o *Order) GetGoodsBuyNum(goodsIds []int, userId int64) ([]GoodsBuyNum, error) {
	var goodsBuyNumList []GoodsBuyNum

	querydb := db
	querydb = querydb.Table("gin_order AS o")
	querydb = querydb.Select("og.goods_id,SUM(og.goods_num) AS goods_num")
	querydb = querydb.Joins("JOIN gin_order_goods AS og ON o.id = og.order_id")
	querydb = querydb.Where("o.user_id = ? AND og.goods_id in (?) AND order_status != ?", userId, goodsIds, "CANCEL")
	querydb = querydb.Group("og.goods_id")
	querydb = querydb.Scan(&goodsBuyNumList)
	err := querydb.Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return goodsBuyNumList, err
	}

	return goodsBuyNumList, nil
}

// SaveOrder
func (o *Order) SaveOrder() error {
	return db.Create(o).Error
}
