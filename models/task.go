package models

// useless
type ActivityTask struct {
	ActivityId        string // 1-256
	ActivityType      ActivityType
	Input             string
	StartedEventId    int64
	TaskToken         string
	WorkflowExecution WorkflowExecution
}

// useless
type DecisionTask struct {
	TaskToken string
}
