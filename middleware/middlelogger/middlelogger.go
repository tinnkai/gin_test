package middlelogger

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
	return func(c *gin.Context) {
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw

		// 获取唯一id
		requestUniqueId := utils.GetUuid()
		c.Set("request_unique_id", requestUniqueId)

		// 开始时间
		startTime := time.Now()

		c.Next()

		// 结束时间
		endTime := time.Now()
		// 执行时间
		latencyTime := endTime.Sub(startTime).Microseconds()
		// 请求方式
		reqMethod := c.Request.Method
		// 请求路由
		reqUri := c.Request.RequestURI
		// 请求参数
		reqPost := c.Request.PostForm
		// 状态码
		statusCode := c.Writer.Status()
		// 请求IP
		clientIP := c.ClientIP()

		requestInfo := map[string]interface{}{
			"reqUniqueId": requestUniqueId,
			"latencyTime": latencyTime,
			"reqMethod":   reqMethod,
			"reqUri":      reqUri,
			"statusCode":  statusCode,
			"clientIP":    clientIP,
			"reqPost":     reqPost,
			"resData":     blw.body.String(),
		}

		logging.LogInfoWithFields("requestInfo", requestInfo)
	}
}
