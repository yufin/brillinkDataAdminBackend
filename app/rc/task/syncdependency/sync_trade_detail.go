package syncdependency

import (
	"encoding/json"
	"github.com/go-admin-team/go-admin-core/sdk"
	"go-admin/app/rc/models"
	"go-admin/app/rc/service/dto"
	cModels "go-admin/common/models"
	"go-admin/pkg/natsclient"
	"strings"
)

type syncTradeDetailProcess struct {
}

func (t syncTradeDetailProcess) Process(contentId int64) error {
	// get content by contentId
	var tbContent models.RcOriginContent
	var dataContent models.RcOriginContent
	db := sdk.Runtime.GetDbByKey(tbContent.TableName())
	err := db.Model(&tbContent).First(&dataContent, contentId).Error
	if err != nil {
		return err
	}

	// pub subject enterprise ident msg to sync waitList task
	if err := t.pubEnterpriseIdentMsg(dataContent.EnterpriseName, dataContent.UscId); err != nil {
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

	for _, key := range t.tradesDetailKeyList() {
		entReport := contentMap[ReportDataKey].(map[string]any)
		tradeDetail := entReport[key].([]any)
		for _, trade := range tradeDetail {
			objectName := t.verifyMapField(trade.(map[string]any), t.contentEnterpriseNameKeyDict()[key])
			// if name not in []string{"其他", "合计"}
			if !strings.Contains(objectName, "其他") && !strings.Contains(objectName, "合计") {

				insertReq := dto.RcTradesDetailInsertReq{
					ContentId:      contentId,
					EnterpriseName: objectName,
					CommodityRatio: t.verifyMapField(trade.(map[string]any), "COMMODITY_RATIO"),
					CommodityName:  t.verifyMapField(trade.(map[string]any), "COMMODITY_NAME"),
					RatioAmountTax: t.verifyMapField(trade.(map[string]any), "RATIO_AMOUNT_TAX"),
					SumAmountTax:   t.verifyMapField(trade.(map[string]any), "SUM_AMOUNT_TAX"),
					DetailType:     t.detailTypeDict()[key],
					StatusCode:     1,
					ControlBy:      cModels.ControlBy{},
				}
				var dataTrade models.RcTradesDetail
				insertReq.Generate(&dataTrade)
				if err = dbTrades.Create(&dataTrade).Error; err != nil {
					return err
				}
				if err := t.pubEnterpriseIdentMsg(insertReq.EnterpriseName, ""); err != nil {
					return err
				}

			}
		}
	}
	return nil
}

func (syncTradeDetailProcess) pubEnterpriseIdentMsg(enterpriseName string, uscId string) error {
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

func (syncTradeDetailProcess) tradesDetailKeyList() []string {
	return []string{
		"customerDetail_12",
		"customerDetail_24",
		"supplierRanking_12",
		"supplierRanking_24",
	}
}

func (syncTradeDetailProcess) detailTypeDict() map[string]int {
	return map[string]int{
		"customerDetail_12":  1,
		"customerDetail_24":  2,
		"supplierRanking_12": 3,
		"supplierRanking_24": 4,
	}
}

func (syncTradeDetailProcess) contentEnterpriseNameKeyDict() map[string]string {
	return map[string]string{
		"customerDetail_12":  "PURCHASER_NAME",
		"customerDetail_24":  "PURCHASER_NAME",
		"supplierRanking_12": "SALES_NAME",
		"supplierRanking_24": "SALES_NAME",
	}
}

func (syncTradeDetailProcess) verifyMapField(m map[string]any, key string) string {
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
