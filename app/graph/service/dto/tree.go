package dto

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type TreeNode struct {
	Id       string         `json:"entityId"`
	RandomId string         `json:"id"`
	Title    string         `json:"title"`
	Labels   []string       `json:"labels"`
	Data     map[string]any `json:"data"`
	Children []*TreeNode    `json:"children"`
}

func (t *TreeNode) SetChild(parentId string, neoChild neo4j.Node) bool {
	//if neoChild.Props["id"] == t.Id {
	// if added: the node targetNode would only show once.
	//	return true
	//}
	if t.Id == parentId {
		if t.Children != nil {
			for _, child := range t.Children {
				if child.Id == neoChild.Props["id"].(string) {
					return true
				}
			}
			(*t).Children = append((*t).Children, SerializeTreeNode(neoChild))
			return true
		} else {
			(*t).Children = []*TreeNode{SerializeTreeNode(neoChild)}
			return true
		}
	} else {
		if t.Children != nil {
			for _, child := range t.Children {
				r := child.SetChild(parentId, neoChild)
				if r {
					return true
				}
			}
		}
	}
	return false
}

func SerializeTreeNode(neoNode neo4j.Node) *TreeNode {
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

	return &TreeNode{
		Id:       neoNode.Props["id"].(string),
		RandomId: uuid.New().String(),
		Title:    title,
		Labels:   neoNode.Labels,
		Data:     data,
	}
}

func SerializeTreeFromPath(pNeoPath *[]neo4j.Path) TreeNode {
	root := SerializeTreeNode((*pNeoPath)[0].Nodes[0])
	//pRoot := &root
	for _, path := range *pNeoPath {
		for _, neoRel := range path.Relationships {
			pParentNode := GetNodeByElementId(&path.Nodes, neoRel.StartElementId)
			pChildNode := GetNodeByElementId(&path.Nodes, neoRel.EndElementId)
			if pParentNode != nil && pChildNode != nil {
				root.SetChild(pParentNode.Props["id"].(string), *pChildNode)
			}
		}
	}
	return *root
	//return *pRoot
}
