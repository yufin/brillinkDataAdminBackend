package service

import (
	//"github.com/go-admin-team/go-admin-core/logger"

	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	"go-admin/common/middleware"
	"go-admin/common/service"
	"reflect"
	"runtime"
)

type SysPost struct {
	service.Service
}

// GetPage 获取SysPost列表
func (e *SysPost) GetPage(r *dto.SysPostGetPageReq) (list *[]models.SysPost, count *int64, err error) {
	list, count, err = service.GetPage[models.SysPost](e, r)
	if err != nil {
		err = errors.WithStack(err)
		e.GetLog().Errorf("database operation failed:%s \r", err)
		return
	}
	return
}

// Get 获取SysPost对象
func (e *SysPost) Get(r *dto.SysPostGetReq) (model *models.SysPost, err error) {
	model, err = service.Get[models.SysPost](e, r)
	if err != nil {
		err = errors.WithStack(err)
		e.GetLog().Errorf("database operation failed:%s \r", err)
		return
	}
	return
}

// Insert 创建SysPost对象
func (e *SysPost) Insert(c *gin.Context, r *dto.SysPostInsertReq) (err error) {
	model := new(models.SysPost)
	model, err = service.Insert[models.SysPost](e, r)
	if err != nil {
		e.GetLog().Errorf("database operation failed:%s \r", err)
		return
	}
	after, _ := json.Marshal(model)
	middleware.SetContextOperateLog(c,
		"新增",
		runtime.FuncForPC(reflect.ValueOf(e.Remove).Pointer()).Name()+
			fmt.Sprintf("数据，ID：%v", model.GetId()),
		"{}",
		string(after),
		"职位",
	)
	return
}

// Update 修改SysPost对象
func (e *SysPost) Update(c *gin.Context, r *dto.SysPostUpdateReq) (err error) {
	var before, after []byte
	before, after, ok, err := service.GetAndUpdate[models.SysPost](e, r)
	if err != nil {
		e.GetLog().Errorf("database operation failed:%s \r", err)
		return
	}
	if ok {
		middleware.SetContextOperateLog(c,
			"修改",
			runtime.FuncForPC(reflect.ValueOf(e.Remove).Pointer()).Name()+
				fmt.Sprintf("数据，ID：%v", r.GetId()),
			string(before),
			string(after),
			"职位",
		)
	}
	return
}

// Remove 删除SysPost
func (e *SysPost) Remove(c *gin.Context, r *dto.SysPostDeleteReq) (err error) {
	ok, err := service.Delete[models.SysPost](e, r)
	if err != nil {
		e.GetLog().Errorf("database operation failed:%s \r", err)
		return
	}
	if ok {
		middleware.SetContextOperateLog(c,
			"删除",
			runtime.FuncForPC(reflect.ValueOf(e.Remove).Pointer()).Name()+
				fmt.Sprintf("数据，ID：%v", r.GetId()),
			"{}",
			"{}",
			"职位",
		)
	}
	return
}

func (e *SysPost) GetList(r *dto.SysPostGetPageReq) (list *[]models.SysPost, err error) {
	list, err = service.GetList[models.SysPost](e, r)
	if err != nil {
		e.GetLog().Errorf("database operation failed:%s \r", err)
		return
	}
	return
}
