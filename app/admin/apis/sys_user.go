package apis

import (
	"go-admin/common/exception"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"

	"go-admin/common/actions"
	"go-admin/common/apis"
	"go-admin/common/jwtauth/user"
	_ "go-admin/common/response/antd"
)

type SysUser struct {
	apis.Api
}

// GetPage
// @Summary 列表用户信息数据
// @Description 获取JSON
// @Tags 用户
// @Param username query string false "username"
// @Success 200 {string} {object} antd.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/sys-user [get]
// @Security Bearer
func (e SysUser) GetPage(c *gin.Context) {
	s := service.SysUser{}
	req := dto.SysUserGetPageReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		panic(exception.New(exception.GetPageUserFail, err))
		return
	}

	//数据权限检查
	p := actions.GetPermission(c)

	list := make([]dto.SysUserGetPageResp, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		panic(exception.New(exception.GetPageUserFail, err))
		return
	}
	e.PageOK(list, count, req.GetPageIndex(), req.GetPageSize())
}

// Get
// @Summary 获取用户
// @Description 获取JSON
// @Tags 用户
// @Param userId path int true "用户编码"
// @Success 200 {object} antd.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/sys-user/{userId} [get]
// @Security Bearer
func (e SysUser) Get(c *gin.Context) {
	s := service.SysUser{}
	req := dto.SysUserById{}
	resp := dto.SysUserGetResp{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, nil).
		MakeService(&s.Service).
		Errors
	if err != nil {
		panic(exception.New(exception.GetUserFail, err))
		return
	}
	var object models.SysUser
	//数据权限检查
	p := actions.GetPermission(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		panic(exception.New(exception.GetUserFail, err))
		return
	}
	resp.Generate(&object)
	e.OK(resp)
}

// Insert
// @Summary 创建用户
// @Description 获取JSON
// @Tags 用户
// @Accept  application/json
// @Product application/json
// @Param data body dto.SysUserInsertReq true "用户数据"
// @Success 200 {object} antd.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/sys-user [post]
// @Security Bearer
func (e SysUser) Insert(c *gin.Context) {
	s := service.SysUser{}
	req := dto.SysUserInsertReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		panic(exception.New(exception.InsertUserFail, err))
		return
	}
	// 设置创建人
	req.SetCreateBy(user.GetUserId(c))
	err = s.Insert(c, &req)
	if err != nil {
		panic(exception.New(exception.InsertUserFail, err))
		return
	}
	e.OK(req.GetId())
}

// Update
// @Summary 修改用户数据
// @Description 获取JSON
// @Tags 用户
// @Accept  application/json
// @Product application/json
// @Param data body dto.SysUserUpdateReq true "body"
// @Success 200 {object} antd.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/sys-user/{userId} [put]
// @Security Bearer
func (e SysUser) Update(c *gin.Context) {
	s := service.SysUser{}
	req := dto.SysUserUpdateReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		panic(exception.New(exception.UpdateUserFail, err))
		return
	}

	req.SetUpdateBy(user.GetUserId(c))

	//数据权限检查
	p := actions.GetPermission(c)

	err = s.Update(c, &req, p)
	if err != nil {
		e.Logger.Error(err)
		panic(exception.New(exception.UpdateUserFail, err))
		return
	}
	e.OK(req.GetId())
}

// Delete
// @Summary 删除用户数据
// @Description 删除数据
// @Tags 用户
// @Param userId path int true "userId"
// @Success 200 {object} antd.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/sys-user/{userId} [delete]
// @Security Bearer
func (e SysUser) Delete(c *gin.Context) {
	s := service.SysUser{}
	req := dto.SysUserById{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON).
		MakeService(&s.Service).
		Errors
	if err != nil {
		panic(exception.New(exception.DeleteUserFail, err))
		return
	}

	// 设置编辑人
	req.SetUpdateBy(user.GetUserId(c))

	// 数据权限检查
	p := actions.GetPermission(c)

	err = s.Remove(c, &req, p)
	if err != nil {
		panic(exception.New(exception.DeleteUserFail, err))
		return
	}
	e.OK(req.GetId())
}

// InsetAvatar
// @Summary 修改头像
// @Description 获取JSON
// @Tags 个人中心
// @Accept multipart/form-data
// @Param file formData file true "file"
// @Success 200 {object} antd.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/user/avatar [post]
// @Security Bearer
func (e SysUser) InsetAvatar(c *gin.Context) {
	s := service.SysUser{}
	req := dto.UpdateSysUserAvatarReq{}
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors
	if err != nil {
		panic(exception.New(exception.InsetUserAvatarFail, err))
		return
	}
	// 数据权限检查
	p := actions.GetPermission(c)
	form, _ := c.MultipartForm()
	files := form.File["upload[]"]
	guid := uuid.New().String()
	filPath := "static/uploadfile/" + guid + ".jpg"
	for _, file := range files {
		e.Logger.Debugf("upload avatar file: %s", file.Filename)
		// 上传文件至指定目录
		err = e.SaveUploadedFile(c, file, filPath)
		if err != nil {
			panic(exception.New(exception.InsetUserAvatarFail, err))
			return
		}
	}
	req.UserId = p.UserId
	req.Avatar = "/" + filPath

	err = s.UpdateAvatar(c, &req, p)
	if err != nil {
		panic(exception.New(exception.InsetUserAvatarFail, err))
		return
	}
	e.OK(filPath)
}

// UpdateStatus 修改用户状态
// @Summary 修改用户状态
// @Description 获取JSON
// @Tags 用户
// @Accept  application/json
// @Product application/json
// @Param data body dto.UpdateSysUserStatusReq true "body"
// @Success 200 {object} antd.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/user/status [put]
// @Security Bearer
func (e SysUser) UpdateStatus(c *gin.Context) {
	s := service.SysUser{}
	req := dto.UpdateSysUserStatusReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON, nil).
		MakeService(&s.Service).
		Errors
	if err != nil {
		panic(exception.New(exception.UpdateUserStatusFail, err))
		return
	}

	req.SetUpdateBy(user.GetUserId(c))

	//数据权限检查
	p := actions.GetPermission(c)

	err = s.UpdateStatus(c, &req, p)
	if err != nil {
		panic(exception.New(exception.UpdateUserStatusFail, err))
		return
	}
	e.OK(req.GetId())
}

// ResetPwd 重置用户密码
// @Summary 重置用户密码
// @Description 获取JSON
// @Tags 用户
// @Accept  application/json
// @Product application/json
// @Param data body dto.ResetSysUserPwdReq true "body"
// @Success 200 {object} antd.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/user/pwd/reset [put]
// @Security Bearer
func (e SysUser) ResetPwd(c *gin.Context) {
	s := service.SysUser{}
	req := dto.ResetSysUserPwdReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON).
		MakeService(&s.Service).
		Errors
	if err != nil {
		panic(exception.New(exception.UpdateUserResetPwdFail, err))
		return
	}

	req.SetUpdateBy(user.GetUserId(c))

	//数据权限检查
	p := actions.GetPermission(c)
	if user.GetUserId(c) == 1 {
		req.Password = "123456"
	}
	err = s.ResetPwd(c, &req, p)
	if err != nil {
		panic(exception.New(exception.UpdateUserResetPwdFail, err))
		return
	}
	e.OK(req.GetId())
}

// UpdatePwd
// @Summary 修改密码
// @Description 获取JSON
// @Tags 用户
// @Accept  application/json
// @Product application/json
// @Param data body dto.PassWord true "body"
// @Success 200 {object} antd.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/user/pwd/set [put]
// @Security Bearer
func (e SysUser) UpdatePwd(c *gin.Context) {
	s := service.SysUser{}
	req := dto.UpdateSysUserPwdReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		panic(exception.New(exception.UpdateUserPwdFail, err))
		return
	}
	req.UserId = user.GetUserId(c)
	// 数据权限检查
	p := actions.GetPermission(c)
	if req.UserId == 1 {
		req.Password = "123456"
	}
	err = s.UpdatePwd(c, &req, p)
	if err != nil {
		panic(exception.New(exception.UpdateUserPwdFail, err))
		return
	}
	e.OK(nil)
}

// GetProfile
// @Summary 获取个人中心数据
// @Description 获取JSON
// @Tags 个人中心
// @Success 200 {object} antd.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/user/profile [get]
// @Security Bearer
func (e SysUser) GetProfile(c *gin.Context) {
	s := service.SysUser{}
	req := dto.SysUserById{}
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors
	if err != nil {
		panic(exception.New(exception.GetUserProfileFail, err))
		return
	}

	req.Id = user.GetUserId(c)

	sysUser := models.SysUser{}
	roles := make([]models.SysRole, 0)
	posts := make([]models.SysPost, 0)
	err = s.GetProfile(&req, &sysUser, &roles, &posts)
	if err != nil {
		panic(exception.New(exception.GetUserProfileFail, err))
		return
	}
	e.OK(gin.H{
		"user":  sysUser,
		"roles": roles,
		"posts": posts,
	})
}

// GetInfo
// @Summary 获取个人信息
// @Description 获取JSON
// @Tags 个人中心
// @Success 200 {object} antd.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/getinfo [get]
// @Security Bearer
func (e SysUser) GetInfo(c *gin.Context) {
	req := dto.SysUserById{}
	s := service.SysUser{}
	r := service.SysRole{}
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&r.Service).
		MakeService(&s.Service).
		Errors
	if err != nil {
		panic(exception.WithStatus(http.StatusUnauthorized, exception.GetUserInfoFail, err))
		return
	}
	p := actions.GetPermission(c)
	var roles = make([]string, 1)
	roles[0] = user.GetRoleKey(c)
	var permissions = make([]string, 1)
	permissions[0] = "*:*:*"
	var buttons = make([]string, 1)
	buttons[0] = "*:*:*"

	var mp = make(map[string]interface{})
	mp["roles"] = roles
	if user.GetRoleKey(c) == "admin" || user.GetRoleName(c) == "系统管理员" {
		mp["permissions"] = permissions
		mp["buttons"] = buttons
	} else {
		list, _ := r.GetById(user.GetRoleId(c))
		mp["permissions"] = list
		mp["buttons"] = list
	}
	sysUser := models.SysUser{}
	req.Id = user.GetUserId(c)
	err = s.Get(&req, p, &sysUser)
	if err != nil {
		panic(exception.WithStatus(http.StatusUnauthorized, exception.LoginFail, err))
		return
	}
	mp["introduction"] = " am a super administrator"
	mp["avatar"] = "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif"
	if sysUser.Avatar != "" {
		mp["avatar"] = sysUser.Avatar
	}
	mp["userName"] = sysUser.NickName
	mp["userId"] = sysUser.UserId
	mp["deptId"] = sysUser.DeptId
	mp["name"] = sysUser.NickName
	mp["code"] = 200
	e.OK(mp)
}
