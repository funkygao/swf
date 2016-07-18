package manager

import (
	"github.com/funkygao/swf/models"
	"github.com/funkygao/swf/services"
)

type Service interface {
	services.Service

	RegisterActivityType(t *models.ActivityType) (err error)
	RegisterWorkflowType(t *models.WorkflowType) (err error)

	SaveWorkflowExecution(*models.StartWorkflowExecutionInput, *models.StartWorkflowExecutionOutput) error
	LoadWorkflowExecution(runId int64) (r *models.WorkflowExecution, err error)

	GetTypeByTaskToken(token string) (*models.WorkflowType, *models.ActivityType)

	GetActivityType(t *models.ActivityType) (r *models.ActivityType, err error)
	GetWorkflowType(t *models.WorkflowType) (r *models.WorkflowType, err error)
}

var Default Service
