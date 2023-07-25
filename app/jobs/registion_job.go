package jobs

import (
	"github.com/go-admin-team/go-admin-core/logger"
	gTask "go-admin/app/graph/task"
	"go-admin/app/rc/task/dataverification"
	"go-admin/app/rc/task/pdfsnapshot"
	"go-admin/app/rc/task/rdm"
	"go-admin/app/rc/task/reportbuilder"
	"go-admin/app/rc/task/syncdependency"
)

// 需要将定义的struct 添加到字典中；
// 字典 key 可以配置到 自动任务 调用目标 中；
func InitJob() {
	jobList = map[string]JobsExec{
		"ExamplesOne":         ExamplesOne{},
		"SyncOriginContent":   syncdependency.SyncOriginContentTask{},
		"SyncDependencyTable": syncdependency.DependencyTableSyncTask{},
		"SyncWaitList":        syncdependency.SyncWaitListTask{},
		"VerifyContentReady":  dataverification.VerifyContentReadyTask{},
		"CollateContent":      reportbuilder.CollateContentTask{},
		"SyncGraph":           gTask.SyncGraphTask{},
		"DecisionFlow":        rdm.AhpRdmTask{},
		"reportSnapshot":      pdfsnapshot.ReportSnapshotTask{},
		//"WipeMsg":            task.WipeMsgTask{},
	}
}

// 新添加的job 必须按照以下格式定义，并实现Exec函数
type ExamplesOne struct {
}

func (t ExamplesOne) Exec(arg interface{}) error {
	str := "JobCore ExamplesOne success"
	// TODO: 这里需要注意 Examples 传入参数是 string 所以 arg.(string)；请根据对应的类型进行转化；
	switch arg.(type) {

	case string:
		if arg.(string) != "" {
			logger.Info(str, arg.(string))
		} else {
			logger.Warn(str, "arg is nil")
		}
		break
	}
	return nil
}
