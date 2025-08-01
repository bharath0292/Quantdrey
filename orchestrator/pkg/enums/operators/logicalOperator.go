package operators

import "fmt"

type LogicalOperator string

const (
	LogicalOperatorAnd LogicalOperator = "and"
	LogicalOperatorOr  LogicalOperator = "or"
)

func (s LogicalOperator) String() string {
	return string(s)
}

func ParseLogicalOperator(s string) (LogicalOperator, error) {
	switch s {
	case string(LogicalOperatorAnd):
		return LogicalOperatorAnd, nil
	case string(LogicalOperatorOr):
		return LogicalOperatorOr, nil
	default:
		return "", fmt.Errorf("invalid logicaloperator: %s", s)
	}
}
