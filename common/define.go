package common

import (
	"github.com/sirupsen/logrus"
	"time"
)

const (
	Version           = "1.0.3"
	NameSpace         = "cw_YunHong_exporter"
	Exporter          = "exporter"
	DefaultSocketPath = "./ipc.state.report"
)

var LogPathMap = map[string]string{
	"darwin":  "/tmp/" + NameSpace + ".log",
	"linux":   "/var/log/gse/" + NameSpace + ".log",
	"windows": "C:\\gse\\logs\\" + NameSpace + ".log",
}

var LogLevelMap = map[string]logrus.Level{
	"ERROR": logrus.ErrorLevel,
	"WARN":  logrus.WarnLevel,
	"INFO":  logrus.InfoLevel,
	"DEBUG": logrus.DebugLevel,
}

var (
	DefaultLogMaxSize    = 10   // 每个日志文件最大10MB
	DefaultLogMaxBackups = 3    // 保留最近的3个日志文件
	DefaultLogMaxAge     = 7    // 保留最近7天的日志
	DefaultLogCompress   = true // 是否压缩旧日志

	DefaultPeriod = time.Minute * 5
)
