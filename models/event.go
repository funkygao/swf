package models

import (
	"encoding/json"
	"time"
)

// An event represents a discrete change in your workflow execution's
// state, such as a new activity being scheduled or a running activity being completed.
type HistoryEvent struct {
	EventId   int64
	EventTime time.Time
	EventType string

	ActivityTaskCompletedEventAttributes *ActivityTaskCompletedEventAttributes
	ActivityTaskScheduledEventAttributes *ActivityTaskScheduledEventAttributes
	ActivityTaskStartedEventAttributes   *ActivityTaskStartedEventAttributes
	ActivityTaskTimedOutEventAttributes  *ActivityTaskTimedOutEventAttributes

	DecisionTaskCompletedEventAttributes *DecisionTaskCompletedEventAttributes
	DecisionTaskScheduledEventAttributes *DecisionTaskScheduledEventAttributes
	DecisionTaskStartedEventAttributes   *DecisionTaskStartedEventAttributes
	DecisionTaskTimedOutEventAttributes  *DecisionTaskTimedOutEventAttributes

	WorkflowExecutionStartedEventAttributes         *WorkflowExecutionStartedEventAttributes
	WorkflowExecutionCompletedEventAttributes       *WorkflowExecutionCompletedEventAttributes
	WorkflowExecutionFailedEventAttributes          *WorkflowExecutionFailedEventAttributes
	WorkflowExecutionCancelRequestedEventAttributes *WorkflowExecutionCancelRequestedEventAttributes
	WorkflowExecutionTimedOutEventAttributes        *WorkflowExecutionTimedOutEventAttributes
	WorkflowExecutionTerminatedEventAttributes      *WorkflowExecutionTerminatedEventAttributes
}

type HistoryEvents []HistoryEvent

func (this *HistoryEvents) AppendEvent(evt ...HistoryEvent) *HistoryEvents {
	for _, e := range evt {
		*this = append(*this, e)
	}
	return this
}

func (this *HistoryEvents) Bytes() []byte {
	b, _ := json.Marshal(this)
	return b
}

func (this *HistoryEvents) From(b []byte) {
	json.Unmarshal(b, this)
}

func (this *HistoryEvents) NextEventId() int64 {
	return this.Last().EventId + 1
}

func (this *HistoryEvents) First() HistoryEvent {
	return (*this)[0]
}

func (this *HistoryEvents) Last() HistoryEvent {
	return (*this)[len(*this)-1]
}

func NewEvent(id int64, t time.Time, typ string) *HistoryEvent {
	return &HistoryEvent{
		EventId:   id,
		EventTime: t,
		EventType: typ,
	}
}

func (this *HistoryEvent) Bytes() []byte {
	b, _ := json.Marshal(this)
	return b
}

func (this *HistoryEvent) From(payload []byte) error {
	return json.Unmarshal(payload, this)
}

const (
	// @enum EventType
	EventTypeWorkflowExecutionStarted = "WorkflowExecutionStarted"
	// @enum EventType
	EventTypeWorkflowExecutionCancelRequested = "WorkflowExecutionCancelRequested"
	// @enum EventType
	EventTypeWorkflowExecutionCompleted = "WorkflowExecutionCompleted"
	// @enum EventType
	EventTypeCompleteWorkflowExecutionFailed = "CompleteWorkflowExecutionFailed"
	// @enum EventType
	EventTypeWorkflowExecutionFailed = "WorkflowExecutionFailed"
	// @enum EventType
	EventTypeFailWorkflowExecutionFailed = "FailWorkflowExecutionFailed"
	// @enum EventType
	EventTypeWorkflowExecutionTimedOut = "WorkflowExecutionTimedOut"
	// @enum EventType
	EventTypeWorkflowExecutionCanceled = "WorkflowExecutionCanceled"
	// @enum EventType
	EventTypeCancelWorkflowExecutionFailed = "CancelWorkflowExecutionFailed"
	// @enum EventType
	EventTypeWorkflowExecutionContinuedAsNew = "WorkflowExecutionContinuedAsNew"
	// @enum EventType
	EventTypeContinueAsNewWorkflowExecutionFailed = "ContinueAsNewWorkflowExecutionFailed"
	// @enum EventType
	EventTypeWorkflowExecutionTerminated = "WorkflowExecutionTerminated"
	// @enum EventType
	EventTypeDecisionTaskScheduled = "DecisionTaskScheduled"
	// @enum EventType
	EventTypeDecisionTaskStarted = "DecisionTaskStarted"
	// @enum EventType
	EventTypeDecisionTaskCompleted = "DecisionTaskCompleted"
	// @enum EventType
	EventTypeDecisionTaskTimedOut = "DecisionTaskTimedOut"
	// @enum EventType
	EventTypeActivityTaskScheduled = "ActivityTaskScheduled"
	// @enum EventType
	EventTypeScheduleActivityTaskFailed = "ScheduleActivityTaskFailed"
	// @enum EventType
	EventTypeActivityTaskStarted = "ActivityTaskStarted"
	// @enum EventType
	EventTypeActivityTaskCompleted = "ActivityTaskCompleted"
	// @enum EventType
	EventTypeActivityTaskFailed = "ActivityTaskFailed"
	// @enum EventType
	EventTypeActivityTaskTimedOut = "ActivityTaskTimedOut"
	// @enum EventType
	EventTypeActivityTaskCanceled = "ActivityTaskCanceled"
	// @enum EventType
	EventTypeActivityTaskCancelRequested = "ActivityTaskCancelRequested"
	// @enum EventType
	EventTypeRequestCancelActivityTaskFailed = "RequestCancelActivityTaskFailed"
	// @enum EventType
	EventTypeWorkflowExecutionSignaled = "WorkflowExecutionSignaled"
	// @enum EventType
	EventTypeMarkerRecorded = "MarkerRecorded"
	// @enum EventType
	EventTypeRecordMarkerFailed = "RecordMarkerFailed"
	// @enum EventType
	EventTypeTimerStarted = "TimerStarted"
	// @enum EventType
	EventTypeStartTimerFailed = "StartTimerFailed"
	// @enum EventType
	EventTypeTimerFired = "TimerFired"
	// @enum EventType
	EventTypeTimerCanceled = "TimerCanceled"
	// @enum EventType
	EventTypeCancelTimerFailed = "CancelTimerFailed"
	// @enum EventType
	EventTypeStartChildWorkflowExecutionInitiated = "StartChildWorkflowExecutionInitiated"
	// @enum EventType
	EventTypeStartChildWorkflowExecutionFailed = "StartChildWorkflowExecutionFailed"
	// @enum EventType
	EventTypeChildWorkflowExecutionStarted = "ChildWorkflowExecutionStarted"
	// @enum EventType
	EventTypeChildWorkflowExecutionCompleted = "ChildWorkflowExecutionCompleted"
	// @enum EventType
	EventTypeChildWorkflowExecutionFailed = "ChildWorkflowExecutionFailed"
	// @enum EventType
	EventTypeChildWorkflowExecutionTimedOut = "ChildWorkflowExecutionTimedOut"
	// @enum EventType
	EventTypeChildWorkflowExecutionCanceled = "ChildWorkflowExecutionCanceled"
	// @enum EventType
	EventTypeChildWorkflowExecutionTerminated = "ChildWorkflowExecutionTerminated"
	// @enum EventType
	EventTypeSignalExternalWorkflowExecutionInitiated = "SignalExternalWorkflowExecutionInitiated"
	// @enum EventType
	EventTypeSignalExternalWorkflowExecutionFailed = "SignalExternalWorkflowExecutionFailed"
	// @enum EventType
	EventTypeExternalWorkflowExecutionSignaled = "ExternalWorkflowExecutionSignaled"
	// @enum EventType
	EventTypeRequestCancelExternalWorkflowExecutionInitiated = "RequestCancelExternalWorkflowExecutionInitiated"
	// @enum EventType
	EventTypeRequestCancelExternalWorkflowExecutionFailed = "RequestCancelExternalWorkflowExecutionFailed"
	// @enum EventType
	EventTypeExternalWorkflowExecutionCancelRequested = "ExternalWorkflowExecutionCancelRequested"
	// @enum EventType
	EventTypeLambdaFunctionScheduled = "LambdaFunctionScheduled"
	// @enum EventType
	EventTypeLambdaFunctionStarted = "LambdaFunctionStarted"
	// @enum EventType
	EventTypeLambdaFunctionCompleted = "LambdaFunctionCompleted"
	// @enum EventType
	EventTypeLambdaFunctionFailed = "LambdaFunctionFailed"
	// @enum EventType
	EventTypeLambdaFunctionTimedOut = "LambdaFunctionTimedOut"
	// @enum EventType
	EventTypeScheduleLambdaFunctionFailed = "ScheduleLambdaFunctionFailed"
	// @enum EventType
	EventTypeStartLambdaFunctionFailed = "StartLambdaFunctionFailed"
)
