package models

import (
	"time"
)

const (
	// @enum ChildPolicy
	ChildPolicyTerminate = "TERMINATE"
	// @enum ChildPolicy
	ChildPolicyRequestCancel = "REQUEST_CANCEL"
	// @enum ChildPolicy
	ChildPolicyAbandon = "ABANDON"
)

type WorkflowType struct {
	Name               string `db:"name"`
	Version            string `db:"ver"`
	Domain             string `db:"domain"`
	Description        string `db:"desc"`
	Cluster            string `db:"cluster"`
	DefaultChildPolicy string `db:"child_policy"`

	DefaultExecutionStartToCloseTimeout time.Duration `db:"-"`
	DefaultTaskStartToCloseTimeout      time.Duration `db:"-"`
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
	RunId      int64  // 1-64 in len
}
