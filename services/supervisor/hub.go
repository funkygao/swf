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

		evts := models.NewHistoryEvents()
		evts.AppendEvent(*evt)

		evt = models.NewEvent(evts.NextEventId(), time.Now(), models.EventTypeDecisionTaskScheduled)
		evt.DecisionTaskScheduledEventAttributes = &models.DecisionTaskScheduledEventAttributes{}
		evts.AppendEvent(*evt)

		// save the history

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

		for _, decision := range m.Decisions {
			log.Debug("%#v", decision)

			switch decision.DecisionType {
			case models.DecisionTypeScheduleActivityTask:
				// how to use m.TaskToken
				// how to make up the history
				// what about ActivityId
				task := this.tasks[m.TaskToken]

				var msg models.PollForActivityTaskOutput
				msg.TaskToken = m.TaskToken
				msg.Input = decision.ScheduleActivityTaskDecisionAttributes.Input

				// fetch history by taskToken
				//taskToken -> task.d.WorkflowExecution.RunId -> history

				var evts models.HistoryEvents
				evt := models.NewEvent(id, time.Now(), models.EventTypeActivityTaskScheduled)
				evt.ActivityTaskScheduledEventAttributes = &models.ActivityTaskScheduledEventAttributes{}
				evt.ActivityTaskScheduledEventAttributes.Input = decision.ScheduleActivityTaskDecisionAttributes.Input
				evts.AppendEvent(evt)

				// save the history?

				this.dispatchWorker(decision.ScheduleActivityTaskDecisionAttributes.ActivityType, msg.Bytes())

			case models.DecisionTypeCompleteWorkflowExecution:
				log.Debug("task[%s] closed", m.TaskToken)
				delete(this.tasks, m.TaskToken)

			default:
				log.Warn("not implemented %T", decision)
			}
		}

		log.Debug("-> %#v", m)

	case *models.RespondActivityTaskCompletedInput:
		// fire ActivityTaskCompleted Event
		// fire DecisionTaskScheduled Event

		out := &models.RespondActivityTaskCompletedOutput{}

		// fetch history by task token

		var evts models.HistoryEvents

		evt := models.NewEvent(id, time.Now(), models.EventTypeActivityTaskCompleted)
		evt.ActivityTaskCompletedEventAttributes = &models.ActivityTaskCompletedEventAttributes{}
		evt.ActivityTaskCompletedEventAttributes.Result = m.Result
		evts.AppendEvent(evt)

		evt = models.NewEvent(id, time.Now(), models.EventTypeDecisionTaskScheduled)
		evt.DecisionTaskScheduledEventAttributes = &models.DecisionTaskScheduledEventAttributes{}
		evts.AppendEvent(evt)

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
