package models

type WorkflowType struct {
	Name    string // 1-256 in len
	Version string // 1-64 in len
}

type WorkflowExecution struct {
	WorkflowId string // 1-256 in len
	RunId      string // 1-64 in len
}
