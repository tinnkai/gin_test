package controllers

import (
	"gin_test/pkg/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ReqTest struct {
	AccessToken string `json:"access_token" form:"access_token"`
	UserName    string `json:"user_name" form:"user_name" binding:"required"` // 带校验方式
	Password    string `json:"password" form:"password" form:"password"`
}

type HongdongController struct {
}

// @Summary 活动列表
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /hongdong/list [post,get]
func (this *HongdongController) List(c *gin.Context) {
	appG := app.Gin{C: c}
	var req ReqTest

	err := c.ShouldBind(&req)
	if err != nil {
		appG.Response(http.StatusBadRequest, app.INVALID_PARAMS, err.Error(), nil, false)
		return
	}
	appG.Response(http.StatusOK, app.SUCCESS, "", req, false)
	return
}

// @Summary 活动详情
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /hongdong/detail [post,get]
func (this *HongdongController) Detail(c *gin.Context) {
	appG := app.Gin{C: c}
	var req ReqTest

	err := c.ShouldBind(&req)
	if err != nil {
		appG.Response(http.StatusBadRequest, app.INVALID_PARAMS, err.Error(), nil, false)
		return
	}
	appG.Response(http.StatusOK, app.SUCCESS, "", req, false)
	return
}