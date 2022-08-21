package models

import (
	"bytes"
	"encoding/json"
	"github.com/bitly/go-simplejson"
	log "github.com/go-admin-team/go-admin-core/logger"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"io/ioutil"
	"regexp"
	"strings"

	"github.com/go-admin-team/go-admin-core/sdk"
	"github.com/go-admin-team/go-admin-core/sdk/runtime"
	"github.com/go-admin-team/go-admin-core/storage"

	"go-admin/common/models"
)

type SysApi struct {
	Id        int    `json:"id" gorm:"primaryKey;autoIncrement;comment:主键编码"`
	Handle    string `json:"handle" gorm:"size:128;comment:handle"`
	Name      string `json:"name" gorm:"size:128;comment:标题"`
	Path      string `json:"path" gorm:"size:128;comment:地址"`
	Method    string `json:"method" gorm:"size:16;comment:请求类型"`
	Type      string `json:"type" gorm:"size:16;comment:接口类型"`
	Project   string `json:"project" gorm:"size:128;comment:项目"`
	Bus       string `json:"bus" gorm:"size:128;comment:业务模块"`
	Func      string `json:"func" gorm:"size:128;comment:func"`
	IsHistory bool   `json:"isHistory" gorm:"comment:是否历史接口"`
	models.ModelTime
	models.ControlBy
}

func (SysApi) TableName() string {
	return "sys_api"
}

func (e *SysApi) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *SysApi) GetId() interface{} {
	return e.Id
}

func SaveSysApi(message storage.Messager) (err error) {
	var rb []byte
	rb, err = json.Marshal(message.GetValues())
	if err != nil {
		err = errors.WithStack(err)
		return err
	}

	var l runtime.Routers
	err = json.Unmarshal(rb, &l)
	if err != nil {
		err = errors.WithStack(err)
		return err
	}
	dbList := sdk.Runtime.GetDb()
	for _, d := range dbList {
		for _, v := range l.List {
			if v.HttpMethod != "HEAD" ||
				strings.Contains(v.RelativePath, "/swagger/") ||
				strings.Contains(v.RelativePath, "/static/") ||
				strings.Contains(v.RelativePath, "/form-generator/") ||
				strings.Contains(v.RelativePath, "/sys/tables") {
				apiName := getApiName(v.RelativePath, v.HttpMethod)
				model := new(SysApi)
				Project, Bus, Func := analysisRouter(v)
				err := d.Where(SysApi{Path: v.RelativePath, Method: v.HttpMethod}).
					Attrs(SysApi{Handle: v.Handler, Project: Project, Bus: Bus, Func: Func, Name: apiName}).
					FirstOrCreate(model).
					Error
				if err != nil {
					err = errors.WithStack(err)
					return err
				}
				if model.Func == "" {
					err = updateApi(v, model, d)
					if err != nil {
						err = errors.WithStack(err)
						return err
					}
				}
			}
		}
		err = checkApi(&l.List, d)
		if err != nil {
			return err
		}
	}
	return nil
}

func getApiName(path, method string) string {
	// 根据接口方法注释里的@Summary填充接口名称，适用于代码生成器
	// 可在此处增加配置路径前缀的if判断，只对代码生成的自建应用进行定向的接口名称填充
	jsonFile, err := ioutil.ReadFile("docs/swagger.json")
	if err != nil {
		log.Error(err)
		return ""
	}
	jsonData, err := simplejson.NewFromReader(bytes.NewReader(jsonFile))
	if err != nil {
		log.Error(err)
		return ""
	}
	urlPath := path
	idPatten := "(.*)/:(\\w+)" // 正则替换，把:id换成{id}
	reg, _ := regexp.Compile(idPatten)
	if reg.MatchString(urlPath) {
		urlPath = reg.ReplaceAllString(path, "${1}/{${2}}") // 把:id换成{id}
	}
	apiTitle, _ := jsonData.Get("paths").Get(urlPath).Get(strings.ToLower(method)).Get("summary").String()
	return apiTitle
}

func checkApi(routers *[]runtime.Router, db *gorm.DB) (err error) {
	list := new([]SysApi)
	err = db.Find(list).Error
	if err != nil {
		err = errors.WithStack(err)
		return err
	}

	for _, api := range *list {
		bl := true
		for _, h := range *routers {
			if api.Path == h.RelativePath && api.Method == h.HttpMethod {
				bl = false
			}
		}
		if bl {
			err = db.Model(api).Update("is_history", true).Error
			if err != nil {
				err = errors.WithStack(err)
				return err
			}
		}
		if api.Name == "" {
			apiTitle := getApiName(api.Path, api.Method)
			if apiTitle == "" {
				continue
			}
			err = db.Model(api).Update("title", apiTitle).Error
			if err != nil {
				err = errors.WithStack(err)
				return err
			}
		}
	}
	return
}

func updateApi(v runtime.Router, model *SysApi, d *gorm.DB) (err error) {
	model.Project, model.Bus, model.Func = analysisRouter(v)
	if model.Project == "" || model.Bus == "" {
		return
	}
	err = d.Save(model).Error
	if err != nil {
		err = errors.WithStack(err)
		return err
	}
	return
}

func analysisRouter(v runtime.Router) (Project string, Bus string, Func string) {
	if strings.Index(v.Handler, "/app/") != -1 {
		temp := strings.Split(v.Handler, "/app/")
		if len(temp) <= 1 {
			return
		}
		projectX := strings.Split(temp[1], "/")
		if len(projectX) <= 1 {
			return
		}
		Project = projectX[0]
		busX := strings.Split(projectX[1], ".")
		if len(busX) <= 1 {
			return
		}
		Bus = busX[1]
		Func = busX[len(busX)-1]
		if strings.Contains(Func, "-fm") {
			Func = strings.Replace(Func, "-fm", "", -1)
		}
	}
	return Project, Bus, Func
}
