package syncdependency

import (
	"encoding/json"
	"github.com/go-admin-team/go-admin-core/sdk"
	"github.com/pkg/errors"
	"go-admin/app/rc/models"
	"go-admin/utils"
	"gorm.io/gorm"
)

type sellingStaSyncProcess struct {
}

func (t sellingStaSyncProcess) Process(contentId int64) error {
	var dataContent models.RcOriginContent
	dbContent := sdk.Runtime.GetDbByKey(dataContent.TableName())
	err := dbContent.Model(&dataContent).First(&dataContent, contentId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
	}

	var tbSst models.RcSellingSta
	dbSst := sdk.Runtime.GetDbByKey(tbSst.TableName())
	if err := dbSst.Unscoped().Where("content_id = ?", contentId).Delete(&models.RcSellingSta{}).Error; err != nil {
		return errors.Wrapf(err, "delete rc_selling_sta error at sellingStaSyncProcess.Process, contentId: %d", contentId)
	}

	var contentMap map[string]any
	if err := json.Unmarshal([]byte(dataContent.Content), &contentMap); err != nil {
		// return nil
		return errors.Wrapf(err, "unmarshal content error at sellingStaSyncProcess.Process, contentId: %d", contentId)
	}

	sellingStaArray := contentMap[ReportDataKey].(map[string]any)["sellingSTA"].([]any)
	for _, v := range sellingStaArray {
		v := v.(map[string]any)
		insertReq := models.RcSellingSta{
			ContentId: contentId,
			Cgje:      t.safeGetString(v, "CGJE"),
			Jezb:      t.safeGetString(v, "JEZB"),
			Ssspdl:    t.safeGetString(v, "SSSPDL"),
			Ssspxl:    t.safeGetString(v, "SSSPXL"),
			Ssspzl:    t.safeGetString(v, "SSSPZL"),
		}
		insertReq.Id = utils.NewFlakeId()
		if err := dbSst.Create(&insertReq).Error; err != nil {
			return err
		}
	}
	return nil
}

func (t sellingStaSyncProcess) safeGetString(m map[string]any, key string) string {
	if v, ok := m[key]; ok {
		return v.(string)
	}
	return ""
}
