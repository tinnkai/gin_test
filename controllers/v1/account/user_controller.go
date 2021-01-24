package account

import (
	"gin_test/pkg/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
}

func (c *User) GetUserInfo(ctx *gin.Context) {
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

func (c *User) GetUserInfoName(ctx *gin.Context) {
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
