package app

import (
	"github.com/gin-gonic/gin"

	"gin_test/pkg/logging"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Response setting gin.JSON
func (g *Gin) Response(httpCode, errCode int, errMsg string, data interface{}, isWriteLog bool) {
	if errMsg == "" {
		errMsg = GetMsg(errCode)
	}

	responseData := Response{
		Code: errCode,
		Msg:  errMsg,
		Data: data,
	}
	g.C.JSON(httpCode, responseData)

	// 是否记录错误日志
	if isWriteLog {
		// 获取请求唯一id
		requestUniqueId, _ := g.C.Get("request_unique_id")
		logging.LogErrorWithFields(responseData, logging.Fields{
			"reqUniqueId": requestUniqueId,
		})
	}

	return
}
