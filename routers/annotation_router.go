package routers

import (
	"gin_test/controllers"

	"github.com/gin-gonic/gin"
	"github.com/xxjwxc/ginrpc"
	"github.com/xxjwxc/ginrpc/api"
)

// @Summary 注解路由
func AnnotationRouter(router *gin.Engine) {
	// 注册注解路由
	annotation := ginrpc.New(ginrpc.WithCtx(func(c *gin.Context) interface{} {
		return api.NewCtx(c)
	}), ginrpc.WithDebug(true))

	annotation.Register(router, new(controllers.HongdongController))

}
