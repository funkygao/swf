package models

type decisionType int

type Decision struct {
}

const (
	CancelTimer decisionType = iota
	CancelWorkflowExecution
	CompleteWorkflowExecution
	ContinueAsNewWorkflowExecution
	FailWorkflowExecution
	RecordMarker
	RequestCancelActivityTask
	RequestCancelExternalWorkflowExecution
	ScheduleActivityTask
	SignalExternalWorkflowExecution
	StartChildWorkflowExecution
	StartTimer
)
