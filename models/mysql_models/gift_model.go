package mysql_models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Gift struct {
	Id         int       `gorm:"column(id);auto"`
	Name       string    `gorm:"column(name);size(200)"`
	StartTime  time.Time `gorm:"column(start_time);type(datetime);null"`
	EndTime    time.Time `gorm:"column(end_time);type(datetime);null"`
	Describe   string    `gorm:"column(describe);null" description:"描述"`
	Status     string    `gorm:"column(status);size(20);null" description:"ENABELD | DISABLED"`
	IsDel      string    `gorm:"column(is_del);size(1);null" description:"Y|N"`
	UpdateTime time.Time `gorm:"column(update_time);type(datetime);null"`
	CreateTime time.Time `gorm:"column(create_time);type(datetime);null"`
}

// 订单赠品
type OrderGift struct {
	RelationGoodsId int     `json:"relationGoodsId"`
	GiveGoodsId     int     `json:"giveGoodsId"`
	Name            string  `json:"name"`
	Image           string  `json:"image"`
	Type            string  `json:"-"`
	Amount          float64 `json:"amount"`
	Num             int     `json:"num"`
}

func (g *Gift) TableName() string {
	return "gin_gift"
}

// 获取赠品列表
func GetOrderGiftList(goodsIds []int) ([]OrderGift, error) {
	var orderGift []OrderGift
	start_time := time.Now().Format("2006-01-02 15:04:05")
	end_time := time.Now().Format("2006-01-02 15:04:05")
	/*db.Table("gift").Where("start_time <= ? AND end_time >= ?", start_time, end_time)
	err := db.Preload("GiftRelationGoods", "goods_id IN ?", goodsId).Scan(&checkoutGoodsGift).Error*/
	querydb := db
	querydb = querydb.Table("gin_gift AS g")
	querydb = querydb.Select("goods.name,goods.image,goods.goods_type AS `type`,goods.sell_amount AS amount,ggg.num,ggg.relation_goods_id,ggg.give_goods_id")
	querydb = querydb.Joins("JOIN gin_gift_relation_goods AS grg ON g.id = grg.gift_id AND grg.goods_id IN (?)", goodsIds)
	querydb = querydb.Joins("JOIN gin_gift_give_goods AS ggg ON grg.gift_id = grg.gift_id AND grg.goods_id = ggg.relation_goods_id")
	querydb = querydb.Joins("JOIN gin_goods AS goods ON ggg.give_goods_id = goods.id")
	querydb = querydb.Where("g.start_time <= ? AND g.end_time >= ?", start_time, end_time)
	querydb = querydb.Scan(&orderGift)
	err := querydb.Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return orderGift, nil
}
