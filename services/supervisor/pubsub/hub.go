package pubsub

import (
	"net/http"
	"time"

	"github.com/funkygao/gafka/cmd/kateway/api/v1"
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
	return this.client.AddTopic(cluster, "appid", topic, ver)
}

func (this *Supervisor) notify(topic, ver string, msg []byte) {
	this.client.Pub("", msg, api.PubOption{
		Topic: topic,
		Ver:   ver,
	})
}

func (this *Supervisor) notifyWorker(topic, ver string, msg []byte) {

}

func (this *Supervisor) notifyDecider() {

}

func (this *Supervisor) recvNotification() {
	cf := api.SubOption{
		AppId: this.cf.Appid,
		Topic: this.cf.Queue(),
		Ver:   this.cf.version(),
		Group: this.cf.Group(),
	}
	for {
		if err := this.client.Sub(cf, func(statusCode int, msg []byte) error {
			select {
			case <-this.quit:
				return api.ErrSubStop
			default:
			}

			if statusCode != http.StatusOK {
				log.Error("http: %s", http.StatusText(statusCode))
				return nil
			}

			this.notificationCh <- msg
			return nil

		}); err != nil {
			log.Error("recv notification: %v", err)
			time.Sleep(time.Minute)
		} else {
			break
		}
	}

}
