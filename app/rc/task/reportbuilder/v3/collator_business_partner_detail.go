package v3

import (
	"encoding/json"
	"github.com/go-admin-team/go-admin-core/sdk"
	"github.com/pkg/errors"
	"go-admin/app/rc/models"
	eModels "go-admin/app/spider/models"
	"gorm.io/gorm"
	"strings"
	"time"
)

type nameIdent struct {
	EnterpriseName string
	UscId          string
}

type CommodityPropResult struct {
	Id             int64
	DetailType     int
	ContentId      int64
	EnterpriseName string
	MajorityRatio  float64
}

type ClaBusinessPartnerDetail struct {
	content   *[]byte
	contentId int64
}

func (s *ClaBusinessPartnerDetail) SetContent(content *[]byte, contentId int64) {
	s.content = content
	s.contentId = contentId
}

func (c *ClaBusinessPartnerDetail) Collating() error {
	// query tradesDetail by contentId with distinct name
	var tbTrades models.RcTradesDetail
	dbTrades := sdk.Runtime.GetDbByKey(tbTrades.TableName())
	var nameIdents []nameIdent
	err := dbTrades.
		Model(models.RcTradesDetail{}).
		Select("usc_id, enterprise_wait_list.enterprise_name as enterprise_name").
		Joins("left join enterprise_wait_list on rc_trades_detail.enterprise_name = enterprise_wait_list.enterprise_name").
		Where("rc_trades_detail.content_id = ?", c.contentId).
		Scan(&nameIdents).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
	}

	var tbContent models.RcOriginContent
	dbContent := sdk.Runtime.GetDbByKey(tbContent.TableName())
	err = dbContent.Model(models.RcOriginContent{}).First(&tbContent, c.contentId).Scan(&tbContent).Error
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
	nameDataMap := make(map[string]*SubjoinData)
	for _, ni := range nameIdents {
		sjdP, err := c.collateSubJoinData(ni.UscId, evalTime)
		if err != nil {
			return err
		}
		nameDataMap[ni.EnterpriseName] = sjdP
	}

	if err := c.collateContent(c.contentId, nameDataMap, evalTime); err != nil {
		return err
	}
	return nil
}

func (c *ClaBusinessPartnerDetail) collateContent(contentId int64, nameDataMap map[string]*SubjoinData, evalTime time.Time) error {
	// iter content and add data to
	var contentMap map[string]any
	err := json.Unmarshal(*c.content, &contentMap)
	if err != nil {
		return err
	}

	contentMap["impExpEntReport"].(map[string]any)["impJsonDate"] = evalTime.Format("2006-01-02")
	for _, key := range []string{"customerDetail_12", "customerDetail_24", "supplierRanking_12", "supplierRanking_24"} {
		dt := map[string]int{
			"customerDetail_12":  1,
			"customerDetail_24":  2,
			"supplierRanking_12": 3,
			"supplierRanking_24": 4,
		}[key]
		commodityProp, err := c.calculateCommodityProportion(contentId, dt)
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
				var subjoinData *SubjoinData
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

	*c.content = csb
	// insert processed content into db
	//var tbProc models.RcProcessedContent
	//dbProc := sdk.Runtime.GetDbByKey(tbProc.TableName())
	//insertReq := models.RcProcessedContent{
	//	ContentId:  contentId,
	//	Content:    string(csb),
	//	StatusCode: 1,
	//}
	//insertReq.Id = utils.NewFlakeId()
	//
	//if err := dbProc.Model(&models.RcProcessedContent{}).Create(&insertReq).Error; err != nil {
	//	return err
	//}
	return nil
}

func (c *ClaBusinessPartnerDetail) calculateCommodityProportion(contentId int64, detailType int) (*[]CommodityPropResult, error) {
	tb := models.RcTradesDetail{}
	db := sdk.Runtime.GetDbByKey(tb.TableName())
	var result []CommodityPropResult

	err := db.Raw(`select id, detail_type, content_id, enterprise_name, cast(replace(commodity_ratio, '%', '') as decimal(10,2)) /
       100. * cast(replace(sum_amount_tax, ',', '') as decimal(10, 2)) / amount_sum as majority_ratio
    from rc_trades_detail
    left outer join
        (select SUBSTRING_INDEX(SUBSTRING_INDEX(commodity_name, '*', 2), '*', -1) as major_commodity,
                sum(cast(replace(sum_amount_tax, ',', '') as decimal (10,2))) as amount_sum
        from rc_trades_detail where content_id = ? and detail_type = ?
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

func (c *ClaBusinessPartnerDetail) collateSubJoinData(uscId string, evalTime time.Time) (*SubjoinData, error) {
	cidP, err := c.collateCompanyInfoDetail(uscId)
	if err != nil {
		return nil, err
	}
	rtdP, err := c.collateRankingTagDetail(uscId, evalTime)
	if err != nil {
		return nil, err
	}
	atdP, err := c.collateAuthorizedTagDetail(uscId, evalTime)
	if err != nil {
		return nil, err
	}
	ptP, err := c.collateProductTags(uscId)
	if err != nil {
		return nil, err
	}
	itP, err := c.collateIndustryTags(uscId)
	if err != nil {
		return nil, err
	}

	sd := SubjoinData{
		MajorCommodityProportion: nil,
		IndustryTag:              itP,
		AuthorizedTag:            atdP,
		CompanyInfo:              cidP,
		ProductTag:               ptP,
		RankingTag:               rtdP,
	}

	return &sd, nil
}

func (c *ClaBusinessPartnerDetail) collateProductTags(uscId string) (*[]string, error) {
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

func (c *ClaBusinessPartnerDetail) collateIndustryTags(uscId string) (*[]string, error) {
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

func (c *ClaBusinessPartnerDetail) collateCompanyInfoDetail(uscId string) (*CompanyInfoDetail, error) {
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
	return &CompanyInfoDetail{
		EstablishDate:  estDate,
		EnterpriseType: data.EnterpriseType,
		CapitalPaidIn:  data.PaidInCapital,
		HomePage:       data.UrlHomepage,
	}, nil
}

func (c *ClaBusinessPartnerDetail) collateRankingTagDetail(uscId string, evalTime time.Time) (*[]RankingTagDetail, error) {
	var tb eModels.EnterpriseRanking
	db := sdk.Runtime.GetDbByKey(tb.TableName())
	data := make([]RankingTagDetail, 0)

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

func (c *ClaBusinessPartnerDetail) collateAuthorizedTagDetail(uscId string, evalTime time.Time) (*[]AuthorizedTagDetail, error) {
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
	result := make([]AuthorizedTagDetail, 0)
	for _, v := range data {
		result = append(result, AuthorizedTagDetail{
			Authority:      v.CertificationAuthority,
			AuthClass:      v.CertificationLevel,
			TagTitle:       v.CertificationTitle,
			AuthorizedDate: v.CertificationDate.Format("2006-01-02"),
		})
	}
	return &result, nil
}

// collateSubjectCompanyTags collate subject company tags
func (c *ClaBusinessPartnerDetail) collateSubjectCompanyTags(uscId string, contentId int64, evalTime time.Time) (*SubjectCompanyTags, error) {
	it, err := c.collateIndustryTags(uscId)
	if err != nil {
		return nil, err
	}
	at, err := c.collateAuthorizedTagDetail(uscId, evalTime)
	if err != nil {
		return nil, err
	}
	rt, err := c.collateRankingTagDetail(uscId, evalTime)
	if err != nil {
		return nil, err
	}
	poros, err := c.collateSubjectEnterpriseProductProportion(contentId)
	if err != nil {
		return nil, err
	}
	return &SubjectCompanyTags{
		IndustryTag:       it,
		AuthorizedTag:     at,
		RankingTag:        rt,
		ProductProportion: poros,
	}, nil
}

func (c *ClaBusinessPartnerDetail) collateSubjectEnterpriseProductProportion(contentId int64) (*[]ProductProportion, error) {
	var tbSst models.RcSellingSta
	dbSst := sdk.Runtime.GetDbByKey(tbSst.TableName())
	props := make([]ProductProportion, 0)
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
