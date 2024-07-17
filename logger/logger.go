package logger

import (
	"github.com/natefinch/lumberjack"
	"io"
	"runtime"
	"strings"
	"winstack_collect/common"

	"github.com/sirupsen/logrus"
)

var defaultLog *logrus.Logger
var defaultLogPath = common.NameSpace + ".log"
var defaultLogLevel = logrus.WarnLevel

func init() {
	// 根据系统来判断存储的日志文件的路径
	system := strings.ToLower(runtime.GOOS)
	if v, ok := common.LogPathMap[system]; ok {
		defaultLogPath = v
	}
	defaultLog := logrus.New()
	defaultLog.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
	// 设置默认的日志等级
	defaultLog.SetLevel(defaultLogLevel)
	defaultLog.SetOutput(&lumberjack.Logger{
		Filename:   defaultLogPath,
		MaxSize:    common.DefaultLogMaxSize,
		MaxBackups: common.DefaultLogMaxBackups,
		MaxAge:     common.DefaultLogMaxAge,
		Compress:   common.DefaultLogCompress,
	})
}

func SetLevel(level logrus.Level) {
	defaultLog.SetLevel(level)
}

func SetOutput(output io.Writer) {
	defaultLog.SetOutput(output)
}

// Debug 相关
func Debug(args ...interface{}) {
	defaultLog.Debug(args...)
}

func Debugln(args ...interface{}) {
	defaultLog.Debugln(args...)
}

func Debugf(format string, args ...interface{}) {
	defaultLog.Debugf(format, args...)
}

// Info 相关
func Info(args ...interface{}) {
	defaultLog.Info(args...)
}

func Infoln(args ...interface{}) {
	defaultLog.Infoln(args...)
}

func Infof(format string, args ...interface{}) {
	defaultLog.Infof(format, args...)
}

// Warn 相关
func Warn(args ...interface{}) {
	defaultLog.Warn(args...)
}

func Warnln(args ...interface{}) {
	defaultLog.Warnln(args...)
}

func Warnf(format string, args ...interface{}) {
	defaultLog.Warnf(format, args...)
}

// Error 相关
func Error(args ...interface{}) {
	defaultLog.Error(args...)
}

func Errorln(args ...interface{}) {
	defaultLog.Errorln(args...)
}

func Errorf(format string, args ...interface{}) {
	defaultLog.Errorf(format, args...)
}
