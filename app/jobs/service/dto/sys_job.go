package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"go-admin/app/jobs/models"
	"time"

	"go-admin/common/dto"
	common "go-admin/common/models"
)

type SysJobGetPageReq struct {
	dto.Pagination `search:"-"`

	JobId          int64     `form:"jobId"  search:"type:exact;column:job_id;table:sys_job" comment:""`
	JobName        string    `form:"jobName"  search:"type:exact;column:job_name;table:sys_job" comment:""`
	JobGroup       string    `form:"jobGroup"  search:"type:exact;column:job_group;table:sys_job" comment:""`
	JobType        string    `form:"jobType"  search:"type:exact;column:job_type;table:sys_job" comment:""`
	CronExpression string    `form:"cronExpression"  search:"type:exact;column:cron_expression;table:sys_job" comment:""`
	InvokeTarget   string    `form:"invokeTarget"  search:"type:exact;column:invoke_target;table:sys_job" comment:""`
	Status         string    `form:"status"  search:"type:exact;column:status;table:sys_job" comment:""`
	EntryId        int64     `form:"entryId"  search:"type:exact;column:entry_id;table:sys_job" comment:""`
	CreatedAt      time.Time `form:"createdAt"  search:"type:exact;column:created_at;table:sys_job" comment:""`
	UpdatedAt      time.Time `form:"updatedAt"  search:"type:exact;column:updated_at;table:sys_job" comment:""`
	SysJobPageOrder
}

type SysJobPageOrder struct {
	JobId     int64     `form:"jobIdOrder"  search:"type:order;column:job_id;table:sys_job"`
	JobName   string    `form:"jobNameOrder"  search:"type:order;column:job_name;table:sys_job"`
	JobGroup  string    `form:"jobGroupOrder"  search:"type:order;column:job_group;table:sys_job"`
	JobType   int64     `form:"jobTypeOrder"  search:"type:order;column:job_type;table:sys_job"`
	Status    int64     `form:"statusOrder"  search:"type:order;column:status;table:sys_job"`
	EntryId   int64     `form:"entryIdOrder"  search:"type:order;column:entry_id;table:sys_job"`
	CreatedAt time.Time `form:"createdAtOrder"  search:"type:order;column:created_at;table:sys_job"`
}

func (m *SysJobGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type SysJobInsertReq struct {
	JobId          int64  `json:"-" comment:""` //
	JobName        string `json:"jobName" comment:""`
	JobGroup       string `json:"jobGroup" comment:""`
	JobType        int64  `json:"jobType" comment:""`
	CronExpression string `json:"cronExpression" comment:""`
	InvokeTarget   string `json:"invokeTarget" comment:""`
	Args           string `json:"args" comment:""`
	MisfirePolicy  int64  `json:"misfirePolicy" comment:""`
	Concurrent     int64  `json:"concurrent" comment:""`
	Status         int64  `json:"status" comment:""`
	EntryId        int    `json:"entryId" comment:""`
	common.ControlBy
}

func (s *SysJobInsertReq) Generate(model *models.SysJob) {
	model.JobId = s.JobId
	model.JobName = s.JobName
	model.JobGroup = s.JobGroup
	model.JobType = s.JobType
	model.CronExpression = s.CronExpression
	model.InvokeTarget = s.InvokeTarget
	model.Args = s.Args
	model.MisfirePolicy = s.MisfirePolicy
	model.Concurrent = s.Concurrent
	model.Status = s.Status
	model.EntryId = s.EntryId
	model.CreateBy = s.CreateBy
}

func (s *SysJobInsertReq) GetId() interface{} {
	return s.JobId
}

type SysJobUpdateReq struct {
	JobId          int64  `json:"jobId" comment:""` //
	JobName        string `json:"jobName" comment:""`
	JobGroup       string `json:"jobGroup" comment:""`
	JobType        int64  `json:"jobType" comment:""`
	CronExpression string `json:"cronExpression" comment:""`
	InvokeTarget   string `json:"invokeTarget" comment:""`
	Args           string `json:"args" comment:""`
	MisfirePolicy  int64  `json:"misfirePolicy" comment:""`
	Concurrent     int64  `json:"concurrent" comment:""`
	Status         int64  `json:"status" comment:""`
	EntryId        int    `json:"entryId" comment:""`
	common.ControlBy
}

func (s *SysJobUpdateReq) Generate(model *models.SysJob) {
	model.JobId = s.JobId
	model.JobName = s.JobName
	model.JobGroup = s.JobGroup
	model.JobType = s.JobType
	model.CronExpression = s.CronExpression
	model.InvokeTarget = s.InvokeTarget
	model.Args = s.Args
	model.MisfirePolicy = s.MisfirePolicy
	model.Concurrent = s.Concurrent
	model.Status = s.Status
	model.EntryId = s.EntryId
	model.UpdateBy = s.UpdateBy
}

func (s *SysJobUpdateReq) GetId() interface{} {
	return s.JobId
}

// SysJobGetReq 功能获取请求参数
type SysJobGetReq struct {
	JobId int64 `uri:"id"`
}

func (s *SysJobGetReq) GetId() interface{} {
	return s.JobId
}

// SysJobDeleteReq 功能删除请求参数
type SysJobDeleteReq struct {
	Id int `uri:"id"`
}

func (s *SysJobDeleteReq) GetId() interface{} {
	return s.Id
}

// -- -----

type SysJobSearch struct {
	dto.Pagination `search:"-"`
	JobId          int    `form:"jobId" search:"type:exact;column:job_id;table:sys_job"`
	JobName        string `form:"jobName" search:"type:icontains;column:job_name;table:sys_job"`
	JobGroup       string `form:"jobGroup" search:"type:exact;column:job_group;table:sys_job"`
	CronExpression string `form:"cronExpression" search:"type:exact;column:cron_expression;table:sys_job"`
	InvokeTarget   string `form:"invokeTarget" search:"type:exact;column:invoke_target;table:sys_job"`
	Status         int    `form:"status" search:"type:exact;column:status;table:sys_job"`
}

func (m *SysJobSearch) GetNeedSearch() interface{} {
	return *m
}

func (m *SysJobSearch) Bind(ctx *gin.Context) error {
	log := api.GetRequestLogger(ctx)
	err := ctx.ShouldBind(m)
	if err != nil {
		log.Errorf("Bind error: %s", err)
	}
	return err
}

func (m *SysJobSearch) Generate() dto.Index {
	o := *m
	return &o
}

type SysJobControl struct {
	JobId          int64  `json:"jobId"`
	JobName        string `json:"jobName" validate:"required"` // 名称
	JobGroup       string `json:"jobGroup"`                    // 任务分组
	JobType        int64  `json:"jobType"`                     // 任务类型
	CronExpression string `json:"cronExpression"`              // cron表达式
	InvokeTarget   string `json:"invokeTarget"`                // 调用目标
	Args           string `json:"args"`                        // 目标参数
	MisfirePolicy  int64  `json:"misfirePolicy"`               // 执行策略
	Concurrent     int64  `json:"concurrent"`                  // 是否并发
	Status         int64  `json:"status"`                      // 状态
	EntryId        int    `json:"entryId"`                     // job启动时返回的id
}

func (s *SysJobControl) Bind(ctx *gin.Context) error {
	return ctx.ShouldBind(s)
}

func (s *SysJobControl) Generate() dto.Control {
	cp := *s
	return &cp
}

func (s *SysJobControl) GenerateM() (common.ActiveRecord, error) {
	return &models.SysJob{
		JobId:          s.JobId,
		JobName:        s.JobName,
		JobGroup:       s.JobGroup,
		JobType:        s.JobType,
		CronExpression: s.CronExpression,
		InvokeTarget:   s.InvokeTarget,
		Args:           s.Args,
		MisfirePolicy:  s.MisfirePolicy,
		Concurrent:     s.Concurrent,
		Status:         s.Status,
		EntryId:        s.EntryId,
	}, nil
}

func (s *SysJobControl) GetId() interface{} {
	return s.JobId
}

type SysJobById struct {
	dto.ObjectById
}

func (s *SysJobById) Generate() dto.Control {
	cp := *s
	return &cp
}

func (s *SysJobById) GenerateM() (common.ActiveRecord, error) {
	return &models.SysJob{}, nil
}

type SysJobItem struct {
	JobId          int    `json:"jobId"`
	JobName        string `json:"jobName" validate:"required"` // 名称
	JobGroup       string `json:"jobGroup"`                    // 任务分组
	JobType        int    `json:"jobType"`                     // 任务类型
	CronExpression string `json:"cronExpression"`              // cron表达式
	InvokeTarget   string `json:"invokeTarget"`                // 调用目标
	Args           string `json:"args"`                        // 目标参数
	MisfirePolicy  int    `json:"misfirePolicy"`               // 执行策略
	Concurrent     int    `json:"concurrent"`                  // 是否并发
	Status         int    `json:"status"`                      // 状态
	EntryId        int    `json:"entryId"`                     // job启动时返回的id
}
