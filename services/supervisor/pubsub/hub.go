package pubsub

import (
	"net/http"
	"time"

	"github.com/funkygao/gafka/cmd/kateway/api/v1"
	log "github.com/funkygao/log4go"
	"github.com/funkygao/swf/models"
)

func (this *Supervisor) NotifySupervisor(sth interface{}) {
	switch m := sth.(type) {
	case *models.RegisterWorkflowTypeInput:

	case *models.RegisterActivityTypeInput:

	case *models.StartWorkflowExecutionInput:

	default:
		log.Error("unkown type: %T", m)

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
		Topic: this.cf.supervisorQueue(),
		Ver:   this.cf.version(),
		Group: this.cf.supervisorGroup(),
	}
	for {
		if err := this.client.Sub(cf, func(statusCode int, msg []byte) error {
			select {
			case <-this.quit:
				return api.ErrSubStop
			default:
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

func (this *Supervisor) recvDecisions() {
	cf := api.SubOption{
		AppId: this.cf.Appid,
		Topic: this.cf.decisionQueue(),
		Ver:   this.cf.version(),
		Group: this.cf.decisionGroup(),
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

			this.decisionCh <- msg

			return nil
		}); err != nil {
			log.Error("recv decisions: %v", err)
			time.Sleep(time.Minute)
		} else {
			break
		}
	}
}
