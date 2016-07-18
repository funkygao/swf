package mysql

import (
	"github.com/funkygao/swf/models"
)

func (this *Manager) RegisterWorkflowType(t *models.WorkflowType) (err error) {
	return this.db.Model(t).Insert()
}

func (this *Manager) RegisterActivityType(t *models.ActivityType) (err error) {
	return this.db.Model(t).Insert()
}

func (this *Manager) GetActivityType(t *models.ActivityType) (r *models.ActivityType, err error) {
	return
}

func (this *Manager) GetWorkflowType(t *models.WorkflowType) (r *models.WorkflowType, err error) {
	return
}
