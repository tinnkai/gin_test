package routers

import (
	v1_order "gin_test/controllers/v1/order"
	v2_order "gin_test/controllers/v2/order"
	"gin_test/middleware/jwt"

	"github.com/gin-gonic/gin"
)

// 订单相关路由
func OrderRouter(router *gin.Engine) {
	// Order group: v1
	v1 := router.Group("/order/v1")
	v1.Use(jwt.JWT())
	{
		var userOrder v1_order.UserOrderController
		v1.POST("/list", userOrder.List)
		v1.POST("/detail", userOrder.Detail)

		var order v1_order.OrderController
		v1.POST("/checkout", order.Checkout)
		v1.POST("/saveorder", order.SaveOrder)
	}

	// Order group: v2
	v2 := router.Group("/order/v2")
	{
		var userOrder v2_order.UserOrderController
		v2.POST("/list", userOrder.List)
		v2.POST("/detail", userOrder.Detail)
	}
}
