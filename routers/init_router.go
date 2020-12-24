package routers

import (
	"github.com/gin-gonic/gin"
)

// @Summary 初始化路由
func InitRouter(router *gin.Engine) {

	// 普通方式路由
	NormalRouter(router)

	// 用户路由
	AccountRouter(router)

	// 订单相关路由
	OrderRouter(router)

	// 注解路由
	AnnotationRouter(router)
}
