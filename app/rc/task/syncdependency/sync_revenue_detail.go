package syncdependency

import (
	"encoding/json"
	"github.com/go-admin-team/go-admin-core/sdk"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"go-admin/app/rc/models"
	"go-admin/utils"
	"gorm.io/gorm"
	"strconv"
	"strings"
	"time"
)

type revenueDetailSyncProcess struct {
}

func (t revenueDetailSyncProcess) Process(contentId int64) error {
	var dataContent models.RcOriginContent
	dbContent := sdk.Runtime.GetDbByKey(dataContent.TableName())
	err := dbContent.Model(&dataContent).First(&dataContent, contentId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return errors.Wrapf(err, "get origin content error at monthlyPurchaseSyncProcess.Process, contentId: %d", contentId)
	}

	var tb models.RcRevenueDetail
	db := sdk.Runtime.GetDbByKey(tb.TableName())
	if err := db.Unscoped().Where("content_id = ?", contentId).Delete(&models.RcRevenueDetail{}).Error; err != nil {
		return errors.Wrap(err, "delete rc_revenue_detail error at revenueDetailSyncProcess.Process")
	}

	var contentMap map[string]any
	if err := json.Unmarshal([]byte(dataContent.Content), &contentMap); err != nil {
		return errors.Wrapf(err, "unmarshal content error at revenueDetailSyncProcess.Process, contentId: %d", contentId)
	}

	revenueDetailArray := contentMap[ReportDataKey].(map[string]any)["lrbDetail"].([]any)
	for _, v := range revenueDetailArray {
		v := v.(map[string]any)
		ssnrq, err := t.getYearValue(v, "SSNRQ")
		if err != nil {
			return errors.Wrapf(err, "get year value error at revenueDetailSyncProcess.Process, contentId: %d", contentId)
		}
		snrq, err := t.getYearValue(v, "SNRQ")
		if err != nil {
			return errors.Wrapf(err, "get year value error at revenueDetailSyncProcess.Process, contentId: %d", contentId)
		}
		rq, err := t.getYearValue(v, "RQ")
		if err != nil {
			return errors.Wrapf(err, "get year value error at revenueDetailSyncProcess.Process, contentId: %d", contentId)
		}
		val2013, err := t.safeGetDecimal(v, "Y2013")
		if err != nil {
			return errors.Wrap(err, "get year value error at revenueDetailSyncProcess.Process")
		}
		val2014, err := t.safeGetDecimal(v, "Y2014")
		if err != nil {
			return errors.Wrap(err, "get year value error at revenueDetailSyncProcess.Process")
		}
		val2015, err := t.safeGetDecimal(v, "Y2015")
		if err != nil {
			return errors.Wrap(err, "get year value error at revenueDetailSyncProcess.Process")
		}
		xh := t.safeGetString(v, "XH")
		var xhInt int64
		if xh != "" {
			xhInt, err = strconv.ParseInt(xh, 10, 64)
			if err != nil {
				return errors.Wrapf(err, "parse int error at revenueDetailSyncProcess.Process, contentId: %d", contentId)
			}
		}
		var xhP *int
		if xhInt != 0 {
			vXh := int(xhInt)
			xhP = &vXh
		}

		xm := t.safeGetString(v, "XM")

		insertReq1 := models.RcRevenueDetail{
			ContentId:   contentId,
			Seq:         xhP,
			Field:       xm,
			Val:         val2013,
			PeriodStart: *ssnrq,
			PeriodEnd:   time.Date(ssnrq.Year(), 12, 31, 0, 0, 0, 0, time.Local),
		}

		insertReq2 := models.RcRevenueDetail{
			ContentId:   contentId,
			Seq:         xhP,
			Field:       xm,
			Val:         val2014,
			PeriodStart: *snrq,
			PeriodEnd:   time.Date(snrq.Year(), 12, 31, 0, 0, 0, 0, time.Local),
		}

		insertReq3 := models.RcRevenueDetail{
			ContentId:   contentId,
			Seq:         xhP,
			Field:       xm,
			Val:         val2015,
			PeriodStart: time.Date(rq.Year(), 1, 1, 0, 0, 0, 0, time.Local),
			PeriodEnd:   *rq,
		}

		for _, req := range []models.RcRevenueDetail{insertReq1, insertReq2, insertReq3} {
			req.Id = utils.NewFlakeId()
			if err := db.Create(&req).Error; err != nil {
				return errors.Wrapf(err, "create rc_revenue_detail error at revenueDetailSyncProcess.Process, req: %+v", req)
			}
		}
	}

	return nil
}

func (t revenueDetailSyncProcess) safeGetDecimal(m map[string]any, key string) (decimal.NullDecimal, error) {
	if v, ok := m[key]; ok {
		if v != nil {
			d, err := decimal.NewFromString(v.(string))
			if err != nil {
				return decimal.NullDecimal{}, errors.Wrapf(err, "parse decimal error at revenueDetailSyncProcess.safeGetDecimal, key: %s", key)
			}
			return decimal.NullDecimal{Decimal: d, Valid: true}, nil
		}
		return decimal.NullDecimal{Valid: false}, nil
	}
	return decimal.NullDecimal{}, nil
}

func (t revenueDetailSyncProcess) safeGetString(m map[string]any, key string) string {
	if v, ok := m[key]; ok {
		if v != nil {
			return v.(string)
		}
		return ""
	}
	return ""
}

func (t revenueDetailSyncProcess) getYearValue(m map[string]any, key string) (*time.Time, error) {
	if v, ok := m[key]; ok {
		if v != nil {
			if strings.Contains(v.(string), "年") {
				year := strings.Replace(v.(string), "年", "", -1)
				yearInt, err := strconv.ParseInt(year, 10, 64)
				if err != nil {
					return nil, errors.Wrapf(err, "parse year error at revenueDetailSyncProcess.getYearValue, key: %s", key)
				}
				r := time.Date(int(yearInt), 1, 1, 0, 0, 0, 0, time.Local)
				return &r, nil
			}
			r, err := time.Parse("2006-01-02", v.(string))
			if err != nil {
				return nil, errors.Wrapf(err, "parse date error at revenueDetailSyncProcess.getYearValue, key: %s", key)
			}
			return &r, nil
		}
	}
	return nil, nil
}
