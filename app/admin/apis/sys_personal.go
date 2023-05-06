package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
	"go-admin/common/exception"
	"strconv"

	"go-admin/common/apis"
	"go-admin/common/jwtauth/user"
	_ "go-admin/common/response/antd"
)

type Personal struct {
	apis.Api
}

// CurrentUser 获取当前用户信息
// @Summary 获取当前用户信息
// @Description 获取JSON
// @Tags 个人中心/UserCenter
// @Accept  application/json
// @Product application/json
// @Success 200 {object} antd.Response{data=dto.Personal} "{"code": 200, "data": [...]}"
// @Router /api/v1/user/current [get]
// @Security Bearer
func (e Personal) CurrentUser(c *gin.Context) {
	s := service.SysPersonal{}
	err := e.MakeContext(c).MakeOrm().MakeService(&s.Service).Errors
	if err != nil {
		err = errors.WithStack(err)
		e.Logger.Error(err)
		panic(err)
		return
	}
	access, err := s.GetAccess(user.GetRoleId(c))
	if err != nil {
		err = errors.WithStack(err)
		e.Logger.Error(err)
		panic(err)
		return
	}

	userM, err := s.GetInfo(user.GetUserId(c))
	if err != nil {
		err = errors.WithStack(err)
		e.Logger.Error(err)
		panic(err)
		return
	}

	var userInfo = dto.Personal{
		Name:      userM.NickName,
		Avatar:    userM.Avatar,
		Userid:    strconv.Itoa(userM.UserId),
		Email:     userM.Email,
		Signature: "海纳百川，有容乃大",
		Title:     "交互专家",
		Group:     "蚂蚁金服－某某某事业群－某某平台部－某某技术部－UED",
		Tags: []dto.Tag{
			{
				Key:   "0",
				Label: "很有想法的",
			},
			{
				Key:   "1",
				Label: "专注设计",
			},
			{
				Key:   "2",
				Label: "辣~",
			},
			{
				Key:   "3",
				Label: "大长腿",
			},
			{
				Key:   "4",
				Label: "川妹子",
			},
			{
				Key:   "5",
				Label: "海纳百川",
			},
		},
		NotifyCount: 12,
		UnreadCount: 11,
		Country:     "China",
		Access:      "admin",
		AccessList:  access,
		Geographic: dto.Geographic{
			Province: dto.Province{
				Label: "浙江省",
				Key:   "330000",
			},
			City: dto.City{
				Label: "杭州市",
				Key:   "330100",
			},
		},
		Address: "西湖区工专路 77 号",
		Phone:   "0752-268888888",
		Mobile:  userM.Phone,
	}
	e.OK(userInfo)
}

func (e Personal) UpdateCurrent(c *gin.Context) {
	s := service.SysPersonal{}
	req := dto.UpdatePersonalReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(exception.WithMsg(50000, "UpdateCurrentFail", err))
		return
	}

	req.UserId = int(user.GetUserId(c))
	req.SetUpdateBy(user.GetUserId(c))

	//数据权限检查
	p := actions.GetPermissionFromContext(c)

	err = s.Update(c, &req, p)
	if err != nil {
		e.Logger.Error(err)
		return
	}
	e.OK(req.GetId())
}

func (e Personal) UserOutLogin(c *gin.Context) {
	var mp = make(map[string]interface{})
	mp["data"] = "ok"
	mp["success"] = true
	c.JSON(200, mp)
}
