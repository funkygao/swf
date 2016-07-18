package supervisor

import (
	log "github.com/funkygao/log4go"
	"github.com/funkygao/swf/models"
)

func (this *Supervisor) Fire(input interface{}) {
	switch m := input.(type) {
	case *models.StartWorkflowExecutionInput:
		// fire WorkflowExecutionStarted Event
		// fire DecisionTaskScheduled Event
		// and schedules the 1st decision task
		// generate runId
		log.Debug("-> %#v", m)

	case *models.RespondActivityTaskCompletedInput:
		// fire ActivityTaskCompleted Event
		// fire DecisionTaskScheduled Event
		log.Debug("-> %#v", m)

	case *models.RespondDecisionTaskCompletedInput:
		// got ScheduleActivityTask Decision
		// fire ActivityTaskScheduled Event
		log.Debug("-> %#v", m)

	default:
		log.Error("-> unkown type: %T", m)

	}

}

func (this *Supervisor) AddTopic(cluster, topic, ver string) error {
	return this.m.AddTopic(cluster, "appid", topic, ver)
}

func (this *Supervisor) notify(topic, ver string, msg []byte) {
	this.m.Pub("appid", topic, ver, msg)
}

func (this *Supervisor) notifyWorker(topic, ver string, msg []byte) {

}

func (this *Supervisor) notifyDecider() {

}

func (this *Supervisor) recvNotification() {

}
