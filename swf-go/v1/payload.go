package swfapi

import (
	"encoding/json"

	"github.com/funkygao/swf/models"
)

type RegisterActivityTypeInput struct {
	models.ActivityType
	Description string
}

type RegisterActivityTypeOutput struct {
	// empty
}

func (this *RegisterActivityTypeOutput) from(payload []byte) {

}

type RegisterWorkflowTypeInput struct {
	models.WorkflowType
	Description string
}

type RegisterWorkflowTypeOutput struct {
	// empty
}

func (this *RegisterWorkflowTypeOutput) from(payload []byte) {

}

type StartWorkflowExecutionInput struct {
	Input        string
	Queue        string
	WorkflowId   string
	WorkflowType models.WorkflowType
}

type StartWorkflowExecutionOutput struct {
	RunId string `json:"run_id"`
}

func (this *StartWorkflowExecutionOutput) from(payload []byte) {
	json.Unmarshal(payload, this)
}
