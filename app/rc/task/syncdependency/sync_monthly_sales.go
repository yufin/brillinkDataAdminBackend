package syncdependency

import (
	"encoding/json"
	"github.com/go-admin-team/go-admin-core/sdk"
	"github.com/pkg/errors"
	"go-admin/app/rc/models"
	"go-admin/utils"
	"gorm.io/gorm"
)

type monthlySalesSyncProcess struct {
}

func (t monthlySalesSyncProcess) Process(contentId int64) error {
	var dataContent models.RcOriginContent
	dbContent := sdk.Runtime.GetDbByKey(dataContent.TableName())
	err := dbContent.Model(&dataContent).First(&dataContent, contentId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return errors.Wrapf(err, "get origin content error at monthlySalesSyncProcess.Process, contentId: %d", contentId)
	}

	var tbMs models.RcMonthlySales
	dbMs := sdk.Runtime.GetDbByKey(tbMs.TableName())
	if err := dbMs.Unscoped().Where("content_id = ?", contentId).Delete(&models.RcMonthlySales{}).Error; err != nil {
		return errors.Wrapf(err, "delete rc_monthly_sales error at monthlySalesSyncProcess.Process, contentId: %d", contentId)
	}

	var contentMap map[string]any
	if err := json.Unmarshal([]byte(dataContent.Content), &contentMap); err != nil {
		return errors.Wrapf(err, "unmarshal content error at monthlySalesSyncProcess.Process, contentId: %d", contentId)
	}

	monthlySalesArray := contentMap[ReportDataKey].(map[string]any)["ydxsqkDetail"].([]any)
	for _, v := range monthlySalesArray {
		v := v.(map[string]any)
		insertReq := models.RcMonthlySales{
			ContentId:       contentId,
			AttributedMonth: t.safeGetString(v, "MONTH"),
			Nxsr:            t.safeGetString(v, "NXSR"),
			Ckxssr:          t.safeGetString(v, "CKXSSR"),
			Sbkjzsr:         t.safeGetString(v, "SBKJZSR"),
			Fpkjsr:          t.safeGetString(v, "FPKJSR"),
		}
		insertReq.Id = utils.NewFlakeId()
		if err := dbMs.Create(&insertReq).Error; err != nil {
			return errors.Wrapf(err, "create rc_monthly_sales error at monthlySalesSyncProcess.Process, contentId: %d", contentId)
		}
	}
	return nil
}

func (t monthlySalesSyncProcess) safeGetString(m map[string]any, key string) string {
	if v, ok := m[key]; ok {
		return v.(string)
	}
	return ""
}
