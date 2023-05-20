package utils

import (
	"encoding/json"
	log "github.com/go-admin-team/go-admin-core/logger"
	"os"
)

func ReadFile(path string) []byte {
	content, err := os.ReadFile(path)
	if err != nil {
		log.Error(err)
		panic(err)
	}
	return content
}

func ParseJsonBytes(contentMap *map[string]any) {
	content := ReadFile("/Users/yufei/Github/graphAdmin/graphAdminBackend/91350200MA31RD1P7H_import_export_enterprises_report_json.json")
	err := json.Unmarshal(content, &contentMap)
	if err != nil {
		log.Error("Parse Json String Error!", err)
		panic(err)
	}
}
