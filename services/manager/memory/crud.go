package memory

import (
	"github.com/funkygao/swf/models"
	"github.com/funkygao/swf/services/manager"
	"github.com/hashicorp/go-memdb"
)

func (this *Manager) RegisterWorkflowType(t *models.WorkflowType) (err error) {
	this.update(func(txn *memdb.Txn) error {
		if err = txn.Insert("WorkflowType", t); err != nil {
			return err
		}
		return nil
	})

	return
}

func (this *Manager) RegisterActivityType(t *models.ActivityType) (err error) {
	this.update(func(txn *memdb.Txn) error {
		if err = txn.Insert("ActivityType", t); err != nil {
			return err
		}
		return nil
	})

	return
}

func (this *Manager) GetWorkflowType(t *models.WorkflowType) (r *models.WorkflowType, err error) {
	this.view(func(txn *memdb.Txn) error {
		var raw interface{}
		raw, err = txn.First("WorkflowType", "id", t.Name, t.Version)
		if raw == nil {
			return manager.ErrNotFound
		}
		if err == nil {
			r = raw.(*models.WorkflowType)
		}
		return err
	})

	return
}

func (this *Manager) GetActivityType(t *models.ActivityType) (r *models.ActivityType, err error) {
	this.view(func(txn *memdb.Txn) error {
		var raw interface{}
		raw, err = txn.First("ActivityType", "id", t.Name, t.Version)
		if raw == nil {
			return manager.ErrNotFound
		}
		if err == nil {
			r = raw.(*models.ActivityType)
		}
		return err
	})

	return
}

func (this *Manager) update(fn func(txn *memdb.Txn) error) error {
	txn := this.db.Txn(true)
	if err := fn(txn); err != nil {
		txn.Abort()
		return err
	}

	txn.Commit()
	return nil
}

func (this *Manager) view(fn func(txn *memdb.Txn) error) error {
	txn := this.db.Txn(false)
	if err := fn(txn); err != nil {
		return err
	}

	txn.Abort()
	return nil
}
