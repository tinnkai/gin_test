package mysql_models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Promotion struct {
	Id               int              `gorm:"column(id);auto"`
	Name             string           `gorm:"column(name);size(200)"`
	PreheatStartTime time.Time        `gorm:"column(preheat_start_time);type(datetime)"`
	PreheatEndTime   time.Time        `gorm:"column(preheat_end_time);type(datetime);null"`
	StartTime        time.Time        `gorm:"column(start_time);type(datetime);null"`
	EndTime          time.Time        `gorm:"column(end_time);type(datetime);null"`
	Image            string           `gorm:"column(image);size(255);null"`
	Limit            int              `gorm:"column(limit);null"`
	Scope            string           `gorm:"column(scope);size(20);null" description:"ALL | GOODS"`
	Describe         string           `gorm:"column(describe);null" description:"描述"`
	Status           string           `gorm:"column(status);size(20);null" description:"ENABELD | DISABLED"`
	IsDel            string           `gorm:"column(is_del);size(1);null" description:"Y|N"`
	UpdateTime       time.Time        `gorm:"column(update_time);type(datetime);null"`
	CreateTime       time.Time        `gorm:"column(create_time);type(datetime);null"`
	PromotionGoodss  []PromotionGoods `gorm:"FOREIGNKEY:PromotionId`
}

// 订单确认页商品赠品
type OrderPromotion struct {
	PromotionId int     `json:"promotionId"`
	GoodsId     int     `json:"goodsId"`
	Amount      float64 `json:"amount"`
}

func (t *Promotion) TableName() string {
	return "gin_promotion"
}

// 获取订单使用的促销价格信息
func GetOrderPromotionList(goodsIds []int) ([]OrderPromotion, error) {
	var orderPromotion []OrderPromotion
	start_time := time.Now().Format("2006-01-02 15:04:05")
	end_time := time.Now().Format("2006-01-02 15:04:05")

	/*err := db.Table(promotion.TableName()).Debug().
	Preload("PromotionGoodss", "goods_id IN (?)", goodsIds, func(db *gorm.DB) *gorm.DB {
		return db.Select("promotion_id,goods_id,amount")
	}).Where("start_time <= ? AND end_time >= ?", start_time, end_time).
	Scan(&orderPromotion).Error
	*/
	err := db.Table("gin_promotion AS p").
		Select("promotion_id,goods_id,amount").
		Joins("JOIN gin_promotion_goods AS pg ON p.id = pg.promotion_id").
		Where("p.start_time <= ? AND p.end_time >= ? AND pg.goods_id IN (?)", start_time, end_time, goodsIds).
		Scan(&orderPromotion).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return orderPromotion, err
	}

	return orderPromotion, nil
}
