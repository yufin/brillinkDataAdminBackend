package task

import (
	"github.com/go-admin-team/go-admin-core/sdk"
	"github.com/pkg/errors"
	"go-admin/app/spider/models"
	"go-admin/app/spider/service/dto"
)

func SyncTaskDetail(uscId string) error {
	if uscId == "" || uscId == "-" {
		return errors.New("uscId is empty")
	}

	taskTopics := []string{
		"enterprise_info",
		"enterprise_certification",
		"enterprise_industry",
		"enterprise_ranking",
		"enterprise_product",
	}
	existingTopics := make([]string, 0)
	var tb models.TaskDetail
	db := sdk.Runtime.GetDbByKey(tb.TableName())
	err := db.Model(&tb).
		Where("usc_id = ?", uscId).
		Pluck("topic", &existingTopics).
		Error
	if err != nil {
		return err
	}
	var missingTopics []string
	for _, topic := range taskTopics {
		found := false
		for _, existingTopic := range existingTopics {
			if topic == existingTopic {
				found = true
				break
			}
		}
		if !found {
			missingTopics = append(missingTopics, topic)
		}
	}

	if len(missingTopics) > 0 {
		for _, topic := range missingTopics {
			insertReq := dto.TaskDetailInsertReq{
				Topic:      topic,
				UscId:      uscId,
				StatusCode: 1,
				Priority:   9,
				Comment:    "data_collection",
			}
			var data models.TaskDetail
			insertReq.Generate(&data)
			err = db.Model(&tb).Create(&data).Error
			if err != nil {
				return err
			}
		}
	}
	return nil
}
