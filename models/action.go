package models

const (
	OpRegisterActivityType = "RegisterActivityType"
	OpRegisterDomain       = "RegisterDomain"
	OpRegisterWorkflowType = "RegisterWorkflowType"

	OpSignalWorkflowExecution        = "SignalWorkflowExecution"
	OpStartWorkflowExecution         = "StartWorkflowExecution"
	OpRequestCancelWorkflowExecution = "RequestCancelWorkflowExecution"
	OpTerminateWorkflowExecution     = "TerminateWorkflowExecution"

	OpPollForActivityTask         = "PollForActivityTask"
	OpPollForDecisionTask         = "PollForDecisionTask"
	OpRecordActivityTaskHeartbeat = "RecordActivityTaskHeartbeat"

	OpRespondActivityTaskCanceled  = "RespondActivityTaskCanceled"
	OpRespondActivityTaskCompleted = "RespondActivityTaskCompleted"
	OpRespondActivityTaskFailed    = "RespondActivityTaskFailed"
	OpRespondDecisionTaskCompleted = "RespondDecisionTaskCompleted"

	OpCountClosedWorkflowExecutions = "CountClosedWorkflowExecutions"
	OpCountOpenWorkflowExecutions   = "CountOpenWorkflowExecutions"
	OpCountPendingActivityTasks     = "CountPendingActivityTasks"
	OpCountPendingDecisionTasks     = "CountPendingDecisionTasks"
	OpDeprecateActivityType         = "DeprecateActivityType"
	OpDeprecateDomain               = "DeprecateDomain"
	OpDeprecateWorkflowType         = "DeprecateWorkflowType"
	OpDescribeActivityType          = "DescribeActivityType"
	OpDescribeDomain                = "DescribeDomain"
	OpDescribeWorkflowExecution     = "DescribeWorkflowExecution"
	OpDescribeWorkflowType          = "DescribeWorkflowType"
	OpGetWorkflowExecutionHistory   = "GetWorkflowExecutionHistory"
	OpListActivityTypes             = "ListActivityTypes"
	OpListClosedWorkflowExecutions  = "ListClosedWorkflowExecutions"
	OpListDomains                   = "ListDomains"
	OpListOpenWorkflowExecutions    = "ListOpenWorkflowExecutions"
	OpListWorkflowTypes             = "ListWorkflowTypes"
)
