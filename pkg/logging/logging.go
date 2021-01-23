package logging

import (
	"fmt"
	"gin_test/pkg/file"
	"gin_test/pkg/setting"
	"runtime"
	"time"

	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var Log = logrus.New()

// 安装使用
func Setup() {

	writerInfo := setInfoRotatelogs()
	writerError := setErrorRotatelogs()
	pathMap := lfshook.WriterMap{
		logrus.InfoLevel:  writerInfo,
		logrus.ErrorLevel: writerError,
		logrus.PanicLevel: writerError,
	}

	// 设置将日志输出到标准输出（默认的输出为stderr,标准错误）
	//Log.SetOutput(writerError)

	// 设置日志级别为warn以上
	//logger.SetLevel(logrus.WarnLevel)
	Log.Hooks.Add(lfshook.NewHook(
		pathMap,
		&logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		},
	))
}

func setInfoRotatelogs() *rotatelogs.RotateLogs {
	path := "runtime/logs/info/"
	// 创建目录
	err := file.MkDir(path)
	if err != nil {
		panic("Log directory creation failed")
	}
	writer, _ := rotatelogs.New(
		path+"info_log_%Y%m%d.log",
		//rotatelogs.WithLinkName(path),             // 生成软链，指向最新日志文件
		// 文件最大保存时间（-1相当于禁用）
		rotatelogs.WithMaxAge(-1),
		// 最大保存个数
		rotatelogs.WithRotationCount(2),
		// 日志切割时间间隔
		rotatelogs.WithRotationTime(200*time.Second),
	)

	return writer
}

func setErrorRotatelogs() *rotatelogs.RotateLogs {
	path := "runtime/logs/error/"
	// 创建目录
	err := file.MkDir(path)
	if err != nil {
		panic("Log directory creation failed")
	}
	writer, _ := rotatelogs.New(
		path+"error_log_%Y%m%d.log",
		//rotatelogs.WithLinkName(path), // 生成软链，指向最新日志文件
		// 文件最大保存时间（-1相当于禁用）
		rotatelogs.WithMaxAge(-1),
		// 最大保存个数
		rotatelogs.WithRotationCount(5),
		rotatelogs.WithRotationTime(24*time.Hour), // 日志切割时间间隔
	)

	return writer
}

type Fields logrus.Fields

func LogFatal(v interface{}) {
	Log.Fatal(v)
}
func LogFatalWithFields(v interface{}, f Fields) {
	entry := Log.WithFields(logrus.Fields(f))
	entry.Fatal(v)
}
func LogError(v interface{}) {
	Log.Error(v)
}
func LogErrorWithFields(v interface{}, f Fields) {
	f["file"] = callers()
	entry := Log.WithFields(logrus.Fields(f))
	entry.Error(v)
}
func LogWarning(v interface{}) {
	Log.Warning(v)
}
func LogWarningWithFields(v interface{}, f Fields) {
	entry := Log.WithFields(logrus.Fields(f))
	entry.Warning(v)
}
func LogInfo(v interface{}) {
	Log.Info(v)
}
func LogInfoWithFields(v interface{}, f Fields) {
	entry := Log.WithFields(logrus.Fields(f))
	entry.Info(v)
}
func LogDebug(v interface{}) {
	Log.Debug(v)
}
func LogDebugWithFields(v interface{}, f Fields) {
	entry := Log.WithFields(logrus.Fields(f))
	entry.Debug(v)
}
func LogTrace(v interface{}) {
	Log.Trace(v)
}
func LogTraceWithFields(v interface{}, f Fields) {
	entry := Log.WithFields(logrus.Fields(f))
	entry.Trace(v)
}

// 获取错误追踪
func callers() (files string) {
	// 获取错误层级
	skip := setting.LogSetting.Skip

	pc := make([]uintptr, 5) // at least 1 entry needed
	n := runtime.Callers(skip, pc)
	frames := runtime.CallersFrames(pc[:n])
	for {
		frame, more := frames.Next()
		files += fmt.Sprintf("%s:%d (%s),", frame.File, frame.Line, frame.Function)
		if !more {
			break
		}
	}
	return
}
