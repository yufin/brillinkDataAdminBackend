package service

import (
	"context"
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"go-admin/app/graph/models"
	"go-admin/app/graph/util"
)

func GetNodeById(ctx context.Context, id string) []neo4j.Node {
	cypher := "MATCH (n {id: $id}) RETURN n;"
	result := models.CypherQuery(ctx, cypher, map[string]any{"id": id})

	nodeArr := make([]neo4j.Node, 0)
	if len(result) == 0 {
		return nodeArr
	}
	n, found := result[0].Get("n")
	if !found {
		return nodeArr
	}
	return append(nodeArr, n.(neo4j.Node))
}

func GetChildrenById(ctx context.Context, id string, expectRel []string) []neo4j.Node {
	expectRelStmt := util.GetRelConstraintStmt(expectRel, "r", true)
	cypher := fmt.Sprintf(
		"MATCH (n{id: $id})-[r]->(m) %s return m;", expectRelStmt)
	result := models.CypherQuery(ctx, cypher, map[string]any{"id": id})
	var resp []neo4j.Node
	for _, child := range result {
		childNode, found := child.Get("m")
		if found {
			resp = append(resp, childNode.(neo4j.Node))
		}
	}
	return resp
}
