package apis

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-admin-team/go-admin-core/sdk"

	"go-admin/app/worklist/models"
	"go-admin/app/worklist/service"
	"go-admin/app/worklist/service/dto"

	"go-admin/common/apis"
	"go-admin/common/global"
	"go-admin/common/jwtauth/user"
)

type TbTodo struct {
	apis.Api
}

func (e TbTodo) GetPage(c *gin.Context) {
	s := service.TbTodo{}
	req := dto.TbTodoGetPageReq{}
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
	//数据权限检查
	//p := actions.GetPermission(c)
	list := make([]models.TbTodo, 0)
	var count int64
	err = s.GetPage(&req, &list, &count)
	if err != nil {
		panic(err)
		return
	}
	e.ListWithPage(list, count, req.GetPageIndex(), req.GetPageSize())
}

func (e TbTodo) GetTotal(c *gin.Context) {
	s := service.TbTodo{}
	req := dto.TbTodoGetPageReq{}
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
	//数据权限检查
	//p := actions.GetPermission(c)
	var normal int64
	var active int64
	var success int64
	var exception int64
	err = s.GetTotal(&req, &normal, &active, &exception, &success)
	if err != nil {
		panic(err)
		return
	}
	mp := make(map[string]interface{}, 3)
	mp["normal"] = normal
	mp["active"] = active
	mp["success"] = success
	mp["exception"] = exception
	e.OK(mp)
}

func (e TbTodo) Get(c *gin.Context) {
	req := dto.TbTodoGetReq{}
	s := service.TbTodo{}
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
	var object models.TbTodo

	err = s.Get(&req, &object).Error
	if err != nil {
		panic(err)
		return
	}
	list := make([]models.TbTodo, 0)
	err = s.GetList(&list)
	if err != nil {
		panic(err)
		return
	}
	mp := make(map[string]interface{}, 0)
	mp["list"] = list
	e.OK(mp)
}

func (e TbTodo) Insert(c *gin.Context) {
	s := service.TbTodo{}
	req := dto.TbTodoInsertReq{}
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
		panic(err)
		return
	}
	q := sdk.Runtime.GetMemoryQueue("")
	mp := make(map[string]interface{}, 0)
	mp["avatar"] = "https://gw.alipayobjects.com/zos/rmsportal/ThXAXghbEsBCCSDihZxY.png"
	mp["title"] = "用户【" + user.GetUserName(c) + "】创建了" + req.Title + "."
	mp["type"] = "notification"
	mp["datetime"] = time.Now()
	message, err := sdk.Runtime.GetStreamMessage("", string(global.Notice), mp)
	if err != nil {
		log.Printf("GetStreamMessage error, %s \n", err.Error())
		//日志报错错误，不中断请求
	} else {
		err = q.Append(message)
		if err != nil {
			log.Printf("Append message error, %s \n", err.Error())
		}
	}
	reqPage := dto.TbTodoGetPageReq{}
	reqPage.PageIndex = 1
	reqPage.PageSize = 10
	list := make([]models.TbTodo, 0)
	var count int64
	err = s.GetPage(&reqPage, &list, &count)
	if err != nil {
		panic(err)
		return
	}
	e.ListWithPage(list, count, 1, 10)
}

func (e TbTodo) Update(c *gin.Context) {
	s := service.TbTodo{}
	req := dto.TbTodoUpdateReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON, nil).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err.Error(), "")
		return
	}
	req.SetUpdateBy(user.GetUserId(c))
	err = s.Update(&req)
	if err != nil {
		e.Error(500, err.Error(), "")
		return
	}
	q := sdk.Runtime.GetMemoryQueue("")
	mp := make(map[string]interface{}, 0)
	mp["avatar"] = "https://gw.alipayobjects.com/zos/rmsportal/ThXAXghbEsBCCSDihZxY.png"
	mp["title"] = "`" + user.GetUserName(c) + "`修改了" + req.Title + "."
	mp["type"] = "notification"
	mp["datetime"] = time.Now()
	message, err := sdk.Runtime.GetStreamMessage("", string(global.Notice), mp)
	if err != nil {
		log.Printf("GetStreamMessage error, %s \n", err.Error())
		//日志报错错误，不中断请求
	} else {
		err = q.Append(message)
		if err != nil {
			log.Printf("Append message error, %s \n", err.Error())
		}
	}
	reqPage := dto.TbTodoGetPageReq{}
	list := make([]models.TbTodo, 0)
	var count int64
	err = s.GetPage(&reqPage, &list, &count)
	if err != nil {
		e.Error(500, err.Error(), "")
		return
	}
	e.ListWithPage(list, count, 1, 10)
}

func (e TbTodo) Delete(c *gin.Context) {
	s := service.TbTodo{}
	req := dto.TbTodoDeleteReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON, nil).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err.Error(), "")
		return
	}
	req.SetUpdateBy(user.GetUserId(c))

	err = s.Remove(&req)
	if err != nil {
		e.Error(500, err.Error(), "")
		return
	}
	reqPage := dto.TbTodoGetPageReq{}
	list := make([]models.TbTodo, 0)
	var count int64
	err = s.GetPage(&reqPage, &list, &count)
	if err != nil {
		e.Error(500, err.Error(), "")
		return
	}
	e.ListWithPage(list, count, 1, 10)
}
