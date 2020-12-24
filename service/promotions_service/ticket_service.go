package promotions_service

import (
	models "gin_test/models/mysql_models"
	"gin_test/pkg/app"
	"gin_test/pkg/errors"
	"gin_test/pkg/utils"
	"sort"
)

// 获取优惠券列表
func GetTicketList(ticketId int, userId int64, orderGoods []models.OrderGoodsInfo) ([]models.OrderTicket, error) {
	// 获取用户优惠券列表
	ticketList, err := models.GetTicketListByUserId(userId)
	if err != nil {
		return ticketList, errors.Newf(app.ERROR_ORDER_TICKET_SELECT_FAIL, "", "")
	}
	// 优惠券是否使用标识
	isUsed := false
	for ti, tv := range ticketList {

		// 获取符合条件的商品信息
		intersect := make(map[int]float64)
		var orderGoodsTotal float64
		for _, gv := range orderGoods {

			if tv.Scope == "GOODS" {
				if _, ok := tv.TicketGoodsIds[gv.Id]; !ok {
					//不存在
					continue
				}
			} else if tv.Scope == "GOODSCATEGORY" {
				// 按商品分类
				if _, ok := tv.TicketGoodsCategoryIds[gv.Id]; !ok {
					//不存在
					continue
				}
			}
			// 设置为可用
			ticketList[ti].Status = "available"

			orderGoodsTotal += utils.Mul(gv.Amount, float64(gv.Num))
			intersect[gv.Id] = orderGoodsTotal
		}

		// 优惠券金额（面值）
		ticketAmount := tv.Amount
		// 优惠券已选中的不能选择其他优惠券，符合条件的商品金额必须大于等于优惠券最低使用金额
		if isUsed == false && orderGoodsTotal >= tv.MinAmount {
			// 当用户自定义使用优惠券，不匹配的直接跳过
			if ticketId > 0 && tv.RecordId != ticketId {
				continue
			}
			ticketList[ti].IsSelect = "yes"
			// 计算均摊抵扣，循环抵扣
			for gi, gv := range orderGoods {
				if _, ok := intersect[gv.Id]; !ok {
					continue
				}
				//存在
				if gv.TotalAmount >= ticketAmount {
					// 精度计算保留量为小数
					orderGoods[gi].TotalAmount = utils.Sub(gv.TotalAmount, ticketAmount)
				} else {
					ticketAmount = utils.Sub(ticketAmount, gv.TotalAmount)
					orderGoods[gi].TotalAmount = 0.00
				}
			}
			// 设为已使用
			isUsed = true
		}

	}
	// 按状态排序
	sort.SliceStable(ticketList, func(i, j int) bool {
		return ticketList[i].Status < ticketList[j].Status
	})
	return ticketList, nil
}

// 获取优惠券信息
func GetTicketInfo(ticketId int, userId int64, orderGoods []models.OrderGoodsInfo) (models.OrderTicket, error) {
	var ticketInfo models.OrderTicket

	// 没有使用优惠券直接返回
	if ticketId < 1 {
		return ticketInfo, nil
	}

	// 获取用户优惠券列表
	var ticket models.Ticket
	ticketInfo, err := ticket.GetTicketInfoById(ticketId, userId)
	if err != nil {
		return ticketInfo, err
	}

	// 获取符合条件的商品信息
	intersect := make(map[int]float64)
	var orderGoodsTotal float64
	for _, gv := range orderGoods {

		if ticketInfo.Scope == "GOODS" {
			if _, ok := ticketInfo.TicketGoodsIds[gv.Id]; !ok {
				//不存在
				continue
			}
		} else if ticketInfo.Scope == "GOODSCATEGORY" {
			// 按商品分类
			if _, ok := ticketInfo.TicketGoodsCategoryIds[gv.Id]; !ok {
				//不存在
				continue
			}
		}

		orderGoodsTotal += utils.Mul(gv.Amount, float64(gv.Num))
		intersect[gv.Id] = orderGoodsTotal
	}

	// 优惠券金额（面值）
	ticketAmount := ticketInfo.Amount
	// 符合条件的商品金额必须大于等于优惠券最低使用金额
	if orderGoodsTotal >= ticketInfo.MinAmount {

		// 计算均摊抵扣，循环抵扣
		for gi, gv := range orderGoods {
			if _, ok := intersect[gv.Id]; !ok {
				continue
			}
			//存在
			if gv.TotalAmount >= ticketAmount {
				// 精度计算保留量为小数
				orderGoods[gi].TotalAmount = utils.Sub(gv.TotalAmount, ticketAmount)
			} else {
				ticketAmount = utils.Sub(ticketAmount, gv.TotalAmount)
				orderGoods[gi].TotalAmount = 0.00
			}
		}
	}

	return ticketInfo, nil
}
