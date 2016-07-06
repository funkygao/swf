package manager

import (
	"github.com/funkygao/swf/models"
	"github.com/funkygao/swf/services"
)

type Service interface {
	services.Service

	RegisterWorkflowType()
	RegisterActivityType()

	GetWorkflowType(models.WorkflowType)
	GetActivityType(models.ActivityType)
}

var Default Service
