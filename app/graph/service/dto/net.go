package dto

import (
	"encoding/json"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"strings"
)

type LinkNode struct {
	Id     string         `json:"id"`
	Labels []string       `json:"labels"`
	Title  string         `json:"title"`
	Data   map[string]any `json:"data"`
}

type Net struct {
	Nodes []LinkNode `json:"nodes"`
	Edges []Edge     `json:"edges"`
}

type Edge struct {
	SourceId string         `json:"source"`
	TargetId string         `json:"target"`
	Id       string         `json:"id"`
	Label    string         `json:"label"`
	Data     map[string]any `json:"data"`
}

func SerializeLinkNode(neoNode neo4j.Node) LinkNode {
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

	return LinkNode{
		Id:     neoNode.Props["id"].(string),
		Labels: neoNode.Labels,
		Title:  title,
		Data:   data,
	}
}

// SerializeEdge warn: Returned instance of Edge Has no SourceId and TargetId
func SerializeEdge(neoRel neo4j.Relationship) Edge {
	copyProps := make(map[string]any)
	for k, v := range neoRel.Props {
		copyProps[k] = v
	}

	var data map[string]any
	if dataJson, ok := neoRel.Props["data"]; ok {
		err := json.Unmarshal([]byte(dataJson.(string)), &data)
		if err != nil {
			data = make(map[string]any)
		}
		delete(copyProps, "data")
	} else {
		data = make(map[string]any)
	}

	delete(copyProps, "id")
	for k, v := range copyProps {
		data[k] = v
	}

	return Edge{
		Id:    neoRel.Props["id"].(string),
		Label: neoRel.Type,
		Data:  data,
	}
}

func SerializeNetFromPath(pNeoPath *[]neo4j.Path) Net {
	nodes := make([]LinkNode, 0)
	edges := make([]Edge, 0)
	var nodeIds []string

	for _, path := range *pNeoPath {

		for _, neoNode := range path.Nodes {
			if !strings.Contains(strings.Join(nodeIds, ","), neoNode.Props["id"].(string)) {
				nodes = append(nodes, SerializeLinkNode(neoNode))
				nodeIds = append(nodeIds, neoNode.Props["id"].(string))
			}
		}

		for _, neoRel := range path.Relationships {
			edge := SerializeEdge(neoRel)
			pStartNode := GetNodeByElementId(&path.Nodes, neoRel.StartElementId)
			pEndNode := GetNodeByElementId(&path.Nodes, neoRel.EndElementId)
			if pStartNode != nil {
				edge.SourceId = (*pStartNode).Props["id"].(string)
			}
			if pEndNode != nil {
				edge.TargetId = (*pEndNode).Props["id"].(string)
			}
			edges = append(edges, edge)
		}
	}
	return Net{
		Nodes: nodes,
		Edges: edges,
	}
}

func GetNodeByElementId(neoNodes *[]neo4j.Node, elementId string) *neo4j.Node {
	for _, neoNode := range *neoNodes {
		if neoNode.GetElementId() == elementId {
			return &neoNode
		}
	}
	return nil
}
