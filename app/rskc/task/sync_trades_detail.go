package task

import (
	"encoding/json"
	log "github.com/go-admin-team/go-admin-core/logger"
	"go-admin/app/rskc/models"
	"go-admin/app/rskc/service"
	"go-admin/app/rskc/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
	cModels "go-admin/common/models"
	"strings"
)

// content-statusCode: 1:待解析录入其他表,2:解析并录入完成,3.数据匹配并录入完成
// tradeDetail-statusCode: 1.待确认企业数据已采集，2.待采集，已经同步至waitList, 3.确认采集完成, 4.匹配并录入完成

// SyncTradesDetail 同步tags4trades
// @Description end with statusCode: content-2, tradeDetail-1
func SyncTradesDetail(s *service.RskcTradesDetail, sContent *service.RskcOriginContent, p *actions.DataPermission) error {
	// check the tables where statusCode = 1
	req := dto.RskcOriginContentGetPageReq{
		Pagination: cDto.Pagination{PageSize: 100, PageIndex: 1},
		StatusCode: 1,
	}
	var count int64
	contentList := make([]models.RskcOriginContent, 0)
	err := sContent.GetPage(&req, p, &contentList, &count)
	if err != nil {
		log.Errorf("Task SyncTradesDetail func-SyncTradesDetail Failed:%s \r\n", err)
		return err
	}

	// iter rows , parse content as jsonObject
	// delete all record in trades_detail if any by contentId,
	// get customer, supplier list
	// insert base field to trades_detail with status_code 1
	for _, content := range contentList {
		err = ParseTradesDetailAndInsert(content.Content, content.Id, content.Id, s, sContent, p)
		if err != nil {
			log.Errorf("Task SyncTradesDetail func-ParseTradesDetailAndInsert Failed:%s \r\n", err)
			return err
		}
	}

	// query trades_detail by status_code = 1, iter rows check if enterprise_name is in table enterprise with full data collected.
	// if collected set statusCode=3, if not sync to waitList set statusCode=2

	// query statusCode=3, match and insert tag&info field , then set statusCode to 4

	return err
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

// ParseTradesDetailAndInsert process of sync tradesDetail from content
// @describe:parse trades detail from content and insert with statusCode = 1, then update content table with statusCode=2 by contentId
func ParseTradesDetailAndInsert(contentJsonStr string, Id int64, contentId int64, tradeDetailService *service.RskcTradesDetail, contentService *service.RskcOriginContent, p *actions.DataPermission) error {
	// delete all records in trades_detail if any by contentId
	func() {
		deleteList := make([]models.RskcTradesDetail, 0)
		var count int64
		err := tradeDetailService.GetPage(&dto.RskcTradesDetailGetPageReq{
			Pagination: cDto.Pagination{PageSize: 1000, PageIndex: 1},
			ContentId:  contentId}, p, &deleteList, &count)
		if err != nil {
			log.Errorf("Task SyncTradesDetail func-ParseTradesDetailAndInsert Failed:%s \r\n", err)
			return
		}
		if len(deleteList) > 0 {
			var ids []int64
			for _, v := range deleteList {
				ids = append(ids, v.Model.Id)
			}
			delReq := dto.RskcTradesDetailDeleteReq{
				Ids: ids,
			}
			err = tradeDetailService.Remove(&delReq, p)
			if err != nil {
				log.Errorf("Task SyncTradesDetail func-ParseTradesDetailAndInsert Failed:%s \r\n", err)
				return
			}
		}
	}()

	var contentMap map[string]any
	err := json.Unmarshal([]byte(contentJsonStr), &contentMap)
	if err != nil {
		log.Errorf("Task SyncTradesDetail func-ParseTradesDetailAndInsert Failed:%s \r\n", err)
		return nil
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
				if err = tradeDetailService.Insert(&insertReq); err != nil {
					log.Errorf("Task SyncTradesDetail func-ParseTradesDetailAndInsert Failed:%s \r\n", err)
					return err
				}
			}

			//exists := func(q *dto.RskcTradesDetailInsertReq, sd *service.RskcTradesDetail, pe *actions.DataPermission) bool {
			//	getPageReq := dto.RskcTradesDetailGetPageReq{
			//		ContentId:      q.ContentId,
			//		EnterpriseName: q.EnterpriseName,
			//		DetailType:     q.DetailType,
			//	}
			//	var countExist int64
			//	var existList []models.RskcTradesDetail
			//	err := sd.GetPage(&getPageReq, pe, &existList, &countExist)
			//	if err != nil {
			//		log.Errorf("Task SyncTradesDetail func-ParseTradesDetailAndInsert Failed:%s \r\n", err)
			//		return false
			//	}
			//	return countExist > 0
			//}(&insertReq, tradeDetailService, p)
			//
			//if !exists {
			//	if err = tradeDetailService.Insert(&insertReq); err != nil {
			//		log.Errorf("Task SyncTradesDetail func-ParseTradesDetailAndInsert Failed:%s \r\n", err)
			//		return err
			//	}
			//}

		}
	}
	// finish sync tradesDetail, update content statusCode = 2
	contentStatusCodeQ := dto.RskcOriginContentUpdateReq{Id: Id, StatusCode: 2}
	contentStatusCodeQ.SetUpdateBy(0)
	err = contentService.Update(&contentStatusCodeQ, p)
	if err != nil {
		log.Errorf("Task SyncTradesDetail func-ParseTradesDetailAndInsert Failed:%s \r\n", err)
		return err
	}
	return nil
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

func CheckEnterpriseCollection() error {
	// todo: check enterprise_name in table enterprise with full data collected if not exist: sync to waitList, set statusCode=2 if collected set statusCode=3
	return nil
}

func SyncTradesDetailExtraInfo() error {
	// todo : query trades_details with contentId check if any row with statusCode not in (4), if not, update content table with statusCode=4 by contentId
	return nil
}
