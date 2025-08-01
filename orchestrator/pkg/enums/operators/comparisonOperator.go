package operators

import "fmt"

type ComparisonOperator string

const (
	ComparisonOperatorEqual          ComparisonOperator = "eq"
	ComparisonOperatorNotEqual       ComparisonOperator = "neq"
	ComparisonOperatorGreater        ComparisonOperator = "gt"
	ComparisonOperatorGreaterOrEqual ComparisonOperator = "gteq"
	ComparisonOperatorLess           ComparisonOperator = "lt"
	ComparisonOperatorLessOrEqual    ComparisonOperator = "lteq"
)

func (s ComparisonOperator) String() string {
	return string(s)
}

func ParseComparisonOperator(s string) (ComparisonOperator, error) {
	switch s {
	case string(ComparisonOperatorEqual):
		return ComparisonOperatorEqual, nil
	case string(ComparisonOperatorNotEqual):
		return ComparisonOperatorNotEqual, nil
	case string(ComparisonOperatorGreater):
		return ComparisonOperatorGreater, nil
	case string(ComparisonOperatorGreaterOrEqual):
		return ComparisonOperatorGreaterOrEqual, nil
	case string(ComparisonOperatorLess):
		return ComparisonOperatorLess, nil
	case string(ComparisonOperatorLessOrEqual):
		return ComparisonOperatorLessOrEqual, nil
	default:
		return "", fmt.Errorf("invalid comparisonOperator: %s", s)
	}
}

func (s ComparisonOperator) ComparisonOperatorCheck(lhs float64, rhs float64) bool {
	switch s {
	case ComparisonOperatorEqual:
		return lhs == rhs
	case ComparisonOperatorNotEqual:
		return lhs != rhs
	case ComparisonOperatorGreater:
		return lhs > rhs
	case ComparisonOperatorGreaterOrEqual:
		return lhs >= rhs
	case ComparisonOperatorLess:
		return lhs < rhs
	case ComparisonOperatorLessOrEqual:
		return lhs <= rhs
	}
	return false
}
