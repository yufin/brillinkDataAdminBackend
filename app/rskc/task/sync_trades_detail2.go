package task

import (
	"encoding/binary"
	"fmt"
	"github.com/nats-io/nats.go"
	"go-admin/common/natsclient"
	"time"
)

type SyncTradesDetailTask struct {
}

func (t SyncTradesDetailTask) Exec(arg interface{}) error {

	return SyncTask()
}

func SyncTask() error {
	for {
		msgs, err := natsclient.SubContentNew.Fetch(1, nats.MaxWait(5*time.Second))
		if err != nil {
			if err == nats.ErrTimeout {
				return nil
			} else {
				return err
			}
		}
		for _, msg := range msgs {
			contentId := int64(binary.BigEndian.Uint64(msg.Data))
			fmt.Println("contentId: ", contentId)
			msg.AckSync()
		}
	}

}
