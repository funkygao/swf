package models

import (
	"encoding/json"
)

type Decision struct {
	DecisionType string

	CompleteWorkflowExecutionDecisionAttributes *CompleteWorkflowExecutionDecisionAttributes
	RequestCancelActivityTaskDecisionAttributes *RequestCancelActivityTaskDecisionAttributes
	ScheduleActivityTaskDecisionAttributes      *ScheduleActivityTaskDecisionAttributes
}

type ScheduleActivityTaskDecisionAttributes struct {
	ActivityId   string
	ActivityType ActivityType
	Input        string
}

type RequestCancelActivityTaskDecisionAttributes struct {
	ActivityId string
}

type CompleteWorkflowExecutionDecisionAttributes struct {
	Result string
}

func (this *Decision) Bytes() []byte {
	b, _ := json.Marshal(this)
	return b
}

const (
	// @enum DecisionType
	DecisionTypeScheduleActivityTask = "ScheduleActivityTask"
	// @enum DecisionType
	DecisionTypeRequestCancelActivityTask = "RequestCancelActivityTask"
	// @enum DecisionType
	DecisionTypeCompleteWorkflowExecution = "CompleteWorkflowExecution"
	// @enum DecisionType
	DecisionTypeFailWorkflowExecution = "FailWorkflowExecution"
	// @enum DecisionType
	DecisionTypeCancelWorkflowExecution = "CancelWorkflowExecution"
	// @enum DecisionType
	DecisionTypeContinueAsNewWorkflowExecution = "ContinueAsNewWorkflowExecution"
	// @enum DecisionType
	DecisionTypeRecordMarker = "RecordMarker"
	// @enum DecisionType
	DecisionTypeStartTimer = "StartTimer"
	// @enum DecisionType
	DecisionTypeCancelTimer = "CancelTimer"
	// @enum DecisionType
	DecisionTypeSignalExternalWorkflowExecution = "SignalExternalWorkflowExecution"
	// @enum DecisionType
	DecisionTypeRequestCancelExternalWorkflowExecution = "RequestCancelExternalWorkflowExecution"
	// @enum DecisionType
	DecisionTypeStartChildWorkflowExecution = "StartChildWorkflowExecution"
	// @enum DecisionType
	DecisionTypeScheduleLambdaFunction = "ScheduleLambdaFunction"
)
