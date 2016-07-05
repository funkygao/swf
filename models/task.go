package models

type ActivityTask struct {
	activityId        string // 1-256
	activityType      ActivityType
	input             string
	startedEventId    int64
	taskToken         string
	workflowExecution WorkflowExecution
}

type DecisionTask struct {
}
