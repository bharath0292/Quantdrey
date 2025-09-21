package strategyenum

type StrategyStatus string

const (
	StatusCreated   StrategyStatus = "CREATED"
	StatusScheduled StrategyStatus = "SCHEDULED"
	StatusRunning   StrategyStatus = "RUNNING"
	StatusStopped   StrategyStatus = "STOPPED"
	StatusCompleted StrategyStatus = "COMPLETED"
	StatusFailed    StrategyStatus = "FAILED"
)

type StopReason string

const (
	StopEndTime   StopReason = "END_TIME"
	StopExitLogic StopReason = "EXIT_LOGIC"
	StopMaxProfit StopReason = "MAX_PROFIT"
	StopMaxLoss   StopReason = "MAX_LOSS"
	StopManual    StopReason = "MANUAL"
	StopError     StopReason = "ERROR"
)
