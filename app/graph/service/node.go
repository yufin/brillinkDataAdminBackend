package service

import (
	"context"
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/pkg/errors"
	"go-admin/app/graph/models"
	"go-admin/app/graph/util"
)

// GetNodeById returns the node By given id.
func GetNodeById(ctx context.Context, id string) ([]neo4j.Node, error) {
	cypher := "MATCH (n {id: $id}) RETURN n;"
	result, err := models.CypherQuery(ctx, cypher, map[string]any{"id": id})
	var nodeArr = make([]neo4j.Node, 0)
	if err != nil {
		return nodeArr, err
	}

	if len(result) == 0 {
		return nodeArr, nil
	}
	n, found := result[0].Get("n")
	if !found {
		return nodeArr, errors.New("key not found")
	}
	nodeArr = append(nodeArr, n.(neo4j.Node))
	return nodeArr, err
}

// GetChildrenById returns the children of a node By given id, expect relationship constrained by expectRel.
func GetChildrenById(ctx context.Context, id string, pageSize int, pageNum int, expectRel []string) ([]neo4j.Node, int64, error) {
	expectRelStmt := util.GetRelConstraintStmt(expectRel, "r", true)
	cypher := fmt.Sprintf(
		"MATCH (n{id: $id})-[r]->(m) %s return m skip $skip limit $limit;", expectRelStmt)
	result, err := models.CypherQuery(ctx, cypher, map[string]any{"id": id, "skip": (pageNum - 1) * pageSize, "limit": pageSize})
	if err != nil {
		return nil, 0, err
	}

	total := func() int64 {
		cypherCount := fmt.Sprintf(
			"MATCH (n{id: $id})-[r]->(m) %s return count(m) as total;", expectRelStmt)
		count, err := CountChildren(ctx, cypherCount, "total", map[string]any{"id": id})
		if err != nil {
			count = 0
		}
		return count
	}()

	var resp []neo4j.Node
	for _, child := range result {
		childNode, found := child.Get("m")
		if found {
			resp = append(resp, childNode.(neo4j.Node))
		}
	}
	return resp, total, nil
}

func CountChildren(ctx context.Context, cypher string, resKey string, params map[string]any) (int64, error) {
	result, err := models.CypherQuery(ctx, cypher, params)
	if err != nil {
		return 0, err
	}
	if len(result) == 0 {
		return 0, errors.New("result is empty")
	}
	count, found := result[0].Get(resKey)
	if !found {
		return 0, errors.New("result key not found")
	}
	return count.(int64), err
}
