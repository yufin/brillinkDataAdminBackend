package syncdependency

import (
	"encoding/json"
	"github.com/go-admin-team/go-admin-core/sdk"
	"github.com/pkg/errors"
	"go-admin/app/rc/models"
	"go-admin/utils"
	"gorm.io/gorm"
)

type monthlyPurchaseSyncProcess struct {
}

func (t monthlyPurchaseSyncProcess) Process(contentId int64) error {
	var dataContent models.RcOriginContent
	dbContent := sdk.Runtime.GetDbByKey(dataContent.TableName())
	err := dbContent.Model(&dataContent).First(&dataContent, contentId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return errors.Wrapf(err, "get origin content error at monthlyPurchaseSyncProcess.Process, contentId: %d", contentId)
	}

	var tbMp models.RcMonthlyPurchase
	dbMp := sdk.Runtime.GetDbByKey(tbMp.TableName())
	if err := dbMp.Unscoped().Where("content_id = ?", contentId).Delete(&models.RcMonthlyPurchase{}).Error; err != nil {
		return errors.Wrap(err, "delete monthly purchase error at monthlyPurchaseSyncProcess.Process")
	}

	var contentMap map[string]any
	if err := json.Unmarshal([]byte(dataContent.Content), &contentMap); err != nil {
		return errors.Wrapf(err, "unmarshal content error at monthlyPurchaseSyncProcess.Process, contentId: %d", contentId)
	}

	monthlyPurchaseArray := contentMap[ReportDataKey].(map[string]any)["ydcgqkSTA"].([]any)
	for _, v := range monthlyPurchaseArray {
		v := v.(map[string]any)
		insertReq := models.RcMonthlyPurchase{
			ContentId:       contentId,
			AttributedMonth: t.safeGetString(v, "LAST_24M"),
			GncgM:           t.safeGetString(v, "GNCG_M"),
			JkcgM:           t.safeGetString(v, "JKCG_M"),
			HjM:             t.safeGetString(v, "HJ_M"),
		}
		insertReq.Id = utils.NewFlakeId()
		if err := dbMp.Create(&insertReq).Error; err != nil {
			return errors.Wrapf(err, "create monthly purchase error at monthlyPurchaseSyncProcess.Process, contentId: %d", contentId)
		}
	}
	return nil
}

func (t monthlyPurchaseSyncProcess) safeGetString(m map[string]any, key string) string {
	if v, ok := m[key]; ok {
		return v.(string)
	}
	return ""
}
