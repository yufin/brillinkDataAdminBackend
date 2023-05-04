package task

import (
	"encoding/json"
	log "github.com/go-admin-team/go-admin-core/logger"
	"go-admin/app/rskc/models"
	"go-admin/app/rskc/service"
	"go-admin/app/rskc/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
)

type SyncTags4TradesTask struct {
}

func (t SyncTags4TradesTask) Exec(args interface{}) error {
	return nil
}

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
	for _, content := range contentList {
		err = ParseTradesDetailAndInsert(content.OriginJsonContent)
		if err != nil {
			log.Errorf("Task SyncTags4Trades func-ParseTradesDetailAndInsert Failed:%s \r\n", err)
			return err
		}
	}

	// delete all record in trades_detail if any by contentId,
	// get customer, supplier list
	// insert base field to trades_detail with status_code 1

	// query trades_detail by status_code = 1, iter rows check if enterprise_name is in table enterprise with full data collected.
	// if collected set statusCode=3, if not sync to waitList set statusCode=2

	// query statusCode=3, match and insert tag&info field , then set statusCode to 4

	// query

	//
}

// ParseTradesDetailAndInsert process of sync tradesDetail from content
func ParseTradesDetailAndInsert(contentJsonStr string) error {
	var contentMap map[string]any
	err := json.Unmarshal([]byte(contentJsonStr), &contentMap)
	if err != nil {
		log.Errorf("Task SyncTags4Trades func-ParseTradesDetailAndInsert Failed:%s \r\n", err)
		return nil
	}

	for key := range []string{"customerDetail_12", "customerDetail_24"} {
		tradeDetail := contentMap[key].([]map[string]any)

	}

	// todo: parse trades detail from content and insert with statusCode = 1, then update content table with statusCode=2 by contentId

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
