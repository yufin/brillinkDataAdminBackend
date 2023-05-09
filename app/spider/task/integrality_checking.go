package task

import (
	log "github.com/go-admin-team/go-admin-core/logger"
	"go-admin/app/spider/models"
	"go-admin/app/spider/service"
	"go-admin/app/spider/service/dto"
	"go-admin/common/actions"
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

type chRes struct {
	count int64
	err   error
}

// CheckIfAllCollected 通过uscId查询4张表中是否都有记录,都存在则返回true,否则返回false
func CheckIfAllCollected(uscId string, sw *ServiceWrap, p *actions.DataPermission) (bool, error) {
	var wg sync.WaitGroup
	countCh := make(chan chRes, 5)

	go func(s *service.EnterpriseInfo, uscId string, list *[]models.EnterpriseInfo, chP *chan chRes) {
		wg.Add(1)
		defer wg.Done()
		var count int64
		if err := s.GetPage(&dto.EnterpriseInfoGetPageReq{UscId: uscId}, p, list, &count); err != nil {
			log.Errorf("CheckIfAllCollectedError %s \r\n", err)
			*chP <- chRes{
				count: 0,
				err:   err,
			}
			return
		}
		*chP <- chRes{
			count: count,
			err:   nil,
		}
	}(sw.InfoS, uscId, &[]models.EnterpriseInfo{}, &countCh)

	go func(s *service.EnterpriseCertification, uscId string, list *[]models.EnterpriseCertification, chP *chan chRes) {
		wg.Add(1)
		defer wg.Done()
		var count int64
		if err := s.GetPage(&dto.EnterpriseCertificationGetPageReq{UscId: uscId}, p, list, &count); err != nil {
			log.Errorf("CheckIfAllCollectedError %s \r\n", err)
			*chP <- chRes{
				count: 0,
				err:   err,
			}
			return
		}
		*chP <- chRes{
			count: count,
			err:   nil,
		}
	}(sw.CertS, uscId, &[]models.EnterpriseCertification{}, &countCh)

	go func(s *service.EnterpriseIndustry, uscId string, list *[]models.EnterpriseIndustry, chP *chan chRes) {
		wg.Add(1)
		defer wg.Done()
		var count int64
		if err := s.GetPage(&dto.EnterpriseIndustryGetPageReq{UscId: uscId}, p, list, &count); err != nil {
			log.Errorf("CheckIfAllCollectedError %s \r\n", err)
			*chP <- chRes{
				count: 0,
				err:   err,
			}
			return
		}
		*chP <- chRes{
			count: count,
			err:   nil,
		}
	}(sw.IndustryS, uscId, &[]models.EnterpriseIndustry{}, &countCh)

	go func(s *service.EnterpriseProduct, uscId string, list *[]models.EnterpriseProduct, chP *chan chRes) {
		wg.Add(1)
		defer wg.Done()
		var count int64
		if err := s.GetPage(&dto.EnterpriseProductGetPageReq{UscId: uscId}, p, list, &count); err != nil {
			log.Errorf("CheckIfAllCollectedError %s \r\n", err)
			*chP <- chRes{
				count: 0,
				err:   err,
			}
			return
		}
		*chP <- chRes{
			count: count,
			err:   nil,
		}
	}(sw.ProductS, uscId, &[]models.EnterpriseProduct{}, &countCh)

	go func(s *service.EnterpriseRanking, uscId string, list *[]models.EnterpriseRanking, chP *chan chRes) {
		wg.Add(1)
		defer wg.Done()
		var count int64
		if err := s.GetPage(&dto.EnterpriseRankingGetPageReq{UscId: uscId}, p, list, &count); err != nil {
			log.Errorf("CheckIfAllCollectedError %s \r\n", err)
			*chP <- chRes{
				count: 0,
				err:   err,
			}
			return
		}
		*chP <- chRes{
			count: count,
			err:   nil,
		}
	}(sw.RankingS, uscId, &[]models.EnterpriseRanking{}, &countCh)

	wg.Wait()
	close(countCh)

	var res chRes
	for range countCh {
		res = <-countCh
		if res.err != nil {
			log.Errorf("CheckIfAllCollectedError %s \r\n", res.err)
			return false, res.err
		}
		if res.count == 0 {
			return false, nil
		}
	}
	return true, nil
}
