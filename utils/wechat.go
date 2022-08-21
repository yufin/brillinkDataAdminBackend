package utils

import (
	"fmt"
	"github.com/go-admin-team/go-admin-core/logger"
	"github.com/go-admin-team/go-admin-core/sdk"
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
)

type TexyMsg struct {
	Msgtype string `json:"msgtype"`
	Text    `json:"text"`
}

type Text struct {
	Content       string   `json:"content"`
	MentionedList []string `json:"mentioned_list"`
}

type MarkdownMsg struct {
	Msgtype  string   `json:"msgtype"`
	Markdown Markdown `json:"markdown"`
}

type Markdown struct {
	Content string `json:"content"`
}

func SendWechatAlert(url, msg string) {
	var m TexyMsg
	m.Msgtype = "text"
	m.Text.Content = msg
	m.Text.MentionedList = append(m.Text.MentionedList, "@all")

	client := resty.New()
	resp, err := client.R().SetHeader("Content-Type", "application/json").SetBody(m).Post(url)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(resp.String())
	fmt.Println(resp.Request.Body)
}

func SendWechatAlertWithErr(requestId, method, uri, ip, abInfo, abSource, abFunc string) {
	url, ok := sdk.Runtime.GetConfig("sys_wechat_webhook").(string)
	if !ok || url == "" {
		logger.Warn(errors.New("config sys_wechat_webhook value is nil"))
		return
	}
	var m MarkdownMsg
	m.Msgtype = "markdown"
	m.Markdown.Content = "接口请求错误反馈请求id:<font color=\"warning\">" + requestId + "</font>，请相关同事注意。\n         " +
		">请求方式:<font color=\"warning\">" + method + "</font>\n" +
		">请求地址:<font color=\"warning\">" + uri + "</font>\n" +
		">IP信息:<font color=\"warning\">" + ip + "</font>\n" +
		">异常信息:<font color=\"comment\">" + abInfo + "</font>\n" +
		">异常来源:<font color=\"comment\">" + abSource + "</font>\n" +
		">异常方法:<font color=\"comment\">" + abFunc + "</font>"

	client := resty.New()
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(m).
		Post(url)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(resp.String())
	fmt.Println(resp.Request.Body)
}
