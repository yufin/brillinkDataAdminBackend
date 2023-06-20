package task

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/buger/jsonparser"
	log "github.com/go-admin-team/go-admin-core/logger"
	"github.com/go-admin-team/go-admin-core/sdk"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"go-admin/app/rskc/models"
	"go-admin/app/rskc/service/dto"
	cModels "go-admin/common/models"
	"go-admin/config"
	"go-admin/utils"
	"io"
	"net/http"
	"strconv"
	"time"
)

type DecisionReqClient struct {
}

func (t DecisionReqClient) requestUrl() string {
	return config.ExtConfig.Vzoom.DecisionEngine.Uri +
		fmt.Sprintf("/decision-engine/decision/task/sync/%s/%s", t.SceneCode(), t.ProductCode())
}

func (t DecisionReqClient) requestMethod() string {
	return "POST"
}

func (t DecisionReqClient) SceneCode() string {
	return "SCENE_1"
}

func (t DecisionReqClient) ProductCode() string {
	return "LH_AHP_SCR"
}

func (t DecisionReqClient) request(url string, jsonPayload []byte) (int, []byte, error) {
	req, err := http.NewRequest(t.requestMethod(), url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return 0, []byte{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "*/*")

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		log.Errorf("Send request error: url: %s, err: %v", url, err)
		return 0, []byte{}, err
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error("Error reading response body:", err)
		return 0, []byte{}, err
	}
	return resp.StatusCode, respBody, nil
}

func requestDecisionEngine(paramId int64) error {
	var dataParam models.RcDecisionParam
	dbParam := sdk.Runtime.GetDbByKey(dataParam.TableName())
	err := dbParam.Model(&dataParam).
		First(&dataParam, paramId).
		Error
	if err != nil {
		return err
	}
	var decisionReqParam dto.RcDecisionParamDecisionRequestBody
	inputParam := dto.DecisionInputParam{
		ApplyTime: time.Now().Format("2006-01-02"),
		OrderNo:   strconv.FormatInt(paramId, 10),
	}
	decisionReqParam.Assignment(&dataParam, &inputParam)
	decisionReqBody := map[string]any{
		"param": decisionReqParam,
	}
	bodyBytes, err := json.Marshal(decisionReqBody)
	if err != nil {
		return err
	}
	cli := DecisionReqClient{}
	var (
		statusCode int
		resp       []byte
		respCode   string
		respMsg    string
	)
	statusCode, resp, err = cli.request(cli.requestUrl(), bodyBytes)
	if err != nil {
		return err
	}
	if statusCode != 200 {
		return errors.Errorf("request statusCode: %d, err: url: %s, err: %v", statusCode, cli.requestUrl(), err)
	}
	respCode, err = jsonparser.GetString(resp, "code")
	if err != nil {
		return err
	}
	respMsg, _ = jsonparser.GetString(resp, "msg")
	if respCode != "000000" {
		return errors.Errorf("decision flow resp Code != 000000, msg:%s", respMsg)
	}
	if err := saveDecisionResult(resp, paramId); err != nil {
		return err
	}
	return nil
}

func saveDecisionResult(resp []byte, paramId int64) error {
	var (
		taskId       string
		finalResult  string
		ahpScore     float64
		fxSwJxccClnx string
		lhQylx       int64
	)

	msg, err := jsonparser.GetString(resp, "msg")
	if err != nil {
		return err
	}
	taskId, err = jsonparser.GetString(resp, "data", "object", "taskId")
	if err != nil {
		return err
	}
	finalResult, err = jsonparser.GetString(resp, "data", "object", "result", "final_result")
	if err != nil {
		return err
	}
	ahpScore, err = jsonparser.GetFloat(resp, "data", "object", "result", "AHP_SCORE")
	if err != nil {
		return err
	}
	fxSwJxccClnx, err = jsonparser.GetString(resp, "data", "object", "result", "fx_sw_jxcc_clnx")
	if err != nil {
		return err
	}
	lhQylx, err = jsonparser.GetInt(resp, "data", "object", "result", "lh_qylx")
	if err != nil {
		return err
	}
	decisionResult := models.RcDecisionResult{
		Model:        cModels.Model{Id: utils.NewFlakeId()},
		ParamId:      paramId,
		TaskId:       taskId,
		FinalResult:  finalResult,
		AhpScore:     decimal.NullDecimal{Decimal: decimal.NewFromFloat(ahpScore), Valid: true},
		FxSwJxccClnx: fxSwJxccClnx,
		LhQylx:       int(lhQylx),
		Msg:          msg,
	}
	dbRes := sdk.Runtime.GetDbByKey(decisionResult.TableName())
	if err := dbRes.Create(&decisionResult).Error; err != nil {
		return err
	}
	return nil
}

// updateDependencyDataToParam update dependency data to RcDecisionParam
func updateDependencyDataToParam(contentId int64) error {
	var dtParam models.RcDecisionParam
	dbParam := sdk.Runtime.GetDbByKey(dtParam.TableName())
	err := dbParam.Model(&models.RcDecisionParam{}).
		Where("content_id = ?", contentId).
		Order("updated_at desc").
		First(&dtParam).Error
	if err != nil {
		return err
	}
	var dataDepd models.RcDependencyData
	dbDepd := sdk.Runtime.GetDbByKey(dataDepd.TableName())
	err = dbDepd.Model(&dataDepd).
		Where("content_id = ?", contentId).
		Order("updated_at desc").
		First(&dataDepd).Error
	if err != nil {
		return err
	}
	dtParam.LhQylx = dataDepd.LhQylx
	dtParam.LhCylwz = dataDepd.LhCylwz
	dtParam.MdQybq = dataDepd.LhQybq
	dtParam.GsGdct = dataDepd.LhGdct
	dtParam.ZxYhsxqk = dataDepd.LhYhsx
	dtParam.ZxDsfsxqk = dataDepd.LhSfsx
	err = dbParam.
		Save(&dtParam).
		Error
	if err != nil {
		return err
	}

	// request decision engine
	if err := requestDecisionEngine(dtParam.Id); err != nil {
		return err
	}

	return nil
}
