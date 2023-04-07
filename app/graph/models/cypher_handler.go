package models

import (
	"context"
	log "github.com/go-admin-team/go-admin-core/logger"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"go-admin/common/database"
)

func CypherQuery(ctx context.Context, cypher string, params map[string]any) ([]neo4j.Record, error) {
	//if ctx == nil {
	//	ctx = context.Background()
	//}
	db := database.Neo4jDriverP
	session := (*db).NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})

	defer func(session neo4j.SessionWithContext, ctx context.Context) {
		err := session.Close(ctx)
		if err != nil {
			log.Errorf("Error closing Neo4j session: %v", err)
		}
	}(session, ctx)

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
