package account

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"gin_test/pkg/app"
	"gin_test/pkg/utils"
	"gin_test/service/account_service"
	"gin_test/validates"
)

type Auth struct {
}

// @Summary Get Auth
// @Produce  json
// @Param phone query string true "phone"
// @Param password query string true "password"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /auth [post]
func (c *Auth) GetAuth(ctx *gin.Context) {
	appG := app.Gin{Ctx: ctx}

	// 绑定验证参数
	vAuth := new(validates.Auth)
	vError := vAuth.Bind(ctx)
	if vError != nil {
		errString := strings.Join(vError, ",")
		appG.Response(http.StatusOK, app.INVALID_PARAMS, errString, nil, true)
		return
	}

	// 验证登录
	user, err := account_service.CheckLogin(vAuth.Phone, vAuth.Password)
	if err != nil {
		appG.Response(http.StatusInternalServerError, app.ERROR_AUTH_CHECK_TOKEN_FAIL, "", nil, false)
		return
	}

	// 创建token
	authUser := &utils.AuthUser{
		Id:       user.Id,
		Username: user.Username,
		Group:    user.Group,
	}
	token, err := utils.GenerateToken(authUser)
	if err != nil {
		appG.Response(http.StatusInternalServerError, app.ERROR_AUTH_TOKEN, "", nil, false)
		return
	}
	//appG.C.Set()
	appG.Response(http.StatusOK, app.SUCCESS, "", map[string]string{
		"token": token,
	}, false)
}
