package v3

import (
	"encoding/binary"
	log "github.com/go-admin-team/go-admin-core/logger"
	"github.com/go-admin-team/go-admin-core/sdk"
	"github.com/nats-io/nats.go"
	"github.com/pkg/errors"
	"go-admin/app/rc/models"
	"go-admin/pkg/natsclient"
	"sync/atomic"
	"time"
)

var rbv3Running int32

type reportCollator interface {
	Collating() error
	SetContent(content *[]byte, contentId int64)
}

type ReportBuilderV3 struct {
}

func (r ReportBuilderV3) Exec(arg interface{}) error {
	if atomic.LoadInt32(&rbv3Running) == 1 {
		log.Info("ReportBuilderV3任务已经在执行中，跳过本次调度")
		return nil
	}
	atomic.StoreInt32(&rbv3Running, 1)
	defer atomic.StoreInt32(&rbv3Running, 0)

	for {
		msgs, err := natsclient.SubContentProcessNew.Fetch(1, nats.MaxWait(5*time.Second))
		if err != nil {
			if err == nats.ErrTimeout {
				return nil
			} else {
				return errors.Wrap(err, "err at natsclient.SubContentProcessNew.Fetch")
			}
		}
		for _, msg := range msgs {
			contentId := int64(binary.BigEndian.Uint64(msg.Data))
			err := r.pipeline(contentId)
			if err != nil {
				return errors.Wrap(err, "err at r.Process")
			} else {
				err = msg.AckSync()
				if err != nil {
					return errors.Wrap(err, "err at msg.AckSync")
				}
			}
		}
	}
}

func (r ReportBuilderV3) pipeline(contentId int64) error {
	collators := []reportCollator{
		&ClaBusinessInfo{},
		&ClaRiskIndexes{},
		&ClaSubjCompanyTag{},
		&ClaBusinessPartnerDetail{},
		&ClaMonthlyTradingSta{},
		&ClaSellingDetailSummary{},
	}

	modelRoc := models.RcOriginContent{}
	db := sdk.Runtime.GetDbByKey(modelRoc.TableName())
	err := db.Model(&models.RcOriginContent{}).First(&modelRoc, contentId).Error
	if err != nil {
		return errors.Wrap(err, "err at db.Model(&models.RcOriginContent{}).First")
	}

	contentBytes := []byte(modelRoc.Content)

	for _, collator := range collators {
		collator.SetContent(&contentBytes, contentId)
		err := collator.Collating()
		if err != nil {
			return errors.Wrap(err, "err at collator.Collating")
		}
	}

	if err := r.saveProcessedContent(&contentBytes, contentId); err != nil {
		return errors.Wrap(err, "err at r.saveProcessedContent")
	}
	if err := r.updateStatusCodeOnSuccess(contentId); err != nil {
		return errors.Wrap(err, "err at r.updateStatusCodeOnSuccess")
	}
	return nil
}

func (r ReportBuilderV3) saveProcessedContent(content *[]byte, contentId int64) error {
	// TODO: FinishMe
	return nil
}

func (r ReportBuilderV3) updateStatusCodeOnSuccess(contentId int64) error {
	// TODO: FinishMe
	return nil
}
