package logger

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go-gin-seed/pkg/setting"
	"io"
	"os"
	"time"
)

var Log = logrus.New()

func Debug(format string, args ...interface{}) {
	if args == nil {
		Log.Debug(format)
	} else {
		Log.Debugf(format, args)
	}
}
func Info(format string, args ...interface{}) {
	Log.Infof(format, args)
}
func Warn(format string, args ...interface{}) {
	Log.Warnf(format, args)
}
func Error(format string, args ...interface{}) {
	Log.Errorf(format, args)
}
func Fatal(format string, args ...interface{}) {
	Log.Fatalf(format, args)
}
func Panic(format string, args ...interface{}) {
	if args == nil {
		Log.Panic(format)
	} else {
		Log.Panicf(format, args)
	}
}

func Init() {
	Log.SetFormatter(&logrus.TextFormatter{TimestampFormat: "2021-01-01 12:30:30.999"})
	Log.SetOutput(os.Stdout)
	logName := setting.AppSetting.LogPathFile
	var f *os.File
	var err error
	if _, err := os.Stat(logName); os.IsNotExist(err) {
		f, err = os.Create(logName)
	} else {
		f, err = os.OpenFile(logName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	}

	if err != nil {
		fmt.Println("open log file failed")
		panic(err)
	}
	writers := []io.Writer{
		f,
		os.Stdout}
	//write to console and file
	fileAndStdoutWriter := io.MultiWriter(writers...)
	Log.SetOutput(fileAndStdoutWriter)
	switch setting.AppSetting.LogLevel {
	case "trace":
		Log.Level = logrus.TraceLevel
	case "debug":
		Log.Level = logrus.DebugLevel
	case "info":
		Log.Level = logrus.InfoLevel
	case "warn":
		Log.Level = logrus.WarnLevel
	case "error":
		Log.Level = logrus.ErrorLevel
	case "fatal":
		Log.Level = logrus.FatalLevel
	case "panic":
		Log.Level = logrus.PanicLevel
	default:
		fmt.Println("log level not match")
		return
	}
}

// 日志记录
func LogHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()
		// 处理请求
		c.Next()
		// 结束时间
		endTime := time.Now()
		// 执行时间
		latencyTime := fmt.Sprintf("%6v", endTime.Sub(startTime))
		// 请求方式
		reqMethod := c.Request.Method
		// 请求路由
		reqUri := c.Request.RequestURI
		// 状态码
		statusCode := c.Writer.Status()
		// 请求IP
		clientIP := c.ClientIP()
		//日志格式
		Log.WithFields(logrus.Fields{
			"http_status": statusCode,
			"total_time":  latencyTime,
			"ip":          clientIP,
			"method":      reqMethod,
			"uri":         reqUri,
		}).Info("access")
	}
}
