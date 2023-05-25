package task

import (
	"encoding/binary"
	"github.com/go-admin-team/go-admin-core/sdk"
	"github.com/nats-io/nats.go"
	"github.com/pkg/errors"
	"go-admin/app/rskc/models"
	sModels "go-admin/app/spider/models"
	sDto "go-admin/app/spider/service/dto"
	"go-admin/common/natsclient"
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
			tradesId := int64(binary.BigEndian.Uint64(msg.Data))
			err := syncWaitListFromTrades(tradesId)
			if err != nil {
				return err
			} else {
				msg.AckSync()
			}
		}
	}
}

func syncWaitListFromTrades(tradesId int64) error {
	var tbTrades models.RskcTradesDetail
	dbTrades := sdk.Runtime.GetDbByKey(tbTrades.TableName())

	var dataTrades models.RskcTradesDetail

	err := dbTrades.Model(&tbTrades).First(&dataTrades, tradesId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
	}

	var tbWait sModels.EnterpriseWaitList
	//var dataWait sModels.EnterpriseWaitList
	dbWait := sdk.Runtime.GetDbByKey(tbTrades.TableName())
	var count int64
	err = dbWait.Model(&tbWait).
		Where("enterprise_name = ?", dataTrades.EnterpriseName).
		Count(&count).Error
	if err != nil {
		return err
	}
	if count == 0 {
		// query enterprise_info by enterprise_name, if exist, get UscId and insert with StatusCode=2
		var tbInfo sModels.EnterpriseInfo
		dbInfo := sdk.Runtime.GetDbByKey(tbInfo.TableName())
		err = dbInfo.Model(&tbInfo).
			Where("enterprise_title = ?", dataTrades.EnterpriseName).
			Order("updated_at desc").
			First(&tbInfo).
			Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		var uscId string
		var statusCode = 1
		if tbInfo.InfoId != 0 && tbInfo.UscId != "" {
			uscId = tbInfo.UscId
			statusCode = 3
		}

		// insert into wait_list
		insertReq := sDto.EnterpriseWaitListInsertReq{
			EnterpriseName: dataTrades.EnterpriseName,
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
