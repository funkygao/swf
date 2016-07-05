package models

// union
type HistoryEvent struct {
}

type historyEventType int

const (
	ActivityTaskCancelRequested historyEventType = iota
	ActivityTaskCanceled
	ActivityTaskCompleted
	ActivityTaskFailed
	ActivityTaskScheduled
	ActivityTaskStarted
	ActivityTaskTimedOut
	CancelTimerFailed
	CancelWorkflowExecutionFailed
	ChildWorkflowExecutionCanceled
	DecisionTaskScheduled
	DecisionTaskStarted
	// ...
)
