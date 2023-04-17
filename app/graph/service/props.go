package service

import (
	"context"
	"fmt"
	"go-admin/app/graph/models"
)

// GetCompanyTitleAutoComplete Return a list of company title that match the keyword(keyWord is Company.title or Company.titleShort)
func GetCompanyTitleAutoComplete(ctx context.Context, keyWord string, pageSize int, pageNum int) ([]any, error) {
	cypher := fmt.Sprintf("match (:Tag)-[:ATTACH_TO]-(n:Company) "+
		"where n.title =~ '(?i).*%s.*' "+
		"with distinct n "+
		"with {title: n.title, id: n.id} as res skip $skip limit $limit "+
		"with collect(res) as propList "+
		"return propList", keyWord)

	result, err := models.CypherQuery(ctx, cypher, map[string]any{"skip": pageSize * (pageNum - 1), "limit": pageSize})
	if err != nil {
		return nil, err
	}
	r, _ := result[0].Get("propList")
	return r.([]any), nil
}

// CountCompanyTitleAutoComplete Return the number of company title that match the keyword(keyWord is Company.title or Company.titleShort)
func CountCompanyTitleAutoComplete(ctx context.Context, keyWord string) (int64, error) {
	cypher := fmt.Sprintf("match (:tag)-[:ATTACH_TO]-(n:Company) "+
		"where n.title =~ '(?i).*%s.*' "+
		"with DISTINCT n "+
		"return count(n) as total", keyWord)
	result, err := models.CypherQuery(ctx, cypher, map[string]any{})
	if err != nil {
		return -1, err
	}
	total, _ := result[0].Get("total")
	return total.(int64), nil
}

// GetLabelTitleAutoComplete Return a list of label title that match the keyword(keyWord is Label.title)
func GetLabelTitleAutoComplete(ctx context.Context, keyWord string, pageSize int, pageNum int) ([]any, error) {
	cypher := fmt.Sprintf("match ()-[:CLASSIFY_OF]->(n:Tag) "+
		"where n.title =~ '(?i).*%s.*' "+
		"with DISTINCT n "+
		"with {title: n.title, id: n.id} as res skip $skip limit $limit "+
		"with collect(res) as propList "+
		"return propList", keyWord)
	result, err := models.CypherQuery(ctx, cypher, map[string]any{"skip": pageSize * (pageNum - 1), "limit": pageSize})
	if err != nil {
		return nil, err
	}
	r, _ := result[0].Get("propList")
	return r.([]any), nil
}

// CountLabelTitleAutoComplete Return the number of label title that match the keyword(keyWord is Label.title)
func CountLabelTitleAutoComplete(ctx context.Context, keyWord string) (int64, error) {
	cypher := fmt.Sprintf("match ()-[:CLASSIFY_OF]->(n:Tag) "+
		"where n.title =~ '(?i).*%s.*' "+
		"with DISTINCT n "+
		"return count(n) as total", keyWord)
	result, err := models.CypherQuery(ctx, cypher, map[string]any{})
	if err != nil {
		return -1, err
	}
	total, _ := result[0].Get("total")
	return total.(int64), nil
}
