package models

import (
	"encoding/json"
)

type decisionType int

type Decision struct {
	DecisionType string
}

func (this *Decision) Bytes() []byte {
	b, _ := json.Marshal(this)
	return b
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
