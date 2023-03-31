package util

import (
	"fmt"
	"strings"
)

func GetExpectLabelsConstraintStmt(labels []string) string {
	if len(labels) == 0 {
		return ""
	}
	var sb strings.Builder
	for _, label := range labels {
		sb.WriteString(":" + label)
	}
	return sb.String()
}

func GetUnexpectRelConstraintStmt(expectRel []string, relKey string, expect bool) string {
	if len(expectRel) == 0 {
		return ""
	}
	var expectStmt string
	if expect {
		expectStmt = ""
	} else {
		expectStmt = "NOT"
	}
	var sb strings.Builder
	relsArray := make([]string, 0, len(expectRel))
	for _, rel := range expectRel {
		relsArray = append(relsArray, fmt.Sprintf("'%s'", rel))
	}
	relsStmt := strings.Join(relsArray, ", ")
	sb.WriteString(fmt.Sprintf(" where %s type(%s) in [%s] ", expectStmt, relKey, relsStmt))
	return sb.String()
}

func GetPropsStmt(kwargs map[string]any, fixPropsExpr string) string {
	var sb strings.Builder
	if len(kwargs) > 0 {
		props := make([]string, 0, len(kwargs))
		for k, _ := range kwargs {
			props = append(props, fmt.Sprintf("%s: $%s", k, k))
		}
		propsStr := strings.Join(props, ", ")
		if fixPropsExpr != "" {
			propsStr = ", " + propsStr
		}
		sb.WriteString(fmt.Sprintf(" {%s%s}", fixPropsExpr, propsStr))
	}
	return sb.String()
}
