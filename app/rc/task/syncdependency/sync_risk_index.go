package syncdependency

import (
	"encoding/json"
	"github.com/go-admin-team/go-admin-core/sdk"
	"github.com/pkg/errors"
	"go-admin/app/rc/models"
	"go-admin/utils"
	"gorm.io/gorm"
)

// TODO: Finish this

type riskIndexSyncProcess struct {
}

func (p riskIndexSyncProcess) Process(contentId int64) error {
	var dataContent models.RcOriginContent
	dbContent := sdk.Runtime.GetDbByKey(dataContent.TableName())
	err := dbContent.Model(&dataContent).First(&dataContent, contentId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
	}

	var dtSsi models.RcRiskIndex
	dbSsi := sdk.Runtime.GetDbByKey(dtSsi.TableName())
	if err := dbSsi.Where("content_id = ?", contentId).Delete(&models.RcRiskIndex{}).Error; err != nil {
		return err
	}

	var contentMap map[string]any
	if err := json.Unmarshal([]byte(dataContent.Content), &contentMap); err != nil {
		return err
	}
	riskIndexes := contentMap[ReportDataKey].(map[string]any)["riskIndexes"].([]any)
	for _, v := range riskIndexes {
		v := v.(map[string]any)
		insertReq := models.RcRiskIndex{
			ContentId: contentId,
			RiskDesc:  p.safeGetString(v, "RISK_DEC"),
			Index:     p.safeGetString(v, "INDEX_DEC"),
			Value:     p.safeGetString(v, "INDEX_VALUE"),
			Flag:      p.safeGetString(v, "INDEX_FLAG"),
		}
		insertReq.Id = utils.NewFlakeId()
		if err := dbSsi.Create(&insertReq).Error; err != nil {
			return err
		}
	}
	return nil
}

func (p riskIndexSyncProcess) safeGetString(m map[string]any, key string) string {
	if v, ok := m[key]; ok {
		return v.(string)
	}
	return ""
}
