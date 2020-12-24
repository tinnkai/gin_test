package order_service

import (
	"encoding/json"
	"fmt"
	models "gin_test/models/mysql_models"
	"gin_test/pkg/app"
	"gin_test/pkg/errors"
	"gin_test/service/promotions_service"
	"strconv"
	"time"
)

// 普通订单确认页接口
type normalCheckOutInterface interface {
	baseCheckOutInterface
	GetOrderGoodsGiftList([]models.OrderGoodsInfo) error
}

// 普通下单接口
type normalSaveOrderInterface interface {
	baseSaveOrderInterface
	GetOrderGoodsGiftList([]models.OrderGoodsInfo) error
}

// Normal post 信息
type NormalPost struct {
	// 基本信息
	BasePost
	// 地址信息
	AddressId int `form:"addressId"`
	// 优惠券id
	TicketId int `form:"ticketId"`
}

// 普通订单
type NormalOrder struct {
	NormalPost
	NormalOrderInfo
}

// 普通订单信息
type NormalOrderInfo struct {
	BaseOrderInfo
}

// 基础校验
func (this *NormalOrder) Validate() error {
	if len(this.NormalPost.GoodsInfo) == 0 {
		return errors.Newf(app.ERROR_ORDER_GOODS_NUM_FAIL, "", "")
	}

	for _, v := range this.NormalPost.GoodsInfo {
		if v.Num < 1 {
			return errors.Newf(app.ERROR_ORDER_GOODS_NUM_FAIL, fmt.Sprintf("商品数量不能小于%d", 1), "")
		}
	}

	return nil
}

// 校验商品
func (this *NormalOrder) ValidateGoods(goodsList []models.Goods, userId int64) error {
	goodsIds := make([]int, len(this.NormalPost.GoodsInfo))
	// 参数商品信息跟实际商品信息是否一致
	if len(this.NormalPost.GoodsInfo) != len(goodsList) {
		return errors.Newf(app.ERROR_ORDER_GOODS_PARAMS_FAIL, "", "")
	}
	for gi, gv := range goodsList {

		// 商品是否存在
		if gv.Id < 1 {
			return errors.Newf(app.ERROR_ORDER_GOODS_NOT_EMPTY, "", "")
		}

		// 商品类型是否正确
		if gv.GoodsType != "GOODS" {
			return errors.Newf(app.ERROR_ORDER_GOODS_TYPE_FAIL, "", "")
		}

		// 商品是否一致
		for _, pgv := range this.NormalPost.GoodsInfo {
			if pgv.Id == gv.Id {
				// 库存是否充足
				if gv.StockNum < 1 || gv.StockNum < pgv.Num {
					return errors.Newf(app.ERROR_ORDER_GOODS_STOCKS, "", "")
				}
			}
		}

		goodsIds[gi] = gv.Id
	}

	// 获取已购买数量
	buyLimitList, err := BuyLimit(goodsIds, userId)
	if err != nil {
		return errors.Newf(app.ERROR_ORDER_GOODS_LIMIT, "", "")
	}
	for _, gv := range goodsList {
		// 商品是否一致
		for _, pgv := range this.NormalPost.GoodsInfo {
			if pgv.Id != gv.Id {
				continue
			}
			// 限购
			if (buyLimitList[pgv.Id] + pgv.Num) > gv.LimitNum {
				return errors.Newf(app.ERROR_ORDER_GOODS_LIMIT, "", "")
			}
		}
	}
	return nil
}

// 订单确认
func (this *NormalOrder) CheckOut(userInfo models.User) (NormalOrderInfo, error) {
	// 用户id
	userId := userInfo.Id

	checkOutData := this.NormalOrderInfo

	// 初始化接口
	var io normalCheckOutInterface
	io = this

	// post数据校验
	err := io.Validate()
	if err != nil {
		return checkOutData, err
	}

	// 普通订单提交信息
	normalPost := this.NormalPost

	// 获取商品信息
	goodsList, err := io.GetGoodsList()

	// 验证商品信息
	err = io.ValidateGoods(goodsList, userInfo.Id)
	if err != nil {
		return checkOutData, err
	}

	// 获取订单商品信息
	checkOutData.GoodsList = io.GetOrderGoodsList(goodsList)

	// 查询赠品
	err = io.GetOrderGoodsGiftList(checkOutData.GoodsList)
	if err != nil {
		return checkOutData, errors.NewErrf(err, app.ERROR_ORDER_FAIL, "查询赠品失败", "")
	}

	// 获取促销信息
	var servicePromotionPrice promotions_service.ServicePromotionPrice
	err = servicePromotionPrice.PromotionPriceInfo(checkOutData.GoodsList)
	if err != nil {
		return checkOutData, errors.NewErrf(err, app.ERROR_ORDER_FAIL, "促销信息查询失败", "")
	}

	// 获取优惠券
	orderTicketList, err := promotions_service.GetTicketList(normalPost.TicketId, userId, checkOutData.GoodsList)
	checkOutData.TicketList = orderTicketList

	// 获取收货地址信息
	areas, err := models.GetAreasInfoById(normalPost.AddressId, userId)

	// 获取运费
	FreightInfo, err := models.GetFreightInfoByCityId(areas.CityCode)
	checkOutData.FreightInfo = FreightInfo

	// 计算订单金额
	checkOutData.CalculationOrderAmount(checkOutData.GoodsList)

	// 计算促销金额
	return checkOutData, nil
}

// 下单
func (this *NormalOrder) SaveOrder(userInfo models.User) (NormalOrderInfo, error) {
	userId := userInfo.Id
	nowTime := time.Now()

	normalSaveOrder := this.NormalOrderInfo

	// 初始化下单接口
	var io normalSaveOrderInterface
	io = this

	// post数据校验
	err := io.Validate()
	if err != nil {
		return normalSaveOrder, err
	}

	// 基础
	normalPost := this.NormalPost
	// 基础
	basePost := normalPost.BasePost

	// 获取商品信息
	goodsList, err := basePost.GetGoodsList()

	// 验证商品信息
	err = io.ValidateGoods(goodsList, userInfo.Id)
	if err != nil {
		return normalSaveOrder, err
	}

	// 订单商品列表
	normalSaveOrder.GoodsList = basePost.GetOrderGoodsList(goodsList)

	// 查询赠品
	err = basePost.GetOrderGoodsGiftList(normalSaveOrder.GoodsList)
	if err != nil {
		return normalSaveOrder, errors.NewErrf(err, app.ERROR_ORDER_FAIL, "查询赠品失败", "")
	}

	// 获取促销信息
	var servicePromotionPrice promotions_service.ServicePromotionPrice
	err = servicePromotionPrice.PromotionPriceInfo(normalSaveOrder.GoodsList)
	if err != nil {
		return normalSaveOrder, errors.NewErrf(err, app.ERROR_ORDER_FAIL, "促销信息查询失败", "")
	}

	// 使用优惠券
	orderTicketInfo, err := promotions_service.GetTicketInfo(normalPost.TicketId, userId, normalSaveOrder.GoodsList)

	// 收货地址地区信息
	areas, err := models.GetAreasInfoById(normalPost.AddressId, userId)

	// 获取运费
	freightInfo, err := models.GetFreightInfoByCityId(areas.CityCode)
	normalSaveOrder.FreightInfo = freightInfo

	// 计算订单金额
	normalSaveOrder.CalculationOrderAmount(normalSaveOrder.GoodsList)

	// 创建订单号
	orderNo := strconv.FormatInt(time.Now().UnixNano(), 10)

	// 保存订单
	modelOrder := models.Order{
		OrderNo:           orderNo,
		UserId:            userId,
		OrderStatus:       "CREATE",
		PayStatus:         "NO_PAY",
		AftermarketStatus: "NO",
		TotalAmount:       normalSaveOrder.TotalAmount,
		RealAmount:        normalSaveOrder.RealAmount,
		PayAmount:         normalSaveOrder.PayAmount,
		PromotionAmount:   normalSaveOrder.PromotionAmount,
		UpdateTime:        nowTime,
		CreateTime:        nowTime,
	}

	// 保存订单
	err = modelOrder.SaveOrder()
	if err != nil {
		return normalSaveOrder, errors.NewErrf(err, app.ERROR_ORDER_SAVE_FAIL, "订单保存失败", "")
	}
	orderId := modelOrder.Id

	// 保存订单商品
	var modelOrderGoodsList []models.OrderGoods
	for _, ogv := range normalSaveOrder.GoodsList {
		modelOrderGoodsList = append(modelOrderGoodsList, models.OrderGoods{
			OrderId:        orderId,
			UserId:         userId,
			GoodsId:        ogv.Id,
			GoodsType:      ogv.Type,
			GoodsNum:       ogv.Num,
			GoodsAmount:    ogv.Amount,
			DeliveryStatus: "NO",
			UpdateTime:     nowTime,
			CreateTime:     nowTime,
		})
	}

	var orderGoodsModel models.OrderGoodsModel
	err = orderGoodsModel.SaveOrderGoods(modelOrderGoodsList)
	if err != nil {
		return normalSaveOrder, errors.NewErrf(err, app.ERROR_ORDER_SAVE_FAIL, "订单商品保存失败", "")
	}

	// 保存订单赠品
	var modelOrderGoodsGiftList []models.OrderGoodsGift
	for _, og := range normalSaveOrder.GoodsList {
		for _, ogv := range og.GiftList {
			modelOrderGoodsGiftList = append(modelOrderGoodsGiftList, models.OrderGoodsGift{
				OrderId:        orderId,
				GoodsId:        ogv.RelationGoodsId,
				GiftGoodsId:    ogv.GiveGoodsId,
				Num:            ogv.Num,
				Amount:         ogv.Amount,
				DeliveryStatus: "NO",
				UpdateTime:     nowTime,
				CreateTime:     nowTime,
			})
		}
	}

	err = models.SaveOrderGoodsGift(modelOrderGoodsGiftList)
	if err != nil {
		return normalSaveOrder, errors.NewErrf(err, app.ERROR_ORDER_SAVE_FAIL, "订单赠品保存失败", "")
	}

	// 使用优惠券跟保存优惠券使用记录
	if orderTicketInfo.Id > 0 {
		// 使用优惠券
		ticketRecord := &models.TicketRecord{
			OrderId:    orderId,
			UseTime:    nowTime,
			Status:     "USED",
			UpdateTime: nowTime,
		}
		ticketRecord.UpdateRecordInfoById(orderTicketInfo.Id, userId)

		// 转换成json格式
		ticketSnapshotcontent, _ := json.Marshal(orderTicketInfo)
		// 记录快照
		orderTicketSnapshot := models.OrderTicketSnapshot{
			OrderId:    orderId,
			UserId:     userId,
			RecordId:   normalPost.TicketId,
			Content:    string(ticketSnapshotcontent),
			UpdateTime: nowTime,
			CreateTime: nowTime,
		}

		// 保存订单
		err = orderTicketSnapshot.SaveOrderTicketSnapshot()
		if err != nil {
			return normalSaveOrder, errors.NewErrf(err, app.ERROR_ORDER_SAVE_FAIL, "订单优惠券快照保存失败", "")
		}
	}

	return normalSaveOrder, nil
}
