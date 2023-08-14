package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"

	"go-admin/common/apis"
	"go-admin/common/jwtauth/user"
	_ "go-admin/common/response/antd"
)

type SysDictData struct {
	apis.Api
}

// GetPage
// @Summary      字典数据列表
// @Description  获取JSON
// @Tags         字典数据
// @Param        status     query     string         false  "status"
// @Param        dictCode   query     string         false  "dictCode"
// @Param        dictType   query     string         false  "dictType"
// @Param        pageSize   query     int            false  "页条数"
// @Param        pageIndex  query     int            false  "页码"
// @Success      200        {object}  antd.Response  "{"code": 200, "data": [...]}"
// @Router       /api/v1/dict/data [get]
// @Security     Bearer
func (e SysDictData) GetPage(c *gin.Context) {
	s := service.SysDictData{}
	req := dto.SysDictDataGetPageReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.Form).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(err)
		return
	}

	list := make([]models.SysDictData, 0)
	var count int64
	err = s.GetPage(&req, &list, &count)
	if err != nil {
		e.Error(500, "查询失败", "")
		return
	}

	e.PageOK(list, count, req.GetPageIndex(), req.GetPageSize())
}

// Get
// @Summary      通过编码获取字典数据
// @Description  获取JSON
// @Tags         字典数据
// @Param        dictCode  path      int            true  "字典编码"
// @Success      200       {object}  antd.Response  "{"code": 200, "data": [...]}"
// @Router       /api/v1/dict/data/{dictCode} [get]
// @Security     Bearer
func (e SysDictData) Get(c *gin.Context) {
	s := service.SysDictData{}
	req := dto.SysDictDataGetReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, nil).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(err)
		return
	}

	var object models.SysDictData

	err = s.Get(&req, &object)
	if err != nil {
		e.Logger.Warnf("Get error: %s", err.Error())
		e.Error(500, "查询失败", "")
		return
	}

	e.OK(object)
}

// Insert
// @Summary      添加字典数据
// @Description  获取JSON
// @Tags         字典数据
// @Accept       application/json
// @Product      application/json
// @Param        data  body      dto.SysDictDataInsertReq  true  "data"
// @Success      200   {object}  antd.Response             "{"code": 200, "message": "添加成功"}"
// @Router       /api/v1/dict/data [post]
// @Security     Bearer
func (e SysDictData) Insert(c *gin.Context) {
	s := service.SysDictData{}
	req := dto.SysDictDataInsertReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(err)
		return
	}
	req.SetCreateBy(user.GetUserId(c))
	err = s.Insert(&req)
	if err != nil {
		e.Error(500, "创建失败", "")
		return
	}

	e.OK(req.GetId())
}

// Update
// @Summary      修改字典数据
// @Description  获取JSON
// @Tags         字典数据
// @Accept       application/json
// @Product      application/json
// @Param        data  body      dto.SysDictDataUpdateReq  true  "body"
// @Success      200   {object}  antd.Response             "{"code": 200, "message": "修改成功"}"
// @Router       /api/v1/dict/data/{dictCode} [put]
// @Security     Bearer
func (e SysDictData) Update(c *gin.Context) {
	s := service.SysDictData{}
	req := dto.SysDictDataUpdateReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON, nil).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(err)
		return
	}
	req.SetUpdateBy(user.GetUserId(c))
	err = s.Update(&req)
	if err != nil {
		e.Error(500, "更新失败", "")
		return
	}
	e.OK(req.GetId())
}

// Delete
// @Summary      删除字典数据
// @Description  删除数据
// @Tags         字典数据
// @Param        dictCode  body      dto.SysDictDataDeleteReq  true  "body"
// @Success      200       {object}  antd.Response             "{"code": 200, "message": "删除成功"}"
// @Router       /api/v1/dict/data [delete]
// @Security     Bearer
func (e SysDictData) Delete(c *gin.Context) {
	s := service.SysDictData{}
	req := dto.SysDictDataDeleteReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON, nil).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(err)
		return
	}
	req.SetUpdateBy(user.GetUserId(c))
	err = s.Remove(&req)
	if err != nil {
		e.Error(500, "删除失败", "")
		return
	}
	e.OK(req.GetId())
}

// GetAll 数据字典根据key获取 业务页面使用
// @Summary      数据字典根据key获取
// @Description  数据字典根据key获取
// @Tags         字典数据
// @Param        dictType  query     int                                              true  "dictType"
// @Success      200       {object}  antd.Response{data=[]dto.SysDictDataGetAllResp}  "{"code": 200, "data": [...]}"
// @Router       /api/v1/dict-data/option-select [get]
// @Security     Bearer
func (e SysDictData) GetAll(c *gin.Context) {
	s := service.SysDictData{}
	req := dto.SysDictDataGetPageReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(err)
		return
	}
	list := make([]models.SysDictData, 0)
	err = s.GetAll(&req, &list)
	if err != nil {
		e.Error(500, "查询失败", "")
		return
	}
	l := make([]dto.SysDictDataGetAllResp, 0)
	for _, i := range list {
		d := dto.SysDictDataGetAllResp{}
		e.Translate(i, &d)
		l = append(l, d)
	}

	e.OK(l)
}
