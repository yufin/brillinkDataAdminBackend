package task

import (
	"encoding/binary"
	"fmt"
	log "github.com/go-admin-team/go-admin-core/logger"
	"github.com/go-admin-team/go-admin-core/sdk"
	"github.com/nats-io/nats.go"
	"go-admin/app/rskc/models"
	"go-admin/pkg/natsclient"
	"sync/atomic"
	"time"
)

type DecisionFlowTask struct {
}

var decisionRunning int32

func (t DecisionFlowTask) Exec(arg interface{}) error {
	if atomic.LoadInt32(&decisionRunning) == 1 {
		log.Info("DecisionFlow任务已经在执行中，跳过本次调度")
		return nil
	}
	atomic.StoreInt32(&decisionRunning, 1)
	defer atomic.StoreInt32(&decisionRunning, 0)

	if err := pubIdsToRequestDecision(); err != nil {
		log.Errorf("selectWaitForRequest Failed:%v \r\n", err)
		return err
	}

	totalPending, _, err := natsclient.SubToRequestDecisionNew.Pending()
	if err == nil {
		fmt.Println("DecisionFlowTask msg totalPending:", totalPending)
	}

	msgs, err := natsclient.SubToRequestDecisionNew.Fetch(1, nats.MaxWait(5*time.Second))
	if err != nil {
		if err == nats.ErrTimeout {
			return nil
		} else {
			return err
		}
	}

	for _, msg := range msgs {
		contentId := int64(binary.BigEndian.Uint64(msg.Data))
		exists, err := CheckContentIdExist(contentId)
		if err != nil {
			return err
		} else {
			if !exists {
				if err := msg.AckSync(); err != nil {
					return err
				}
				break
			}
		}
		log.Infof("开始请求决策引擎: contentId = %d\r\n", contentId)
		if err := updateDependencyDataToParam(contentId); err != nil {
			log.Errorf("requestDecisionEngine Failed:%v \r\n", err)
			return err
		}
		if err := msg.AckSync(); err != nil {
			return err
		}
	}
	return nil
}

func pubIdsToRequestDecision() error {
	// select tbParam where has no fk id in tbResult, and has same contentId in tbData
	var tbParam models.RcDecisionParam
	db := sdk.Runtime.GetDbByKey(tbParam.TableName())
	contentIds := make([]int64, 0)
	err := db.
		Table(tbParam.TableName()).
		Raw(`select t.content_id as content_id
				from
					(select content_id, param_id
					from rc_decision_param rdp
					left join  rc_decision_result rdr on rdp.id = rdr.param_id
					where param_id is null and rdp.status_code = 0) t
				left join rc_dependency_data rdd
				on rdd.content_id = t.content_id where rdd.content_id is not null`).
		Pluck("content_id", &contentIds).Error
	if err != nil {
		return err
	}
	if len(contentIds) == 0 {
		return nil
	}
	for _, id := range contentIds {
		msg := make([]byte, 8)
		binary.BigEndian.PutUint64(msg, uint64(id))
		_, err := natsclient.TaskJs.Publish(natsclient.TopicToRequestDecisionNew, msg)
		if err != nil {
			return err
		}
		err = db.Model(models.RcDecisionParam{}).
			Where("content_id = ?", id).
			Update("status_code", 1).
			Error
		if err != nil {
			return err
		}
	}
	return nil
}
