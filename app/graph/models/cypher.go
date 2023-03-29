package models

import (
	"context"
	log "github.com/go-admin-team/go-admin-core/logger"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"go-admin/app/graph/util"
	"go-admin/common/database"
)

func CypherQuery(ctx context.Context, cypher string, params map[string]any) []neo4j.Record {
	if ctx == nil {
		ctx = context.Background()
	}
	db := database.Neo4jDriverP
	session := (*db).NewSession(ctx, neo4j.SessionConfig{})
	util.PanicOnClosureError(ctx, session)
	result, err := session.Run(ctx, cypher, params)

	var output []neo4j.Record
	for result.Next(ctx) {
		record := result.Record()
		output = append(output, *record)
	}
	if err = result.Err(); err != nil {
		log.Errorf("Cypher exc Failed. %v Cypher:%v,", err, cypher)
	}
	return output
}
