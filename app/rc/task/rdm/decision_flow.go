package rdm

import (
	"encoding/binary"
	"fmt"
	log "github.com/go-admin-team/go-admin-core/logger"
	"github.com/go-admin-team/go-admin-core/sdk"
	"github.com/nats-io/nats.go"
	"go-admin/app/rc/models"
	"go-admin/pkg/natsclient"
	"go-admin/utils"
	"sync/atomic"
	"time"
)

// todo:缺少contentId的校验逻辑
var decisionRunning int32

type RdmPipeline interface {
	Pipeline() error
}

type AhpRdmTask struct {
}

func (t AhpRdmTask) Exec(arg interface{}) error {
	if atomic.LoadInt32(&decisionRunning) == 1 {
		log.Info("DecisionFlow任务已经在执行中，跳过本次调度")
		return nil
	}
	atomic.StoreInt32(&decisionRunning, 1)
	defer atomic.StoreInt32(&decisionRunning, 0)

	if err := SyncDefaultDependencyParam(); err != nil {
		log.Errorf("SyncDefaultDependencyParam Failed:%v \r\n", err)
		return err
	}

	if err := pubIdsForRdm(); err != nil {
		log.Errorf("selectWaitForRequest Failed:%v \r\n", err)
		return err
	}

	totalPending, _, err := natsclient.SubToRequestDecisionNew.Pending()
	if err == nil {
		fmt.Println("DecisionFlowTask msg totalPending:", totalPending)
	}
	for {
		msgs, err := natsclient.SubToRequestDecisionNew.Fetch(1, nats.MaxWait(5*time.Second))
		if err != nil {
			if err == nats.ErrTimeout {
				return nil
			} else {
				return err
			}
		}
		for _, msg := range msgs {
			depId := int64(binary.BigEndian.Uint64(msg.Data))
			log.Infof("开始请求决策引擎: depId = %d\r\n", depId)

			var ahp RdmPipeline
			ahp = PySidecarAhpRdm{depId: depId, AppType: 1}
			if err := ahp.Pipeline(); err != nil {
				log.Errorf("PySidecarAhpRdm.Pipeline depId:%v Failed:%v \r\n", depId, err)
				return err
			}
			if err := msg.AckSync(); err != nil {
				return err
			}
		}
	}
}

func SyncDefaultDependencyParam() error {
	// 自动同步 rc_dependency_data default with null contentId
	var modelRdd models.RcDependencyData
	db := sdk.Runtime.GetDbByKey(modelRdd.TableName())

	userIds := make([]int64, 0)
	err := db.Model(&modelRdd).
		Distinct("create_by").
		Pluck("create_by", &userIds).
		Error
	if err != nil {
		return err
	}

	var tbContent models.RcOriginContent
	dbRoc := sdk.Runtime.GetDbByKey(tbContent.TableName())
	for _, userId := range userIds {
		userId := userId
		uscIds := make([]string, 0)
		err := db.Model(&modelRdd).
			Distinct("usc_id").
			Where("create_by = ?", userId).
			Pluck("usc_id", &uscIds).
			Error
		if err != nil {
			return err
		}

		var modelContentInfo models.RcOriginContentInfo
		for _, uscId := range uscIds {
			uscId := uscId
			contentInfos := make([]models.RcOriginContentInfo, 0)
			err = dbRoc.
				Model(&modelContentInfo).
				Select("id").
				Where("usc_id = ?", uscId).
				Order("created_at").
				Scan(&contentInfos).
				Error
			if err != nil {
				return err
			}
			// make sure the rows query by createBy and uscId has same amount with contentIds
			rddList := make([]models.RcDependencyData, 0)
			err = db.Model(&modelRdd).
				Where("usc_id = ?", uscId).
				Where("create_by = ?", userId).
				Order("created_at").
				Scan(&rddList).
				Error
			if err != nil {
				return err
			}

			// get contentId which not in rddList.contentId
			rocNotExists := make([]models.RcOriginContentInfo, 0)
			for _, cItem := range contentInfos {
				cid := cItem.Id
				found := false
				for _, rddItem := range rddList {
					rddCid := rddItem.ContentId
					if cid == rddCid {
						found = true
						break
					}
				}
				if !found {
					rocNotExists = append(rocNotExists, cItem)
				}
			}
			// get rdd with null contentId
			defaultRdd := rddList[len(rddList)-1]
			if len(rocNotExists) > 0 {
				nullContentIdsRdds := make([]models.RcDependencyData, 0)
				toUpdateRdds := make([]models.RcDependencyData, 0)

				for _, rddNullable := range rddList {
					if rddNullable.ContentId == 0 {
						nullContentIdsRdds = append(nullContentIdsRdds, rddNullable)
					}
				}
				for i, cItem2Update := range rocNotExists {
					if i+1 <= len(nullContentIdsRdds) {
						//nullContentIdsRdds[i].ContentId = cItem2Update.Id
						toUpdateRdd := nullContentIdsRdds[i]
						toUpdateRdd.ContentId = cItem2Update.Id
						toUpdateRdd.AttributedMonth = cItem2Update.YearMonth
						toUpdateRdds = append(toUpdateRdds, toUpdateRdd)
					} else {
						toUpdateRdd := defaultRdd
						toUpdateRdd.Id = utils.NewFlakeId()
						toUpdateRdd.ContentId = cItem2Update.Id
						toUpdateRdd.AttributedMonth = cItem2Update.YearMonth
						toUpdateRdds = append(toUpdateRdds, toUpdateRdd)
					}
				}
				if err := db.Model(&modelRdd).Save(&toUpdateRdds).Error; err != nil {
					return err
				}
			}
		}
	}

	return nil
}

// pubIdsToRequestDecision query rows in dependency data where has not been requested, pub to queue.
func pubIdsForRdm() error {
	var tbDep models.RcDependencyData
	db := sdk.Runtime.GetDbByKey(tbDep.TableName())
	depIds := make([]int64, 0)
	err := db.
		Table(tbDep.TableName()).
		Select("rc_dependency_data.id as dep_id").
		Joins("LEFT JOIN rc_rdm_result rrr ON rc_dependency_data.id = rrr.dep_id").
		Where("rc_dependency_data.content_id IS NOT NULL and rc_dependency_data.content_id IS != 0 AND rrr.dep_id IS NULL").
		Pluck("dep_id", &depIds).
		Error
	if err != nil {
		return err
	}
	if len(depIds) == 0 {
		return nil
	}
	for _, depId := range depIds {
		msg := make([]byte, 8)
		binary.BigEndian.PutUint64(msg, uint64(depId))
		// TODO: try this (idempotent message writes)
		//m := nats.NewMsg(natsclient.TopicToRequestDecisionNew)
		//m.Data = msg
		//m.Header.Set("Nats-Msg-Id", "unique-id-123")
		//_, err := natsclient.TaskJs.PublishMsg(m)

		_, err := natsclient.TaskJs.Publish(natsclient.TopicToRequestDecisionNew, msg)
		if err != nil {
			return err
		}
		err = db.Model(models.RcDependencyData{}).
			Where("id = ?", depId).
			Update("status_code", 1).
			Error
		if err != nil {
			return err
		}
	}
	return nil
}

//func requestDecisionEngine(depId int64, cli DecisionCli) error {
//	var dataDep models.RcDependencyData
//	dbDep := sdk.Runtime.GetDbByKey(dataDep.TableName())
//	err := dbDep.Model(&dataDep).First(&dataDep, depId).Error
//	if err != nil {
//		if errors.Is(err, gorm.ErrRecordNotFound) {
//			return nil
//		}
//		return err
//	}
//
//	var dataParam models.RcDecisionParam
//	dbParam := sdk.Runtime.GetDbByKey(dataParam.TableName())
//	err = dbParam.Model(&dataParam).
//		Where("content_id = ?", dataDep.ContentId).
//		First(&dataParam).
//		Error
//	if err != nil {
//		if errors.Is(err, gorm.ErrRecordNotFound) {
//			return nil
//		}
//		return err
//	}
//
//	// assign param
//	dataParam.LhQylx = dataDep.LhQylx
//	dataParam.LhCylwz = dataDep.LhCylwz
//	dataParam.GsGdct = dataDep.LhGdct
//	dataParam.ZxYhsxqk = dataDep.LhYhsx
//	dataParam.ZxDsfsxqk = dataDep.LhSfsx
//	dataParam.MdQybq = dataDep.LhQybq
//	var decisionReqParam dto.RcDecisionParamDecisionRequestBody
//	inputParam := dto.DecisionInputParam{
//		ApplyTime: time.Now().Format("2006-01-02"),
//		OrderNo:   strconv.FormatInt(depId, 10),
//	}
//	decisionReqParam.Assignment(&dataParam, &inputParam)
//	decisionReqBody := map[string]any{
//		"param": decisionReqParam,
//	}
//	bodyBytes, err := json.Marshal(decisionReqBody)
//	if err != nil {
//		return err
//	}
//	var (
//		statusCode int
//		resp       []byte
//		respCode   string
//		respMsg    string
//	)
//	statusCode, resp, err = cli.Request(bodyBytes)
//	if err != nil {
//		return err
//	}
//	if statusCode != 200 {
//		return errors.Errorf("request statusCode: %d, err: url: %s, err: %v", statusCode, cli.RequestUrl(), err)
//	}
//	respCode, err = jsonparser.GetString(resp, "code")
//	if err != nil {
//		return err
//	}
//	respMsg, _ = jsonparser.GetString(resp, "msg")
//	if respCode != "000000" {
//		return errors.Errorf("decision flow resp Code != 000000, msg:%s", respMsg)
//	}
//	if err := saveDecisionResult(resp, depId); err != nil {
//		return err
//	}
//	return nil
//}
//
//func saveDecisionResult(resp []byte, depId int64) error {
//	var (
//		taskId       string
//		finalResult  string
//		ahpScore     float64
//		fxSwJxccClnx string
//		lhQylx       int64
//	)
//
//	msg, err := jsonparser.GetString(resp, "msg")
//	if err != nil {
//		return err
//	}
//	taskId, err = jsonparser.GetString(resp, "data", "object", "taskId")
//	if err != nil {
//		return err
//	}
//	finalResult, err = jsonparser.GetString(resp, "data", "object", "result", "final_result")
//	if err != nil {
//		return err
//	}
//	ahpScore, err = jsonparser.GetFloat(resp, "data", "object", "result", "AHP_SCORE")
//	if err != nil {
//		return err
//	}
//	fxSwJxccClnx, err = jsonparser.GetString(resp, "data", "object", "result", "fx_sw_jxcc_clnx")
//	if err != nil {
//		return err
//	}
//	lhQylx, err = jsonparser.GetInt(resp, "data", "object", "result", "lh_qylx")
//	if err != nil {
//		return err
//	}
//	decisionResult := models.RcDecisionResult{
//		Model:        cModels.Model{Id: utils.NewFlakeId()},
//		ResId:        depId,
//		TaskId:       taskId,
//		FinalResult:  finalResult,
//		AhpScore:     decimal.NullDecimal{Decimal: decimal.NewFromFloat(ahpScore), Valid: true},
//		FxSwJxccClnx: fxSwJxccClnx,
//		LhQylx:       int(lhQylx),
//		Msg:          msg,
//	}
//	dbRes := sdk.Runtime.GetDbByKey(decisionResult.TableName())
//	if err := dbRes.Create(&decisionResult).Error; err != nil {
//		return err
//	}
//	return nil
//}

//// updateDependencyDataToParam update dependency data to RcDecisionParam
//func updateDependencyDataToParam(contentId int64) error {
//	var dtParam models.RcDecisionParam
//	dbParam := sdk.Runtime.GetDbByKey(dtParam.TableName())
//	err := dbParam.Model(&models.RcDecisionParam{}).
//		Where("content_id = ?", contentId).
//		Order("updated_at desc").
//		First(&dtParam).Error
//	if err != nil {
//		return err
//	}
//	var dataDepd models.RcDependencyData
//	dbDepd := sdk.Runtime.GetDbByKey(dataDepd.TableName())
//	err = dbDepd.Model(&dataDepd).
//		Where("content_id = ?", contentId).
//		Order("updated_at desc").
//		First(&dataDepd).Error
//	if err != nil {
//		return err
//	}
//	dtParam.LhQylx = dataDepd.LhQylx
//	dtParam.LhCylwz = dataDepd.LhCylwz
//	dtParam.MdQybq = dataDepd.LhQybq
//	dtParam.GsGdct = dataDepd.LhGdct
//	dtParam.ZxYhsxqk = dataDepd.LhYhsx
//	dtParam.ZxDsfsxqk = dataDepd.LhSfsx
//	err = dbParam.
//		Save(&dtParam).
//		Error
//	if err != nil {
//		return err
//	}
//
//	// request decision engine
//	if err := requestDecisionEngine(dtParam.Id); err != nil {
//		return err
//	}
//
//	return nil
//}
