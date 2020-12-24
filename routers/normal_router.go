package routers

import (
	"gin_test/controllers"

	"github.com/gin-gonic/gin"
)

// @Summary 普通路由
func NormalRouter(router *gin.Engine) {
	// 用户信息
	router.GET("/user-info", controllers.GetUserInfo)
}
