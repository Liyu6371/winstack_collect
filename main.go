package main

import (
	"flag"
	"fmt"
	"github.com/natefinch/lumberjack"
	"os"
	"strings"
	"winstack_collect/common"
	"winstack_collect/config"
	"winstack_collect/logger"
	"winstack_collect/socket"
)

var (
	fileFlag    = flag.String("c", "", "config file path")
	versionFlag = flag.Bool("v", false, "show version")
)

func main() {
	flag.Parse()
	// 显示版本号
	if *versionFlag {
		fmt.Println(common.Version)
		os.Exit(0)
	}
	// 配置文件地址
	if *fileFlag == "" {
		fmt.Println("config file path is required")
		os.Exit(1)
	}
	c, err := config.ParseConfig(*fileFlag)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// 拿到配置开始定义 logger
	if c.Logger.Level != "" {
		if level, ok := common.LogLevelMap[strings.ToUpper(c.Logger.Level)]; ok {
			logger.SetLevel(level)
		}
	}
	if c.Logger.Path != "" {
		logger.SetOutput(&lumberjack.Logger{
			Filename:   c.Logger.Path,
			MaxSize:    common.DefaultLogMaxSize,
			MaxBackups: common.DefaultLogMaxBackups,
			MaxAge:     common.DefaultLogMaxAge,
			Compress:   common.DefaultLogCompress,
		})
	}
	// 初始化 socket 配置
	socket.InitSocket()
	// 初始化采集服务
	collectService := NewService(c)
	collectService.Run()
}
