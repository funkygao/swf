package pubsub

import (
	"net/http"

	"github.com/funkygao/gafka/cmd/kateway/api/v1"
	log "github.com/funkygao/log4go"
)

func (this *Supervisor) NotifySupervisor() {

}

func (this *Supervisor) notify(topic, ver string, msg []byte) {
	this.client.Pub("", msg, api.PubOption{
		Topic: topic,
		Ver:   ver,
	})
}

func (this *Supervisor) notifyWorker(topic, ver string, msg []byte) {
	this.client.Pub("", msg, api.PubOption{
		Topic: topic,
		Ver:   ver,
	})
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
		} else {
			break
		}
	}
}
