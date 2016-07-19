package server

import (
	log "github.com/funkygao/log4go"
	"github.com/funkygao/swf/models"
	"github.com/funkygao/swf/services/manager"
)

func (this *apiServer) registerWorkflowType(input *models.RegisterWorkflowTypeInput) (
	output *models.RegisterWorkflowTypeOutput, err error) {
	err = manager.Default.RegisterWorkflowType(&input.WorkflowType)
	if err != nil {
		return
	}

	if err = this.supervisor().AddTopic(input.Cluster, input.Topic(), input.Version); err != nil {
		return
	}

	output = &models.RegisterWorkflowTypeOutput{}

	log.Debug("registerWorkflowType %#v -> %#v", input, output)

	return
}

func (this *apiServer) registerActivityType(input *models.RegisterActivityTypeInput) (
	output *models.RegisterActivityTypeOutput, err error) {
	err = manager.Default.RegisterActivityType(&input.ActivityType)
	if err != nil {
		return
	}

	if err = this.supervisor().AddTopic(input.Cluster, input.Topic(), input.Version); err != nil {
		return
	}

	output = &models.RegisterActivityTypeOutput{}

	log.Debug("registerActivityType %#v -> %#v", input, output)

	return
}

func (this *apiServer) startWorkflowExecution(input *models.StartWorkflowExecutionInput) (
	output *models.StartWorkflowExecutionOutput, err error) {
	o, e := this.supervisor().Fire(input)
	if e != nil {
		return nil, e
	}

	output = o.(*models.StartWorkflowExecutionOutput)

	err = manager.Default.SaveWorkflowExecution(input, output)
	if err != nil {
		log.Error(err)
	}

	log.Debug("startWorkflowExecution %#v -> %#v", input, output)

	return
}

func (this *apiServer) pollForDecisionTask(input *models.PollForDecisionTaskInput) (
	output *models.PollForDecisionTaskOutput, err error) {
	// fire ScheduleActivityTask decision
	//this.ctx.pubsub.Sub(opt, func(statusCode int, msg []byte) error {
	//return nil
	//})

	// how to get the WorkflowType?
	output = &models.PollForDecisionTaskOutput{}
	output.TaskToken = ""

	log.Debug("pollForDecisionTask %#v -> %#v", input, output)

	return
}

func (this *apiServer) pollForActivityTask(input *models.PollForActivityTaskInput) (
	output *models.PollForActivityTaskOutput, err error) {
	//this.ctx.pubsub.Sub(opt, func(statusCode int, msg []byte) error {
	//	return nil
	//})

	// how to get the ActivityType?
	output = &models.PollForActivityTaskOutput{}
	output.Input = ""
	output.TaskToken = ""
	output.ActivityId = ""
	output.WorkflowExecution = models.WorkflowExecution{}

	log.Debug("pollForActivityTask %#v -> %#v", input, output)

	return
}

func (this *apiServer) respondActivityTaskCompleted(input *models.RespondActivityTaskCompletedInput) (
	output *models.RespondActivityTaskCompletedOutput, err error) {
	this.supervisor().Fire(input)

	output = &models.RespondActivityTaskCompletedOutput{}

	log.Debug("respondActivityTaskCompleted %#v -> %#v", input, output)

	return
}

func (this *apiServer) respondDecisionTaskCompleted(input *models.RespondDecisionTaskCompletedInput) (
	output *models.RespondDecisionTaskCompletedOutput, err error) {
	this.ctx.supervisor.Fire(input)

	output = &models.RespondDecisionTaskCompletedOutput{}

	log.Debug("respondDecisionTaskCompleted %#v -> %#v", input, output)

	return
}
