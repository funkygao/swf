package pubsub

import (
	"github.com/funkygao/swf/services/supervisor"
)

type Supervisor struct {
	cf *config
}

func New(cf *config) supervisor.Service {
	return &Supervisor{
		cf: cf,
	}
}

func (this *Supervisor) Start() error {
	// watch Decider topics
	// consume supervisor topic
	return nil
}

func (this *Supervisor) Stop() {

}
