package task

import (
	"encoding/json"
	"github.com/go-admin-team/go-admin-core/sdk"
	"github.com/pkg/errors"
	"go-admin/app/rskc/models"
	"go-admin/app/rskc/service/dto"
	"gorm.io/gorm"
)

func syncSellingStaFromContent(contentId int64) error {
	var dataContent models.RskcOriginContent
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
	if err := dbSst.Where("content_id = ?", contentId).Delete(&models.RcSellingSta{}).Error; err != nil {
		return err
	}

	var contentMap map[string]any
	if err := json.Unmarshal([]byte(dataContent.Content), &contentMap); err != nil {
		return nil
	}

	sellingStaArray := contentMap[ReportDataKey].(map[string]any)["sellingSTA"].([]any)
	for _, v := range sellingStaArray {
		v := v.(map[string]any)
		insertReq := dto.RcSellingStaInsertReq{
			ContentId: contentId,
			Cgje:      safeGetString(v, "CGJE"),
			Jezb:      safeGetString(v, "JEZB"),
			Ssspdl:    safeGetString(v, "SSSPDL"),
			Ssspxl:    safeGetString(v, "SSSPXL"),
			Ssspzl:    safeGetString(v, "SSSPZL"),
		}
		var insertSst models.RcSellingSta
		insertReq.Generate(&insertSst)
		if err := dbSst.Create(&insertSst).Error; err != nil {
			return err
		}
	}
	return nil
}

func safeGetString(m map[string]any, key string) string {
	if v, ok := m[key]; ok {
		return v.(string)
	}
	return ""
}
