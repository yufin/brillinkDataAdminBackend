package task

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	log "github.com/go-admin-team/go-admin-core/logger"
	"github.com/go-admin-team/go-admin-core/sdk"
	"github.com/nats-io/nats.go"
	"github.com/pkg/errors"
	"go-admin/app/rc/models"
	"go-admin/app/rc/service/dto"
	cModels "go-admin/common/models"
	"go-admin/pkg/natsclient"
	"gorm.io/gorm"
	"strings"
	"sync"
	"time"
)

type SyncDependencyTableTask struct {
}

var mutexStdt = &sync.Mutex{}

func (t SyncDependencyTableTask) Exec(arg interface{}) error {
	mutexStdt.Lock()
	defer mutexStdt.Unlock()
	return pullContentNew()
}

func pullContentNew() error {
	// TODO: add log at this layer
	for {
		// get total msg count by subscriber
		totalPending, _, err := natsclient.SubContentNew.Pending()
		if err == nil {
			fmt.Println("SyncDependencyTableTask msg totalPending:", totalPending)
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

			exists, err := CheckContentIdExist(contentId)
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
			var err1, err2, err3 error
			var wg sync.WaitGroup
			wg.Add(1)
			go func() {
				defer wg.Done()
				err1 = parseContentToDetails(contentId)
			}()
			wg.Add(1)
			go func() {
				defer wg.Done()
				err2 = syncSellingStaFromContent(contentId)
			}()
			wg.Add(1)
			go func() {
				defer wg.Done()
				err3 = syncDecisionParamFromContent(contentId)
			}()
			wg.Wait()
			if err1 != nil {
				return err
			}
			if err2 != nil {
				return err
			}
			if err3 != nil {
				return err
			}
			if err := markContentAsCompleteAsync(contentId); err != nil {
				return err
			}
			if err := msg.AckSync(); err != nil {
				return err
			}
		}
	}
}

func markContentAsCompleteAsync(contentId int64) error {
	// set status code to 2, which means all dependency data collected.
	var data models.RcOriginContent
	dbContent := sdk.Runtime.GetDbByKey(data.TableName())
	err := dbContent.Model(&data).Where("id = ?", contentId).Update("status_code", 1).Error
	if err != nil {
		return err
	}
	return nil
}

func parseContentToDetails(contentId int64) error {

	// get content by contentId
	var tbContent models.RcOriginContent
	var dataContent models.RcOriginContent
	db := sdk.Runtime.GetDbByKey(tbContent.TableName())
	err := db.Model(&tbContent).First(&dataContent, contentId).Error
	if err != nil {
		return err
	}

	// pub subject enterprise ident msg to sync waitList task
	if err := pubEnterpriseIdentMsg(dataContent.EnterpriseName, dataContent.UscId); err != nil {
		return err
	}

	// delete all row by contentId to avoid duplication
	var tbTrades models.RcTradesDetail
	dbTrades := sdk.Runtime.GetDbByKey(tbTrades.TableName())
	// with unscoped, db exec hard delete instead of soft delete.
	err = dbTrades.
		Unscoped().
		Where("content_id = ?", contentId).
		Delete(&models.RcTradesDetail{}).
		Error
	if err != nil {
		return err
	}

	var contentMap map[string]any
	err = json.Unmarshal([]byte(dataContent.Content), &contentMap)
	if err != nil {
		return err
	}

	for _, key := range tradesDetailKeyList() {
		entReport := contentMap[ReportDataKey].(map[string]any)
		tradeDetail := entReport[key].([]any)
		for _, trade := range tradeDetail {
			objectName := verifyMapField(trade.(map[string]any), contentEnterpriseNameKeyDict()[key])
			// if name not in []string{"其他", "合计"}
			if !strings.Contains(objectName, "其他") && !strings.Contains(objectName, "合计") {

				insertReq := dto.RcTradesDetailInsertReq{
					ContentId:      contentId,
					EnterpriseName: objectName,
					CommodityRatio: verifyMapField(trade.(map[string]any), "COMMODITY_RATIO"),
					CommodityName:  verifyMapField(trade.(map[string]any), "COMMODITY_NAME"),
					RatioAmountTax: verifyMapField(trade.(map[string]any), "RATIO_AMOUNT_TAX"),
					SumAmountTax:   verifyMapField(trade.(map[string]any), "SUM_AMOUNT_TAX"),
					DetailType:     detailTypeDict()[key],
					StatusCode:     1,
					ControlBy:      cModels.ControlBy{},
				}
				var dataTrade models.RcTradesDetail
				insertReq.Generate(&dataTrade)
				if err = dbTrades.Create(&dataTrade).Error; err != nil {
					return err
				}
				if err := pubEnterpriseIdentMsg(insertReq.EnterpriseName, ""); err != nil {
					return err
				}

			}
		}
	}
	return nil
}

func pubEnterpriseIdentMsg(enterpriseName string, uscId string) error {
	m := map[string]string{
		"enterprise_name": enterpriseName,
		"usc_id":          uscId,
	}
	msg, _ := json.Marshal(m)
	_, err := natsclient.TaskJs.Publish(natsclient.TopicTradeNew, msg)
	if err != nil {
		return err
	}
	return nil
}

const ReportDataKey = "impExpEntReport"

func tradesDetailKeyList() []string {
	return []string{
		"customerDetail_12",
		"customerDetail_24",
		"supplierRanking_12",
		"supplierRanking_24",
	}
}

func detailTypeDict() map[string]int {
	return map[string]int{
		"customerDetail_12":  1,
		"customerDetail_24":  2,
		"supplierRanking_12": 3,
		"supplierRanking_24": 4,
	}
}

func contentEnterpriseNameKeyDict() map[string]string {
	return map[string]string{
		"customerDetail_12":  "PURCHASER_NAME",
		"customerDetail_24":  "PURCHASER_NAME",
		"supplierRanking_12": "SALES_NAME",
		"supplierRanking_24": "SALES_NAME",
	}
}

func verifyMapField(m map[string]any, key string) string {
	value, ok := m[key]
	if !ok {
		return "-"
	}
	str, isString := value.(string)
	if isString {
		return str
	}
	return "-"
}

func CheckContentIdExist(contentId int64) (bool, error) {
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
