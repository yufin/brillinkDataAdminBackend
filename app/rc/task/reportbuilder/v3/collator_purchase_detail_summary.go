package v3

import (
	"encoding/json"
	"fmt"
	"github.com/buger/jsonparser"
	"github.com/go-admin-team/go-admin-core/sdk"
	"github.com/pkg/errors"
	"go-admin/app/rc/models"
	"strconv"
	"strings"
	"time"
)

type ClaPurchaseDetailSummary struct {
	content   *[]byte
	contentId int64
}

func (s *ClaPurchaseDetailSummary) SetContent(content *[]byte, contentId int64) {
	s.content = content
	s.contentId = contentId
}

func (s *ClaPurchaseDetailSummary) Collating() error {
	summary := make([]string, 0)

	summary1, err := s.subSummary1()
	if err != nil {
		return err
	}
	summary = append(summary, summary1)

	summary2, err := s.subSummary2()
	if err != nil {
		return err
	}
	summary = append(summary, summary2)

	for i := 1; i <= 5; i++ {
		summary3, err := s.subSummary3(i)
		if err != nil {
			return err
		}
		summary = append(summary, summary3)
	}

	summary4, err := s.subSummary4()
	if err != nil {
		return errors.Wrap(err, "subSummary4 error")
	}
	summary = append(summary, summary4)

	summaryBytes, err := json.Marshal(summary)
	if err != nil {
		return errors.Wrap(err, "marshal summary error")
	}

	content, err := jsonparser.Set(*s.content, summaryBytes, "impExpEntReport", "purchaseDetailSummary")
	if err != nil {
		return errors.Wrap(err, "set summary error")
	}
	*s.content = content

	return nil
}

func (s *ClaPurchaseDetailSummary) subSummary1() (string, error) {
	sql :=
		`select sum(ratio) as top5_ratio_sum
		from (select cast(replace(ratio_amount_tax, '%', '') as float) as ratio 
		      from rc_trades_detail
			  where content_id = ?
				and detail_type = 3
			  order by ratio desc
			  limit ?) t`

	modelTd := models.RcTradesDetail{}
	db := sdk.Runtime.GetDbByKey(modelTd.TableName())

	var top5SumRatio float64
	if err := db.Table(modelTd.TableName()).Raw(sql, s.contentId, 5).First(&top5SumRatio).Error; err != nil {
		return "", err
	}

	var top20SumRatio float64
	if err := db.Table(modelTd.TableName()).Raw(sql, s.contentId, 20).First(&top20SumRatio).Error; err != nil {
		return "", err
	}

	// TODO: 确定是用top5SumRatio还是top20SumRatio
	var desc string
	if top5SumRatio < 30 {
		desc = "很低"
	} else if 30 <= top5SumRatio && top5SumRatio < 50 {
		desc = "较低"
	} else if 50 <= top5SumRatio && top5SumRatio < 70 {
		desc = "比较平均"
	} else if 70 <= top5SumRatio && top5SumRatio < 90 {
		desc = "较高"
	} else {
		desc = "很高"
	}

	return fmt.Sprintf("近12个月前五大供应商合计占比%.2f%%，前20供应商合计占比%.2f%%，供应商集中度%s",
		top5SumRatio, top20SumRatio, desc), nil
}

func (s *ClaPurchaseDetailSummary) subSummary2() (string, error) {
	modelRi := models.RcRiskIndex{}
	db := sdk.Runtime.GetDbByKey(modelRi.TableName())

	var stabRatioStr string
	if err := db.Model(&modelRi).
		Select("index_value").
		Where("content_id = ? and index_dec = ?", s.contentId, "供应商稳定性").
		First(&stabRatioStr).
		Error; err != nil {
		return "", err
	}
	// TODO: 确定index_value的单位
	var desc string
	stabRatio, err := strconv.ParseFloat(stabRatioStr, 64)
	if err != nil {
		return "", nil
	} else {
		stabRatio = stabRatio * 100
		if stabRatio < 10 {
			desc = "很低，新供应商数量很多"
		} else if 10 <= stabRatio && stabRatio < 30 {
			desc = "较低，新供应商数量较多"
		} else if 30 <= stabRatio && stabRatio < 50 {
			desc = "一般，新老供应商约一半"
		} else if 50 <= stabRatio && stabRatio < 70 {
			desc = "较高，新供应商数量较少"
		} else {
			desc = "很高，新供应商数量很少"
		}
	}

	return fmt.Sprintf(
		"近12个月与前12个月重要供应商金额占比重合程度比值为%.2f%%,供应商稳定性%s", stabRatio, desc), err
}

func (s *ClaPurchaseDetailSummary) subSummary3(idx int) (string, error) {

	modelTd := models.RcTradesDetail{}
	db := sdk.Runtime.GetDbByKey(modelTd.TableName())

	var m12TopSup, m12TopSupAtM24 models.RcTradesDetail
	if err := db.Model(&modelTd).
		Where("content_id = ? and detail_type = ?", s.contentId, 3).
		Order("cast(replace(ratio_amount_tax, '%', '') as float) desc").
		Offset(idx - 1).
		Limit(1).
		First(&m12TopSup).
		Error; err != nil {
		return "", err
	}

	if err := db.Model(&modelTd).
		Where("content_id = ? and detail_type = ?", s.contentId, 4).
		Where("enterprise_name = ?", m12TopSup.EnterpriseName).
		First(&m12TopSupAtM24).
		Error; err != nil {
		return "", err
	}

	var ranking int
	if err := db.Table(modelTd.TableName()).Raw(
		`select ranking
			from (select *, row_number() over () as ranking
				  from rc_trades_detail
				  where content_id = ?
					and detail_type = 4
				  order by cast(replace(ratio_amount_tax, '%', '') as float) desc) t
			where t.id = ?;`, s.contentId, m12TopSupAtM24.Id).
		First(&ranking).Error; err != nil {
		return "", err
	}

	fM12Ratio, err := strconv.ParseFloat(strings.Replace(m12TopSup.RatioAmountTax, "%", "", -1), 64)
	if err != nil {
		return "", err
	}
	fM24Ratio, err := strconv.ParseFloat(strings.Replace(m12TopSupAtM24.RatioAmountTax, "%", "", -1), 64)
	if err != nil {
		return "", err
	}
	annualChangedRatio := (fM12Ratio - fM24Ratio) / fM24Ratio * 100
	var annualChangedDesc string
	if annualChangedRatio > 0 {
		if 10 < annualChangedRatio && annualChangedRatio <= 30 {
			annualChangedDesc = "近期采购量<font color= #EF5644>小幅度提高</font>"
		} else if 30 < annualChangedRatio && annualChangedRatio <= 100 {
			annualChangedDesc = "近期采购量<font color= #EF5644>大幅度提高</font>"
		} else if 100 < annualChangedRatio && annualChangedRatio <= 170 {
			annualChangedDesc = "近期采购量<font color= #EF5644>约提高了1倍</font>"
		} else if 170 < annualChangedRatio {
			x := int((annualChangedRatio-100)/70) + 1
			annualChangedDesc = fmt.Sprintf("近期采购单量<font color= #EF5644>约提高了%d倍</font>", x)
		}
	} else if annualChangedRatio < 0 {
		if -10 > annualChangedRatio && annualChangedRatio >= -40 {
			annualChangedDesc = "近期采购量<font color= #67C23A>小幅度减少</font>"
		} else if -40 > annualChangedRatio && annualChangedRatio >= -60 {
			annualChangedDesc = "近期采购量<font color= #67C23A>减少约一半</font>"
		} else if -60 > annualChangedRatio && annualChangedRatio >= -90 {
			annualChangedDesc = "近期采购量<font color= #67C23A>约大幅减少</font>"
		} else if -90 > annualChangedRatio && annualChangedRatio >= -100 {
			annualChangedDesc = "近期基本<font color= #67C23A>没有采购</font>"
		}
	}
	if annualChangedDesc == "" {
		annualChangedDesc = "近期采购量<font color= #E6A23C>基本稳定</font>"
	}

	return fmt.Sprintf("近12个月第%d大供应商为'%s'，采购占比为%s，近24个月此供应商采购占比为%s，采购排名第%d，近12个月内采购占比相较于24个月内变化%.2f%%。%s。",
		idx, m12TopSup.EnterpriseName, m12TopSup.RatioAmountTax, m12TopSupAtM24.RatioAmountTax, ranking, annualChangedRatio, annualChangedDesc), nil
}

func (s *ClaPurchaseDetailSummary) subSummary4() (string, error) {
	roc, err := s.getContent()
	if err != nil {
		return "", errors.Wrap(err, "ClaSellingDetailSummary.subSummary4 getContent error")
	}
	maxTimeStr := roc.YearMonth + "-01"

	maxTime, err := time.Parse("2006-01-02", maxTimeStr)

	minTime, err := time.Parse("2006-01-02", "1900-01-01")
	if err != nil {
		return "", errors.Wrap(err, "parse time error at subSummary4")
	}

	time2000, err := time.Parse("2006-01-02", "1999-12-31")
	if err != nil {
		return "", errors.Wrap(err, "parse time error at subSummary4")
	}
	time2010, err := time.Parse("2006-01-02", "2009-12-31")
	if err != nil {
		return "", errors.Wrap(err, "parse time error at subSummary4")
	}
	time2020, err := time.Parse("2006-01-02", "2019-12-31")
	if err != nil {
		return "", errors.Wrap(err, "parse time error at subSummary4")
	}

	total, err := s.countByEstablishDate(minTime, maxTime)
	if err != nil {
		return "", errors.Wrap(err, "count by establish date error at subSummary4")
	}

	c1, err := s.countByEstablishDate(minTime, time2000)
	if err != nil {
		return "", errors.Wrap(err, "count by establish date error at subSummary4")
	}
	c2, err := s.countByEstablishDate(time2000, time2010)
	if err != nil {
		return "", errors.Wrap(err, "count by establish date error at subSummary4")
	}
	c3, err := s.countByEstablishDate(time2010, time2020)
	if err != nil {
		return "", errors.Wrap(err, "count by establish date error at subSummary4")
	}
	c4, err := s.countByEstablishDate(time2020, maxTime)
	if err != nil {
		return "", errors.Wrap(err, "count by establish date error at subSummary4")
	}

	c1p := float64(c1) / float64(total) * 100
	c2p := float64(c2) / float64(total) * 100
	c3p := float64(c3) / float64(total) * 100
	c4p := float64(c4) / float64(total) * 100

	var desc string
	if c4p < 40. {
		desc = "供应商基本是经营多年的公司"
	} else {
		desc = "供应商近5年成立的较多"
	}

	return fmt.Sprintf("近12个月前20供应商中国内企业客户有%d家，成立时间在2000年之前%d家，占比%.2f%%，2000-2009年有%d家，占比%.2f%%，2010-2019年有%d家，占比%.2f%%，2019年-%d年有%d家，占比%.2f%%，%s。",
		total, c1, c1p, c2, c2p, c3, c3p, maxTime.Year(), c4, c4p, desc), nil

}

func (s *ClaPurchaseDetailSummary) countByEstablishDate(minTime time.Time, maxTime time.Time) (int64, error) {
	modelTd := models.RcTradesDetail{}
	db := sdk.Runtime.GetDbByKey(modelTd.TableName())
	var count int64
	err := db.Table(modelTd.TableName()).Raw(
		`select count(distinct enterprise_info.usc_id) as n
			from (select usc_id, enterprise_wait_list.enterprise_name as enterprise_name
				  from rc_trades_detail
						   left join enterprise_wait_list on rc_trades_detail.enterprise_name = enterprise_wait_list.enterprise_name
				  where rc_trades_detail.content_id = ?
					and rc_trades_detail.detail_type = 3
					and length(usc_id) = 18) t
					 left join
				 enterprise_info on t.usc_id = enterprise_info.usc_id
			where established_date <= ?
			  and established_date > ?;`, s.contentId, maxTime, minTime).
		First(&count).
		Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (s *ClaPurchaseDetailSummary) getContent() (*models.RcOriginContent, error) {
	modelRoc := models.RcOriginContent{}
	db := sdk.Runtime.GetDbByKey(modelRoc.TableName())
	if err := db.Model(&modelRoc).First(&modelRoc, s.contentId).Error; err != nil {
		return nil, errors.Wrapf(err, "get content error at getContent with contentId:%d", s.contentId)
	}
	return &modelRoc, nil
}
