package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-admin-team/go-admin-core/sdk"
	"go-admin/common/exception"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"

	"go-admin/common/apis"
	"go-admin/common/global"
	"go-admin/common/jwtauth/user"
	_ "go-admin/common/response/antd"
)

type SysRole struct {
	apis.Api
}

// GetPage
// @Summary      角色列表数据
// @Description  Get JSON
// @Tags         角色/Role
// @Param        roleName   query     string         false  "roleName"
// @Param        status     query     string         false  "status"
// @Param        roleKey    query     string         false  "roleKey"
// @Param        pageSize   query     int            false  "页条数"
// @Param        pageIndex  query     int            false  "页码"
// @Success      200        {object}  antd.Response  "{"code": 200, "data": [...]}"
// @Router       /api/v1/role [get]
// @Security     Bearer
func (e SysRole) GetPage(c *gin.Context) {
	s := service.SysRole{}
	req := dto.SysRoleGetPageReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.Form).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.New(exception.GetPageRoleFail, err))
		return
	}

	//list := make([]models.SysRole, 0)
	//var count int64

	list, count, err := s.GetPage(&req)
	if err != nil {
		panic(exception.New(exception.GetPageRoleFail, err))
		return
	}

	e.PageOK(list, *count, req.GetPageIndex(), req.GetPageSize())
}

// GetOption
// @Summary      角色下拉框数据
// @Description  Get JSON
// @Tags         角色/Role
// @Success      200  {object}  antd.Response  "{"code": 200, "data": [...]}"
// @Router       /api/v1/role-option [get]
// @Security     Bearer
func (e SysRole) GetOption(c *gin.Context) {
	s := service.SysRole{}
	req := dto.SysRoleGetPageReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.Form).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.New(exception.GetRoleOptionFail, err))
		return
	}
	list, err := s.GetList(&req)
	if err != nil {
		panic(exception.New(exception.GetRoleOptionFail, err))
		return
	}
	l := make([]dto.SysRoleGetListResp, 0)
	for _, i := range *list {
		if i.RoleName == "系统管理员" {
			continue
		}
		d := dto.SysRoleGetListResp{}
		e.Translate(i, &d)
		l = append(l, d)

	}
	e.OK(l)
}

// Get
// @Summary      获取Role数据
// @Description  获取JSON
// @Tags         角色/Role
// @Param        roleId  path      string         false  "roleId"
// @Success      200     {object}  antd.Response  "{"code": 200, "data": [...]}"
// @Router       /api/v1/role/{id} [get]
// @Security     Bearer
func (e SysRole) Get(c *gin.Context) {
	s := service.SysRole{}
	req := dto.SysRoleGetReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, nil).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.New(exception.GetRoleFail, err))
		return
	}

	object, err := s.Get(&req)
	if err != nil {
		panic(exception.New(exception.GetRoleFail, err))
		return
	}

	e.OK(object)
}

// Insert
// @Summary      创建角色
// @Description  获取JSON
// @Tags         角色/Role
// @Accept       application/json
// @Product      application/json
// @Param        data  body      dto.SysRoleInsertReq  true  "data"
// @Success      200   {object}  antd.Response         "{"code": 200, "data": [...]}"
// @Router       /api/v1/role [post]
// @Security     Bearer
func (e SysRole) Insert(c *gin.Context) {
	s := service.SysRole{}
	req := dto.SysRoleInsertReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.New(exception.InsertRoleFail, err))
		return
	}

	// 设置创建人
	req.CreateBy = user.GetUserId(c)
	if req.Status == "" {
		req.Status = "2"
	}
	cb := sdk.Runtime.GetCasbinKey(c.Request.Host)
	err = s.Insert(c, &req, cb)
	if err != nil {
		e.Logger.Error(err)
		panic(exception.New(exception.InsertRoleFail, err))
		return
	}
	_, err = global.LoadPolicy(c)
	if err != nil {
		panic(exception.New(exception.InsertRoleFail, err))
		return
	}
	e.OK(req.GetId())
}

// Update 修改用户角色
// @Summary      修改用户角色
// @Description  获取JSON
// @Tags         角色/Role
// @Accept       application/json
// @Product      application/json
// @Param        data  body      dto.SysRoleUpdateReq  true  "body"
// @Success      200   {object}  antd.Response         "{"code": 200, "data": [...]}"
// @Router       /api/v1/role/{id} [put]
// @Security     Bearer
func (e SysRole) Update(c *gin.Context) {
	s := service.SysRole{}
	req := dto.SysRoleUpdateReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.New(exception.UpdateRoleFail, err))
		return
	}
	cb := sdk.Runtime.GetCasbinKey(c.Request.Host)

	req.SetUpdateBy(user.GetUserId(c))

	err = s.Update(c, &req, cb)
	if err != nil {
		panic(exception.New(exception.UpdateRoleFail, err))
		return
	}
	_, err = global.LoadPolicy(c)
	if err != nil {
		panic(exception.New(exception.UpdateRoleFail, err))
		return
	}
	e.OK(req.GetId())
}

// Delete
// @Summary      删除用户角色
// @Description  删除数据
// @Tags         角色/Role
// @Param        data  body      dto.SysRoleDeleteReq  true  "body"
// @Success      200   {object}  antd.Response         "{"code": 200, "data": [...]}"
// @Router       /api/v1/role [delete]
// @Security     Bearer
func (e SysRole) Delete(c *gin.Context) {
	s := new(service.SysRole)
	req := dto.SysRoleDeleteReq{}
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

	err = s.Remove(c, &req)
	if err != nil {
		e.Logger.Error(err)
		panic(exception.New(exception.DeleteRoleFail, err))
		return
	}
	_, err = global.LoadPolicy(c)
	if err != nil {
		panic(exception.New(exception.DeleteRoleFail, err))
		return
	}
	e.OK(req.GetId())
}

// Update2Status 修改用户角色状态
// @Summary      修改用户角色
// @Description  获取JSON
// @Tags         角色/Role
// @Accept       application/json
// @Product      application/json
// @Param        data  body      dto.UpdateStatusReq  true  "body"
// @Success      200   {object}  antd.Response        "{"code": 200, "data": [...]}"
// @Router       /api/v1/role-status [put]
// @Security     Bearer
func (e SysRole) Update2Status(c *gin.Context) {
	s := service.SysRole{}
	req := dto.UpdateStatusReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.New(exception.UpdateRoleStatusFail, err))
		return
	}
	req.SetUpdateBy(user.GetUserId(c))
	err = s.UpdateStatus(c, &req)
	if err != nil {
		panic(exception.New(exception.UpdateRoleStatusFail, err))
		return
	}
	e.OK(req.GetId())
}

// Update2DataScope 更新角色数据权限
// @Summary      更新角色数据权限
// @Description  获取JSON
// @Tags         角色/Role
// @Accept       application/json
// @Product      application/json
// @Param        data  body      dto.RoleDataScopeReq  true  "body"
// @Success      200   {object}  antd.Response         "{"code": 200, "data": [...]}"
// @Router       /api/v1/role-status/{id} [put]
// @Security     Bearer
func (e SysRole) Update2DataScope(c *gin.Context) {
	s := service.SysRole{}
	req := dto.RoleDataScopeReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON, nil).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.New(exception.UpdateRoleDataScopeFail, err))
		return
	}
	data := &models.SysRole{
		RoleId:    req.RoleId,
		DataScope: req.DataScope,
		DeptIds:   req.DeptIds,
	}
	data.UpdateBy = user.GetUserId(c)
	err = s.UpdateDataScope(c, &req).Error
	if err != nil {
		panic(exception.New(exception.UpdateRoleDataScopeFail, err))
		return
	}
	e.OK(nil)
}
