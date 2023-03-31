package models

import (
	"context"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"go-admin/app/graph/util"
	"go-admin/common/database"
)

func CypherQuery(ctx context.Context, cypher string, params map[string]any) ([]neo4j.Record, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	db := database.Neo4jDriverP
	session := (*db).NewSession(ctx, neo4j.SessionConfig{})
	util.PanicOnClosureError(ctx, session)
	result, err := session.Run(ctx, cypher, params)
	if err != nil {
		return nil, err
	}

	var output []neo4j.Record
	for result.Next(ctx) {
		record := result.Record()
		output = append(output, *record)
	}
	return output, err
}
