package jwt

import (
	"gin_test/pkg/app"
	"gin_test/pkg/setting"
	"gin_test/pkg/utils"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// JWT 参数
type jwtParams struct {
	Token string `form:"token"`
}

// JWT is jwt middleware
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data interface{}

		code := app.SUCCESS
		msg := ""
		// 绑定接收jwt参数
		jwtParams := jwtParams{}
		err := c.ShouldBind(&jwtParams)
		// 未接收到或者接收错误报参数错误
		if err != nil {
			code = app.ERROR_AUTH_EMPTY
		} else {
			// 校验token
			user, err := utils.ParseToken(jwtParams.Token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorEmpty:
					code = app.ERROR_AUTH_EMPTY
				case jwt.ValidationErrorExpired:
					code = app.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
				default:
					code = app.ERROR_AUTH_CHECK_TOKEN_FAIL
				}

				// 获取错误文本
				msg = err.(*jwt.ValidationError).Error()
			} else {
				// 将用户验证信息存储在上下文中
				c.Set(setting.RediskeySetting.AuthUserKey, *user)
			}
		}

		if code != app.SUCCESS {
			appG := app.Gin{C: c}
			appG.Response(http.StatusOK, code, msg, data, false)

			c.Abort()
			return
		}

		c.Next()
	}
}
