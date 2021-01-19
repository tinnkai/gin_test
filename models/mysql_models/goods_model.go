package mysql_models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// GinGoods [...]
type Goods struct {
	Id           int       `gorm:"primaryKey;column:id;type:int(11);not null" json:"-"`
	Name         string    `gorm:"column:name;type:varchar(200);not null" json:"name"`
	MarketAmount float64   `gorm:"column:market_amount;type:decimal(10,2);not null" json:"market_amount"`
	SellAmount   float64   `gorm:"column:sell_amount;type:decimal(10,2);not null" json:"sell_amount"`
	GoodsType    string    `gorm:"column:goods_type;type:varchar(30)" json:"goods_type"` // 主商品:GOODS,配件:PARTS,赠品:GIFT
	StockNum     int       `gorm:"column:stock_num;type:int(11)" json:"stock_num"`       // 库存
	LimitNum     int       `gorm:"column:limit_num;type:int(11)" json:"limit_num"`       // 限购
	Image        string    `gorm:"column:image;type:varchar(255)" json:"image"`          // 图片
	Describe     string    `gorm:"column:describe;type:text" json:"describe"`            // 描述
	Content      string    `gorm:"column:content;type:text" json:"content"`              // 详情
	UpdateTime   time.Time `gorm:"column:update_time;type:datetime" json:"update_time"`
	CreateTime   time.Time `gorm:"column:create_time;type:datetime" json:"create_time"`
}

func (t *Goods) TableName() string {
	return "gin_goods"
}

// 订单商品
type OrderGoodsInfo struct {
	Id             int
	Name           string `json:"name"`
	Image          string `json:"image"`
	Type           string
	OriginalAmount float64
	Amount         float64
	TotalAmount    float64
	Num            int
	GiftList       []OrderGift
}

// GetGoodsInfo Get a single article based on ID
func GetGoodsInfo(id int) (Goods, error) {
	var goodsInfo Goods
	err := db.Where("id = ?", id, 0).First(&goodsInfo).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return goodsInfo, err
	}

	return goodsInfo, nil
}

// 获取商品列表
func GetGoodsList(ids []int) ([]Goods, error) {
	var goodsList []Goods
	err := db.Select("id,name,market_amount,sell_amount,goods_type,stock_num,limit_num").
		Where("id in (?)", ids).Find(&goodsList).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return goodsList, err
	}

	return goodsList, nil
}
