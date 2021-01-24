package app

import (
	"github.com/gin-gonic/gin"

	"gin_test/pkg/logging"
)

type Gin struct {
	Ctx *gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// 返回 gin.JSON
func (g *Gin) Response(httpCode, errCode int, errMsg string, data interface{}, isWriteLog bool) {
	if errMsg == "" {
		errMsg = GetMsg(errCode)
	}

	responseData := Response{
		Code: errCode,
		Msg:  errMsg,
		Data: data,
	}
	g.Ctx.JSON(httpCode, responseData)

	// 是否记录错误日志
	if isWriteLog {
		// 记录错误日志
		responseErrorLog(g.Ctx, responseData)
	}

	return
}

// Response error log
func responseErrorLog(ctx *gin.Context, responseData Response) {
	// 获取请求唯一id
	requestUniqueId, _ := ctx.Get("request_unique_id")
	// 请求方式
	reqMethod := ctx.Request.Method
	// 请求路由
	reqUri := ctx.Request.RequestURI
	// 请求参数(POST)
	reqPost := ctx.Request.PostForm
	// 状态码
	statusCode := ctx.Writer.Status()
	// 请求IP
	clientIP := ctx.ClientIP()

	requestInfo := map[string]interface{}{
		"reqUniqueId": requestUniqueId,
		"reqMethod":   reqMethod,
		"reqUri":      reqUri,
		"statusCode":  statusCode,
		"clientIP":    clientIP,
		"reqPost":     reqPost,
	}
	logging.LogErrorWithFields(responseData, requestInfo)
}
