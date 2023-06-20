package task

import (
	"bytes"
	"encoding/json"
	"fmt"
	log "github.com/go-admin-team/go-admin-core/logger"
	"github.com/go-admin-team/go-admin-core/sdk"
	"go-admin/app/rskc/models"
	"go-admin/app/rskc/service/dto"
	"go-admin/config"
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
	return "LH_APH_SCR"
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
	var decisionReqBody dto.RcDecisionParamDecisionRequestBody
	inputParam := dto.DecisionInputParam{
		ApplyTime: time.Now().Format("2006-01-02"),
		OrderNo:   strconv.FormatInt(paramId, 10),
	}
	decisionReqBody.Assignment(&dataParam, &inputParam)
	bodyBytes, err := json.Marshal(decisionReqBody)
	if err != nil {
		return err
	}
	cli := DecisionReqClient{}
	var (
		statusCode int
		resp       []byte
	)
	statusCode, resp, err = cli.request(cli.requestUrl(), bodyBytes)
	fmt.Println(statusCode, resp)
	return nil
}

// updateDependencyDataToParam update dependency data to RcDecisionParam
func updateDependencyDataToParam(contentId int64) error {
	var tbParam models.RcDecisionParam
	dbParam := sdk.Runtime.GetDbByKey(tbParam.TableName())
	err := dbParam.Model(&tbParam).
		Where("content_id = ?", contentId).
		Order("updated_at desc").
		First(&tbParam).Error
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
	updateReq := dto.RcDecisionParamInsertReq{
		Id:        tbParam.Id,
		LhQylx:    dataDepd.LhQylx,
		LhCylwz:   dataDepd.LhCylwz,
		MdQybq:    dataDepd.LhQybq,
		GsGdct:    dataDepd.LhGdct,
		ZxYhsxqk:  dataDepd.LhYhsx,
		ZxDsfsxqk: dataDepd.LhSfsx,
	}
	var modelParam models.RcDecisionParam
	updateReq.Generate(&modelParam)
	err = dbParam.Model(&modelParam).
		Save(&modelParam).
		Error
	if err != nil {
		return err
	}

	// request decision engine
	if err := requestDecisionEngine(modelParam.Id); err != nil {
		return err
	}

	return nil
}
