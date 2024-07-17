package main

import (
	"flag"
	"fmt"
	"os"
	"winstack_collect/common"
	"winstack_collect/config"
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
	fmt.Println(c)
	//// 拿到配置开始定义 logger
	//if c.Logger.Level != "" {
	//	if level, ok := common.LogLevelMap[strings.ToUpper(c.Logger.Level)]; ok {
	//		logger.SetLevel(level)
	//	}
	//}
	//if c.Logger.Path != "" {
	//	logger.SetOutput(&lumberjack.Logger{
	//		Filename:   c.Logger.Path,
	//		MaxSize:    common.DefaultLogMaxSize,
	//		MaxBackups: common.DefaultLogMaxBackups,
	//		MaxAge:     common.DefaultLogMaxAge,
	//		Compress:   common.DefaultLogCompress,
	//	})
	//}
	//// 开始启动 socket 所有任务共用一个 socket 链接上报数据到 gse
	//// 若初始化 socket 失败也会导致程序直接退出
	//socket.InitSocket()
	//// ch 监听中断信号
	//ch := make(chan os.Signal, 1)
	//signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	//// 周期初始化
	//period := common.DefaultPeriod
	//if c.Period.CollectPeriod != "" {
	//	if p, err := time.ParseDuration(c.Period.CollectPeriod); err == nil {
	//		period = p
	//	}
	//}
	//ticker := time.NewTicker(period)
	//defer ticker.Stop()
	//
	//for {
	//	select {
	//	case <-ticker.C:
	//		fmt.Println("开始作业")
	//	case <-ch:
	//		fmt.Println("结束作业")
	//		return
	//	}
	//}
}
