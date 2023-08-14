package v3

import (
	"encoding/json"
	"github.com/buger/jsonparser"
	"github.com/go-admin-team/go-admin-core/sdk"
	"github.com/pkg/errors"
	"go-admin/app/rc/models"
	eModels "go-admin/app/spider/models"
	"gorm.io/gorm"
	"time"
)

type ClaSubjCompanyTag struct {
	content   *[]byte
	contentId int64
}

func (s *ClaSubjCompanyTag) SetContent(content *[]byte, contentId int64) {
	s.content = content
	s.contentId = contentId
}

func (s *ClaSubjCompanyTag) Collating() error {
	modelRoc := models.RcOriginContent{}
	dbRoc := sdk.Runtime.GetDbByKey(modelRoc.TableName())
	err := dbRoc.Model(&modelRoc).First(&modelRoc, s.contentId).Error
	if err != nil {
		return err
	}
	evalTime, err := time.Parse("2006-01-02", modelRoc.YearMonth+"-01")

	subjTags, err := s.collateSubjectCompanyTags(modelRoc.UscId, evalTime)
	subjTagsBytes, err := json.Marshal(subjTags)
	if err != nil {
		return err
	}
	tempC, err := jsonparser.Set(*s.content, subjTagsBytes, "impExpEntReport", "subjectCompanyTags")
	if err != nil {
		return errors.Wrapf(err, "collateSubjectCompanyTags error, contentId: %d", s.contentId)
	}

	*s.content = tempC

	return nil
}

func (s *ClaSubjCompanyTag) collateSubjectCompanyTags(uscId string, evalTime time.Time) (*SubjectCompanyTags, error) {
	it, err := s.collateIndustryTags(uscId)
	if err != nil {
		return nil, err
	}
	at, err := s.collateAuthorizedTagDetail(uscId, evalTime)
	if err != nil {
		return nil, err
	}
	rt, err := s.collateRankingTagDetail(uscId, evalTime)
	if err != nil {
		return nil, err
	}
	poros, err := s.collateSubjectEnterpriseProductProportion()
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

func (s *ClaSubjCompanyTag) collateSubjectEnterpriseProductProportion() (*[]ProductProportion, error) {
	var tbSst models.RcSellingSta
	dbSst := sdk.Runtime.GetDbByKey(tbSst.TableName())
	props := make([]ProductProportion, 0)
	err := dbSst.Table(tbSst.TableName()).Raw(
		`select SUBSTRING_INDEX(SUBSTRING_INDEX(ssspxl, '*', 2), '*', -1) as category,
				   concat(sum(cast(Replace(jezb, '%', '') as DECIMAL(10, 2))), '%') as proportion,
				   group_concat(
						   SUBSTRING_INDEX(ssspxl, '*', -1),
						   concat('(', jezb, ')') order by jezb desc) as category_detail
			from rc_selling_sta
			where content_id = ?
			  and SSSPDL not in ('合计', '其他')
			group by category
			ORDER BY  REPLACE(proportion,"%","")+0 desc;`, s.contentId).
		Scan(&props).Error
	if err != nil {
		return nil, err
	}
	if len(props) == 0 {
		return nil, nil
	}
	return &props, nil
}

func (s *ClaSubjCompanyTag) collateAuthorizedTagDetail(uscId string, evalTime time.Time) (*[]AuthorizedTagDetail, error) {
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

func (s *ClaSubjCompanyTag) collateRankingTagDetail(uscId string, evalTime time.Time) (*[]RankingTagDetail, error) {
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

func (s *ClaSubjCompanyTag) collateIndustryTags(uscId string) (*[]string, error) {
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
