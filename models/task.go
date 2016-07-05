package structs

type WorkflowExecution struct {
	runId, workflowId string
}

type ActivityType struct {
	name, version string
}

type ActivityTask struct {
	activityId        string // 1-256
	activityType      ActivityType
	input             string
	startedEventId    int64
	taskToken         string
	workflowExecution WorkflowExecution
}

type DecistionTask struct {
}

type Decision struct {
}

type HistoryEvent struct {
}

type WorkflowExecution struct {
	workflowId string // 1-256 in len
	runId      string // 1-64 in len
}

type WorkflowType struct {
	name    string // 1-256 in len
	version string // 1-64 in len
}

type decisionType int

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
