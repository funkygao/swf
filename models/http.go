package models

import (
	"encoding/json"
)

type HttpApiResponse interface {
	Bytes() []byte
}

type HttpApiRequest interface {
	From([]byte) error
}

type RegisterActivityTypeInput struct {
	ActivityType
}

func (this *RegisterActivityTypeInput) From(payload []byte) error {
	return json.Unmarshal(payload, this)
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
}

func (this *RegisterWorkflowTypeInput) From(payload []byte) error {
	return json.Unmarshal(payload, this)
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
	RunId int64 `json:"run_id"`
}

func (this *StartWorkflowExecutionOutput) From(payload []byte) error {
	return json.Unmarshal(payload, this)
}

func (this *StartWorkflowExecutionOutput) Bytes() []byte {
	b, _ := json.Marshal(this)
	return b
}

type TerminateWorkflowExecutionInput struct {
	RunId      string
	WorkflowId string
	Reason     string
	Details    string
}

type TerminateWorkflowExecutionOutput struct {
	// empty
}

type PollForActivityTaskInput struct {
	Identity string
	ActivityType
}

func (this *PollForActivityTaskInput) From(payload []byte) error {
	return json.Unmarshal(payload, this)
}

type PollForActivityTaskOutput struct {
	// The unique ID of the task.
	ActivityId string

	ActivityType ActivityType

	Input string

	// The ID of the ActivityTaskStarted event recorded in the history.
	StartedEventId int64

	// The opaque string used as a handle on the task. This token is used by workers
	// to communicate progress and response information back to the system about
	// the task.
	TaskToken string

	// The workflow execution that started this activity task.
	WorkflowExecution WorkflowExecution
}

func (this *PollForActivityTaskOutput) Bytes() []byte {
	b, _ := json.Marshal(this)
	return b
}

func (this *PollForActivityTaskOutput) From(payload []byte) error {
	return json.Unmarshal(payload, this)
}

type PollForDecisionTaskInput struct {
	Identity string
	Queue    string
}

func (this *PollForDecisionTaskInput) From(payload []byte) error {
	return json.Unmarshal(payload, this)
}

type PollForDecisionTaskOutput struct {
	Events            []HistoryEvent
	StartedEventId    int64
	TaskToken         string
	WorkflowExecution WorkflowExecution
	WorkflowType      WorkflowType
}

func (this *PollForDecisionTaskOutput) Bytes() []byte {
	b, _ := json.Marshal(this)
	return b
}

func (this *PollForDecisionTaskOutput) From(payload []byte) error {
	return json.Unmarshal(payload, this)
}

type RespondDecisionTaskCompletedInput struct {
	Decisions []Decision
	TaskToken string
}

func (this *RespondDecisionTaskCompletedInput) From(payload []byte) error {
	return json.Unmarshal(payload, this)
}

func (this *RespondDecisionTaskCompletedInput) Reset() {
	this.Decisions = make([]Decision, 0)
	this.TaskToken = ""
}

func (this *RespondDecisionTaskCompletedInput) AddDecision(d Decision) {
	this.Decisions = append(this.Decisions, d)
}

type RespondDecisionTaskCompletedOutput struct {
	// empty
}

func (this *RespondDecisionTaskCompletedOutput) Bytes() []byte {
	b, _ := json.Marshal(this)
	return b
}

func (this *RespondDecisionTaskCompletedOutput) From(payload []byte) error {
	return json.Unmarshal(payload, this)
}

type RespondActivityTaskCompletedInput struct {
	Result    string
	TaskToken string
}

func (this *RespondActivityTaskCompletedInput) From(payload []byte) error {
	return json.Unmarshal(payload, this)
}

type RespondActivityTaskCompletedOutput struct {
	// empty
}

func (this *RespondActivityTaskCompletedOutput) Bytes() []byte {
	b, _ := json.Marshal(this)
	return b
}

func (this *RespondActivityTaskCompletedOutput) From(payload []byte) error {
	return json.Unmarshal(payload, this)
}
