package service

import (
	"encoding/json"
	"github.com/go-admin-team/go-admin-core/sdk"
	"github.com/pkg/errors"
	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	"go-admin/common"
	cDto "go-admin/common/dto"
	"go-admin/common/service"
	"time"

	"gorm.io/gorm"
)

type SysConfig struct {
	service.Service
}

// GetPage 获取SysConfig列表
func (e *SysConfig) GetPage(c *dto.SysConfigGetPageReq, list *[]models.SysConfig, count *int64) error {
	err := e.Orm.
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		err = errors.WithStack(err)
		e.Log.Errorf("Service GetSysConfigPage error:%s", err)
		return err
	}
	return nil
}

// Get 获取SysConfig对象
func (e *SysConfig) Get(d *dto.SysConfigGetReq, model *models.SysConfig) error {
	err := e.Orm.First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetSysConfigPage error:%s", err)
		return err
	}
	if err != nil {
		err = errors.WithStack(err)
		e.Log.Errorf("Service GetSysConfig error:%s", err)
		return err
	}
	return nil
}

// Insert 创建SysConfig对象
func (e *SysConfig) Insert(c *dto.SysConfigControl) error {
	var err error
	var data models.SysConfig
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		err = errors.WithStack(err)
		e.Log.Errorf("Service InsertSysConfig error:%s", err)
		return err
	}
	return nil
}

// Update 修改SysConfig对象
func (e *SysConfig) Update(c *dto.SysConfigControl) error {
	var err error
	var model = models.SysConfig{}
	e.Orm.First(&model, c.GetId())
	c.Generate(&model)
	db := e.Orm.Save(&model)
	err = db.Error
	if err != nil {
		err = errors.WithStack(err)
		e.Log.Errorf("Service UpdateSysConfig error:%s", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")

	}
	return nil
}

// SetSysConfig 修改SysConfig对象
func (e *SysConfig) SetSysConfig(c *[]dto.GetSetSysConfigReq) error {
	var err error
	for _, req := range *c {
		var model = models.SysConfig{}
		e.Orm.Where("config_key = ?", req.ConfigKey).First(&model)
		if model.Id != 0 {
			req.Generate(&model)
			db := e.Orm.Save(&model)
			err = db.Error
			if err != nil {
				err = errors.WithStack(err)
				e.Log.Errorf("Service SetSysConfig error:%s", err)
				return err
			}
			if db.RowsAffected == 0 {
				return errors.New("无权更新该数据")
			}
		}
	}
	return nil
}

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

func (e *SysConfig) UpdateForSet(c *[]dto.UpdateSetSysConfigReq) error {
	m := *c
	for _, req := range m {
		if req.ConfigKey == "custom" {
			err := updateSetConfigWithCustom(req, e)
			if err != nil {
				err = errors.WithStack(err)
				return err
			}
		} else {
			var data models.SysConfig
			if err := e.Orm.Where("config_key = ?", req.ConfigKey).
				First(&data).Error; err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					err = errors.New("未找到" + req.ConfigKey + "对应的信息，或者对该数据无权处理！")
				}
				e.Log.Errorf("Service GetSysConfigPage error:%s", err)
				err = errors.WithStack(err)
				return err
			}
			if data.ConfigValue != req.ConfigValue && req.ConfigValue != "*********" {
				data.ConfigValue = req.ConfigValue.(string)
				if err := e.Orm.Model(&data).Update("config_value", req.ConfigValue).Error; err != nil {
					e.Log.Errorf("Service GetSysConfigPage error:%s", err)
					err = errors.WithStack(err)
					return err
				}
			}
		}
	}

	return nil
}

// updateSetConfigWithCustom 更新自定义参数
func updateSetConfigWithCustom(req dto.UpdateSetSysConfigReq, e *SysConfig) error {
	list := req.ConfigValue
	var configList []int64
	err := e.Orm.Model(&models.SysConfig{}).Select("id").Where("config_module = 'custom' ").Find(&configList).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		e.Log.Errorf("获取自定义配置项:%s", err)
		err = errors.WithStack(err)
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
		var d models.SysConfig
		var newConfig models.SysConfig
		db := e.Orm.Model(&d).Where("config_key = ? ", cus.ConfigKey).First(&d)
		err = db.Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			e.Log.Errorf("查询配置项失败:%s", err)
			err = errors.WithStack(err)
			return err
		}
		if errors.Is(err, gorm.ErrRecordNotFound) {
			isGet = false
			cus.Generate(&newConfig)
			err = e.Orm.Model(&d).Create(&newConfig).Error
			if err != nil {
				err = errors.WithStack(err)
				e.Log.Errorf("新建配置项失败:%s", err)
				return err
			}
		}
		if d.ConfigType == "Y" {
			err = errors.New("配置项" + d.ConfigKey + "已存在,请更换其他键名!")
			return err
		}
		if isGet {
			configList = common.DeleteSliceElms[int64](configList, d.Id)
			cus.Generate(&d)
			if err := e.Orm.Save(&d).Error; err != nil {
				e.Log.Errorf("更新配置项:%s", err)
				err = errors.WithStack(err)
				return err
			}
		}
	}
	if len(configList) > 0 {
		err = e.Orm.Model(&models.SysConfig{}).Where("config_module = 'custom' and id in ? ", configList).Update("deleted_at", time.Now()).Error
		if err != nil {
			err = errors.WithStack(err)
			e.Log.Errorf("删除失败:%s", err)
			return err
		}
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

func (e *SysConfig) GetAll(orm *gorm.DB) error {
	var err error
	var list []models.SysConfig
	err = orm.Model(models.SysConfig{}).Find(&list).Error
	if err != nil {
		err = errors.WithStack(err)
		e.Log.Errorf("At Service GetSysConfigByKEY Error:%s", err)
		return err
	}
	for _, config := range list {
		sdk.Runtime.SetConfig(config.ConfigKey, config.ConfigValue)
	}
	return err
}
