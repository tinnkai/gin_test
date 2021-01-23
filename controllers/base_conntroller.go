package controllers

import (
	"github.com/gin-gonic/gin"

	"gin_test/pkg/app"
	"gin_test/pkg/errors"
	"gin_test/pkg/utils"
)

// @Summary Get Auth User Info
func GetAuthUserInfo(c *gin.Context) (utils.AuthUser, error) {
	// redis key
	authUserKey := "AuthUserInfo"
	// 初始化用户信息
	userInfo := utils.AuthUser{}
	v, exist := c.Get(authUserKey)
	if !exist {
		return userInfo, errors.Newf(app.ERROR_AUTH, authUserKey+" not exist", "")
	}

	userInfo, ok := v.(utils.AuthUser)
	if ok {
		return userInfo, nil
	}

	return userInfo, errors.Newf(app.ERROR_AUTH, "用户信息异常", "")
}
