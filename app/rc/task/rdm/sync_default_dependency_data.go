package rdm

import (
	"github.com/go-admin-team/go-admin-core/sdk"
	"go-admin/app/rc/models"
	"go-admin/utils"
)

type defaultDepDataSyncVerifyModel struct {
	RocId     int64
	DepId     int64
	ContentId int64
	UscId     string
	YearMonth string
}

type SyncDefaultDependencyParamProcess struct {
}

func (a SyncDefaultDependencyParamProcess) Process() error {
	// auto sync rc_dependency_data default with null contentId
	var modelRdd models.RcDependencyData
	db := sdk.Runtime.GetDbByKey(modelRdd.TableName())

	createdByList := make([]int64, 0)
	err := db.Model(&modelRdd).
		Distinct("create_by").
		Pluck("create_by", &createdByList).
		Error
	if err != nil {
		return err
	}

	//var tbContent models.RcOriginContent
	//dbRoc := sdk.Runtime.GetDbByKey(tbContent.TableName())
	for _, createdBy := range createdByList {
		uscIds := make([]string, 0)
		err := db.Model(&models.RcDependencyData{}).
			Distinct("usc_id").
			Where("create_by = ?", createdBy).
			Pluck("usc_id", &uscIds).
			Error
		if err != nil {
			return err
		}

		// loop through rdd.uscIds
		for _, uscId := range uscIds {
			var defaultDeps []defaultDepDataSyncVerifyModel
			err = db.Raw(
				`select roc.id as roc_id,rdd.id as dep_id, roc.year_month, content_id, roc.usc_id as usc_id
					from (select id, usc_id, rc_origin_content.year_month, row_number() over (partition by usc_id) as n from rc_origin_content) roc
							 left join (select id,
											   content_id,
											   usc_id,
											   row_number() over (partition by usc_id order by created_at desc ) as nrdd
										from rc_dependency_data where create_by = ?) rdd
									   on roc.usc_id = rdd.usc_id and roc.n = rdd.nrdd
					where roc.usc_id = ? order by rdd.id desc;`, createdBy, uscId).
				Scan(&defaultDeps).
				Error
			if err != nil {
				return err
			}
			for _, dd := range defaultDeps {
				if dd.DepId == 0 {
					dd := dd
					// insert new to rdd with rdd.contentId = rocId
					//which data from rdd where contentId is not null or 0
					dRdd, err := a.getDefaultRdd(createdBy, uscId)
					if err != nil {
						return err
					}
					rdd2Create := models.RcDependencyData{
						ContentId:       dd.RocId,
						AttributedMonth: dd.YearMonth,
						UscId:           dd.UscId,
						LhQylx:          dRdd.LhQylx,
						LhCylwz:         dRdd.LhCylwz,
						LhGdct:          dRdd.LhGdct,
						LhYhsx:          dRdd.LhYhsx,
						LhSfsx:          dRdd.LhSfsx,
						LhQybq:          dRdd.LhQybq,
						AdditionData:    dRdd.AdditionData,
					}
					rdd2Create.CreateBy = createdBy
					rdd2Create.Id = utils.NewFlakeId()
					err = db.Create(&rdd2Create).Error
					if err != nil {
						return err
					}
				} else if dd.ContentId == 0 {
					// update rdd with rdd.contentId = rocId
					if err = db.Model(&models.RcDependencyData{}).
						Where("id = ?", dd.DepId).
						Updates(map[string]interface{}{
							"content_id":       dd.RocId,
							"attributed_month": dd.YearMonth,
						}).Error; err != nil {
						return err
					}
				}
			}
		}
	}
	return nil
}

func (a SyncDefaultDependencyParamProcess) getDefaultRdd(createBy int64, uscId string) (*models.RcDependencyData, error) {
	modelRdd := models.RcDependencyData{}
	db := sdk.Runtime.GetDbByKey(modelRdd.TableName())
	err := db.Model(&modelRdd).
		Where("create_by = ?", createBy).
		Where("usc_id = ?", uscId).
		Where("content_id is not null").
		Order("created_at desc").
		First(&modelRdd).
		Error
	if err != nil {
		return nil, err
	}
	return &modelRdd, nil
}
