package task

import (
	"encoding/binary"
	"encoding/json"
	"github.com/go-admin-team/go-admin-core/sdk"
	"github.com/nats-io/nats.go"
	"go-admin/app/rskc/models"
	"go-admin/app/rskc/service/dto"
	cModels "go-admin/common/models"
	"go-admin/common/natsclient"
	"strings"
	"sync"
	"time"
)

type SyncTradesDetailTask struct {
}

var mutexStdt = &sync.Mutex{}

func (t SyncTradesDetailTask) Exec(arg interface{}) error {
	mutexStdt.Lock()
	defer mutexStdt.Unlock()
	return pullContentNew()
}

func pullContentNew() error {
	// TODO: add log at this layer
	for {
		// get total msg count by subscriber
		//totalPending, _, err := natsclient.SubContentNew.Pending()
		//if err == nil {
		//	fmt.Println("SubContentNew msg totalPending:", totalPending)
		//}

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
			err := parseContentToDetails(contentId)
			if err != nil {
				return err
			} else {
				if err := msg.AckSync(); err != nil {
					return err
				}
			}
		}
	}
}

func markContentAsCompleteAsync(contentId int64) error {
	// set status code to 2, which means all dependency data collected.
	var data models.RskcOriginContent
	dbContent := sdk.Runtime.GetDbByKey(data.TableName())
	err := dbContent.Model(&data).Where("id = ?", contentId).Update("status_code", 1).Error
	if err != nil {
		return err
	}
	return nil
}

func parseContentToDetails(contentId int64) error {
	// get content by contentId
	var tbContent models.RskcOriginContent
	var dataContent models.RskcOriginContent
	db := sdk.Runtime.GetDbByKey(tbContent.TableName())
	err := db.Model(&tbContent).First(&dataContent, contentId).Error
	if err != nil {
		return err
	}

	// delete all row by contentId to avoid duplication
	var tbTrades models.RskcTradesDetail
	dbTrades := sdk.Runtime.GetDbByKey(tbTrades.TableName())
	// with unscoped, db exec hard delete instead of soft delete.
	err = dbTrades.
		Unscoped().
		Where("content_id = ?", contentId).
		Delete(&models.RskcTradesDetail{}).
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

				insertReq := dto.RskcTradesDetailInsertReq{
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
				var dataTrade models.RskcTradesDetail
				insertReq.Generate(&dataTrade)
				if err = dbTrades.Create(&dataTrade).Error; err != nil {
					return err
				} else {
					if err := pubTradeId(dataTrade.Id); err != nil {
						return err
					}
				}
			}
		}
	}

	if err := markContentAsCompleteAsync(contentId); err != nil {
		return err
	}

	return nil
}

func pubTradeId(tradeId int64) error {
	msg := make([]byte, 8)
	binary.BigEndian.PutUint64(msg, uint64(tradeId))
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
