package service

import (
	"context"
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"go-admin/app/graph/models"
	"go-admin/app/graph/util"
	"math"
)

// ExpandPathFromSource returns the path from sourceId to the end of the graph (depth constrained by depth)
func ExpandPathFromSource(ctx context.Context, sourceId string, depth int, limit int) []neo4j.Path {
	depth = int(math.Min(float64(depth), 2))
	cypher := fmt.Sprintf(
		"MATCH (startNode {id: $sourceId}) "+
			"MATCH path = (startNode)-[ *%v]-(endNode) "+
			"RETURN path limit $limit", depth)

	param := map[string]any{"sourceId": sourceId, "limit": limit}
	result, err := models.CypherQuery(ctx, cypher, param)
	if err != nil {
		return nil
	}
	var resp []neo4j.Path
	for _, path := range result {
		path, found := path.Get("path")
		if found {
			resp = append(resp, path.(neo4j.Path))
		}
	}
	return resp
}

// GetPathBetween returns the path between two nodes (filtered by filterStmt, which is cypher stmt) By given sourceId and targetId.
func GetPathBetween(ctx context.Context, sourceId string, targetId string, filterStmt string) []neo4j.Path {
	cypher := fmt.Sprintf(
		"MATCH (s {id: $sourceId}) "+
			"MATCH (t {id: $targetId}) "+
			"MATCH p = (s)-[r*]->(t) "+
			"%s "+
			"RETURN p ", filterStmt)

	param := map[string]any{"sourceId": sourceId, "targetId": targetId}
	result, err := models.CypherQuery(ctx, cypher, param)
	if err != nil {
		return nil
	}
	var resp []neo4j.Path
	for _, path := range result {
		path, found := path.Get("p")
		if found {
			resp = append(resp, path.(neo4j.Path))
		}
	}
	return resp
}

// GetPathToChildren returns the paginated(pageSize counts on relationship) path to children of a node By given sourceId.
func GetPathToChildren(ctx context.Context, sourceId string, pageSize int, pageNum int) ([]neo4j.Path, int64) {
	pageNum = int(math.Max(float64(pageNum), 1))
	cypher := "MATCH p=(n {id: $sourceId})-[r]->(m) return p skip $skip limit $limit;"
	cypherCount := "MATCH p=(n {id: $sourceId})-[r]->(m) return count(p) as total;"
	param := map[string]any{"sourceId": sourceId, "skip": pageSize * (pageNum - 1), "limit": pageSize}
	result, err := models.CypherQuery(ctx, cypher, param)
	total, errCount := models.CypherQuery(ctx, cypherCount, map[string]any{"sourceId": sourceId})
	if err != nil || errCount != nil {
		return nil, 0
	}
	respTotal, _ := total[0].Get("total")
	var resp []neo4j.Path
	for _, path := range result {
		path, found := path.Get("p")
		if found {
			resp = append(resp, path.(neo4j.Path))
		}
	}
	return resp, respTotal.(int64)
}

// GetPathFromSourceByIds returns the path from sourceId to targetIds (filtered by expectLabels and expectRels)
func GetPathFromSourceByIds(
	ctx context.Context, sourceId string, targetIds []string, expectLabels []string, expectRels []string) []neo4j.Path {
	expectLabelsStmt := util.GetExpectLabelsConstraintStmt(expectLabels)
	cypher := fmt.Sprintf("MATCH (rootNode {id: $sourceId}) "+
		"MATCH (targetNode %s) where targetNode.id in $targetIds "+
		"MATCH p=(rootNode)-[r *]->(targetNode) "+
		"WHERE all(rel in relationships(p) WHERE type(rel) in $expectRels) "+
		"return p", expectLabelsStmt)
	param := map[string]any{"sourceId": sourceId, "targetIds": targetIds, "expectRels": expectRels}
	result, err := models.CypherQuery(ctx, cypher, param)
	if err != nil {
		return nil
	}
	var resp []neo4j.Path
	for _, path := range result {
		path, found := path.Get("p")
		if found {
			resp = append(resp, path.(neo4j.Path))
		}
	}
	return resp
}
