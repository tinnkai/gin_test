package order_service

import (
	"gin_test/models/mysql_models"
	models "gin_test/models/mysql_models"
)

// 基本订单确认页接口
type baseCheckOutInterface interface {
	Validate() error
	GetGoodsList() ([]models.Goods, error)
	ValidateGoods(goodsList []models.Goods, userId int64) error
	GetOrderGoodsList([]models.Goods) []models.OrderGoodsInfo
	CalculationOrderAmount(orderGoods []mysql_models.OrderGoodsInfo)
}

// 基本下单接口
type baseSaveOrderInterface interface {
	Validate() error
	GetGoodsList() ([]models.Goods, error)
	ValidateGoods(goodsList []models.Goods, userId int64) error
	GetOrderGoodsList([]models.Goods) []models.OrderGoodsInfo
	CalculationOrderAmount(orderGoods []mysql_models.OrderGoodsInfo)
}

// post基本信息
type BasePost struct {
	// 活动类id
	ActivityId int `form:"activityId"`
	// 商品信息
	GoodsInfo []GoodsParams `form:"-"`
	// 订单类型
	Type string `form:"type"`
	// 留言
	Note string `form:"note"`
}

//用于保存实例化的结构体对象
var OrderTypeStruct map[string]interface{}

// 商品参数信息
type GoodsParams struct {
	// 商品id
	Id int `form:"id"`
	// 商品数量
	Num int `form:"num"`
}

// 基本订单信息
type BaseOrderInfo struct {
	GoodsList   []mysql_models.OrderGoodsInfo `json:"goodsList"`
	TicketList  []mysql_models.OrderTicket    `json:"ticketList"`
	FreightInfo mysql_models.OrderFreight     `json:"freightInfo"`

	// 订单号
	OrderNo         string  `json:"-"`
	OrderType       string  `json:"orderType"`
	TotalAmount     float64 `json:"totalAmount"`
	RealAmount      float64 `json:"realAmount"`
	PayAmount       float64 `json:"payAmount"`
	PromotionAmount float64 `json:"promotionAmount"`
}

// 获取商品信息列表
func (this *BasePost) GetGoodsList() ([]mysql_models.Goods, error) {
	ids := make([]int, len(this.GoodsInfo))

	j := 0
	for _, gv := range this.GoodsInfo {
		ids[j] = int(gv.Id)
		j++
	}

	goodsList, err := mysql_models.GetGoodsList(ids)
	if err != nil {
		return goodsList, err
	}

	return goodsList, nil
}

// 获取订单商品信息列表
func (this *BasePost) GetOrderGoodsList(goodsList []mysql_models.Goods) []mysql_models.OrderGoodsInfo {

	// 订单商品列表
	var orderGoodsList []mysql_models.OrderGoodsInfo

	for _, gv := range goodsList {
		// 从post数据中获取商品数量
		var goodsNum int = 0
		for _, v := range this.GoodsInfo {
			if gv.Id == v.Id {
				goodsNum = v.Num
			}
		}
		orderGoods := mysql_models.OrderGoodsInfo{
			Id:             gv.Id,
			Name:           gv.Name,
			Image:          "",
			Type:           gv.GoodsType,
			OriginalAmount: gv.SellAmount,
			Amount:         gv.SellAmount,
			TotalAmount:    gv.SellAmount * float64(goodsNum),
			Num:            goodsNum,
		}

		orderGoodsList = append(orderGoodsList, orderGoods)
	}

	return orderGoodsList
}

// 查询订单赠品
func (this *BasePost) GetOrderGoodsGiftList(orderGoodsList []mysql_models.OrderGoodsInfo) error {
	// 商品id
	goodsIds := make([]int, len(orderGoodsList))

	// 商品数量
	goodsNums := make(map[int]int)
	for gi, gv := range orderGoodsList {
		goodsIds[gi] = gv.Id
		goodsNums[gv.Id] = gv.Num
	}

	// 获取订单商品赠品
	giftList, err := mysql_models.GetOrderGiftList(goodsIds)
	if err != nil {
		return err
	}

	for ogi, ogv := range orderGoodsList {
		giftInfo := make([]mysql_models.OrderGift, 0)
		for _, giftV := range giftList {
			if ogv.Id == giftV.RelationGoodsId {
				goodsNum := goodsNums[giftV.RelationGoodsId]
				giftV.Num = giftV.Num * goodsNum
				giftInfo = append(giftInfo, giftV)
			}
		}
		orderGoodsList[ogi].GiftList = giftInfo
	}

	return nil
}

// 限购数量
func BuyLimit(goodsIds []int, userId int64) (map[int]int, error) {
	buyGoodsNum := make(map[int]int)

	// 查询订单
	var order mysql_models.Order
	goodsBuyNumList, err := order.GetGoodsBuyNum(goodsIds, userId)
	if err != nil {
		return buyGoodsNum, err
	}

	for _, v := range goodsBuyNumList {
		buyGoodsNum[v.GoodsId] = v.GoodsNum
	}

	return buyGoodsNum, nil
}

// 计算订单金额
func (this *BaseOrderInfo) CalculationOrderAmount(orderGoods []mysql_models.OrderGoodsInfo) {

	for _, v := range orderGoods {
		this.TotalAmount += v.TotalAmount
		this.RealAmount += v.TotalAmount
		this.PayAmount += v.TotalAmount
	}

	// 运费
	this.TotalAmount += this.FreightInfo.Amount
	this.RealAmount += this.FreightInfo.Amount
	this.PayAmount += this.FreightInfo.Amount
}
