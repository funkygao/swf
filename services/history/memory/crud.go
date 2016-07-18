package memory

import (
	"github.com/funkygao/swf/models"
)

func (this *History) SaveWorkflowExecution(*models.StartWorkflowExecutionInput, *models.StartWorkflowExecutionOutput) error {
	return nil
}

func (this *History) LoadWorkflowExecution(runId int64) {

}

func (this *History) RegisterWorkflowType() {

}

func (this *History) RegisterActivityType() {

}

func (this *History) LoadWorkflowTypes() {

}

func (this *History) LoadActivityTypes() {

}
