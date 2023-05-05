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
)

type SyncTradesDetailTask struct {
}

func (t SyncTradesDetailTask) Exec(args interface{}) error {
	return SyncTags4Trades()
}

// content-statusCode: 1:待解析录入其他表,2:解析并录入完成,3.数据匹配并录入完成
// tradeDetail-statusCode: 1.待确认企业数据已采集，2.待采集，已经同步至waitList, 3.采集完成, 4.匹配并录入完成

// SyncTags4Trades 同步tags4trades
// @Description end with statusCode: content-2, tradeDetail-1
func SyncTags4Trades() error {
	// check the tables where statusCode = 1
	req := dto.OriginContentGetPageReq{
		Pagination: cDto.Pagination{PageSize: 100, PageIndex: 1},
		StatusCode: 1,
	}
	var count int64
	contentList := make([]models.OriginContent, 0)
	contentService := service.OriginContent{}
	p := actions.DataPermission{}
	err := contentService.GetPage(&req, &p, &contentList, &count)
	if err != nil {
		log.Errorf("Task SyncTags4Trades func-SyncTags4Trades Failed:%s \r\n", err)
		return err
	}

	// iter rows , parse content as jsonObject
	// delete all record in trades_detail if any by contentId,
	// get customer, supplier list
	// insert base field to trades_detail with status_code 1
	for _, content := range contentList {
		err = ParseTradesDetailAndInsert(content.OriginJsonContent, content.ContentId, &p)
		if err != nil {
			log.Errorf("Task SyncTags4Trades func-ParseTradesDetailAndInsert Failed:%s \r\n", err)
			return err
		}
	}

	// query trades_detail by status_code = 1, iter rows check if enterprise_name is in table enterprise with full data collected.
	// if collected set statusCode=3, if not sync to waitList set statusCode=2

	// query statusCode=3, match and insert tag&info field , then set statusCode to 4

	return err
}

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
func ParseTradesDetailAndInsert(contentJsonStr string, contentId string, p *actions.DataPermission) error {
	tradeDetailService := service.RskcTradesDetail{}
	// delete all records in trades_detail if any by contentId
	func() {
		deleteList := make([]models.RskcTradesDetail, 0)
		var count int64
		err := tradeDetailService.GetPage(&dto.RskcTradesDetailGetPageReq{
			Pagination: cDto.Pagination{PageSize: 1000, PageIndex: 1},
			ContentId:  contentId}, p, &deleteList, &count)
		if err != nil {
			log.Errorf("Task SyncTags4Trades func-ParseTradesDetailAndInsert Failed:%s \r\n", err)
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
				log.Errorf("Task SyncTags4Trades func-ParseTradesDetailAndInsert Failed:%s \r\n", err)
				return
			}
		}
	}()

	var contentMap map[string]any
	err := json.Unmarshal([]byte(contentJsonStr), &contentMap)
	if err != nil {
		log.Errorf("Task SyncTags4Trades func-ParseTradesDetailAndInsert Failed:%s \r\n", err)
		return nil
	}

	for _, key := range tradesDetailKeyList() {
		tradeDetail := contentMap[key].([]map[string]any)
		for _, trade := range tradeDetail {
			insertReq := dto.RskcTradesDetailInsertReq{
				ContentId:      contentId,
				EnterpriseName: trade[contentEnterpriseNameKeyDict()[key]].(string),
				CommodityRatio: trade["COMMODITY_RATIO"].(string),
				CommodityName:  trade["COMMODITY_NAME"].(string),
				RatioAmountTax: trade["RATIO_AMOUNT_TAX"].(string),
				SumAmountTax:   trade["SUM_AMOUNT_TAX"].(string),
				DetailType:     detailTypeDict()[key],
				StatusCode:     1,
				ControlBy:      cModels.ControlBy{},
			}
			if err = tradeDetailService.Insert(&insertReq); err != nil {
				log.Errorf("Task SyncTags4Trades func-ParseTradesDetailAndInsert Failed:%s \r\n", err)
				return err
			}
		}
	}
	// finish sync tradesDetail, update content statusCode = 2
	contentService := service.OriginContent{}
	err = contentService.Update(&dto.OriginContentUpdateReq{ContentId: contentId, StatusCode: 2}, p)
	if err != nil {
		log.Errorf("Task SyncTags4Trades func-ParseTradesDetailAndInsert Failed:%s \r\n", err)
		return err
	}
	return nil
}

func CheckEnterpriseCollection() error {
	// todo: check enterprise_name in table enterprise with full data collected if not exist: sync to waitList, set statusCode=2 if collected set statusCode=3
	return nil
}

func SyncTradesDetailExtraInfo() error {
	// todo : query trades_details with contentId check if any row with statusCode not in (4), if not, update content table with statusCode=4 by contentId
	return nil
}
