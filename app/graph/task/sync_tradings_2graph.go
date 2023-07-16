package task

import (
	log "github.com/go-admin-team/go-admin-core/logger"
	"sync/atomic"
)

type SyncTrades2Graph struct {
}

var syncTradesRunning int32

func (t SyncTrades2Graph) Exec(args interface{}) error {
	if atomic.LoadInt32(&syncTradesRunning) == 1 {
		log.Info("SyncGraph任务已经在执行中，跳过本次调度")
		return nil
	}
	atomic.StoreInt32(&syncTradesRunning, 1)
	defer atomic.StoreInt32(&syncTradesRunning, 0)

	return nil
}
