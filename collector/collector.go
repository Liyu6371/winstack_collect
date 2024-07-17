package collector

import (
	"context"
	"sync"
	"winstack_collect/collector/vmware"
	"winstack_collect/collector/winstack"
	"winstack_collect/config"
	"winstack_collect/logger"
)

// CollectService 定义接口类型
type CollectService interface {
	Run()
}

// Controller 数据采集服务管理
type Controller struct {
	wg              sync.WaitGroup
	WinStackService *winstack.WinStack
	VMWareService   *vmware.VMWare
}

func NewControllerMgr(c *config.Config) *Controller {
	mgr := &Controller{}
	if c.WinStackTask != nil {
		mgr.WinStackService = c.WinStackTask
	}
	return mgr
}

func (c *Controller) Start(ctx context.Context) {
	// 启动 WinStack 采集任务
	c.runService(c.WinStackService)
	// 启动 VMWare 采集任务
	c.runService(c.VMWareService)
	// 接收到上级传递的结束信号
	<-ctx.Done()
	// 等待所有的协程退出
	c.wg.Wait()
	logger.Info("close collector Server")
}

func (c *Controller) runService(service CollectService) {
	if service != nil {
		c.wg.Add(1)
		go func(s CollectService) {
			defer c.wg.Done()
			s.Run()
		}(service)
	}
}
