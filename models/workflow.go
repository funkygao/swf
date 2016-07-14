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

type WorkflowExecution struct {
	WorkflowId string // 1-256 in len
	RunId      string // 1-64 in len
}
