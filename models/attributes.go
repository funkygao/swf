package models

type ActivityTaskCompletedEventAttributes struct {
	Result           string
	ScheduledEventId int64
	StartedEventId   int64
}

type ActivityTaskScheduledEventAttributes struct {
	ActivityId                   string
	ActivityType                 ActivityType
	DecisionTaskCompletedEventId int64
	Input                        string
}

type ActivityTaskStartedEventAttributes struct {
	Identity         string
	ScheduledEventId int64
}

type ActivityTaskTimedOutEventAttributes struct {
	Details     string
	TimeoutType string
}

type DecisionTaskCompletedEventAttributes struct {
	ExecutionContext string
	ScheduledEventId int64
	StartedEventId   int64
}

type DecisionTaskScheduledEventAttributes struct {
	StartToCloseTimeout string
}

type DecisionTaskStartedEventAttributes struct {
	Identity         string
	ScheduledEventId int64
}

type DecisionTaskTimedOutEventAttributes struct {
	ScheduledEventId int64
	StartedEventId   int64
	TimeoutType      string
}

type WorkflowExecutionStartedEventAttributes struct {
	Input                        string
	ExecutionStartToCloseTimeout string
	WorkflowType                 WorkflowType
	TaskStartToCloseTimeout      string
}

type WorkflowExecutionCompletedEventAttributes struct {
	InitiatedEventId  int64
	Result            string
	WorkflowExecution WorkflowExecution
	WorkflowType      WorkflowType
}

type WorkflowExecutionFailedEventAttributes struct {
}

type WorkflowExecutionCancelRequestedEventAttributes struct {
}

type WorkflowExecutionTimedOutEventAttributes struct {
}

type WorkflowExecutionTerminatedEventAttributes struct {
}
