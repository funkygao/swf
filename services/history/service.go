package history

import (
	"github.com/funkygao/swf/models"
	"github.com/funkygao/swf/services"
)

type Service interface {
	services.Service

	SaveWorkflowExecution(*models.StartWorkflowExecutionInput, *models.StartWorkflowExecutionOutput) error
	LoadWorkflowExecution(runId int64)
}

var Default Service
