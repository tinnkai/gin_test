package logger

import (
	"bytes"
	"gin_test/pkg/logging"
	"gin_test/pkg/utils"
	"time"

	"github.com/gin-gonic/gin"
)

// 定义返回日志结构
type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

// 写入
func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// 日志中间件
func Log() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 开始时间
		startTime := time.Now()

		// 获取唯一id
		requestUniqueId := utils.GetUuid()
		ctx.Set("request_unique_id", requestUniqueId)

		// 请求方式
		reqMethod := ctx.Request.Method
		// 请求路由
		reqUri := ctx.Request.RequestURI
		// 请求参数
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

		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: ctx.Writer}
		ctx.Writer = blw

		ctx.Next()

		// 结束时间
		endTime := time.Now()
		// 执行时间
		latencyTime := endTime.Sub(startTime).Microseconds()

		requestInfo["latencyTime"] = latencyTime
		requestInfo["resData"] = blw.body.String()

		logging.LogInfoWithFields("requestInfo", requestInfo)
	}
}
