package models

type WorkflowType struct {
	Name    string // 1-256 in len
	Version string // 1-64 in len
	Cluster string
}

func (this WorkflowType) Topic() string {
	return decisionTopicPrefix + this.Name
}

func (this WorkflowType) Validate() error {
	if this.Name == "" || this.Version == "" || this.Cluster == "" {
		return ErrRequiredMissing
	}

	return nil
}

// A WorkflowExecution can be closed as completed, canceled, failed, or timed out.
//  A workflow execution could be closed by the decider, by the person administering
// the workflow, or by SWF.
type WorkflowExecution struct {
	WorkflowId string // 1-256 in len
	RunId      string // 1-64 in len
}
