package syncdependency

import (
	"github.com/go-admin-team/go-admin-core/sdk"
	"github.com/pkg/errors"
	"go-admin/app/rc/models"
	cModels "go-admin/common/models"
	"go-admin/utils"
	"gorm.io/gorm"
	"time"
)

type dmShareholderSyncProcess struct {
}

func (t dmShareholderSyncProcess) Process(contentId int64) error {
	var dataContent models.RcOriginContent
	dbContent := sdk.Runtime.GetDbByKey(dataContent.TableName())
	err := dbContent.Model(&dataContent).First(&dataContent, contentId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return errors.Wrapf(err, "get origin content error at dmShareholderSyncProcess.Process, contentId: %d", contentId)
	}

	var modelEs models.EnterpriseShareholderWaitList
	var exists int64

	//now, err := time.ParseInLocation("2006-01-02 15:04:05", dataContent.YearMonth+"-01 00:00:00", time.Local)
	now := time.Now()
	if err != nil {
		return errors.Wrapf(err, "parse time error at dmShareholderSyncProcess.Process, contentId: %d", contentId)
	}

	dbDss := sdk.Runtime.GetDbByKey(modelEs.TableName())
	err = dbDss.Model(&models.EnterpriseShareholderWaitList{}).
		Where("valid_start_time <= ?", now).
		Where("valid_end_time > ?", now).
		Where("usc_id = ?", dataContent.UscId).
		Count(&exists).
		Error
	if err != nil {
		return errors.Wrapf(err, "get dm enterprise shareholder error at dmShareholderSyncProcess.Process, contentId: %d", contentId)
	}

	if exists > 0 {
		return nil
	}

	validEnd := now.AddDate(0, 2, 0)

	insertReq := models.EnterpriseShareholderWaitList{
		Model:          cModels.Model{Id: utils.NewFlakeId()},
		UscId:          dataContent.UscId,
		ValidStartTime: now,
		ValidEndTime:   validEnd,
	}
	if err := dbDss.Create(&insertReq).Error; err != nil {
		return errors.Wrapf(err, "create dm enterprise shareholder error at dmShareholderSyncProcess.Process, contentId: %d", contentId)
	}
	return nil
}
