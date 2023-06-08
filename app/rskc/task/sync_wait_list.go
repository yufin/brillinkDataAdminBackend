package task

import (
	"encoding/json"
	"github.com/go-admin-team/go-admin-core/sdk"
	"github.com/nats-io/nats.go"
	"github.com/pkg/errors"
	sModels "go-admin/app/spider/models"
	sDto "go-admin/app/spider/service/dto"
	"go-admin/pkg/natsclient"
	"gorm.io/gorm"
	"sync"
	"time"
)

type SyncWaitListTask struct {
}

var mutexEwl = &sync.Mutex{}

func (t SyncWaitListTask) Exec(arg interface{}) error {
	mutexEwl.Lock()
	defer mutexEwl.Unlock()
	return pullTradesNew()
}

func pullTradesNew() error {
	// TODO: add log at this layer
	for {
		msgs, err := natsclient.SubTradeNew.Fetch(1, nats.MaxWait(5*time.Second))
		if err != nil {
			if err == nats.ErrTimeout {
				return nil
			} else {
				return err
			}
		}
		for _, msg := range msgs {
			m, err := parseEnterpriseIdentMsg(msg)
			if err != nil {
				return err
			}
			err = syncWaitListFromMsg(m)
			if err != nil {
				return err
			} else {
				msg.AckSync()
			}
		}
	}
}

func parseEnterpriseIdentMsg(msg *nats.Msg) (map[string]string, error) {
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

func syncWaitListFromMsg(enterpriseIdentMap map[string]string) error {

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
