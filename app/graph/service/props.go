package service

import (
	"context"
	"fmt"
	"go-admin/app/graph/models"
)

func GetCompanyTitleAutoComplete(ctx context.Context, keyWord string, pageSize int, pageNum int) []any {
	cypher := fmt.Sprintf("match (:Label)-[:ATTACH_TO]-(n:Company) "+
		"where n.title =~ '(?i).*%s.*'or n.titleShort =~ '(?i).*%s.*' "+
		"with DISTINCT n, CASE WHEN n.title IS NOT NULL THEN n.title ELSE n.titleShort END AS title "+
		"with {title: title, id: n.id} as res skip $skip limit $limit "+
		"with collect(res) as propList "+
		"return propList", keyWord, keyWord)

	result := models.CypherQuery(ctx, cypher, map[string]any{"skip": pageSize * (pageNum - 1), "limit": pageSize})
	r, _ := result[0].Get("propList")
	return r.([]any)
}

func CountCompanyTitleAutoComplete(ctx context.Context, keyWord string) int64 {
	cypher := fmt.Sprintf("match (:Label)-[:ATTACH_TO]-(n:Company) "+
		"where n.title =~ '(?i).*%s.*'or n.titleShort =~ '(?i).*%s.*' "+
		"with DISTINCT n "+
		"return count(n) as total", keyWord, keyWord)
	result := models.CypherQuery(ctx, cypher, map[string]any{})
	total, _ := result[0].Get("total")
	return total.(int64)
}

func GetLabelTitleAutoComplete(ctx context.Context, keyWord string, pageSize int, pageNum int) []any {
	cypher := fmt.Sprintf("match ()-[:CLASSIFY_OF]->(n:Label) "+
		"where n.title =~ '(?i).*%s.*' "+
		"with DISTINCT n "+
		"with {title: n.title, id: n.id} as res skip $skip limit $limit "+
		"with collect(res) as propList "+
		"return propList", keyWord)
	result := models.CypherQuery(ctx, cypher, map[string]any{"skip": pageSize * (pageNum - 1), "limit": pageSize})
	r, _ := result[0].Get("propList")
	return r.([]any)
}

func CountLabelTitleAutoComplete(ctx context.Context, keyWord string) int64 {
	cypher := fmt.Sprintf("match ()-[:CLASSIFY_OF]->(n:Label) "+
		"where n.title =~ '(?i).*%s.*' "+
		"with DISTINCT n "+
		"return count(n) as total", keyWord)
	result := models.CypherQuery(ctx, cypher, map[string]any{})
	total, _ := result[0].Get("total")
	return total.(int64)
}
