package task

import (
	"github.com/nats-io/nats.go"
	"go-admin/common/natsclient"
	"time"
)

type WipeMsgTask struct {
}

func (t WipeMsgTask) Exec(arg interface{}) error {
	return wipeMsg()
}

func wipeMsg() error {
	subs := []*nats.Subscription{
		natsclient.SubContentProcessNew,
		natsclient.SubContentNew,
		natsclient.SubTradeNew,
	}
	for _, sub := range subs {

		msgs, err := sub.Fetch(1, nats.MaxWait(1*time.Second))
		if err != nil {
			if err == nats.ErrTimeout {
				return nil
			} else {
				return err
			}
		}
		for _, msg := range msgs {
			if err := msg.AckSync(); err != nil {
				return err
			}
		}
	}
	return nil
}
