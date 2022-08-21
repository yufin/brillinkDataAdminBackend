package service

import (
	"encoding/json"
	"fmt"
	"github.com/go-admin-team/go-admin-core/logger"
	"github.com/pkg/errors"
	cDto "go-admin/common/dto"
	"gorm.io/gorm"
)

type Service struct {
	Orm   *gorm.DB
	Msg   string
	MsgID string
	Log   *logger.Helper
	Error error
}

func (s *Service) GetOrm() *gorm.DB {
	return s.Orm
}
func (s *Service) GetMsg() string {
	return s.Msg
}
func (s *Service) GetMsgID() string {
	return s.MsgID
}
func (s *Service) GetLog() *logger.Helper {
	return s.Log
}
func (s *Service) GetError() error {
	return s.Error
}

func (s *Service) AddError(err error) error {
	if s.Error == nil {
		s.Error = err
	} else if err != nil {
		s.Error = fmt.Errorf("%+v$$$%+v", s.Error, err)
	}
	return s.Error
}

func (s *Service) Errorf(err error, msg string) error {
	return fmt.Errorf("%v;%+v", msg, err)
}

type iService interface {
	GetOrm() *gorm.DB
	GetMsg() string
	GetMsgID() string
	GetLog() *logger.Helper
	GetError() error
}

type iSearch interface {
	GetNeedSearch() interface{}
	GetPageSize() int
	GetPageIndex() int
}

type iList interface {
	GetNeedSearch() interface{}
}

type iGet interface {
	GetId() interface{}
}

// 针对 Insert、Update、Delete等操作
type action[T any] interface {
	Generate(*T)
	GetId() interface{}
}

type iUpdate interface {
	Generate() interface{}
	GetId() interface{}
}

// GetPage 分页
func GetPage[T any](service iService, search iSearch, scope ...func(db *gorm.DB) *gorm.DB) (*[]T, *int64, error) {
	list := new([]T)
	var count int64
	db := service.GetOrm().Model(new(T))
	for _, f := range scope {
		db = db.Scopes(f)
	}
	db = db.Scopes(
		cDto.MakeCondition(search.GetNeedSearch()),
		cDto.Paginate(search.GetPageSize(), search.GetPageIndex()),
	)
	err := db.Find(list).Limit(-1).Offset(-1).
		Count(&count).Error
	if err != nil {
		err = errors.WithStack(err)
		return list, &count, err
	}
	return list, &count, err
}

// Get 根据id获取数据
func Get[T any](service iService, get iGet, scope ...func(db *gorm.DB) *gorm.DB) (*T, error) {
	model := new(T)
	db := service.GetOrm().Model(new(T))
	for _, f := range scope {
		db = db.Scopes(f)
	}
	db = db.First(model, get.GetId())
	err := db.Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.Wrap(err, "查看对象不存在或无权查看")
		service.GetLog().Errorf("db error:%s", err)
		return model, err
	}
	if err != nil {
		err = errors.WithStack(err)
		service.GetLog().Errorf("db error:%s", err)
		return model, err
	}
	return model, err
}

// GetById 通过主键获取对象数据
func GetById[T any](service iService, get interface{}, scope ...func(db *gorm.DB) *gorm.DB) (*T, error) {
	model := new(T)
	db := service.GetOrm().Model(new(T))
	for _, f := range scope {
		db = db.Scopes(f)
	}
	db = db.First(model, get)
	err := db.Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.Wrap(err, "查看对象不存在或无权查看")
		service.GetLog().Errorf("db error:%s", err)
		return model, err
	}
	if err != nil {
		err = errors.WithStack(err)
		service.GetLog().Errorf("db error:%s", err)
		return model, err
	}
	return model, err
}

func Insert[T any](service iService, insert action[T]) (*T, error) {
	model := new(T)
	insert.Generate(model)
	err := service.GetOrm().Create(model).Error
	if err != nil {
		err = errors.WithStack(err)
		service.GetLog().Errorf("db error:%s", err)
		return model, err
	}
	return model, err
}

func GetAndUpdate[T any](service iService, update action[T], scope ...func(db *gorm.DB) *gorm.DB) (before, after []byte, ok bool, err error) {
	model := new(T)
	tx := service.GetOrm().Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	tx.Model(new(T))
	for _, f := range scope {
		tx = tx.Scopes(f)
	}
	tx = tx.First(&model, update.GetId())
	before, err = json.Marshal(model)
	update.Generate(model)
	// TODO: 需要修改为指定字段更新，目前更新存在风险
	tx = tx.Save(&model)
	if tx.Error != nil {
		err = errors.WithStack(tx.Error)
		service.GetLog().Errorf("database operation failed:%s", err)
		return before, after, false, err
	}
	if tx.RowsAffected == 0 {
		return before, after, false, errors.New("无权更新该数据")
	}
	after, err = json.Marshal(*model)
	return before, after, true, err
}

func UpdateForReq[T any](service iService, update iUpdate, scope ...func(db *gorm.DB) *gorm.DB) (before, after []byte, ok bool, err error) {
	model := new(T)
	db := service.GetOrm().Model(model)
	for _, f := range scope {
		db = db.Scopes(f)
	}
	db = db.Updates(update.Generate())
	if err = db.Error; err != nil {
		err = errors.WithStack(err)
		return
	}
	if db.RowsAffected == 0 {
		err = errors.WithStack(db.Error)
		return
	}
	return
}

func Delete[T any](service iService, delete action[T], scope ...func(db *gorm.DB) *gorm.DB) (bool, error) {
	model := new(T)
	delete.Generate(model)
	db := service.GetOrm().Debug()
	for _, f := range scope {
		db = db.Scopes(f)
	}
	db = db.Delete(model, delete.GetId())
	if db.Error != nil {
		err := errors.WithStack(db.Error)
		service.GetLog().Errorf("Delete error: %s", err)
		return false, err
	}
	if db.RowsAffected == 0 {
		err := errors.New("无权删除该数据")
		return false, err
	}
	return true, nil
}

func GetList[T any](service iService, search iList, scope ...func(db *gorm.DB) *gorm.DB) (list *[]T, err error) {
	list = new([]T)
	db := service.GetOrm().Model(new(T))
	for _, f := range scope {
		db = db.Scopes(f)
	}
	if search != nil {
		db = db.
			Scopes(
				cDto.MakeCondition(search.GetNeedSearch()),
			)
	}
	err = db.
		Find(list).Error
	if err != nil {
		err = errors.WithStack(err)
		service.GetLog().Errorf("db error:%s", err)
		return
	}
	return
}
