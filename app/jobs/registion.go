package jobs

import (
	"github.com/go-admin-team/go-admin-core/logger"
	"go-admin/app/rskc/task"
)

// 需要将定义的struct 添加到字典中；
// 字典 key 可以配置到 自动任务 调用目标 中；
func InitJob() {
	jobList = map[string]JobsExec{
		"ExamplesOne":        ExamplesOne{},
		"SyncOriginContent":  task.SyncOriginContentTask{},
		"SyncTradesDetail":   task.SyncTradesDetailTask{},
		"SyncWaitList":       task.SyncWaitListTask{},
		"VerifyContentReady": task.VerifyContentReadyTask{},
		"CollateContent":     task.CollateContentTask{},
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
