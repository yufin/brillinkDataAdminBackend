package task

import (
	"bytes"
	"fmt"
	log "github.com/go-admin-team/go-admin-core/logger"
	"go-admin/config"
	"io"
	"net/http"
	"time"
)

type DecisionReqClient struct {
}

func (t DecisionReqClient) requestUrl() string {
	return config.ExtConfig.Vzoom.DecisionEngine.Uri +
		fmt.Sprintf("/decision-engine/decision/task/sync/%s/%s", t.sceneCode(), t.productCode())
}

func (t DecisionReqClient) requestMethod() string {
	return "POST"
}

func (t DecisionReqClient) sceneCode() string {
	return "SCENE_1"
}

func (t DecisionReqClient) productCode() string {
	return "LH_AHP_SCR"
}

func (t DecisionReqClient) Request(jsonPayload []byte) (int, []byte, error) {
	req, err := http.NewRequest(t.requestMethod(), t.requestUrl(), bytes.NewBuffer(jsonPayload))
	if err != nil {
		return 0, []byte{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "*/*")

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		log.Errorf("Send request error: url: %s, err: %v", t.requestUrl(), err)
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
