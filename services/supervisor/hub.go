package supervisor

import (
	"time"

	log "github.com/funkygao/log4go"
	"github.com/funkygao/swf/models"
	"github.com/funkygao/swf/services/history"
)

func (this *Supervisor) Fire(input interface{}) (output interface{}, err error) {
	switch m := input.(type) {
	case *models.StartWorkflowExecutionInput:
		runId := this.nextId()
		out := &models.StartWorkflowExecutionOutput{
			RunId: runId,
		}

		// WAL
		evt := models.NewEvent(1, time.Now(), models.EventTypeWorkflowExecutionStarted)
		evt.WorkflowExecutionStartedEventAttributes = &models.WorkflowExecutionStartedEventAttributes{}
		evt.WorkflowExecutionStartedEventAttributes.Input = m.Input
		evt.WorkflowExecutionStartedEventAttributes.WorkflowType = m.WorkflowType

		var x models.HistoryEvents
		evts := &x
		evts.AppendEvent(*evt)

		evt = models.NewEvent(evt.EventId+1, time.Now(), models.EventTypeDecisionTaskScheduled)
		evt.DecisionTaskScheduledEventAttributes = &models.DecisionTaskScheduledEventAttributes{}
		evts.AppendEvent(*evt)

		var msg models.PollForDecisionTaskOutput
		msg.Events = *evts
		msg.WorkflowType = m.WorkflowType
		msg.WorkflowExecution.RunId = runId
		msg.WorkflowExecution.WorkflowId = m.WorkflowId

		// dispatch events to decider queue
		this.m.Pub("appid", m.WorkflowType.Topic(), m.WorkflowType.Version, msg.Bytes())

		history.Default.SaveWorkflowExecution(m, out)

		output = out

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

	return

}

func (this *Supervisor) AddTopic(cluster, topic, ver string) error {
	return this.m.AddTopic(cluster, "appid", topic, ver)
}

func (this *Supervisor) notify(topic, ver string, msg []byte) {
	this.m.Pub("appid", topic, ver, msg)
}

func (this *Supervisor) notifyWorker(topic, ver string, msg []byte) {

}

func (this *Supervisor) notifyDecider(w models.WorkflowType) {

}

func (this *Supervisor) recvNotification() {

}
