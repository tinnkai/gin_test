package controllers

import "github.com/gin-gonic/gin"

func GetUserInfo(ctx *gin.Context) {

	name := ctx.Query("name")
	age := ctx.Query("age")

	userInfo := map[string]string{
		"name": name,
		"age":  age,
	}
	ctx.JSON(200, userInfo)
}

func GetUserInfoName(ctx *gin.Context) {

	name := ctx.Query("name")
	age := ctx.Query("age")

	userInfo := map[string]string{
		"name": name,
		"age":  age,
	}
	ctx.JSON(200, userInfo)
}
