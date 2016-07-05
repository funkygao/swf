package models

import (
	"encoding/json"
)

type RegisterActivityTypeInput struct {
	ActivityType
	Description string
}

type RegisterActivityTypeOutput struct {
	// empty
}

func (this *RegisterActivityTypeOutput) From(payload []byte) {

}

type RegisterWorkflowTypeInput struct {
	WorkflowType
	Description string
}

type RegisterWorkflowTypeOutput struct {
	// empty
}

func (this *RegisterWorkflowTypeOutput) From(payload []byte) {

}

type StartWorkflowExecutionInput struct {
	Input        string
	Queue        string
	WorkflowId   string
	WorkflowType WorkflowType
}

func (this *StartWorkflowExecutionInput) From(payload []byte) {
	json.Unmarshal(payload, this)
}

type StartWorkflowExecutionOutput struct {
	RunId string `json:"run_id"`
}

func (this *StartWorkflowExecutionOutput) From(payload []byte) {
	json.Unmarshal(payload, this)
}
