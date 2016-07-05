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

func (this *RegisterActivityTypeOutput) Bytes() []byte {
	b, _ := json.Marshal(this)
	return b
}

func (this *RegisterActivityTypeOutput) From(payload []byte) error {
	return json.Unmarshal(payload, this)
}

type RegisterWorkflowTypeInput struct {
	WorkflowType
	Description string
}

type RegisterWorkflowTypeOutput struct {
	// empty
}

func (this *RegisterWorkflowTypeOutput) Bytes() []byte {
	b, _ := json.Marshal(this)
	return b
}

func (this *RegisterWorkflowTypeOutput) From(payload []byte) error {
	return json.Unmarshal(payload, this)
}

type StartWorkflowExecutionInput struct {
	Input        string
	Queue        string
	WorkflowId   string
	WorkflowType WorkflowType
}

func (this *StartWorkflowExecutionInput) From(payload []byte) error {
	return json.Unmarshal(payload, this)
}

type StartWorkflowExecutionOutput struct {
	RunId string `json:"run_id"`
}

func (this *StartWorkflowExecutionOutput) From(payload []byte) error {
	return json.Unmarshal(payload, this)
}

func (this *StartWorkflowExecutionOutput) Bytes() []byte {
	b, _ := json.Marshal(this)
	return b
}
