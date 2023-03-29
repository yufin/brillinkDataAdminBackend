package apis

import (
	"fmt"
	"go-admin/common/exception"
	"mime/multipart"
	"net/http"
	"strconv"

	vd "github.com/bytedance/go-tagexpr/v2/validator"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-admin-team/go-admin-core/logger"
	"github.com/go-admin-team/go-admin-core/sdk/pkg"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"go-admin/common/response/antd"
	"go-admin/common/service"
)

type Api struct {
	Context *gin.Context
	Logger  *logger.Helper
	Orm     *gorm.DB
	Errors  error
}

// GetLogger 获取上下文提供的日志
func (e Api) GetLogger() *logger.Helper {
	return GetRequestLogger(e.Context)
}

// GetOrm 获取Orm DB
func (e *Api) GetOrm(c *gin.Context) (*gorm.DB, error) {
	db, err := pkg.GetOrm(c)
	if err != nil {
		err = errors.Wrap(err, "数据库连接获取失败")
		panic(exception.WithStatus(http.StatusInternalServerError, exception.GetOrmFail, err))
		return nil, err
	}
	return db, nil
}

// Error 通常错误数据处理
// showType error display type： 0 silent; 1 message.warn; 2 message.error; 4 notification; 9 page
func (e *Api) Error(errCode int, errMsg string, showType string) {
	if showType == "" {
		showType = "2"
	}
	antd.Error(e.Context, strconv.Itoa(errCode), errMsg, showType)
}

func (e *Api) SaveUploadedFile(c *gin.Context, file *multipart.FileHeader, filPath string) error {
	if err := c.SaveUploadedFile(file, filPath); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// OK 通常成功数据处理
func (e *Api) OK(data interface{}) {
	antd.OK(e.Context, data)
}

// PageOK 分页数据处理
func (e *Api) PageOK(result interface{}, total int64, current int, pageSize int) {
	antd.PageOK(e.Context, result, total, current, pageSize)
}

func (e *Api) ListOK(result interface{}) {
	antd.ListOK(e.Context, result)
}

func (e *Api) ListWithPage(result interface{}, total int64, current int, pageSize int) {
	antd.ListWithPage(e.Context, result, total, current, pageSize)
}

// Custom 兼容函数
func (e *Api) Custom(data gin.H) {
	antd.Custum(e.Context, data)
}

// MakeContext 设置http上下文
func (e *Api) MakeContext(c *gin.Context) *Api {
	e.Context = c
	e.Logger = GetRequestLogger(c)
	return e
}

// Bind 参数校验
func (e *Api) Bind(d interface{}, bindings ...binding.Binding) *Api {
	var err error
	if len(bindings) == 0 {
		bindings = constructor.GetBindingForGin(d)
	}
	for i := range bindings {
		if bindings[i] == nil {
			err = e.Context.ShouldBindUri(d)
		} else {
			err = e.Context.ShouldBindWith(d, bindings[i])
		}
		if err != nil && err.Error() == "EOF" {
			e.Logger.Warn("request body is not present anymore. ")
			err = nil
			continue
		}
		if err != nil {
			err = errors.WithStack(err)
			e.AddError(err)
			break
		}
	}
	if err1 := vd.Validate(d); err1 != nil {
		err1 = errors.WithStack(err1)
		e.AddError(err1)
	}
	return e
}

// MakeOrm 设置Orm DB
func (e *Api) MakeOrm() *Api {
	var err error
	if e.Logger == nil {
		err = errors.New("at MakeOrm logger is nil")
		e.AddError(err)
		return e
	}
	db, err := pkg.GetOrm(e.Context)
	if err != nil {
		e.Logger.Error(http.StatusInternalServerError, err, "数据库连接获取失败")
		err = errors.Wrap(err, "数据库连接获取失败")
		panic(exception.WithStatus(http.StatusInternalServerError, exception.GetOrmFail, err))
	}
	e.Orm = db
	return e
}

func (e *Api) MakeService(c *service.Service) *Api {
	c.Log = e.Logger
	c.Orm = e.Orm
	return e
}

func (e *Api) AddError(err error) {
	if e.Errors == nil {
		e.Errors = err
	} else if err != nil {
		e.Logger.Error(err)
		e.Errors = fmt.Errorf("%+v$$$%+v", e.Errors, err)
	}
}

func (e Api) Translate(form, to interface{}) {
	pkg.Translate(form, to)
}
