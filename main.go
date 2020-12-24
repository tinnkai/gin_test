package main

import (
	"fmt"
	"gin_test/models/mysql_models"
	"gin_test/pkg/gredis"
	"gin_test/pkg/logging"
	"gin_test/pkg/setting"
	"gin_test/pkg/utils"
	"gin_test/routers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	// 读取配置文件
	setting.Setup()
	// mysql model
	mysql_models.Setup()
	// 日志
	logging.Setup()
	// redis
	gredis.Setup()
	// mongodb model
	//mongodbmodels.Setup()
	// 安装初始化相关变量
	utils.Setup()
}

func main() {
	// 设置运行模式
	gin.SetMode(setting.ServerSetting.RunMode)
	// 设置读取超时时间
	readTimeout := setting.ServerSetting.ReadTimeout
	// 设置写入超时时间
	writeTimeout := setting.ServerSetting.WriteTimeout
	// 端口
	httpPort := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	// 请求头的最大字节数
	maxHeaderBytes := 1 << 20

	// 初始化 gin
	router := gin.Default()

	// 初始化路由
	routers.InitRouter(router)

	// 服务相关配置
	server := &http.Server{
		Addr:           httpPort,
		Handler:        router,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}
	server.ListenAndServe()
}