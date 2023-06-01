package natsclient

import (
	"fmt"
	log "github.com/go-admin-team/go-admin-core/logger"
	"github.com/go-admin-team/go-admin-core/sdk/pkg"
	"github.com/nats-io/nats.go"
	extConfig "go-admin/config"
)

var NatsConn *nats.Conn

func InitNatsConn() error {
	if !extConfig.ExtConfig.Nats.Activate {
		return nil
	}
	if NatsConn != nil {
		return nil
	}
	log.Info(pkg.Green("Nats Connection Initializing...."))
	nc, err := nats.Connect(extConfig.ExtConfig.Nats.Uri)
	if err != nil {
		log.Errorf(pkg.Red(fmt.Sprintf("Nats Connection initializing Failed. %v", err)))
		return err
	}
	NatsConn = nc
	log.Info(pkg.Green("Nats Connection initialized."))
	return nil
}

func CloseNats() error {
	//if !extConfig.ExtConfig.Nats.Activate {
	//	return nil
	//}
	if NatsConn == nil {
		return nil
	}
	err := NatsConn.Drain()
	if err != nil {
		log.Errorf(pkg.Red(fmt.Sprintf("Nats Connection close Failed. %v", err)))
		return err
	}
	return nil
}
