package service

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk"
	"github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	"go-admin/common"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
	"go-admin/common/middleware"
	"go-admin/common/service"
	"reflect"
	"runtime"
	"time"

	"gorm.io/gorm"
)

type SysConfig struct {
	service.Service
}

// GetPage 获取SysConfig列表
func (e *SysConfig) GetPage(r *dto.SysConfigGetPageReq, p *actions.DataPermission) (list *[]models.SysConfig, count *int64, err error) {
	list, count, err = service.GetPage[models.SysConfig](e, r)
	if err != nil {
		err = errors.WithStack(err)
		e.GetLog().Errorf("database operation failed:%s \r", err)
		return
	}
	return
}

// Get 获取SysConfig对象
func (e *SysConfig) Get(r *dto.SysConfigGetReq, p *actions.DataPermission) (model *models.SysConfig, err error) {
	model, err = service.Get[models.SysConfig](e, r, func(db *gorm.DB) *gorm.DB {
		return db.Scopes(
			actions.Permission(models.SysConfig{}.TableName(), p),
		)
	})
	if err != nil {
		err = errors.WithStack(err)
		e.GetLog().Errorf("database operation failed:%s \r", err)
		return
	}
	return
}

// Insert 创建SysConfig对象
func (e *SysConfig) Insert(c *gin.Context, r *dto.SysConfigControl) (err error) {
	model := new(models.SysConfig)
	model, err = service.Insert[models.SysConfig](e, r)
	if err != nil {
		e.GetLog().Errorf("database operation failed:%s \r", err)
		return
	}
	sdk.Runtime.SetConfig(c.Request.RequestURI, r.ConfigKey, r.ConfigValue)
	after, _ := json.Marshal(model)
	middleware.SetContextOperateLog(c,
		"新增",
		runtime.FuncForPC(reflect.ValueOf(e.Remove).Pointer()).Name()+
			fmt.Sprintf("数据，ID：%v", model.GetId()),
		"{}",
		string(after),
		"系统配置",
	)
	return
}

// Update 修改SysConfig对象
func (e *SysConfig) Update(c *gin.Context, r *dto.SysConfigControl) (err error) {
	var before, after []byte
	before, after, ok, err := service.GetAndUpdate[models.SysConfig](e, r)
	if err != nil {
		e.GetLog().Errorf("database operation failed:%s \r", err)
		return
	}
	sdk.Runtime.SetConfig(c.Request.RequestURI, r.ConfigKey, r.ConfigValue)
	if ok {
		middleware.SetContextOperateLog(c,
			"修改",
			runtime.FuncForPC(reflect.ValueOf(e.Remove).Pointer()).Name()+
				fmt.Sprintf("数据，ID：%v", r.GetId()),
			string(before),
			string(after),
			"系统配置",
		)
	}
	return
}

//// SetSysConfig 修改SysConfig对象
//func (e *SysConfig) SetSysConfig(c *[]dto.GetSetSysConfigReq) error {
//	var err error
//	for _, req := range *c {
//		var model = models.SysConfig{}
//		e.Orm.Where("config_key = ?", req.ConfigKey).First(&model)
//		if model.Id != 0 {
//			req.Generate(&model)
//			db := e.Orm.Save(&model)
//			err = db.Error
//			if err != nil {
//				err = errors.WithStack(err)
//				e.Log.Errorf("Service SetSysConfig error:%s", err)
//				return err
//			}
//			if db.RowsAffected == 0 {
//				return errors.New("无权更新该数据")
//			}
//			sdk.Runtime.SetConfig(c.Request.RequestURI,req.ConfigKey, req.ConfigValue)
//		}
//	}
//	return nil
//}

func (e *SysConfig) GetForSet(unCustom *[]dto.GetSetSysConfigReq, custom *[]dto.GetSetSysConfigCustomResp) error {
	var err error
	var data models.SysConfig

	err = e.Orm.Model(&data).Where("config_module != ?", "custom").
		Find(unCustom).Error
	if err != nil {
		err = errors.WithStack(err)
		e.Log.Errorf("Service GetSysConfigPage error:%s", err)
		return err
	}
	err = e.Orm.Model(&data).Where("config_module = ?", "custom").
		Find(custom).Error
	if err != nil {
		err = errors.WithStack(err)
		e.Log.Errorf("Service GetSysConfigPage error:%s", err)
		return err
	}
	return nil
}

// UpdateForSet 更新配置
func (e *SysConfig) UpdateForSet(c *gin.Context, r *[]dto.UpdateSetSysConfigReq) error {
	m := *r
	for _, req := range m {
		if req.ConfigKey == "custom" {
			err := e.updateSetConfigWithCustom(c, req)
			if err != nil {
				err = errors.WithStack(err)
				return err
			}
		} else {
			var data models.SysConfig
			err := e.Orm.
				Where("config_key = ?", req.ConfigKey).
				First(&data).Error
			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					err = errors.New("未找到" + req.ConfigKey + "对应的信息，或者对该数据无权处理！")
				}
				e.Log.Errorf("Service GetSysConfigPage error:%s", err)
				err = errors.WithStack(err)
				return err
			}
			before, err := json.Marshal(&data)
			if err != nil {
				err = errors.WithStack(err)
				e.Log.Errorf("数据格式化失败:%s", err)
				return err
			}
			if data.ConfigValue != req.ConfigValue && req.ConfigValue != "*********" {
				data.ConfigValue = req.ConfigValue.(string)
				err := e.Orm.Model(&data).
					Update("config_value", req.ConfigValue).
					Error
				if err != nil {
					e.Log.Errorf("Service GetSysConfigPage error:%s", err)
					err = errors.WithStack(err)
					return err
				}
				sdk.Runtime.SetConfig(c.Request.RequestURI, req.ConfigKey, req.ConfigValue)
				after, err := json.Marshal(&data)
				if err != nil {
					err = errors.WithStack(err)
					e.Log.Errorf("数据格式化失败:%s", err)
					return err
				}
				middleware.SetContextOperateLog(c,
					middleware.OperateUpdate,
					fmt.Sprintf("更新配置数据，ID：%v", data.GetId()),
					string(before),
					string(after),
					"系统配置",
				)
			}
		}
	}
	return nil
}

// updateSetConfigWithCustom 更新自定义参数
func (e *SysConfig) updateSetConfigWithCustom(c *gin.Context, req dto.UpdateSetSysConfigReq) error {
	list := req.ConfigValue
	configList, err := service.GetListOutDiff[models.SysConfig, int64](e, nil, func(db *gorm.DB) *gorm.DB {
		return db.Select("id").Where("config_module = 'custom' ")
	})
	if err != nil {
		return err
	}

	for _, cc := range list.([]interface{}) {
		marshal, err := json.Marshal(cc)
		if err != nil {
			err = errors.Wrap(err, "配置项解析失败!")
			return err
		}
		cus := dto.UpdateSetSysConfigCustomReq{}
		err = json.Unmarshal(marshal, &cus)
		if err != nil {
			err = errors.Wrap(err, "配置项解析失败!")
			return err
		}
		var isGet = true
		var oldConfig models.SysConfig
		var newConfig models.SysConfig
		db := e.Orm.Model(&oldConfig)
		if cus.IsInsert && cus.Id != 0 {
			cus.Id = 0
			db = db.Where("config_key = ? ", cus.ConfigKey)
		} else {
			db = db.Where("id = ?", cus.Id)
		}
		db = db.First(&oldConfig)
		err = db.Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			e.Log.Errorf("查询配置项失败:%s", err)
			err = errors.WithStack(err)
			return err
		}
		// 如果没有找到，就新建
		if errors.Is(err, gorm.ErrRecordNotFound) {
			isGet = false
			cus.Generate(&newConfig)
			err = e.Orm.Model(&oldConfig).Create(&newConfig).Error
			if err != nil {
				var mysqlErr *mysql.MySQLError
				if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
					err = errors.New("配置项" + oldConfig.ConfigKey + "已存在,请更换其他键名")
					return err
				}
				err = errors.WithStack(err)
				e.Log.Errorf("新建配置项失败:%s", err)
				return err
			}
			after, err := json.Marshal(&newConfig)
			if err != nil {
				err = errors.WithStack(err)
				e.Log.Errorf("数据格式化失败:%s", err)
				return err
			}
			sdk.Runtime.SetConfig(c.Request.RequestURI, newConfig.ConfigKey, newConfig.ConfigValue)
			middleware.SetContextOperateLog(c,
				"创建",
				fmt.Sprintf("创建自定义参数数据，ID：%v", newConfig.GetId()),
				"{}",
				string(after),
				"系统配置",
			)
		}
		if oldConfig.ConfigType == "Y" {
			err = errors.New("配置项" + oldConfig.ConfigKey + "已存在,请更换其他键名!")
			return err
		}
		if cus.IsInsert && oldConfig.Id != 0 {
			err = errors.New("配置项" + oldConfig.ConfigKey + "已存在,请更换其他键名!")
			return err
		}
		if isGet {
			before, err := json.Marshal(&oldConfig)
			if err != nil {
				err = errors.WithStack(err)
				e.Log.Errorf("数据格式化失败:%s", err)
				return err
			}
			configList = common.DeleteSliceElms[int64](configList, oldConfig.Id)
			cus.Generate(&oldConfig)
			if err := e.Orm.Save(&oldConfig).Error; err != nil {
				e.Log.Errorf("更新配置项:%s", err)
				err = errors.WithStack(err)
				return err
			}
			sdk.Runtime.SetConfig(c.Request.RequestURI, oldConfig.ConfigKey, oldConfig.ConfigValue)
			after, err := json.Marshal(&oldConfig)
			if err != nil {
				err = errors.WithStack(err)
				e.Log.Errorf("数据格式化失败:%s", err)
				return err
			}
			middleware.SetContextOperateLog(c,
				middleware.OperateUpdate,
				fmt.Sprintf("更新自定义参数数据，ID：%v", oldConfig.GetId()),
				string(before),
				string(after),
				"系统配置",
			)
		}
	}
	if len(configList) > 0 {
		err = e.Orm.Model(&models.SysConfig{}).
			Where("config_module = 'custom' and id in ? ", configList).
			Update("deleted_at", time.Now()).
			Error
		if err != nil {
			err = errors.WithStack(err)
			e.Log.Errorf("删除失败:%s", err)
			return err
		}
		middleware.SetContextOperateLog(c,
			"删除",
			fmt.Sprintf("更新自定义参数数据，ID：%v", configList),
			"{}",
			"{}",
			"系统配置",
		)
	}
	return nil
}

// Remove 删除SysConfig
func (e *SysConfig) Remove(d *dto.SysConfigDeleteReq) error {
	var err error
	var data models.SysConfig

	db := e.Orm.Delete(&data, d.Ids)
	if db.Error != nil {
		err = db.Error
		e.Log.Errorf("Service RemoveSysConfig error:%s", err)
		return err
	}
	if db.RowsAffected == 0 {
		err = errors.New("无权删除该数据")
		return err
	}
	return nil
}

// GetWithKey 根据Key获取SysConfig
func (e *SysConfig) GetWithKey(c *dto.SysConfigByKeyReq, resp *dto.GetSysConfigByKEYForServiceResp) error {
	var err error
	var data models.SysConfig
	err = e.Orm.Table(data.TableName()).Where("config_key = ?", c.ConfigKey).First(resp).Error
	if err != nil {
		err = errors.WithStack(err)
		e.Log.Errorf("At Service GetSysConfigByKEY Error:%s", err)
		return err
	}
	return nil
}

func (e *SysConfig) GetWithKeyList(c *dto.SysConfigGetToSysAppReq, list *[]models.SysConfig) error {
	var err error
	err = e.Orm.
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
		).
		Find(list).Error
	if err != nil {
		err = errors.WithStack(err)
		e.Log.Errorf("Service GetSysConfigByKey error:%s", err)
		return err
	}
	return nil
}

func (e *SysConfig) GetAll(mp map[string]*gorm.DB) error {
	var err error
	var list []models.SysConfig
	for key, db := range mp {
		err = db.Model(models.SysConfig{}).Find(&list).Error
		if err != nil {
			err = errors.WithStack(err)
			e.Log.Errorf("At Service GetSysConfigByKEY Error:%s", err)
			return err
		}
		for _, config := range list {
			sdk.Runtime.SetConfig(key, config.ConfigKey, config.ConfigValue)
		}
	}

	return err
}
