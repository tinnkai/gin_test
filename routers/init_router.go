package routers

import (
	_ "gin_test/docs"
	"gin_test/middleware/middlelogger"
	"gin_test/pkg/utils"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @Summary 初始化路由
func InitRouter(router *gin.Engine) {
	// 获取配置环境变量
	configorEnv := utils.GetConfigorEnv()

	// swagger debug run
	if configorEnv != "pro" {
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	// 日志中间件
	router.Use(middlelogger.Log())

	// 普通方式路由
	NormalRouter(router)

	// 用户路由
	AccountRouter(router)

	// 订单相关路由
	OrderRouter(router)

	// 注解路由
	AnnotationRouter(router)
}
