package supervisor

import (
	"time"

	log "github.com/funkygao/log4go"
	"github.com/funkygao/swf/models"
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

		evt = models.NewEvent(evts.NextEventId(), time.Now(), models.EventTypeDecisionTaskScheduled)
		evt.DecisionTaskScheduledEventAttributes = &models.DecisionTaskScheduledEventAttributes{}
		evts.AppendEvent(*evt)

		var msg models.PollForDecisionTaskOutput
		msg.Events = *evts
		msg.TaskToken = this.nextTaskToken()
		msg.WorkflowType = m.WorkflowType
		msg.WorkflowExecution.RunId = runId
		msg.WorkflowExecution.WorkflowId = m.WorkflowId

		this.tasks[msg.TaskToken] = task{
			d: &msg,
		}

		this.dispatchDecider(m.WorkflowType, msg.Bytes())

		output = out

		log.Debug("-> %#v", m)

	case *models.RespondDecisionTaskCompletedInput:
		// got ScheduleActivityTask Decision
		// fire ActivityTaskScheduled Event

		for _, dicision := range m.Decisions {
			log.Debug("%#v", decision)

			switch dicision.DecisionType {
			case models.DecisionTypeScheduleActivityTask:
				task := this.tasks[m.TaskToken]

				var msg models.PollForActivityTaskOutput
				msg.TaskToken = m.TaskToken
				msg.Input = d.Input

				evt := models.NewEvent(id, time.Now(), models.EventTypeActivityTaskScheduled)
				evt.ActivityTaskScheduledEventAttributes = &models.ActivityTaskScheduledEventAttributes{}
				evt.ActivityTaskScheduledEventAttributes.Input = task.d.Events[0].WorkflowExecutionStartedEventAttributes.Input

				this.dispatchWorker(dicision.ScheduleActivityTaskDecisionAttributes.ActivityType, msg.Bytes())

			case models.DecisionTypeCompleteWorkflowExecution:
				log.Debug("task[%s] closed", m.TaskToken)
				delete(this.tasks, m.TaskToken)
			}
		}

		log.Debug("-> %#v", m)

	case *models.RespondActivityTaskCompletedInput:
		// fire ActivityTaskCompleted Event
		// fire DecisionTaskScheduled Event

		out := &models.RespondActivityTaskCompletedOutput{}

		evt := models.NewEvent(id, time.Now(), models.EventTypeActivityTaskCompleted)
		evt.ActivityTaskCompletedEventAttributes = &models.ActivityTaskCompletedEventAttributes{}
		evt.ActivityTaskCompletedEventAttributes.Result = m.Result

		evt = models.NewEvent(id, time.Now(), models.EventTypeDecisionTaskScheduled)
		evt.DecisionTaskScheduledEventAttributes = &models.DecisionTaskScheduledEventAttributes{}

		var msg models.PollForDecisionTaskOutput
		msg.Events = *evts
		msg.TaskToken = this.nextTaskToken()

		this.dispatchDecider(w, msg.Bytes())

		output = out

		log.Debug("-> %#v", m)

	default:
		log.Error("-> unkown type: %T", m)

	}

	return

}

func (this *Supervisor) AddTopic(cluster, topic, ver string) error {
	return this.m.AddTopic(cluster, "appid", topic, ver)
}

func (this *Supervisor) dispatchWorker(w models.ActivityType, msg []byte) {
	this.m.Pub("appid", w.Topic(), w.Version, msg)
}

func (this *Supervisor) dispatchDecider(w models.WorkflowType, msg []byte) {
	this.m.Pub("appid", w.Topic(), w.Version, msg)

}

func (this *Supervisor) recvNotification() {

}
