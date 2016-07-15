package memory

import (
	"github.com/funkygao/swf/models"
)

func (this *History) Save() {
	//tx := this.db.Txn(true)
}

func (this *History) Get() {

}

func (this *History) SaveWorkflowExecution(*models.StartWorkflowExecutionInput, *models.StartWorkflowExecutionOutput) error {
	return nil
}
