package server

import (
	"io"
	"net/http"
	"strconv"

	"github.com/funkygao/gafka/mpool"
	log "github.com/funkygao/log4go"
	"github.com/funkygao/swf/models"
	"github.com/funkygao/swf/services/manager"
	"github.com/funkygao/swf/services/supervisor"
	"github.com/funkygao/swf/utils"
	"github.com/julienschmidt/httprouter"
)

func (this *apiServer) registerWorkflowType(input *models.RegisterWorkflowTypeInput) (
	output *models.RegisterWorkflowTypeOutput, err error) {
	manager.Default.RegisterWorkflowType(&input.WorkflowType)
	if err = supervisor.Default.AddTopic(input.Cluster, input.Topic(), input.Version); err != nil {
		return
	}

	output = &models.RegisterWorkflowTypeOutput{}

	log.Debug("registerWorkflowType %#v -> %#v", input, output)

	return
}

func (this *apiServer) registerActivityType(input *models.RegisterActivityTypeInput) (
	output *models.RegisterActivityTypeOutput, err error) {
	manager.Default.RegisterActivityType(&input.ActivityType)
	if err = supervisor.Default.AddTopic(input.Cluster, input.Topic(), input.Version); err != nil {
		return
	}

	output = &models.RegisterActivityTypeOutput{}

	log.Debug("registerActivityType %#v -> %#v", input, output)

	return
}

func (this *apiServer) startWorkflowExecution(input *models.StartWorkflowExecutionInput) (
	output *models.StartWorkflowExecutionOutput, err error) {
	// fire WorkflowExecutionStarted Event
	// fire DecisionTaskScheduled Event
	// and schedules the 1st decision task
	// generate runId
	supervisor.Default.NotifySupervisor(input)

	var runId int64
	runId, err = this.ctx.idgen.Next()
	if err != nil {
		return
	}

	output = &models.StartWorkflowExecutionOutput{
		RunId: strconv.FormatInt(runId, 10),
	}

	log.Debug("startWorkflowExecution %#v -> %#v", input, output)

	return
}

func (this *apiServer) pollForDecisionTask(input *models.PollForDecisionTaskInput) (
	output *models.PollForDecisionTaskOutput, err error) {
	// fire ScheduleActivityTask decision
	this.ctx.pubsub.Sub(opt, func(statusCode int, msg []byte) error {

	})

	output = &models.PollForDecisionTaskOutput{}

	log.Debug("pollForDecisionTask %#v -> %#v", input, output)

	return
}

func (this *apiServer) pollForActivityTask(input *models.PollForActivityTaskInput) (
	output *models.PollForActivityTaskOutput, err error) {
	this.ctx.pubsub.Sub(opt, func(statusCode int, msg []byte) error {

	})

	output = &models.PollForActivityTaskOutput{}

	log.Debug("pollForActivityTask %#v -> %#v", input, output)

	return
}

func (this *apiServer) respondActivityTaskCompleted(input *models.RespondActivityTaskCompletedInput) (
	output *models.RespondActivityTaskCompletedOutput, err error) {
	// fire ActivityTaskCompleted Event
	// fire DecisionTaskScheduled Event
	supervisor.Default.NotifySupervisor(input)

	output = &models.RespondActivityTaskCompletedOutput{}

	log.Debug("respondActivityTaskCompleted %#v -> %#v", input, output)

	return
}

func (this *apiServer) respondDecisionTaskCompleted(input *models.RespondDecisionTaskCompletedInput) (
	output *models.RespondDecisionTaskCompletedOutput, err error) {
	// ScheduleActivityTask Decision
	// fire ActivityTaskScheduled Event
	supervisor.Default.NotifySupervisor(input)

	output = &models.RespondDecisionTaskCompletedOutput{}

	log.Debug("respondDecisionTaskCompleted %#v -> %#v", input, output)

	return
}
