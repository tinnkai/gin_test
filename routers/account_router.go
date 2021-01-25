package routers

import (
	v1_account "gin_test/controllers/v1/account"
	"gin_test/middleware/local_auth"

	"github.com/gin-gonic/gin"
)

// 用户相关路由
func AccountRouter(router *gin.Engine) {
	// 无登录认证组 V1
	v1AccountGroup := router.Group("/account/v1")
	{
		var account v1_account.Auth
		v1AccountGroup.GET("/auth", account.GetAuth)
	}

	// 登录认证组 V1
	v1AccountGroupAuth := v1AccountGroup.Use(local_auth.CheckLogin())
	{
		var user v1_account.User
		v1AccountGroupAuth.GET("/user-info", user.GetUserInfo)
	}

}
