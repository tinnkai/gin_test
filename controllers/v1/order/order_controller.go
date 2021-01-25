package order

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gin_test/controllers"
	"gin_test/models/mysql_models"
	"gin_test/pkg/app"
	"gin_test/pkg/errors"
	"gin_test/service/order_service"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
}

// @Summary 普通订单确认页
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /order/v1/checkout [post]
func (this *OrderController) Checkout(ctx *gin.Context) {
	appG := app.Gin{Ctx: ctx}
	aa := []int{
		1, 2, 3,
	}
	fmt.Println(aa[4])
	// 获取登录信息
	authUserInfo := controllers.GetAuthUserInfo(ctx)
	if authUserInfo.UserId < 1 {
		appG.Response(http.StatusOK, app.ERROR_AUTH, "", "", false)
		return
	}

	// 获取用户信息
	userInfo, err := mysql_models.UserRepository.GetUserInfoById(authUserInfo.UserId, "id,username,password,phone,`group`")
	if err != nil {
		appG.Response(http.StatusOK, app.ERROR, err.Error(), "", false)
		return
	}

	// 初始化post信息
	pastData := new(order_service.NormalPost)
	err = ctx.Bind(pastData)
	if err != nil {
		appG.Response(http.StatusOK, app.ERROR, err.Error(), "", false)
		return
	}
	// 商品参数信息
	goodsInfo := ctx.PostForm("goodsInfo")
	err = json.Unmarshal([]byte(goodsInfo), &pastData.GoodsInfo)
	if err != nil {
		appG.Response(http.StatusOK, app.ERROR, err.Error(), "", false)
		return
	}

	normalOrder := new(order_service.NormalOrder)
	normalOrder.NormalPost = *pastData

	// 确认订单
	checkOutData, err := normalOrder.CheckOut(userInfo)
	if err != nil {
		errInfo := errors.GetErrorContext(err)
		appG.Response(http.StatusOK, errInfo.Code, errInfo.Message, errInfo.Field, true)
		return
	}

	appG.Response(http.StatusOK, app.SUCCESS, "", checkOutData, false)
}

// @Summary 普通下单
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /order/v1/saveorder [post]
func (this *OrderController) SaveOrder(ctx *gin.Context) {
	appG := app.Gin{Ctx: ctx}
	// 获取登录用户信息
	// 获取登录信息
	authUserInfo := controllers.GetAuthUserInfo(ctx)
	if authUserInfo.UserId < 1 {
		appG.Response(http.StatusOK, app.ERROR_AUTH, "", "", false)
		return
	}

	// 初始化post信息
	pastData := new(order_service.NormalPost)
	err := ctx.Bind(pastData)
	if err != nil {
		appG.Response(http.StatusOK, app.ERROR, err.Error(), "", false)
		return
	}
	// 商品信息
	goodsInfo := ctx.PostForm("goodsInfo")
	err = json.Unmarshal([]byte(goodsInfo), &pastData.GoodsInfo)
	if err != nil {
		appG.Response(http.StatusOK, app.ERROR, err.Error(), "", false)
		return
	}

	// 获取用户信息
	userInfo, err := mysql_models.UserRepository.GetUserInfoById(authUserInfo.UserId, "id,username,password,phone,`group`")
	if err != nil {
		appG.Response(http.StatusOK, app.ERROR, err.Error(), "", false)
		return
	}

	normalOrder := new(order_service.NormalOrder)
	normalOrder.NormalPost = *pastData

	// 获取checkout信息
	saveData, err := normalOrder.SaveOrder(userInfo)
	if err != nil {
		errInfo := errors.GetErrorContext(err)
		appG.Response(http.StatusOK, errInfo.Code, errInfo.Message, errInfo.Field, false)
		return
	}

	appG.Response(http.StatusOK, app.SUCCESS, "", saveData, false)
}
