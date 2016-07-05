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

	// http://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-workflow-exec-lifecycle.html
	var (
		resp models.HttpApiResponse
		err  error
	)
	switch api {
	case models.OpRegisterWorkflowType:
		input := &models.RegisterWorkflowTypeInput{}
		if this.unmarshalRequest(input, &payload.Body, w) {
			return
		}

		resp, err = this.registerWorkflowType(input)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

	case models.OpRegisterActivityType:
		input := &models.RegisterActivityTypeInput{}
		if this.unmarshalRequest(input, &payload.Body, w) {
			return
		}

		resp, err = this.registerActivityType(input)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

	case models.OpStartWorkflowExecution:
		input := &models.StartWorkflowExecutionInput{}
		if this.unmarshalRequest(input, &payload.Body, w) {
			return
		}

		resp, err = this.startWorkflowExecution(input)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

	case models.OpPollForActivityTask:
		input := &models.PollForActivityTaskInput{}
		if this.unmarshalRequest(input, &payload.Body, w) {
			return
		}

		resp, err = this.pollForActivityTask(input)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

	case models.OpPollForDecisionTask:
		input := &models.PollForDecisionTaskInput{}
		if this.unmarshalRequest(input, &payload.Body, w) {
			return
		}

		resp, err = this.pollForDecisionTask(input)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

	case models.OpRespondActivityTaskCompleted:
		input := &models.RespondActivityTaskCompletedInput{}
		if this.unmarshalRequest(input, &payload.Body, w) {
			return
		}

		resp, err = this.respondActivityTaskCompleted(input)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

	case models.OpRespondDecisionTaskCompleted:
		input := &models.RespondDecisionTaskCompletedInput{}
		if this.unmarshalRequest(input, &payload.Body, w) {
			return
		}

		resp, err = this.respondDecisionTaskCompleted(input)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

	case models.OpRecordActivityTaskHeartbeat:

	case models.OpTerminateWorkflowExecution:

	default:
		this.notFoundHandler(w, r)
	}

	w.Write(resp.Bytes())
}

func (this *apiServer) registerWorkflowType(input *models.RegisterWorkflowTypeInput) (
	output *models.RegisterWorkflowTypeOutput, err error) {
	output = &models.RegisterWorkflowTypeOutput{}
	return
}

func (this *apiServer) registerActivityType(input *models.RegisterActivityTypeInput) (
	output *models.RegisterActivityTypeOutput, err error) {
	output = &models.RegisterActivityTypeOutput{}
	return
}

func (this *apiServer) startWorkflowExecution(input *models.StartWorkflowExecutionInput) (
	output *models.StartWorkflowExecutionOutput, err error) {
	// fire WorkflowExecutionStarted Event
	// fire DecisionTaskScheduled Event
	// and schedules the 1st decision task
	output = &models.StartWorkflowExecutionOutput{}
	return
}

func (this *apiServer) pollForDecisionTask(input *models.PollForDecisionTaskInput) (
	output *models.PollForDecisionTaskOutput, err error) {
	// fire ScheduleActivityTask decision
	output = &models.PollForDecisionTaskOutput{}
	return
}

func (this *apiServer) pollForActivityTask(input *models.PollForActivityTaskInput) (
	output *models.PollForActivityTaskOutput, err error) {
	output = &models.PollForActivityTaskOutput{}
	return
}

func (this *apiServer) respondActivityTaskCompleted(input *models.RespondActivityTaskCompletedInput) (
	output *models.RespondActivityTaskCompletedOutput, err error) {
	// fire ActivityTaskCompleted Event
	// fire DecisionTaskScheduled Event
	output = &models.RespondActivityTaskCompletedOutput{}
	return
}

func (this *apiServer) respondDecisionTaskCompleted(input *models.RespondDecisionTaskCompletedInput) (
	output *models.RespondDecisionTaskCompletedOutput, err error) {
	// ScheduleActivityTask Decision
	// fire ActivityTaskScheduled Event
	output = &models.RespondDecisionTaskCompletedOutput{}
	return
}

func (this *apiServer) unmarshalRequest(input models.HttpApiRequest, payload *[]byte, w http.ResponseWriter) (errFoundAndHandled bool) {
	if err := input.From(*payload); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return true
	}

	return false
}
