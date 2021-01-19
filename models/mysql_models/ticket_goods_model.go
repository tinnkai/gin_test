package mysql_models

import (
	"time"
)

type TicketGoods struct {
	Id              int       `gorm:"column(id);auto"`
	TicketId        int       `gorm:"column(ticket_id);null"`
	GoodsId         int       `gorm:"column(goods_id);null"`
	GoodsCategoryId int       `gorm:"column(goods_category_id);null"`
	UpdateTime      time.Time `gorm:"column(update_time);type(datetime);null"`
	CreateTime      time.Time `gorm:"column(create_time);type(datetime);null"`
}

func (t *TicketGoods) TableName() string {
	return "gin_ticket_goods"
}

// 查询优惠券商品
func (t TicketGoods) GetTicketGoodsListByTicketId(ticketId int) (map[int]int, map[int]int, error) {
	var ticketGoodsList []TicketGoods
	err := db.Where("ticket_id = ?", ticketId).
		Find(&ticketGoodsList).Error

	goodsIdsMap := make(map[int]int)
	goodsCategoryIdsMap := make(map[int]int)
	if len(ticketGoodsList) > 0 {
		for _, v := range ticketGoodsList {
			goodsIdsMap[v.GoodsId] = v.GoodsId
			goodsCategoryIdsMap[v.GoodsCategoryId] = v.GoodsCategoryId
		}
	}
	return goodsIdsMap, goodsCategoryIdsMap, err
}
