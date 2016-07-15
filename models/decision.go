package models

type decisionType int

type Decision struct {
}

const (
	CancelTimer decisionType = iota

	CancelWorkflowExecution
	CompleteWorkflowExecution // close the workflow execution
	FailWorkflowExecution
	ContinueAsNewWorkflowExecution

	RecordMarker
	RequestCancelActivityTask
	RequestCancelExternalWorkflowExecution
	ScheduleActivityTask
	SignalExternalWorkflowExecution
	StartChildWorkflowExecution
	StartTimer
)
