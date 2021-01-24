package order

import (
	"gin_test/pkg/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserOrderController struct {
}

// @Summary 订单列表
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /order/v2/list [post]
func (this *UserOrderController) List(ctx *gin.Context) {
	appRes := app.Gin{Ctx: ctx}
	appRes.Response(http.StatusOK, app.SUCCESS, "v2_order_list", nil, false)
	return
}

// @Summary 订单详情
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /order/v2/detail [post]
func (this *UserOrderController) Detail(ctx *gin.Context) {
	appRes := app.Gin{Ctx: ctx}
	appRes.Response(http.StatusOK, app.SUCCESS, "v2_order_detail", nil, false)
	return
}
