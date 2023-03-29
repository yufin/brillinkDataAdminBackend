package dto

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type TreeNode struct {
	Id       string         `json:"id"`
	Title    string         `json:"title"`
	Labels   []string       `json:"labels"`
	Data     map[string]any `json:"data"`
	EntityId string         `json:"entityId"`
	Children []TreeNode     `json:"children"`
}

func SerializeTreeNode(neoNode neo4j.Node) TreeNode {
	copyProps := make(map[string]any)
	for k, v := range neoNode.Props {
		copyProps[k] = v
	}
	var data map[string]any
	if dataJson, ok := neoNode.Props["data"]; ok {
		err := json.Unmarshal([]byte(dataJson.(string)), &data)
		if err != nil {
			data = make(map[string]any)
		}
		delete(copyProps, "data")
	} else {
		data = make(map[string]any)
	}

	var title string
	if titleTemp, ok := neoNode.Props["title"]; ok {
		title = titleTemp.(string)
		delete(copyProps, "title")
	} else {
		title = neoNode.Props["titleShort"].(string)
		delete(copyProps, "titleShort")
	}
	delete(copyProps, "id")
	for k, v := range copyProps {
		data[k] = v
	}

	return TreeNode{
		EntityId: neoNode.Props["id"].(string),
		Id:       uuid.New().String(),
		Title:    title,
		Labels:   neoNode.Labels,
		Data:     data,
	}
}
