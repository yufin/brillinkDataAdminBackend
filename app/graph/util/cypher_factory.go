package util

import (
	"fmt"
	"strings"
)

func GetLabelsConstraintStmt(constraintLabels []string, nodeKey string, expect bool) string {
	if len(constraintLabels) == 0 {
		return ""
	}
	var expectStmt string
	if expect {
		expectStmt = ""
	} else {
		expectStmt = "NOT "
	}
	var sb strings.Builder
	labelsArr := make([]string, 0)
	for _, label := range constraintLabels {
		labelsArr = append(labelsArr, fmt.Sprintf("'%s'", label))
	}
	labelsStmt := strings.Join(labelsArr, ",")
	sb.WriteString(fmt.Sprintf("WHERE any(label IN labels(%s) WHERE %slabel IN [%s]) ", nodeKey, expectStmt, labelsStmt))
	return sb.String()
}

func GetRelConstraintStmt(constraintRel []string, relKey string, expect bool) string {
	if len(constraintRel) == 0 {
		return ""
	}
	var expectStmt string
	if expect {
		expectStmt = ""
	} else {
		expectStmt = "NOT"
	}
	var sb strings.Builder
	relsArray := make([]string, 0)
	for _, rel := range constraintRel {
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
