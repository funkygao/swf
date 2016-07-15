# swf
A aws swf open source implementation in golang.

### Terms

- worker which is registered as ActivityType 
  - in: task
  - out:
- decider which is registered as WorkflowType 
  - in:  Event
  - out: Decision
- supervisor
  - in:
  - out:

workflowId
runId


workflow history example:

    Invoice0001
    
    Start Workflow Execution
    
    Schedule Verify Order
    Start Verify Order Activity
    Complete Verify Order Activity
    
    Schedule Charge Credit Card
    Start Charge Credit Card Activity
    Complete Charge Credit Card Activity
    
    Schedule Ship Order
    Start Ship Order Activity

    Complete Ship Order Activity
    
    Schedule Record Order Completion
    Start Record Order Completion Activity
    Complete Record Order Completion Activity
    
    Close Workflow


### Features

- both sequential and parallel worker
- single decider
- reusable worker for multiple workflows


### Actions

#### performed by worker

- PollForActivityTask
- RespondActivityTaskCompleted
- RespondActivityTaskFailed
- RespondActivityTaskCanceled
- RecordActivityTaskHeartbeat

#### performed by decider

- PollForDecisionTask
- RespondDecisionTaskCompleted

#### related to workflow execution

- StartWorkflowExecution
- RequestCancelWorkflowExecution
- SignalWorkflowExecution
- TerminateWorkflowExecution

#### related to administration

- RegisterActivityType
- DeprecateActivityType

- RegisterWorkflowType
- DeprecateWorkflowType

- RequestCancelWorkflowExecution
- TerminateWorkflowExecution
