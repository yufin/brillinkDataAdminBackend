package v3

import (
	"fmt"
	"github.com/buger/jsonparser"
	"github.com/go-admin-team/go-admin-core/sdk"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"go-admin/app/rc/models"
	"strconv"
	"strings"
	"time"
)

// ClaRevenueDetail 8.3 利润表分析/lrbDetail
type ClaRevenueDetail struct {
	content   *[]byte
	contentId int64
}

func (s *ClaRevenueDetail) SetContent(content *[]byte, contentId int64) {
	s.content = content
	s.contentId = contentId
}

type paramKey struct {
	period string
	val    string
}

func (s *ClaRevenueDetail) Collating() error {
	path := []string{"impExpEntReport", "lrbDetail"}
	var errIter error

	c := *s.content
	idx := 0
	idxP := &idx
	param4Iter := []paramKey{
		{"SSNRQ", "Y2013"},
		{"SNRQ", "Y2014"},
		{"RQ", "Y2015"},
	}
	_, err := jsonparser.ArrayEach(c,
		func(value []byte, dt jsonparser.ValueType, offset int, err error) {
			if dt != jsonparser.Object {
				errIter = errors.Errorf(
					"value is not object while iter lrbDetail whith contentId:%d", s.contentId)
				return
			}
			field, err := jsonparser.GetString(value, "XM")
			if err != nil {
				errIter = errors.Wrapf(err, "get field error while iter lrbDetail GetString XM whith contentId:%d", s.contentId)
				return
			}
			if field == "其中：利息费用" {
				for _, param := range param4Iter {
					val, err := jsonparser.GetUnsafeString(value, param.val)
					if err != nil {
						errIter = errors.Wrapf(err, "get %s error while iter lrbDetail GetString XM whith contentId:%d", param.val, s.contentId)
						return
					}
					period, err := jsonparser.GetUnsafeString(value, param.period)
					if err != nil {
						errIter = errors.Wrapf(err, "get %s error while iter lrbDetail GetString XM whith contentId:%d", param.period, s.contentId)
						return
					}
					err = s.setModify(*idxP, "财务费用", val, period)
					if err != nil {
						errIter = errors.Wrapf(err, "set %s error while iter lrbDetail GetString XM whith contentId:%d", param.val, s.contentId)
						return
					}
				}
				*idxP += 1
			}
		}, path...)

	if errIter != nil {
		return errors.Wrapf(errIter, "ArrayEach error whild iter lrbDetail whith contentId:%d", s.contentId)
	}

	if err != nil {
		return errors.Wrapf(err, "ArrayEach return error while iter lrbDetail whith contentId:%d", s.contentId)
	}

	return nil
}

func (s *ClaRevenueDetail) setModify(idx int, toSetKey string, val string, periodStartStr string) error {
	if val != "0" {
		return nil
	}

	y, err := s.convertYear(periodStartStr)
	if y == nil {
		return nil
	}
	if err != nil {
		return errors.Wrap(err, "convertYear error at ClaRevenueDetail.setContent")
	}

	finExp, err := s.getDetailByField("财务费用", *y)
	if err != nil {
		return errors.Wrapf(err, "getDetailByField error at ClaRevenueDetail.setContent, field: %s, periodStart: %s", "财务费用", *y)
	}

	if !finExp.Val.Valid || !finExp.Val.Decimal.Equal(decimal.Zero) {
		toSet := finExp.Val.Decimal.String()
		idxStr := fmt.Sprintf("[%d]", idx)
		b, err := jsonparser.Set(*s.content, []byte(toSet), "impExpEntReport", "lrbDetail", idxStr, toSetKey)
		if err != nil {
			return errors.Wrapf(err, "jsonparser.Set error at ClaRevenueDetail.setContent, idxStr: %s, toSetKey: %s, toSet: %s", idxStr, toSetKey, toSet)
		}
		*s.content = b
	}
	return nil
}

func (s *ClaRevenueDetail) getDetailByField(field string, periodStart time.Time) (*models.RcRevenueDetail, error) {
	tb := models.RcRevenueDetail{}
	db := sdk.Runtime.GetDbByKey(tb.TableName())
	var data models.RcRevenueDetail
	err := db.Model(models.RcRevenueDetail{}).
		Where("content_id = ? and field = ? and period_start  = ?", s.contentId, field, periodStart).
		First(&data).Error
	if err != nil {
		return nil, errors.Wrapf(
			err, "get detail by field error at ClaRevenueDetail.getDetailByField, field: %s, periodStart: %s", field, periodStart)
	}
	return &data, nil
}

func (s *ClaRevenueDetail) getPeriodScope() ([]time.Time, error) {
	tb := models.RcRevenueDetail{}
	db := sdk.Runtime.GetDbByKey(tb.TableName())

	scope := make([]time.Time, 0)

	err := db.Table(tb.TableName()).
		Select("period_start").
		Where("content_id = ? and field = ?", s.contentId, "财务费用").
		Order("period_start desc").
		Scan(scope).Error
	if err != nil {
		return nil, err
	}

	return scope, nil
}

func (s *ClaRevenueDetail) convertYear(v string) (*time.Time, error) {
	if v == "null" || v == "" {
		return nil, nil
	}
	if strings.Contains(v, "年") {
		year := strings.Replace(v, "年", "", -1)
		yearInt, err := strconv.ParseInt(year, 10, 64)
		if err != nil {
			return nil, errors.Wrapf(err, "parse year error at ClaRevenueDetail.convertYear, v: %s", v)
		}
		r := time.Date(int(yearInt), 1, 1, 0, 0, 0, 0, time.Local)
		return &r, nil
	}
	r, err := time.Parse("2006-01-02", v)
	if err != nil {
		return nil, errors.Wrapf(err, "parse year error at ClaRevenueDetail.convertYear, v: %s", v)
	}
	return &r, nil
}
