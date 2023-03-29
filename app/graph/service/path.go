package service

import (
	"context"
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"go-admin/app/graph/models"
)

func ExpandPathFromSource(ctx context.Context, sourceId string, depth int, limit int) []neo4j.Path {
	cypher := fmt.Sprintf(
		"MATCH (startNode {id: $sourceId}) "+
			"WITH startNode "+
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
		"MATCH (startNode {id: $sourceId}) "+
			"MATCH (targetNode {id: $targetId}) "+
			"MATCH p = (startNode)-[r *]-(endNode) "+
			"%s "+
			"RETURN p ", filterStmt)

	param := map[string]any{"sourceId": sourceId, "targetId": targetId}
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
