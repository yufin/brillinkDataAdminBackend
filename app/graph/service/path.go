package service

import (
	"context"
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"go-admin/app/graph/models"
	"math"
)

func ExpandPathFromSource(ctx context.Context, sourceId string, depth int, limit int) []neo4j.Path {
	depth = int(math.Min(float64(depth), 2))
	cypher := fmt.Sprintf(
		"MATCH (startNode {id: $sourceId}) "+
			"MATCH path = (startNode)-[ *%v]-(endNode) "+
			"RETURN path limit $limit", depth)

	param := map[string]any{"sourceId": sourceId, "limit": limit}
	result := models.CypherQuery(ctx, cypher, param)
	var resp []neo4j.Path
	for _, path := range result {
		path, found := path.Get("path")
		if found {
			resp = append(resp, path.(neo4j.Path))
		}
	}
	return resp
}

func GetPathBetween(ctx context.Context, sourceId string, targetId string, filterStmt string) []neo4j.Path {
	cypher := fmt.Sprintf(
		"MATCH (s {id: $sourceId}) "+
			"MATCH (t {id: $targetId}) "+
			"MATCH p = (s)-[r*]->(t) "+
			"%s "+
			"RETURN p ", filterStmt)

	param := map[string]any{"sourceId": sourceId, "targetId": targetId}
	result := models.CypherQuery(ctx, cypher, param)
	var resp []neo4j.Path
	for _, path := range result {
		path, found := path.Get("p")
		if found {
			resp = append(resp, path.(neo4j.Path))
		}
	}
	return resp
}

func GetPathToChildren(ctx context.Context, sourceId string, pageSize int, pageNum int) ([]neo4j.Path, int64) {
	pageNum = int(math.Max(float64(pageNum), 1))
	cypher := "MATCH p=(n {id: $sourceId})-[r]->(m) return p skip $skip limit $limit;"
	cypherCount := "MATCH p=(n {id: $sourceId})-[r]->(m) return count(p) as total;"
	param := map[string]any{"sourceId": sourceId, "skip": pageSize * (pageNum - 1), "limit": pageSize}
	result := models.CypherQuery(ctx, cypher, param)
	total := models.CypherQuery(ctx, cypherCount, map[string]any{"sourceId": sourceId})
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
