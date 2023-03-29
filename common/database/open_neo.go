package database

import (
	"context"
	"fmt"
	log "github.com/go-admin-team/go-admin-core/logger"
	"github.com/go-admin-team/go-admin-core/sdk/pkg"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	extConfig "go-admin/config"
)

var Neo4jDriverP *neo4j.DriverWithContext = nil

//func init() {
//	initNeo4jDriver()
//}

func initNeo4jDriver() {
	log.Info(pkg.Green("Neo4J Driver initializing...."))
	neo4jDriver, err := neo4j.NewDriverWithContext(
		extConfig.ExtConfig.Graph.Neo4j.Uri,
		neo4j.BasicAuth(
			extConfig.ExtConfig.Graph.Neo4j.Username,
			extConfig.ExtConfig.Graph.Neo4j.Password,
			""),
	)
	if err != nil {
		log.Error(pkg.Red(fmt.Sprintf("Neo4j Driver Failed to initialized. %v", err)))
		panic(err)
	} else {
		ctx := context.Background()
		connErr := neo4jDriver.VerifyConnectivity(ctx)
		if connErr != nil {
			log.Errorf(pkg.Red("Neo4j Driver UNABLE to connect. " + connErr.Error()))
			//err := neo4jDriver.Close(context.Background())
			//if err != nil {
			//	log.Errorf("Neo4j Driver UNABLE to close. %v", err)
			//}
			panic(connErr)
		} else {
			log.Info(pkg.Green("Neo4j Driver initialized."))
		}
	}
	Neo4jDriverP = &neo4jDriver
}

func GetNeo4jDriver() *neo4j.DriverWithContext {
	if Neo4jDriverP == nil {
		initNeo4jDriver()
	}
	return Neo4jDriverP
}

//func CypherQuery(ctx context.Context, cypher string, params map[string]any) []neo4j.Record {
//	if ctx == nil {
//		ctx = context.Background()
//	}
//	db := GetNeo4jDriver()
//	session := (*db).NewSession(ctx, neo4j.SessionConfig{})
//	util.PanicOnClosureError(ctx, session)
//	result, err := session.Run(ctx, cypher, params)
//
//	var output []neo4j.Record
//	for result.Next(ctx) {
//		record := result.Record()
//		output = append(output, *record)
//	}
//	if err = result.Err(); err != nil {
//		log.Errorf("Cypher exc Failed. %v Cypher:%v,", err, cypher)
//	}
//	return output
//}
