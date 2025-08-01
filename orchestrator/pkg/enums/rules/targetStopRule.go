package rules

import "fmt"

type TargetStopRule string

const (
	TargetStopRulePercentage TargetStopRule = "percentage"
	TargetStopRuleAbsolute   TargetStopRule = "absolute"
)

func (s TargetStopRule) String() string {
	return string(s)
}

func ParseTargetStopRule(s string) (TargetStopRule, error) {
	switch s {
	case string(TargetStopRulePercentage):
		return TargetStopRulePercentage, nil
	case string(TargetStopRuleAbsolute):
		return TargetStopRuleAbsolute, nil
	default:
		return "", fmt.Errorf("invalid targetStopType: %s", s)
	}
}
