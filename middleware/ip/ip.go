package ip

import (
	"gin_test/pkg/app"
	"gin_test/pkg/setting"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ip白名单验证
func IpWhiteListCheck() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取ip：如果通过 nginx 需要配置一下参数
		// proxy_set_header X-Forward-For $remote_addr;
		// proxy_set_header X-real-ip $remote_addr;
		ip := ctx.ClientIP()

		flag := false
		// 获取ip白名单
		whiteList := setting.IpSetting.WhiteList
		for _, item := range whiteList {
			if item == ip {
				flag = true
			}
		}
		if flag == false {
			appG := app.Gin{Ctx: ctx}
			appG.Response(http.StatusOK, app.ERROR_IP_CHECK_FAIL, "", nil, false)

			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
