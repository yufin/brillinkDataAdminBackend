package task

import (
	"github.com/gin-gonic/gin"
	log "github.com/go-admin-team/go-admin-core/logger"
	"go-admin/app/spider/models"
	"go-admin/app/spider/service"
	"go-admin/app/spider/service/dto"
	"go-admin/common/actions"
	"go-admin/common/apis"
	"sync"
)

func CheckIntegrity() {
	/*
		分为更新与首次录入两种逻辑
		录入机制：
		TODO: 待录入查询机制：1.

		TODO: 1.enterprise_industry by: uscId查询接口 order by updatedTime desc
		TODO: 2.enterprise_info by: uscId查询接口 order by updatedTime desc
		TODO: 3.enterprise_product by: uscId查询接口 order by updatedTime desc
		TODO: 4.enterprise_ranking by: uscId查询接口 order by updatedTime desc

		TODO: 定时任务：回写trades_detail表中公司的数据:通过trades_detail.enterpriseName查询wait_list表中对应的uscId, 通过uscId查询四个数据表.若都存在更改trades_detail表statusCode = 3

		TODO: 定时任务：查询trades_detail表中statusCode=3的行，查询对应数据表回写至trades_detail表中,更新statusCode=4

		TODO: 定时任务：通过contentId查询trades_detail表,若所有记录都为4则更新content的statusCode=3
	*/
}

type ServiceWrap struct {
	WaitListS *service.EnterpriseWaitList
	IndustryS *service.EnterpriseIndustry
	InfoS     *service.EnterpriseInfo
	ProductS  *service.EnterpriseProduct
	RankingS  *service.EnterpriseRanking
	CertS     *service.EnterpriseCertification
}

func GenServiceWrap(e *apis.Api, c *gin.Context) (*ServiceWrap, error) {
	sWait := service.EnterpriseWaitList{}
	if err := e.MakeContext(c).MakeOrm().MakeService(&sWait.Service).Errors; err != nil {
		return nil, err
	}
	sInd := service.EnterpriseIndustry{}
	if err := e.MakeContext(c).MakeOrm().MakeService(&sInd.Service).Errors; err != nil {
		return nil, err
	}
	sProd := service.EnterpriseProduct{}
	if err := e.MakeContext(c).MakeOrm().MakeService(&sProd.Service).Errors; err != nil {
		return nil, err
	}
	sRank := service.EnterpriseRanking{}
	if err := e.MakeContext(c).MakeOrm().MakeService(&sRank.Service).Errors; err != nil {
		return nil, err
	}
	sCert := service.EnterpriseCertification{}
	if err := e.MakeContext(c).MakeOrm().MakeService(&sCert.Service).Errors; err != nil {
		return nil, err
	}
	sInf := service.EnterpriseInfo{}
	if err := e.MakeContext(c).MakeOrm().MakeService(&sInf.Service).Errors; err != nil {
		return nil, err
	}
	sw := &ServiceWrap{
		WaitListS: &sWait,
		IndustryS: &sInd,
		InfoS:     &sInf,
		ProductS:  &sProd,
		RankingS:  &sRank,
		CertS:     &sCert,
	}
	return sw, nil
}

type CollectionStatus struct {
	UscId     string `json:"-"`
	Count     int64  `json:"count"`
	TableName string `json:"tableName"`
	Err       error  `json:"-"`
}

func GetDataCollectionDetailByUscId(uscId string, p *actions.DataPermission, sw *ServiceWrap) (*[]CollectionStatus, error) {
	var wg sync.WaitGroup
	countCh := make(chan CollectionStatus, 5)

	wg.Add(1)
	go func(s *service.EnterpriseInfo, uscId string, list *[]models.EnterpriseInfo, chP *chan CollectionStatus) {
		defer wg.Done()
		var count int64
		if err := s.GetPage(&dto.EnterpriseInfoGetPageReq{UscId: uscId}, p, list, &count); err != nil {
			log.Errorf("CheckIfAllCollectedError %s \r\n", err)
			*chP <- CollectionStatus{
				Count:     0,
				Err:       err,
				TableName: "enterprise_info",
				UscId:     uscId,
			}
			return
		}
		*chP <- CollectionStatus{
			Count:     count,
			Err:       nil,
			TableName: "enterprise_info",
			UscId:     uscId,
		}
	}(sw.InfoS, uscId, &[]models.EnterpriseInfo{}, &countCh)

	wg.Add(1)
	go func(s *service.EnterpriseCertification, uscId string, list *[]models.EnterpriseCertification, chP *chan CollectionStatus) {
		defer wg.Done()
		var count int64
		if err := s.GetPage(&dto.EnterpriseCertificationGetPageReq{UscId: uscId}, p, list, &count); err != nil {
			log.Errorf("CheckIfAllCollectedError %s \r\n", err)
			*chP <- CollectionStatus{
				Count:     0,
				Err:       err,
				TableName: "enterprise_certification",
				UscId:     uscId,
			}
			return
		}
		*chP <- CollectionStatus{
			Count:     count,
			Err:       nil,
			TableName: "enterprise_certification",
			UscId:     uscId,
		}
	}(sw.CertS, uscId, &[]models.EnterpriseCertification{}, &countCh)

	wg.Add(1)
	go func(s *service.EnterpriseIndustry, uscId string, list *[]models.EnterpriseIndustry, chP *chan CollectionStatus) {
		defer wg.Done()
		var count int64
		if err := s.GetPage(&dto.EnterpriseIndustryGetPageReq{UscId: uscId}, p, list, &count); err != nil {
			log.Errorf("CheckIfAllCollectedError %s \r\n", err)
			*chP <- CollectionStatus{
				Count:     0,
				Err:       err,
				TableName: "enterprise_industry",
				UscId:     uscId,
			}
			return
		}
		*chP <- CollectionStatus{
			Count:     count,
			Err:       nil,
			TableName: "enterprise_industry",
			UscId:     uscId,
		}
	}(sw.IndustryS, uscId, &[]models.EnterpriseIndustry{}, &countCh)

	wg.Add(1)
	go func(s *service.EnterpriseProduct, uscId string, list *[]models.EnterpriseProduct, chP *chan CollectionStatus) {
		defer wg.Done()
		var count int64
		if err := s.GetPage(&dto.EnterpriseProductGetPageReq{UscId: uscId}, p, list, &count); err != nil {
			log.Errorf("CheckIfAllCollectedError %s \r\n", err)
			*chP <- CollectionStatus{
				Count:     0,
				Err:       err,
				TableName: "enterprise_product",
				UscId:     uscId,
			}
			return
		}
		*chP <- CollectionStatus{
			Count:     count,
			Err:       nil,
			TableName: "enterprise_product",
			UscId:     uscId,
		}
	}(sw.ProductS, uscId, &[]models.EnterpriseProduct{}, &countCh)

	wg.Add(1)
	go func(s *service.EnterpriseRanking, uscId string, list *[]models.EnterpriseRanking, chP *chan CollectionStatus) {
		defer wg.Done()
		var count int64
		if err := s.GetPage(&dto.EnterpriseRankingGetPageReq{UscId: uscId}, p, list, &count); err != nil {
			log.Errorf("CheckIfAllCollectedError %s \r\n", err)
			*chP <- CollectionStatus{
				Count:     0,
				Err:       err,
				TableName: "enterprise_ranking",
				UscId:     uscId,
			}
			return
		}
		*chP <- CollectionStatus{
			Count:     count,
			Err:       nil,
			TableName: "enterprise_ranking",
			UscId:     uscId,
		}
	}(sw.RankingS, uscId, &[]models.EnterpriseRanking{}, &countCh)

	wg.Wait()

	close(countCh)

	statusList := make([]CollectionStatus, 0)
	chLen := len(countCh)
	var res CollectionStatus
	var i int
	for i = 0; i < chLen; i++ {
		res = <-countCh
		if res.Err != nil {
			log.Errorf("CheckIfAllCollectedError %s \r\n", res.Err)
			return nil, res.Err
		}
		statusList = append(statusList, res)
	}
	return &statusList, nil
}

// CheckIfAllCollected 通过uscId查询4张表中是否都有记录,都存在则返回true,否则返回false
func CheckIfAllCollected(statusList *[]CollectionStatus) bool {
	for _, res := range *statusList {
		if res.Count == 0 {
			return false
		}
	}
	return true
}
