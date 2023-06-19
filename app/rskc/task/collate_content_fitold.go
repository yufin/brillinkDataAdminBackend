package task

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/go-admin-team/go-admin-core/sdk"
	"github.com/nats-io/nats.go"
	"github.com/pkg/errors"
	"go-admin/app/rskc/models"
	"go-admin/app/rskc/service/dto"
	eModels "go-admin/app/spider/models"
	"go-admin/pkg/natsclient"
	"gorm.io/gorm"
	"strings"
	"time"
)

type CollateContentTask struct {
}

func (t CollateContentTask) Exec(arg interface{}) error {
	return pullToProcessNew()
}

func pullToProcessNew() error {
	for {
		// get total msg count by subscriber
		totalPending, _, err := natsclient.SubContentProcessNew.Pending()
		if err == nil {
			fmt.Println("SubContentProcessNew msg totalPending:", totalPending)
		}
		msgs, err := natsclient.SubContentProcessNew.Fetch(1, nats.MaxWait(5*time.Second))
		if err != nil {
			if err == nats.ErrTimeout {
				return nil
			} else {
				return err
			}
		}
		for _, msg := range msgs {
			contentId := int64(binary.BigEndian.Uint64(msg.Data))
			err := collateDependencyForContent(contentId)
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

type nameIdent struct {
	EnterpriseName string
	UscId          string
}

func collateDependencyForContent(contentId int64) error {
	// query tradesDetail by contentId with distinct name
	var tbTrades models.RskcTradesDetail
	dbTrades := sdk.Runtime.GetDbByKey(tbTrades.TableName())
	var nameIdents []nameIdent
	err := dbTrades.
		Model(models.RskcTradesDetail{}).
		Select("usc_id, enterprise_wait_list.enterprise_name as enterprise_name").
		Joins("left join enterprise_wait_list on rskc_trades_detail.enterprise_name = enterprise_wait_list.enterprise_name").
		Where("rskc_trades_detail.content_id = ?", contentId).
		Scan(&nameIdents).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
	}

	var tbContent models.RskcOriginContent
	dbContent := sdk.Runtime.GetDbByKey(tbContent.TableName())
	err = dbContent.Model(models.RskcOriginContent{}).First(&tbContent, contentId).Scan(&tbContent).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
	}
	evalTime, err := time.Parse("2006-01-02", tbContent.YearMonth+"-01")
	if err != nil {
		return err
	}
	nameDataMap := make(map[string]*dto.SubjoinData)
	for _, ni := range nameIdents {
		sjdP, err := collateSubJoinData(ni.UscId, evalTime)
		if err != nil {
			return err
		}
		nameDataMap[ni.EnterpriseName] = sjdP
	}

	if err := collateContent(contentId, nameDataMap, evalTime); err != nil {
		return err
	}
	return nil
}

func collateContent(contentId int64, nameDataMap map[string]*dto.SubjoinData, evalTime time.Time) error {
	// iter content and add data to
	var data models.RskcOriginContent
	db := sdk.Runtime.GetDbByKey(data.TableName())
	err := db.Model(&models.RskcOriginContent{}).First(&data, contentId).Error
	if err != nil {
		return err
	}
	var contentMap map[string]any
	err = json.Unmarshal([]byte(data.Content), &contentMap)
	if err != nil {
		return err
	}

	// 组装subjectCompanyTags
	sct, err := collateSubjectCompanyTags(data.UscId, contentId, evalTime)
	if err != nil {
		return err
	}
	contentMap["impExpEntReport"].(map[string]any)["subjectCompanyTags"] = sct.GenMap()

	contentMap["impExpEntReport"].(map[string]any)["impJsonDate"] = evalTime.Format("2006-01-02")
	for _, key := range []string{"customerDetail_12", "customerDetail_24", "supplierRanking_12", "supplierRanking_24"} {
		dt := map[string]int{
			"customerDetail_12":  1,
			"customerDetail_24":  2,
			"supplierRanking_12": 3,
			"supplierRanking_24": 4,
		}[key]
		commodityProp, err := calculateCommodityProportion(contentId, dt)
		if err != nil {
			return err
		}
		for idx, item := range contentMap["impExpEntReport"].(map[string]any)[key].([]interface{}) {
			mItem, ok := item.(map[string]any)
			if ok {
				companyNameTemp := func() any {
					keyCompany := map[string]string{
						"customerDetail_12":  "PURCHASER_NAME",
						"customerDetail_24":  "PURCHASER_NAME",
						"supplierRanking_12": "SALES_NAME",
						"supplierRanking_24": "SALES_NAME",
					}[key]
					return mItem[keyCompany]
				}()
				// determine if companyName is empty
				companyName, ok := companyNameTemp.(string)
				var subjoinData *dto.SubjoinData
				if ok {
					subjoinData = nameDataMap[companyName]
				}
				if subjoinData != nil {
					// 规则处理：企业名称包含‘供应链’时不匹配标签
					contentMap["impExpEntReport"].(map[string]any)[key].([]interface{})[idx].(map[string]any)["companyInfo"] = subjoinData.CompanyInfo.GenMap()
					if !strings.Contains(companyName, "供应链") {
						contentMap["impExpEntReport"].(map[string]any)[key].([]interface{})[idx].(map[string]any)["industryTag"] = subjoinData.IndustryTag
						contentMap["impExpEntReport"].(map[string]any)[key].([]interface{})[idx].(map[string]any)["productTag"] = subjoinData.ProductTag
						contentMap["impExpEntReport"].(map[string]any)[key].([]interface{})[idx].(map[string]any)["rankingTag"] = func() *[]map[string]any {
							rt := make([]map[string]any, 0)
							if subjoinData.RankingTag == nil {
								return nil
							}
							for _, r := range *subjoinData.RankingTag {
								rt = append(rt, *r.GenMap())
							}
							return &rt
						}()
						contentMap["impExpEntReport"].(map[string]any)[key].([]any)[idx].(map[string]any)["authorizedTag"] = func() *[]map[string]any {
							rt := make([]map[string]any, 0)
							if subjoinData.AuthorizedTag == nil {
								return nil
							}
							for _, at := range *subjoinData.AuthorizedTag {
								rt = append(rt, at.GenMap())
							}
							return &rt
						}()
						contentMap["impExpEntReport"].(map[string]any)[key].([]any)[idx].(map[string]any)["major_commodity_proportion"] = func() *float64 {
							if commodityProp == nil {
								return nil
							}
							for _, v := range *commodityProp {
								if v.EnterpriseName == companyName {
									return &v.MajorityRatio
								}
							}
							return nil
						}()
					} else {
						contentMap["impExpEntReport"].(map[string]any)[key].([]any)[idx].(map[string]any)["major_commodity_proportion"] = nil
						contentMap["impExpEntReport"].(map[string]any)[key].([]any)[idx].(map[string]any)["authorizedTag"] = nil
						contentMap["impExpEntReport"].(map[string]any)[key].([]interface{})[idx].(map[string]any)["industryTag"] = nil
						contentMap["impExpEntReport"].(map[string]any)[key].([]interface{})[idx].(map[string]any)["productTag"] = nil
						contentMap["impExpEntReport"].(map[string]any)[key].([]interface{})[idx].(map[string]any)["rankingTag"] = nil
					}
				}
			}
		}
	}

	csb, err := json.Marshal(contentMap)
	if err != nil {
		return err
	}
	// insert processed content into db
	var tbProc models.RskcProcessedContent
	dbProc := sdk.Runtime.GetDbByKey(tbProc.TableName())
	insertReq := dto.RskcProcessedContentInsertReq{
		ContentId:  contentId,
		Content:    string(csb),
		StatusCode: 1,
	}
	var pc models.RskcProcessedContent
	insertReq.Generate(&pc)
	if err := dbProc.Model(&models.RskcProcessedContent{}).Create(&pc).Error; err != nil {
		return err
	}
	return nil
}

type CommodityPropResult struct {
	Id             int64
	DetailType     int
	ContentId      int64
	EnterpriseName string
	MajorityRatio  float64
}

func calculateCommodityProportion(contentId int64, detailType int) (*[]CommodityPropResult, error) {
	tb := models.RskcTradesDetail{}
	db := sdk.Runtime.GetDbByKey(tb.TableName())
	var result []CommodityPropResult

	err := db.Raw(`select id, detail_type, content_id, enterprise_name, cast(replace(commodity_ratio, '%', '') as decimal(10,2)) /
       100. * cast(replace(sum_amount_tax, ',', '') as decimal(10, 2)) / amount_sum as majority_ratio
    from rskc_trades_detail
    left outer join
        (select SUBSTRING_INDEX(SUBSTRING_INDEX(commodity_name, '*', 2), '*', -1) as major_commodity,
                sum(cast(replace(sum_amount_tax, ',', '') as decimal (10,2))) as amount_sum
        from rskc_trades_detail where content_id = ? and detail_type = ?
        group by major_commodity order by amount_sum desc limit 1) as grouped
    on SUBSTRING_INDEX(SUBSTRING_INDEX(commodity_name, '*', 2), '*', -1) = grouped.major_commodity
    where  content_id = ? and detail_type = ?;`,
		contentId, detailType, contentId, detailType).
		Scan(&result).
		Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func collateSubJoinData(uscId string, evalTime time.Time) (*dto.SubjoinData, error) {
	cidP, err := collateCompanyInfoDetail(uscId)
	if err != nil {
		return nil, err
	}
	rtdP, err := collateRankingTagDetail(uscId, evalTime)
	if err != nil {
		return nil, err
	}
	atdP, err := collateAuthorizedTagDetail(uscId, evalTime)
	if err != nil {
		return nil, err
	}
	ptP, err := collateProductTags(uscId)
	if err != nil {
		return nil, err
	}
	itP, err := collateIndustryTags(uscId)
	if err != nil {
		return nil, err
	}

	sd := dto.SubjoinData{
		MajorCommodityProportion: nil,
		IndustryTag:              itP,
		AuthorizedTag:            atdP,
		CompanyInfo:              cidP,
		ProductTag:               ptP,
		RankingTag:               rtdP,
	}

	return &sd, nil
}

func collateProductTags(uscId string) (*[]string, error) {
	var data eModels.EnterpriseProduct
	db := sdk.Runtime.GetDbByKey(data.TableName())
	err := db.Model(&eModels.EnterpriseProduct{}).
		Where("usc_id = ?", uscId).
		Order("created_at DESC").
		First(&data).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	var result []string

	if data.ProductData == "" {
		return nil, nil
	}
	err = json.Unmarshal([]byte(data.ProductData), &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func collateIndustryTags(uscId string) (*[]string, error) {
	var data eModels.EnterpriseIndustry
	db := sdk.Runtime.GetDbByKey(data.TableName())
	err := db.Model(&data).
		Where("usc_id = ?", uscId).
		Order("created_at DESC").
		First(&data).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	var result []string
	if data.IndustryData == "" {
		return nil, nil
	}
	err = json.Unmarshal([]byte(data.IndustryData), &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func collateCompanyInfoDetail(uscId string) (*dto.CompanyInfoDetail, error) {
	var data eModels.EnterpriseInfo
	db := sdk.Runtime.GetDbByKey(data.TableName())
	err := db.Model(&eModels.EnterpriseInfo{}).
		Where("usc_id = ?", uscId).
		Order("created_at DESC").
		First(&data).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	estDate := ""
	if data.EstablishedDate != nil {
		estDate = data.EstablishedDate.Format("2006-01-02")
	}
	return &dto.CompanyInfoDetail{
		EstablishDate:  estDate,
		EnterpriseType: data.EnterpriseType,
		CapitalPaidIn:  data.PaidInCapital,
	}, nil
}

func collateRankingTagDetail(uscId string, evalTime time.Time) (*[]dto.RankingTagDetail, error) {
	var tb eModels.EnterpriseRanking
	db := sdk.Runtime.GetDbByKey(tb.TableName())
	data := make([]dto.RankingTagDetail, 0)

	err := db.Model(&eModels.EnterpriseRanking{}).
		Select("distinct ranking_list.list_title as tag_title, "+
			"ranking_list.list_published_date as date_published, "+
			"ranking_list.list_source as autority, "+
			"ranking_position as ranking, "+
			"list_participants_total as total").
		Joins("left join ranking_list on ranking_list.id = enterprise_ranking.list_id").
		Where("usc_id = ?", uscId).
		Where("list_published_date <= ?", evalTime).
		Order("date_published DESC").
		Limit(20).
		Scan(&data).
		Error
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, nil
	}
	return &data, nil
}

func collateAuthorizedTagDetail(uscId string, evalTime time.Time) (*[]dto.AuthorizedTagDetail, error) {
	var tb eModels.EnterpriseCertification
	db := sdk.Runtime.GetDbByKey(tb.TableName())
	data := make([]eModels.EnterpriseCertification, 0)
	err := db.Model(&tb).Raw(
		`SELECT *
			FROM (
				SELECT *, ROW_NUMBER() OVER(PARTITION BY certification_source ORDER BY created_at DESC) as rn
				FROM enterprise_certification
				WHERE usc_id = ?
				AND certification_date <= ?
			) sub
			WHERE rn = 1;`, uscId, evalTime).
		Scan(&data).
		Error
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, nil
	}
	result := make([]dto.AuthorizedTagDetail, 0)
	for _, v := range data {
		result = append(result, dto.AuthorizedTagDetail{
			Authority:      v.CertificationAuthority,
			AuthClass:      v.CertificationLevel,
			TagTitle:       v.CertificationTitle,
			AuthorizedDate: v.CertificationDate.Format("2006-01-02"),
		})
	}
	return &result, nil
}

// collateSubjectCompanyTags collate subject company tags
func collateSubjectCompanyTags(uscId string, contentId int64, evalTime time.Time) (*dto.SubjectCompanyTags, error) {
	it, err := collateIndustryTags(uscId)
	if err != nil {
		return nil, err
	}
	at, err := collateAuthorizedTagDetail(uscId, evalTime)
	if err != nil {
		return nil, err
	}
	rt, err := collateRankingTagDetail(uscId, evalTime)
	if err != nil {
		return nil, err
	}
	poros, err := collateSubjectEnterpriseProductProportion(contentId)
	if err != nil {
		return nil, err
	}
	return &dto.SubjectCompanyTags{
		IndustryTag:       it,
		AuthorizedTag:     at,
		RankingTag:        rt,
		ProductProportion: poros,
	}, nil
}

func collateSubjectEnterpriseProductProportion(contentId int64) (*[]dto.ProductProportion, error) {
	var tbSst models.RcSellingSta
	dbSst := sdk.Runtime.GetDbByKey(tbSst.TableName())
	props := make([]dto.ProductProportion, 0)
	err := dbSst.Raw(
		`select SUBSTRING_INDEX(SUBSTRING_INDEX(ssspxl, '*', 2), '*', -1) as category,
        concat(sum(cast(Replace(jezb, '%', '') as DECIMAL(10,2))), '%')as proportion
        from rc_selling_sta where content_id = ? and  SSSPDL not in ('合计', '其他')
        group by category`, contentId).
		Scan(&props).Error
	if err != nil {
		return nil, err
	}
	if len(props) == 0 {
		return nil, nil
	}
	return &props, nil
}
