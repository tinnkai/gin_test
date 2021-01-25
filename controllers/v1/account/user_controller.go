package account

import (
	"fmt"
	"gin_test/controllers"
	"gin_test/pkg/app"
	"gin_test/service/account_service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
}

func (this *User) GetUserInfo(ctx *gin.Context) {
	appG := app.Gin{Ctx: ctx}

	userInfo := controllers.GetAuthUserInfo(ctx)

	userInfo, _ = account_service.CheckLogin(3)
	fmt.Println(userInfo)

	//appG response
	appG.Response(http.StatusOK, app.SUCCESS, "", userInfo, false)
	return
}

func (this *User) GetUserInfoName(ctx *gin.Context) {
	appG := app.Gin{Ctx: ctx}

	name := ctx.Query("name")
	age := ctx.Query("age")

	userInfo := map[string]string{
		"name": name,
		"age":  age,
	}
	//appG response
	appG.Response(http.StatusOK, app.SUCCESS, "", userInfo, false)
	return
}
