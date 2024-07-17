package main

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"winstack_collect/collector"
	"winstack_collect/config"
	"winstack_collect/socket"
)

type Server interface {
	Start(ctx context.Context)
}

type Service struct {
	wg         sync.WaitGroup
	ctx        context.Context
	cancel     context.CancelFunc
	c          *config.Config
	gse        *socket.GseClient     // gse implement Server interface
	collectMgr *collector.Controller // collectMgr implement Server interface
}

func NewService(c *config.Config) *Service {
	ctx, cancel := context.WithCancel(context.Background())
	return &Service{
		wg:         sync.WaitGroup{},
		ctx:        ctx,
		cancel:     cancel,
		c:          c,
		gse:        socket.GetGseClient(),
		collectMgr: collector.NewControllerMgr(c),
	}
}

func (svc *Service) Run() {
	// 非测试模式下
	if !svc.c.TestModel {
		svc.runServer(svc.gse)
	}
	// 启动云监控采集服务
	svc.runServer(svc.collectMgr)

	// 阻塞并且监听中断信号
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	svc.wg.Wait()
	svc.cancel()
}

func (svc *Service) runServer(server Server) {
	if server != nil {
		svc.wg.Add(1)
		go func(s Server) {
			defer svc.wg.Done()
			server.Start(svc.ctx)
		}(server)
	}
}
