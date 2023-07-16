package gotenbergclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	log "github.com/go-admin-team/go-admin-core/logger"
	"net/http"
	"time"
)

type GtbCli interface {
	ChromiumConvert(url string) (*http.Response, error)
	request(url string, method string, body []byte) (*http.Response, error)
}

type GtbClient struct {
	Uri string
}

func NewGtbClient(uri string) *GtbClient {
	return &GtbClient{
		Uri: uri,
	}
}

func (g GtbClient) ChromiumConvert(url string) (*http.Response, error) {
	api := g.Uri + "/forms/chromium/convert/url"
	body := map[string]any{"url": url}
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	resp, err := g.request(api, "POST", bodyBytes)
	//defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("response code is not 200: %d", resp.StatusCode))
	}
	return resp, nil
}

func (g GtbClient) request(url string, method string, body []byte) (*http.Response, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "*/*")
	httpCli := &http.Client{Timeout: 5 * time.Second}
	resp, err := httpCli.Do(req)
	if err != nil {
		log.Errorf("Send request error: url: %s, err: %v", url, err)
		return nil, err
	}
	return resp, nil
}
