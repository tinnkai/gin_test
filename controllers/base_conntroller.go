package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"gin_test/models/mysql_models"
	"gin_test/pkg/app"
)

var (
	userId   int64
	username string
	phone    int64
)

// 获取登录用户信息
func GetAuthUserInfo(ctx *gin.Context) mysql_models.AuthUserInfo {
	fmt.Println(userId)
	if userId < 1 {
		getUserBaseInfo(ctx)
	}
	// 初始化用户信息
	userInfo := mysql_models.AuthUserInfo{
		UserId:   userId,
		Username: username,
		Phone:    phone,
	}

	return userInfo
}

// 获取登录用户id
func GetAuthUserId(ctx *gin.Context) int64 {
	if userId < 1 {
		getUserBaseInfo(ctx)
	}
	// 用户id
	return userId
}

// 获取登录用户名
func GetAuthUsername(ctx *gin.Context) string {
	if userId < 1 {
		getUserBaseInfo(ctx)
	}
	// 用户名
	return username
}

// 获取登录用户手机号
func GetAuthPhone(ctx *gin.Context) int64 {
	if userId < 1 {
		getUserBaseInfo(ctx)
	}
	// 手机号码
	return phone
}

// 从上下文中获取用户基本信息
func getUserBaseInfo(ctx *gin.Context) {
	appG := app.Gin{Ctx: ctx}
	// 上下文中获取用户信息
	authUserKey := "UserInfo"
	// 初始化用户信息
	userInfo := mysql_models.AuthUserInfo{}
	v, exist := ctx.Get(authUserKey)
	fmt.Println(v)
	if !exist {
		appG.Response(http.StatusOK, app.ERROR_AUTH_EMPTY, "", "", false)
		return
	}

	userInfo, ok := v.(mysql_models.AuthUserInfo)
	if !ok || userInfo.UserId < 1 {
		appG.Response(http.StatusOK, app.ERROR, "", "", false)
		return
	}

	// 用户id
	userId = userInfo.UserId
	// 用户名
	username = userInfo.Username
	// 手机号
	phone = userInfo.Phone
}
