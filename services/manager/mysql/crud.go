package mysql

import (
	"github.com/funkygao/swf/models"
	"github.com/go-ozzo/ozzo-dbx"
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

func (this *Manager) SaveWorkflowExecution(in *models.StartWorkflowExecutionInput, out *models.StartWorkflowExecutionOutput) error {
	w := &models.WorkflowExecution{}
	w.WorkflowType = in.WorkflowType
	w.RunId = out.RunId
	w.WorkflowId = in.WorkflowId
	return this.db.Model(w).Insert()
}

func (this *Manager) LoadWorkflowExecution(runId int64) (r *models.WorkflowExecution, err error) {
	err = this.db.Select().Where(dbx.HashExp{"run_id": runId}).One(r)
	return
}
