package v3

import (
	"fmt"
	"github.com/buger/jsonparser"
	"github.com/go-admin-team/go-admin-core/sdk"
	"go-admin/app/rc/models"
	spModels "go-admin/app/spider/models"
)

type ClaBusinessInfo struct {
	content   *[]byte
	contentId int64
}

func (s *ClaBusinessInfo) Collating() error {
	modelRoc := models.RcOriginContentInfo{}
	dbRoc := sdk.Runtime.GetDbByKey(modelRoc.TableName())
	err := dbRoc.Model(&modelRoc).First(&modelRoc, s.contentId).Error
	if err != nil {
		return err
	}

	modelInfo := spModels.EnterpriseInfo{}
	db := sdk.Runtime.GetDbByKey(modelInfo.TableName())
	err = db.Model(&modelInfo).
		Where("usc_id = ?", modelRoc.UscId).
		First(&modelInfo).
		Error
	if err != nil {
		return err
	}

	tempC, err := jsonparser.Set(*s.content, []byte(fmt.Sprintf(`"%s"`, modelInfo.PaidInCapital)), "impExpEntReport", "businessInfo", "capitalPaidIn")
	if err != nil {
		return err
	}
	*s.content = tempC
	return nil
}

func (s *ClaBusinessInfo) SetContent(content *[]byte, contentId int64) {
	s.content = content
	s.contentId = contentId
}
