package syncdependency

import (
	"encoding/binary"
	"fmt"
	log "github.com/go-admin-team/go-admin-core/logger"
	"github.com/go-admin-team/go-admin-core/sdk"
	"github.com/nats-io/nats.go"
	"github.com/pkg/errors"
	"go-admin/app/rc/models"
	"go-admin/pkg/natsclient"
	"gorm.io/gorm"
	"sync"
	"sync/atomic"
	"time"
)

type DependencyTableSyncTask struct {
}

var dtstRunning int32

type DependencySyncProcess interface {
	Process(contentId int64) error
}

func (t DependencyTableSyncTask) Exec(arg interface{}) error {
	if atomic.LoadInt32(&dtstRunning) == 1 {
		log.Info("DependencyTableSyncTask任务已经在执行中，跳过本次调度")
		return nil
	}
	atomic.StoreInt32(&dtstRunning, 1)
	defer atomic.StoreInt32(&dtstRunning, 0)

	return t.pipeline()
}

func (t DependencyTableSyncTask) pipeline() error {
	// set sync process for task
	processes := []DependencySyncProcess{
		sellingStaSyncProcess{},
		syncTradeDetailProcess{},
		decisionParamSyncProcess{},
		riskIndexSyncProcess{},
	}
	concurrencyLimit := 3

	for {
		// get total msg count by subscriber
		totalPending, _, err := natsclient.SubContentNew.Pending()
		if err == nil {
			fmt.Println("DependencyTableSyncTask msg totalPending:", totalPending)
		}

		msgs, err := natsclient.SubContentNew.Fetch(1, nats.MaxWait(5*time.Second))
		if err != nil {
			if err == nats.ErrTimeout {
				return nil
			} else {
				return err
			}
		}

		for _, msg := range msgs {
			contentId := int64(binary.BigEndian.Uint64(msg.Data))

			exists, err := t.CheckContentIdExist(contentId)
			if err != nil {
				return err
			} else {
				if !exists {
					if err := msg.AckSync(); err != nil {
						return err
					}
					break
				}
			}

			log.Infof("开始解析并同步依赖数据: contentId = %d\r\n", contentId)

			// control concurrency
			var wg sync.WaitGroup
			limitCh := make(chan struct{}, concurrencyLimit)
			errCh := make(chan error)
			done := make(chan struct{})
			for i, _ := range processes {
				limitCh <- struct{}{}
				wg.Add(1)
				i := i
				go func(index int) {
					defer wg.Done()
					err := processes[i].Process(contentId)
					if err != nil {
						log.Errorf("sync dependencies process error: %s, contentId=%s\r\n", err, contentId)
						errCh <- err
					}
					<-limitCh
				}(i)
			}

			go func() {
				wg.Wait()
				close(done)
			}()

			select {
			case <-done:
				if err := t.markContentAsCompleteAsync(contentId); err != nil {
					return err
				}
				if err := msg.AckSync(); err != nil {
					return err
				}
			case err := <-errCh:
				return err
			}

			//var err1, err2, err3 error
			//var wg sync.WaitGroup
			//wg.Add(1)
			//go func() {
			//	defer wg.Done()
			//	syd := syncTradeDetailProcess{}
			//	err1 = syd.Process(contentId)
			//}()
			//wg.Add(1)
			//go func() {
			//	defer wg.Done()
			//	sss := sellingStaSyncProcess{}
			//	err2 = sss.Process(contentId)
			//}()
			//wg.Add(1)
			//go func() {
			//	defer wg.Done()
			//	dps := decisionParamSyncProcess{}
			//	err3 = dps.Process(contentId)
			//}()
			//wg.Wait()
			//if err1 != nil {
			//	return err
			//}
			//if err2 != nil {
			//	return err
			//}
			//if err3 != nil {
			//	return err
			//}
			//if err := t.markContentAsCompleteAsync(contentId); err != nil {
			//	return err
			//}
			//if err := msg.AckSync(); err != nil {
			//	return err
			//}
		}
	}
}

func (t DependencyTableSyncTask) markContentAsCompleteAsync(contentId int64) error {
	// set status code to 2, which means all dependency data collected.
	var data models.RcOriginContent
	dbContent := sdk.Runtime.GetDbByKey(data.TableName())
	err := dbContent.Model(&data).Where("id = ?", contentId).Update("status_code", 1).Error
	if err != nil {
		return err
	}
	return nil
}

func (t DependencyTableSyncTask) CheckContentIdExist(contentId int64) (bool, error) {
	var data models.RcOriginContent
	db := sdk.Runtime.GetDbByKey(data.TableName())
	err := db.Model(&data).Where("id = ?", contentId).First(&data).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
