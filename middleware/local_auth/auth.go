package local_auth

import (
	"fmt"
	"gin_test/pkg/app"
	"gin_test/service/account_service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 本地登录参数
type localLoginParams struct {
	UserId int64 `form:"user_id"`
}

// 验证登录
func CheckLogin() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		errCode := app.SUCCESS
		errMsg := ""
		// 绑定接收登录参数
		localLoginParams := localLoginParams{}
		err := ctx.ShouldBind(&localLoginParams)

		// 未接收到或者接收错误报参数错误
		if err != nil {
			errCode = app.ERROR_AUTH_EMPTY
		} else {
			// 校验用户id
			userInfo, err := account_service.CheckLogin(localLoginParams.UserId)
			if err != nil {
				errCode = app.ERROR_AUTH
				errMsg = fmt.Sprintf("登录失败: %v", err)
			} else {
				// 将用户验证信息存储在上下文中
				// value, err := json.Marshal(userInfo)
				// if err != nil {
				// 	errCode = app.ERROR_AUTH
				// 	errMsg = fmt.Sprintf("登录失败: %v", err)
				// }
				ctx.Set("UserInfo", userInfo)
			}
		}

		// 当错误码不为 1 返回错误
		if errCode != app.SUCCESS {
			appG := app.Gin{Ctx: ctx}
			appG.Response(http.StatusOK, errCode, errMsg, nil, false)

			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
