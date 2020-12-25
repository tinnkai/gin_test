package routers

import (
	_ "gin_test/docs"
	"gin_test/pkg/setting"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @Summary 初始化路由
func InitRouter(router *gin.Engine) {
	// swagger debug run
	if setting.ServerSetting.RunMode == "debug" {
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	// 普通方式路由
	NormalRouter(router)

	// 用户路由
	AccountRouter(router)

	// 订单相关路由
	OrderRouter(router)

	// 注解路由
	AnnotationRouter(router)
}
