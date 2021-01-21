package utils

import (
	"fmt"
	"gin_test/pkg/setting"
	"os"
	"strconv"
	"time"

	"github.com/segmentio/ksuid"
)

// 获取配置环境变量
func GetConfigorEnv() string {
	configorEnv := os.Getenv("CONFIGOR_ENV")
	if configorEnv == "" {
		configorEnv = "pro"
	}

	return configorEnv
}

// 获取分布式id(雪花算法)
func GetUuid() string {
	id := ksuid.New()
	return id.String()
}

// 获取当前时间
func NowDateTime() string {
	return time.Now().Format(setting.AppSetting.DateTimeFormat)
}

// 获取当前日期
func NowDate() string {
	return time.Now().Format(setting.AppSetting.DateFormat)
}

// 无时区时间类型
type CustomDatetime time.Time

// 返回json时去除时间中的时区
func (c CustomDatetime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(c).Format(setting.AppSetting.DateTimeFormat))
	return []byte(stamp), nil
}

// 加法
func Add(value1, value2 float64) float64 {
	res, err := strconv.ParseFloat(fmt.Sprintf("%.2f", (value1+value2)), 64)
	if err != nil {
		panic("加法计算错误")
	}
	return res
}

// 减法
func Sub(value1, value2 float64) float64 {
	res, err := strconv.ParseFloat(fmt.Sprintf("%.2f", (value1-value2)), 64)
	if err != nil {
		panic("减法计算错误")
	}
	return res
}

// 乘法
func Mul(value1, value2 float64) float64 {
	res, err := strconv.ParseFloat(fmt.Sprintf("%.2f", (value1*value2)), 64)
	if err != nil {
		panic("乘法计算错误")
	}
	return res
}

// 除法
func Div(value1, value2 float64) float64 {
	res, err := strconv.ParseFloat(fmt.Sprintf("%.2f", (value1/value2)), 64)
	if err != nil {
		panic("除法计算错误")
	}
	return res
}
