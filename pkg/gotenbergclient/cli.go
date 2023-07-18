package gotenbergclient

import (
	"bytes"
	"errors"
	"fmt"
	log "github.com/go-admin-team/go-admin-core/logger"
	"mime/multipart"
	"net/http"
)

type GtbCli interface {
	ChromiumConvert(url string) (*http.Response, error)
	request(url string, payload map[string]any) (*http.Response, error)
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

	payload := map[string]interface{}{
		"url": url,
	}
	resp, err := g.request(api, payload)
	//defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("response code is not 200: %d", resp.StatusCode))
	}
	return resp, nil
}

func (g GtbClient) request(url string, payload map[string]any) (*http.Response, error) {
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	for k, v := range payload {
		fw, _ := w.CreateFormField(k)
		_, err := fw.Write([]byte(v.(string)))
		if err != nil {
			panic(err)
		}
	}
	w.Close()
	resp, err := http.Post(url, w.FormDataContentType(), &b)
	if err != nil {
		log.Errorf("Send request error: url: %s, err: %v", url, err)
		return nil, err
	}
	return resp, nil
}
