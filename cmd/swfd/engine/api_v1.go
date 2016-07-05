package engine

import (
	"io"
	"net/http"

	"github.com/funkygao/gafka/mpool"
	log "github.com/funkygao/log4go"
	"github.com/funkygao/swf/models"
	"github.com/funkygao/swf/utils"
	"github.com/julienschmidt/httprouter"
)

func (this *apiServer) handleApiV1(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	payloadLen := int(r.ContentLength)
	payload := mpool.NewMessage(payloadLen)
	payload.Body = payload.Body[0:payloadLen]
	defer payload.Free()

	if _, err := io.ReadAtLeast(r.Body, payload.Body, payloadLen); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	api := r.Header.Get("X-Swf-Api")
	log.Debug("%s %s(%s) %s", api, r.RemoteAddr, utils.HttpRemoteIp(r), string(payload.Body))

	switch api {
	case models.OpStartWorkflowExecution:
		input := &models.StartWorkflowExecutionInput{}
		if err := input.From(payload.Body); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		this.startWorkflowExecution(input)

	case models.OpRegisterWorkflowType:

	case models.OpRegisterActivityType:

	case models.OpPollForActivityTask:
	case models.OpPollForDecisionTask:

	case models.OpRecordActivityTaskHeartbeat:

	case models.OpRespondActivityTaskCompleted:
	case models.OpRespondDecisionTaskCompleted:
	case models.OpTerminateWorkflowExecution:

	default:
		this.notFoundHandler(w, r)
	}
}

func (this *apiServer) startWorkflowExecution(input *models.StartWorkflowExecutionInput) (
	output *models.StartWorkflowExecutionOutput, err error) {
	return nil, nil

}

// When a decider schedules an activity task, it provides the data (which you determine) that the activity worker needs to perform the activity task.
// Amazon SWF inserts this data into the activity task before sending it to the activity worker.
//
// The execution state for a workflow execution is stored in its workflow history.
// There can be only one decision task open at any time for a given workflow execution.
// Every time a state change occurs for a workflow execution, Amazon SWF schedules a decision task.

/*
RegisterDomain(name, description, workflowExecutionRetentionPeriodInDays string)

RegisterWorkflowType(domain, name, version, description, defaultTaskList string)

RegisterActivityType(domain, name, version, description, defaultTaskList string)

// workflowId unique across StartWorkflowExecution
StartWorkflowExecution(domain, taskList, workflowType, workflowId, input string, tagList []string) (runId string)

PollForDecisionTask(domain, taskList, identity string, maximumPageSize int, reverseOrder bool) (events []HistoryEvent, previousStartedEventId int64, startedEventId int64, taskToken string, we WorkflowExecution)
RespondDecisionTaskCompleted(decisions []Decision, executionContext string, taskToken string)

PollForActivityTask(domain, taskList, identity string) (activityId string, activityType, input, startedEventId, taskToken, runId, workflowId string)
RespondActivityTaskCompleted(result, taskToken string)

*/
