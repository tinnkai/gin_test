package utils

import (
	"fmt"
	"strconv"
	"time"
)

// 获取当前时间
func NowDateTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// 获取当前日期
func NowDate() string {
	return time.Now().Format("2006-01-02")
}

type Datetime time.Time

func (t Datetime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(t).Format("2006-01-02 15:04:05"))
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
