package syncdependency

import (
	"encoding/json"
	log "github.com/go-admin-team/go-admin-core/logger"
	"github.com/go-admin-team/go-admin-core/sdk"
	"github.com/nats-io/nats.go"
	"github.com/pkg/errors"
	sModels "go-admin/app/spider/models"
	sDto "go-admin/app/spider/service/dto"
	"go-admin/pkg/natsclient"
	"gorm.io/gorm"
	"sync/atomic"
	"time"
)

var syncWaitListRunning int32

type SyncWaitListTask struct {
}

func (t SyncWaitListTask) Exec(arg interface{}) error {
	if atomic.LoadInt32(&syncWaitListRunning) == 1 {
		log.Info("SyncWaitListTask任务已经在执行中，跳过本次调度")
		return nil
	}
	atomic.StoreInt32(&syncWaitListRunning, 1)
	defer atomic.StoreInt32(&syncWaitListRunning, 0)

	return t.Process()
}

func (t SyncWaitListTask) Process() error {
	// TODO: add log at this layer
	for {
		msgs, err := natsclient.SubTradeNew.Fetch(1, nats.MaxWait(5*time.Second))
		if err != nil {
			if errors.Is(err, nats.ErrTimeout) {
				return nil
			} else {
				return err
			}
		}
		for _, msg := range msgs {
			m, err := t.parseEnterpriseIdentMsg(msg)
			if err != nil {
				return err
			}
			err = t.syncWaitListFromMsg(m)
			if err != nil {
				return err
			} else {
				err = msg.AckSync()
				if err != nil {
					return err
				}
			}
		}
	}
}

func (t SyncWaitListTask) parseEnterpriseIdentMsg(msg *nats.Msg) (map[string]string, error) {
	var m map[string]string
	err := json.Unmarshal(msg.Data, &m)
	if err != nil {
		return nil, err
	}
	_, ok := m["enterprise_name"]
	if !ok {
		return nil, errors.Errorf("key enterprise_name not found in msg: %s", string(msg.Data))
	}
	_, ok = m["usc_id"]
	if !ok {
		return nil, errors.Errorf("key usc_id not found in msg: %s", string(msg.Data))
	}
	return m, nil
}

func (t SyncWaitListTask) syncWaitListFromMsg(enterpriseIdentMap map[string]string) error {

	enterpriseName := enterpriseIdentMap["enterprise_name"]
	uscIdInput := enterpriseIdentMap["usc_id"]

	var tbWait sModels.EnterpriseWaitList
	dbWait := sdk.Runtime.GetDbByKey(tbWait.TableName())
	var count int64
	err := dbWait.Model(&tbWait).
		Where("enterprise_name = ?", enterpriseName).
		Count(&count).Error
	if err != nil {
		return err
	}

	if count == 0 {
		// query enterprise_info by enterprise_name, if exists, get UscId and insert with StatusCode=3
		var tbInfo sModels.EnterpriseInfo
		dbInfo := sdk.Runtime.GetDbByKey(tbInfo.TableName())

		var queryCond = "enterprise_title = ?"
		var queryArg = enterpriseName
		if len(uscIdInput) == 18 {
			queryCond = "usc_id = ?"
			queryArg = uscIdInput
		}
		err = dbInfo.Model(&tbInfo).
			Where(queryCond, queryArg).
			Order("updated_at desc").
			First(&tbInfo).
			Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		var statusCode = 1
		var uscId string
		if len(uscIdInput) == 18 {
			uscId = uscIdInput
		}

		if tbInfo.InfoId != 0 {
			statusCode = 3
			if uscId == "" {
				// 已匹配info,uscIdInput非空
				uscId = tbInfo.UscId
				statusCode = 3
			}
		} else {
			if uscId != "" {
				// 未匹配info,uscIdInput非空
				statusCode = 2 // wait for data collection
			}
		}

		// insert into wait_list
		insertReq := sDto.EnterpriseWaitListInsertReq{
			EnterpriseName: enterpriseName,
			UscId:          uscId,
			Priority:       9,
			StatusCode:     statusCode,
		}
		var dataInsert sModels.EnterpriseWaitList
		insertReq.Generate(&dataInsert)
		err = dbWait.Model(&tbWait).Create(&dataInsert).Error
		if err != nil {
			return err
		}
	}
	return nil
}
