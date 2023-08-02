package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-admin-team/go-admin-core/sdk/pkg"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"

	"go-admin/common/apis"
	"go-admin/common/jwtauth/user"
)

type SysDept struct {
	apis.Api
}

// GetPage
// @Summary      分页部门列表数据
// @Description  分页列表
// @Tags         部门
// @Param        deptName  query     string         false  "deptName"
// @Param        deptId    query     string         false  "deptId"
// @Param        position  query     string         false  "position"
// @Success      200       {object}  antd.Response  "{"code": 200, "data": [...]}"
// @Router       /api/v1/dept [get]
// @Security     Bearer
func (e SysDept) GetPage(c *gin.Context) {
	s := service.SysDept{}
	req := dto.SysDeptGetPageReq{}
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
	list := make([]models.SysDept, 0)
	list, err = s.SetDeptPage(&req)
	if err != nil {
		panic(err)
		return
	}
	e.OK(list)
}

// Get
// @Summary      获取部门数据
// @Description  获取JSON
// @Tags         部门
// @Param        deptId  path      string         false  "deptId"
// @Success      200     {object}  antd.Response  "{"code": 200, "data": [...]}"
// @Router       /api/v1/dept/{deptId} [get]
// @Security     Bearer
func (e SysDept) Get(c *gin.Context) {
	s := service.SysDept{}
	req := dto.SysDeptGetReq{}
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
	model, err := s.Get(&req)
	if err != nil {
		panic(err)
		return
	}

	e.OK(model)
}

// Insert 添加部门
// @Summary      添加部门
// @Description  获取JSON
// @Tags         部门
// @Accept       application/json
// @Product      application/json
// @Param        data  body      dto.SysDeptInsertReq  true  "data"
// @Success      200   {string}  string                "{"code": 200, "message": "添加成功"}"
// @Success      200   {string}  string                "{"code": -1, "message": "添加失败"}"
// @Router       /api/v1/dept [post]
// @Security     Bearer
func (e SysDept) Insert(c *gin.Context) {
	s := service.SysDept{}
	req := dto.SysDeptInsertReq{}
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

	// 设置创建人
	req.SetCreateBy(user.GetUserId(c))
	err = s.Insert(c, &req)
	if err != nil {
		panic(err)
		return
	}
	e.OK(req.GetId())
}

// Update
// @Summary      修改部门
// @Description  获取JSON
// @Tags         部门
// @Accept       application/json
// @Product      application/json
// @Param        id    path      int                   true  "id"
// @Param        data  body      dto.SysDeptUpdateReq  true  "body"
// @Success      200   {string}  string                "{"code": 200, "message": "添加成功"}"
// @Success      200   {string}  string                "{"code": -1, "message": "添加失败"}"
// @Router       /api/v1/dept/{deptId} [put]
// @Security     Bearer
func (e SysDept) Update(c *gin.Context) {
	s := service.SysDept{}
	req := dto.SysDeptUpdateReq{}
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
	req.SetUpdateBy(user.GetUserId(c))
	err = s.Update(c, &req)
	if err != nil {
		panic(err)
		return
	}
	e.OK(req.GetId())
}

// Delete
// @Summary      删除部门
// @Description  删除数据
// @Tags         部门
// @Param        data  body      dto.SysDeptDeleteReq  true  "body"
// @Success      200   {string}  string                "{"code": 200, "message": "删除成功"}"
// @Success      200   {string}  string                "{"code": -1, "message": "删除失败"}"
// @Router       /api/v1/dept [delete]
// @Security     Bearer
func (e SysDept) Delete(c *gin.Context) {
	s := service.SysDept{}
	req := dto.SysDeptDeleteReq{}
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

	err = s.Remove(c, &req)
	if err != nil {
		panic(err)
		return
	}
	e.OK(req.GetId())
}

// Get2Tree 用户管理 左侧部门树
func (e SysDept) Get2Tree(c *gin.Context) {
	s := service.SysDept{}
	req := dto.SysDeptGetPageReq{}
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
	list := make([]dto.DeptLabel, 0)
	list, err = s.SetDeptTree(&req)
	if err != nil {
		panic(err)
		return
	}
	e.OK(list)
}

// GetToTree antd 用户管理部门树
func (e SysDept) GetToTree(c *gin.Context) {
	s := service.SysDept{}
	req := dto.SysDeptGetPageReq{}
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
	list := make([]dto.DeptTree, 0)
	req.Status = "2"
	list, err = s.SetTree(&req)
	if err != nil {
		panic(err)
		return
	}
	e.OK(list)
}

// GetDeptTreeRoleSelect TODO: 此接口需要调整不应该将list和选中放在一起
func (e SysDept) GetDeptTreeRoleSelect(c *gin.Context) {
	s := service.SysDept{}
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(err)
		return
	}

	id, err := pkg.StringToInt(c.Param("roleId"))
	result, err := s.SetDeptLabel()
	if err != nil {
		panic(err)
		return
	}
	menuIds := make([]int, 0)
	if id != 0 {
		menuIds, err = s.GetWithRoleId(id)
		if err != nil {
			panic(err)
			return
		}
	}
	e.OK(gin.H{
		"depts":       result,
		"checkedKeys": menuIds,
	})
}
