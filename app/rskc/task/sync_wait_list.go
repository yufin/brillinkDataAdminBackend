package task

import (
	log "github.com/go-admin-team/go-admin-core/logger"
	"go-admin/app/rskc/models"
	"go-admin/app/rskc/service"
	"go-admin/app/rskc/service/dto"
	sModels "go-admin/app/spider/models"
	sDto "go-admin/app/spider/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
)
import sService "go-admin/app/spider/service"

type ServicesWrap struct {
	SContent *service.RskcOriginContent
	SDetail  *service.RskcTradesDetail
	SWait    *sService.EnterpriseWaitList
	P        *actions.DataPermission
}

// SyncWaitList 同步企业至等待列表
func SyncWaitList(s *ServicesWrap) error {
	if err := SyncToWaitListFromDetail(s); err != nil {
		return err
	}

	return nil
}

func SyncToWaitListFromDetail(sw *ServicesWrap) error {
	detailReq := dto.RskcTradesDetailGetPageReq{
		Pagination: cDto.Pagination{
			PageSize:  3000,
			PageIndex: 1,
		},
		StatusCode: 1,
	}
	detailList := make([]models.RskcTradesDetail, 0)
	var count int64
	if err := sw.SDetail.GetPage(&detailReq, sw.P, &detailList, &count); err != nil {
		log.Errorf("Task SyncWaitList func-SyncWaitList Failed:%s \r\n", err)
		return err
	}

	if count == 0 {
		return nil
	}

	for _, detail := range detailList {

		err := func(sw *ServicesWrap) error {
			// check if waitList exists with same enterprise_name
			waitReq := sDto.EnterpriseWaitListGetPageReq{
				EnterpriseName: detail.EnterpriseName,
			}
			var wCount int64
			wList := make([]sModels.EnterpriseWaitList, 0)
			if err := sw.SWait.GetPage(&waitReq, sw.P, &wList, &wCount); err != nil {
				log.Errorf("Task SyncWaitList func-SyncToWaitListFromDetail Failed:%s \r\n", err)
				return err
			}

			if wCount == 0 {
				// insert into enterprise_wait_list
				insertReq := sDto.EnterpriseWaitListInsertReq{
					EnterpriseName: detail.EnterpriseName,
					UscId:          "",
					Priority:       9,
					QccUrl:         "",
					StatusCode:     1, // 数据爬取状态码,1.待确认爬取状态,2.待爬取,3.爬取完成,-1.爬取失败
					Source:         "rskc_trades_detail",
				}
				insertReq.SetCreateBy(0)

				if err := sw.SWait.Insert(&insertReq); err != nil {
					log.Errorf("Task SyncWaitList func-SyncToWaitListFromDetail Failed:%s \r\n", err)
					return err
				}
			}
			// update rskc_trade_detail status = 2 (1.待确认企业数据已采集，2.待采集，已经同步至waitList, 3.采集完成, 4.匹配并录入完成)
			updateReq := dto.RskcTradesDetailUpdateReq{Id: detail.Id, StatusCode: 2}
			updateReq.SetUpdateBy(0)
			if err := sw.SDetail.Update(&updateReq, sw.P); err != nil {
				log.Errorf("Task SyncWaitList func-SyncToWaitListFromDetail Failed:%s \r\n", err)
				return err
			}
			return nil
		}(sw)
		if err != nil {
			log.Errorf("Task SyncWaitList func-SyncToWaitListFromDetail Failed:%s \r\n", err)
		}
	}
	return nil
}
