/*
@Author : Akiraka
@Time : 2022/2/21 15:56
@Software: GoLand
*/

package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// PostRequest Akiraka 20220218 请求POST方法
func PostRequest(model interface{}, url string) (body []byte, token string) {
	//请求地址模板
	data, err := json.Marshal(model)
	//创建一个请求
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		// handle error
	}
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	//设置请求头
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	//发送请求
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("发送request失败，失败信息%s\n", err.Error())
	}
	// Akiraka 20220221 接收返回的 Body
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("获取body失败，失败信息%s\n", err.Error())
	}
	// 返回华为云 Token 出去
	token = resp.Header.Get("X-Subject-Token")
	//关闭请求
	defer resp.Body.Close()

	// 返回 body 出去
	return body, token
}

// GetRequest Akiraka 20220221 请求GET方法
func GetRequest(model interface{}, url string) (body []byte, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	//设置请求头
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	//发送请求
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	// Akiraka 20220221 接收返回的 Body
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("获取body失败，失败信息%s\n", err.Error())
	}
	//关闭请求
	defer resp.Body.Close()
	// 返回 body 出去
	return body, err
}
