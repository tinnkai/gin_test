package promotions_service

import models "gin_test/models/mysql_models"

type ServicePromotionPrice struct {
}

// 获取促销价格信息
func (this *ServicePromotionPrice) PromotionPriceInfo(orderGoodsList []models.OrderGoodsInfo) error {
	goodsIds := make([]int, len(orderGoodsList))
	for gi, gv := range orderGoodsList {
		goodsIds[gi] = gv.Id
	}

	// 获取促销价格列表
	orderPromotion, err := models.GetOrderPromotionList(goodsIds)
	if err != nil {
		return err
	}

	for ogi, ov := range orderGoodsList {
		for _, pv := range orderPromotion {
			if ov.Id == pv.GoodsId {
				orderGoodsList[ogi].Amount = pv.Amount
				orderGoodsList[ogi].TotalAmount = pv.Amount * float64(ov.Num)
			}
		}
	}
	return nil
}
