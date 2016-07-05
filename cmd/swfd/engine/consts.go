package engine

const (
	opRegisterActivityType = "RegisterActivityType"
	opRegisterDomain       = "RegisterDomain"
	opRegisterWorkflowType = "RegisterWorkflowType"

	opSignalWorkflowExecution        = "SignalWorkflowExecution"
	opStartWorkflowExecution         = "StartWorkflowExecution"
	opRequestCancelWorkflowExecution = "RequestCancelWorkflowExecution"
	opTerminateWorkflowExecution     = "TerminateWorkflowExecution"

	opPollForActivityTask         = "PollForActivityTask"
	opPollForDecisionTask         = "PollForDecisionTask"
	opRecordActivityTaskHeartbeat = "RecordActivityTaskHeartbeat"

	opRespondActivityTaskCanceled  = "RespondActivityTaskCanceled"
	opRespondActivityTaskCompleted = "RespondActivityTaskCompleted"
	opRespondActivityTaskFailed    = "RespondActivityTaskFailed"
	opRespondDecisionTaskCompleted = "RespondDecisionTaskCompleted"

	opCountClosedWorkflowExecutions = "CountClosedWorkflowExecutions"
	opCountOpenWorkflowExecutions   = "CountOpenWorkflowExecutions"
	opCountPendingActivityTasks     = "CountPendingActivityTasks"
	opCountPendingDecisionTasks     = "CountPendingDecisionTasks"
	opDeprecateActivityType         = "DeprecateActivityType"
	opDeprecateDomain               = "DeprecateDomain"
	opDeprecateWorkflowType         = "DeprecateWorkflowType"
	opDescribeActivityType          = "DescribeActivityType"
	opDescribeDomain                = "DescribeDomain"
	opDescribeWorkflowExecution     = "DescribeWorkflowExecution"
	opDescribeWorkflowType          = "DescribeWorkflowType"
	opGetWorkflowExecutionHistory   = "GetWorkflowExecutionHistory"
	opListActivityTypes             = "ListActivityTypes"
	opListClosedWorkflowExecutions  = "ListClosedWorkflowExecutions"
	opListDomains                   = "ListDomains"
	opListOpenWorkflowExecutions    = "ListOpenWorkflowExecutions"
	opListWorkflowTypes             = "ListWorkflowTypes"
)
