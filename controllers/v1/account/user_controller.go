package account

import (
	"encoding/json"
	"fmt"
	"gin_test/controllers"
	"gin_test/models/mysql_models"
	"gin_test/pkg/app"
	"gin_test/pkg/session"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
}

func (this *User) GetUserInfo(ctx *gin.Context) {
	appG := app.Gin{Ctx: ctx}

	userInfo := controllers.GetAuthUserInfo(ctx)
	err := session.Set(ctx, "UserInfo", userInfo)
	if err != nil {
		//appG response
		appG.Response(http.StatusOK, app.ERROR_AUTH, "", nil, false)
		return
	}

	// 设置一个数字
	err = session.Set(ctx, "count", 1233445)
	if err != nil {
		//appG response
		appG.Response(http.StatusOK, app.ERROR_AUTH, "", nil, false)
		return
	}

	//appG response
	appG.Response(http.StatusOK, app.SUCCESS, "", userInfo, false)
	return
}

func (this *User) GetUser(ctx *gin.Context) {
	appG := app.Gin{Ctx: ctx}
	var err error
	v := session.Get(ctx, "UserInfo")
	userInfo := mysql_models.AuthUserInfo{}
	if v != nil {

		err = json.Unmarshal(v.([]uint8), &userInfo)
		if err != nil {
			//appG response
			appG.Response(http.StatusOK, app.ERROR_AUTH, "", nil, false)
			return
		}
	}

	// 设置一个数字
	v1 := session.Get(ctx, "count")
	if v1 != nil {
		aa := 0
		err = json.Unmarshal(v1.([]uint8), &aa)

		fmt.Println(v1, aa, err)
	}

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
