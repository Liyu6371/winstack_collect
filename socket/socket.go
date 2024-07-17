package socket

import (
	"context"
	"github.com/TencentBlueKing/bkmonitor-datalink/pkg/libgse/gse"
	"os"
	"winstack_collect/common"
	"winstack_collect/config"
	"winstack_collect/logger"
)

var defaultGseClient *gse.GseSimpleClient

func InitSocket() {
	defaultGseClient = gse.NewGseSimpleClient()
	c := config.GetConfig()
	SocketPath := c.Socket.SocketPath
	// 设置默认的socket文件位置
	if c.Socket.SocketPath != "" {
		defaultGseClient.SetAgentHost(SocketPath)
	} else {
		defaultGseClient.SetAgentHost(common.DefaultSocketPath)
		logger.Warnf("use default socket path: %s", common.DefaultSocketPath)
	}
}

// Start 启动 GSE 通信
func Start(ctx context.Context) {
	err := defaultGseClient.Start()
	if err != nil {
		logger.Errorf("unable to start default gse, %s\n", err)
		os.Exit(1)
	}
	logger.Debugln("successfully started default gse")
	// 启动一个协程监听终止事件
	go func() {
		<-ctx.Done()
		defaultGseClient.Close()
		logger.Infoln("successfully closed default gse")
	}()
}

// Send 发送数据到 GSE
func Send(msg gse.GseMsg) {
	if err := defaultGseClient.Send(msg); err != nil {
		logger.Errorf("unable to send message to gse, %s\n", err)
	}
	logger.Debugln("successfully sent message to gse")
}
