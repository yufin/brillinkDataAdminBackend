package utils

import (
	log "github.com/go-admin-team/go-admin-core/logger"
	"github.com/sony/sonyflake"
)

var Flake *sonyflake.Sonyflake = nil

// GetFlake 获取flake实例
func GetFlake() *sonyflake.Sonyflake {
	if Flake == nil {
		st := sonyflake.Settings{}
		Flake = sonyflake.NewSonyflake(st)
		if Flake == nil {
			log.Error("sonyflake not created \r\n")
			panic("sonyflake not created")
		}
	}
	return Flake
}

func NewFlakeId() int64 {
	id, err := GetFlake().NextID()
	if err != nil {
		log.Errorf("sonyflake nextId error:%s \r\n", err)
		panic("sonyflake not created")
	}
	return int64(id)
}
