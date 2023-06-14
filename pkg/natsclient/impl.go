package natsclient

import (
	"fmt"
	log "github.com/go-admin-team/go-admin-core/logger"
	"github.com/go-admin-team/go-admin-core/sdk/pkg"
	"github.com/nats-io/nats.go"
)

var (
	TaskJs                  nats.JetStreamContext
	SubContentNew           *nats.Subscription
	SubTradeNew             *nats.Subscription
	SubContentProcessNew    *nats.Subscription
	SubToRequestDecisionNew *nats.Subscription
)

type TaskStream struct {
}

func (TaskStream) StreamName() string {
	return "TASK"
}

func (TaskStream) Subjects() []string {
	return []string{TopicTaskRskcPrefix + ">"}
}

func (e TaskStream) StreamConfig() *nats.StreamConfig {
	return &nats.StreamConfig{
		Name:      e.StreamName(),
		Retention: nats.WorkQueuePolicy,
		Subjects:  e.Subjects(),
	}
}

func (e TaskStream) InitTaskStream() error {
	if TaskJs == nil {
		js, err := NatsConn.JetStream()
		if err != nil {
			return err
		}
		TaskJs = js
	}
	_, err := TaskJs.AddStream(e.StreamConfig())
	if err != nil {
		log.Error(pkg.Blue(fmt.Sprintf("AddStream error: %v", err)))
		//return err
	}
	//TaskJs.StreamInfo(e.StreamName())
	return nil
}

func (e TaskStream) InitSubscription() error {
	var err error
	if SubContentNew == nil {
		SubContentNew, err = TaskJs.PullSubscribe(TopicContentNew, "sub-content-new", nats.BindStream(e.StreamName()))
		if err != nil {
			log.Error(pkg.Blue(fmt.Sprintf("Add Subscribe error: %v", err)))
		}
	}
	if SubTradeNew == nil {
		SubTradeNew, err = TaskJs.PullSubscribe(TopicTradeNew, "sub-trade-new", nats.BindStream(e.StreamName()))
		if err != nil {
			log.Error(pkg.Blue(fmt.Sprintf("Add Subscribe error: %v", err)))
		}
	}
	if SubContentProcessNew == nil {
		SubContentProcessNew, err = TaskJs.PullSubscribe(TopicContentToProcessNew, "sub-contenttoprocess-new", nats.BindStream(e.StreamName()))
		if err != nil {
			log.Error(pkg.Blue(fmt.Sprintf("Add Subscribe error: %v", err)))
		}
	}
	if SubToRequestDecisionNew == nil {
		SubToRequestDecisionNew, err = TaskJs.PullSubscribe(TopicToRequestDecisionNew, "sub-requestdecision-new", nats.BindStream(e.StreamName()))
		if err != nil {
			log.Error(pkg.Blue(fmt.Sprintf("Add Subscribe error: %v", err)))
		}
	}
	return nil
}

func (e TaskStream) Setup() error {
	err := e.InitTaskStream()
	if err != nil {
		return err
	}
	err = e.InitSubscription()
	if err != nil {
		return err
	}
	return nil
}

func (e TaskStream) OnClose() error {
	if err := SubContentNew.Unsubscribe(); err != nil {
		return err
	}
	if err := SubTradeNew.Unsubscribe(); err != nil {
		return err
	}
	if err := CloseNats(); err != nil {
		return err
	}
	return nil
}
