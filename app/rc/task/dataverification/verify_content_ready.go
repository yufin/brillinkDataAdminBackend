package dataverification

import (
	"encoding/binary"
	"fmt"
	log "github.com/go-admin-team/go-admin-core/logger"
	"github.com/go-admin-team/go-admin-core/sdk"
	"github.com/go-admin-team/go-admin-core/sdk/pkg"
	"go-admin/app/rc/models"
	sModels "go-admin/app/spider/models"
	"go-admin/pkg/natsclient"
)

type VerifyContentReadyTask struct {
}

func (t VerifyContentReadyTask) Exec(arg interface{}) error {
	return iterContentIdtWaitForVerify()
}

// logic: query origin_content where status_code = 1, get contentId
// query trades_details by contentId, get enterprise_name
// query wait_list by enterprise_name with condition usc_id is not null and not "",or "-"
// if exist, get usc_id, check if exist in enterprise_info by usc_id,

// if all enterprise_name from trades_details exist, then mark the contentId as ready,
// generate process_content

type dataCollectionStatus struct {
	ContentId      int64  `json:"content_id"`
	ContentStatus  int    `json:"content_status"`
	EnterpriseName string `json:"enterprise_name"`
	UscId          string `json:"usc_id"`
	IdentStatus    int    `json:"ident_status"`
}

func iterContentIdtWaitForVerify() error {
	var tbContent models.RcOriginContent
	db := sdk.Runtime.GetDbByKey(tbContent.TableName())

	var unprocessedIds []int64
	err := db.Model(&models.RcOriginContent{}).
		Joins("LEFT JOIN rc_processed_content on rc_origin_content.id = rc_processed_content.content_id").
		//Where("rc_processed_content.content_id is NULL").
		Where("rc_origin_content.status_code = 1").
		Pluck("rc_origin_content.id", &unprocessedIds).
		Error
	if err != nil {
		return err
	}

	for _, contentId := range unprocessedIds {
		allPass, err := dataCollectionCheckByContentId(contentId)
		if err != nil {
			return err
		}
		if allPass {
			// pub msg to process content
			err = func() error {
				msg := make([]byte, 8)
				binary.BigEndian.PutUint64(msg, uint64(contentId))
				_, err := natsclient.TaskJs.Publish(natsclient.TopicContentToProcessNew, msg)
				return err
			}()
			if err != nil {
				return err
			}
			if err := markContentAsSend(contentId); err != nil {
				return err
			}
			log.Info(pkg.Blue(fmt.Sprintf("Report with contentId:%d all data checked, ready to Gen Report.", contentId)))
		}
	}
	return nil
}

func dataCollectionCheckByContentId(contentId int64) (bool, error) {
	var tbContent models.RcOriginContent
	db := sdk.Runtime.GetDbByKey(tbContent.TableName())

	var result []dataCollectionStatus
	db.Model(&tbContent).
		Select("rc_origin_content.id as content_id, rc_origin_content.status_code as content_status,"+
			" rc_trades_detail.enterprise_name as enterprise_name, enterprise_wait_list.usc_id as usc_id,"+
			" enterprise_wait_list.status_code as ident_status").
		Joins("right join rc_trades_detail on rc_origin_content.id = rc_trades_detail.content_id").
		Joins("left join enterprise_wait_list on rc_trades_detail.enterprise_name = enterprise_wait_list.enterprise_name").
		Where("content_id = ?", contentId).
		Scan(&result)

	var tbInfo sModels.EnterpriseInfo
	dbInfo := sdk.Runtime.GetDbByKey(tbInfo.TableName())
	for _, r := range result {
		if len(r.UscId) == 18 {
			// check if exist in enterprise_info by usc_id
			var countInfo int64
			err := dbInfo.Model(&tbInfo).Where("usc_id = ?", r.UscId).Count(&countInfo).Error
			if err != nil {
				return false, err
			}
			if countInfo == 0 {
				return false, nil
			}
		} else if r.IdentStatus != 9 {
			return false, nil
		}
	}

	// 检查OriginContent公司主体是否采集
	var contentInfo models.RcOriginContentInfo
	err := db.Model(&contentInfo).
		Where("id = ?", contentId).
		First(&contentInfo).
		Error
	if err != nil {
		return false, err
	}
	var countInfo int64
	err = dbInfo.Model(&tbInfo).Where("usc_id = ?", contentInfo.UscId).Count(&countInfo).Error
	if err != nil {
		return false, err
	}
	if countInfo == 0 {
		return false, nil
	}
	return true, nil
}

func markContentAsSend(contentId int64) error {
	// set status code to 2, which means all dependency data collected.
	var data models.RcOriginContent
	dbContent := sdk.Runtime.GetDbByKey(data.TableName())
	err := dbContent.Model(&data).Where("id = ?", contentId).Update("status_code", 2).Error
	if err != nil {
		return err
	}
	return nil
}

// debug sql: check integrality
//select t.*, ei.usc_id
//from
//(
//select rc_origin_content.id as content_id, rc_origin_content.status_code as content_status,
//rc_trades_detail.enterprise_name as enterprise_name, enterprise_wait_list.usc_id as usc_id,
//enterprise_wait_list.status_code as ident_status
//from rc_origin_content
//right join rc_trades_detail on rc_origin_content.id = rc_trades_detail.content_id
//left join enterprise_wait_list on rc_trades_detail.enterprise_name = enterprise_wait_list.enterprise_name) t
//left join enterprise_info ei on t.usc_id = ei.usc_id
//where content_id = 464191697715211414
