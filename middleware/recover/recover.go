package recover

import (
	"fmt"
	"gin_test/pkg/app"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

// 错误恢复中间件
func ErrorRecover() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 接管恐慌性报错（panic）
		defer func() {
			if err := recover(); err != nil {
				errMsg := "服务运行错误"
				if gin.Mode() == gin.DebugMode {
					//打印错误堆栈信息
					debug.PrintStack()
					// 格式化错误信息
					errMsg = fmt.Sprintf("服务运行错误: %v", err)
				}
				//返回错误
				appG := app.Gin{Ctx: ctx}
				appG.Response(http.StatusInternalServerError, app.ERROR, errMsg, nil, true)
				//终止后续接口调用，不加的话recover到异常后，还会继续执行接口里后续代码
				ctx.Abort()
				return
			}
		}()

		ctx.Next()
	}
}
