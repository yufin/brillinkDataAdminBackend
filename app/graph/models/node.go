package models

import "github.com/neo4j/neo4j-go-driver/v5/neo4j"

type Node struct {
	Props     map[string]any
	Labels    []string
	ElementId string
}

func TransformNode(neoNode neo4j.Node) Node {
	return Node{
		Props:     neoNode.GetProperties(),
		Labels:    neoNode.Labels,
		ElementId: neoNode.GetElementId(),
	}
}

type NodeInf interface {
	GetElementId() string
	GetId() int64
	GetProperties() map[string]any
}
