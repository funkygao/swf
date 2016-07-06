package pubsub

import (
	"github.com/funkygao/swf/services/supervisor"
)

type Supervisor struct {
}

func New() supervisor.Service {
	return &Supervisor{}
}

func (this *Supervisor) Start() error {
	// watch Decider topics
	// consume supervisor topic
	return nil
}

func (this *Supervisor) Stop() {

}
