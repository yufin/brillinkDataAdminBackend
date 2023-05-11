package task

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/rskc/models"
	"go-admin/app/rskc/service"
	"go-admin/app/rskc/service/dto"
	spModels "go-admin/app/spider/models"
	spService "go-admin/app/spider/service"
	spDto "go-admin/app/spider/service/dto"
	"go-admin/common/actions"
	"go-admin/common/apis"
)

type ServiceWrapRskc struct {
	SWait    *spService.EnterpriseWaitList
	STrades  *service.RskcTradesDetail
	SContent *service.RskcOriginContent
}

func (sw *ServiceWrapRskc) GenServiceWrapRskc(e *apis.Api, c *gin.Context) error {
	sWait := spService.EnterpriseWaitList{}
	if err := e.MakeContext(c).MakeOrm().MakeService(&sWait.Service).Errors; err != nil {
		return err
	}
	sTrades := service.RskcTradesDetail{}
	if err := e.MakeContext(c).MakeOrm().MakeService(&sTrades.Service).Errors; err != nil {
		return err
	}
	sContent := service.RskcOriginContent{}
	if err := e.MakeContext(c).MakeOrm().MakeService(&sContent.Service).Errors; err != nil {
		return err
	}
	sw.SWait = &sWait
	sw.STrades = &sTrades
	sw.SContent = &sContent
	return nil
}

// content-statusCode: 1:待解析录入其他表,2:解析并录入完成,3.数据匹配全部完成,并已录入
// tradeDetail-statusCode: 1.未同步至waitList，2.待采集，已经同步至waitList, 3.确认采集完成, 4.匹配并录入完成
// waitList-statusCode 1.待匹配qccUrl&uscId,2.待爬取,3.爬取完成,-1.爬取失败,9非法公司(自动忽略)"`

func SyncTradesDetailStatusFromWaitList(sw *ServiceWrapRskc, p *actions.DataPermission) error {
	rtdReq := dto.RskcTradesDetailGetPageReq{
		StatusCode: 2, // 2.待采集，已经同步至waitList
	}
	rtdReq.PageSize = 99999
	rtdList := make([]models.RskcTradesDetail, 0)
	var rtdCount int64
	if err := sw.STrades.GetPage(&rtdReq, p, &rtdList, &rtdCount); err != nil {
		return err
	}
	for _, detail := range rtdList {
		// query waitList By enterpriseName
		ewlReq := spDto.EnterpriseWaitListGetPageReq{
			EnterpriseName: detail.EnterpriseName,
		}
		ewlReq.PageSize = 99999
		var ewlCount int64
		elwList := make([]spModels.EnterpriseWaitList, 0)
		if err := sw.SWait.GetPage(&ewlReq, p, &elwList, &ewlCount); err != nil {
			return err
		}
		if len(elwList) == 0 {
			// reset rtd statusCode = 1
			rtdUpdateReq := dto.RskcTradesDetailUpdateReq{
				Id:         detail.Id,
				StatusCode: 1,
			}
			if err := sw.STrades.Update(&rtdUpdateReq, p); err != nil {
				return err
			}
		} else {
			// update rtd statusCode = 3
			ewl := elwList[0]
			if ewl.StatusCode == 3 || ewl.StatusCode == 9 {
				// update statusCode = 3
				rtdUpdateReq := dto.RskcTradesDetailUpdateReq{
					Id:         detail.Id,
					StatusCode: 3,
				}
				if err := sw.STrades.Update(&rtdUpdateReq, p); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func SyncDependencies(sw *ServiceWrapRskc, p *actions.DataPermission) error {
	if err := SyncTradesDetailStatusFromWaitList(sw, p); err != nil {
		return err
	}

	// query content where statusCode=2
	contentReq := dto.RskcOriginContentGetPageReq{
		StatusCode: 2,
	}
	contentReq.PageSize = 99999
	var count int64
	contentList := make([]models.RskcOriginContent, 0)
	if err := sw.SContent.GetPage(&contentReq, p, &contentList, &count); err != nil {
		return err
	}

	// query tradeDetail where statusCode=2 by contentId as result
	for _, content := range contentList {
		// query tradesDetail by contentId, check if all tradesDetail.statusCode = 3
		rtdReq := dto.RskcTradesDetailGetPageReq{
			ContentId: content.Id,
		}
		rtdReq.PageSize = 99999
		rtdList := make([]models.RskcTradesDetail, 0)
		var rtdCount int64
		if err := sw.STrades.GetPage(&rtdReq, p, &rtdList, &rtdCount); err != nil {
			return err
		}
		for _, rtd := range rtdList {
			if rtd.StatusCode != 3 {
				break
			}
		}
		// TODO: sync dependencies table content to tradeDetail

		// update content statusCode = 3
		contentUpdateReq := dto.RskcOriginContentUpdateReq{
			Id:         content.Id,
			StatusCode: 3,
		}
		if err := sw.SContent.Update(&contentUpdateReq, p); err != nil {
			return err
		}
	}
	return nil
}

func SyncDependenciesContent() {
}
