package v3

import (
	"encoding/json"
	log "github.com/go-admin-team/go-admin-core/logger"
	"github.com/go-admin-team/go-admin-core/sdk"
	"github.com/pkg/errors"
	"go-admin/app/rc/models"
	"sync/atomic"
)

var rbv34TestRunning int32

type ReportBuilderV34Test struct {
}

func (r ReportBuilderV34Test) Exec(arg interface{}) error {
	if atomic.LoadInt32(&rbv34TestRunning) == 1 {
		log.Info("ReportBuilderV3任务已经在执行中，跳过本次调度")
		return nil
	}
	atomic.StoreInt32(&rbv34TestRunning, 1)
	defer atomic.StoreInt32(&rbv34TestRunning, 0)

	if err := r.pipeline(471122969607810198); err != nil {
		return errors.Wrap(err, "err at r.pipeline")
	}

	return nil
}

func (r ReportBuilderV34Test) pipeline(contentId int64) error {
	collators := []reportCollator{
		&ClaBusinessInfo{},
		//&ClaRiskIndexes{},
		//&ClaSubjCompanyTag{},
		//&ClaBusinessPartnerDetail{},
		//&ClaMonthlyTradingSta{},
		//&ClaRevenueDetail{},
		//&ClaPurchaseDetailSummary{},
		//&ClaSellingDetailSummary{},
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

	return nil
}

func (r ReportBuilderV34Test) DynamicProcess(contentId int64) (map[string]any, error) {
	collators := []reportCollator{
		&ClaBusinessInfo{},
		&ClaRiskIndexes{},
		&ClaSubjCompanyTag{},
		&ClaBusinessPartnerDetail{},
		&ClaMonthlyTradingSta{},
		&ClaRevenueDetail{},
		&ClaPurchaseDetailSummary{},
		&ClaSellingDetailSummary{},
		&ClaFinancialSummary{},
	}
	modelRoc := models.RcOriginContent{}
	db := sdk.Runtime.GetDbByKey(modelRoc.TableName())
	err := db.Model(&models.RcOriginContent{}).First(&modelRoc, contentId).Error
	if err != nil {
		return nil, errors.Wrap(err, "err at db.Model(&models.RcOriginContent{}).First")
	}

	contentBytes := []byte(modelRoc.Content)
	contentP := &contentBytes

	for _, collator := range collators {
		collator := collator
		collator.SetContent(contentP, contentId)
		err := collator.Collating()
		if err != nil {
			return nil, errors.Wrap(err, "err at collator.Collating")
		}
	}

	m := make(map[string]interface{})
	err = json.Unmarshal(*contentP, &m)
	if err != nil {
		return nil, errors.Wrap(err, "err at json.Unmarshal")
	}
	// contentBytes to map

	return m, nil
}
