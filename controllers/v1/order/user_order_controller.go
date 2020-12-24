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
// @Router /order/list [post,get]
func (this *UserOrderController) List(c *gin.Context) {
	appRes := app.Gin{C: c}
	appRes.Response(http.StatusOK, app.SUCCESS, "V1_order_list", nil, false)
	return
}

// @Summary 订单详情
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /order/detail [post,get]
func (this *UserOrderController) Detail(c *gin.Context) {
	appRes := app.Gin{C: c}
	appRes.Response(http.StatusOK, app.SUCCESS, "V1_order_detail", nil, false)
	return
}
