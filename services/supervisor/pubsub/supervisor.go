package pubsub

import (
	"runtime"

	"github.com/funkygao/gafka/cmd/kateway/api/v1"
	"github.com/funkygao/swf/services/supervisor"
)

type Supervisor struct {
	cf *config

	client *api.Client

	decisionCh, notificationCh chan []byte

	quit chan struct{}
}

func New(cf *config) supervisor.Service {
	this := &Supervisor{
		cf:             cf,
		quit:           make(chan struct{}),
		decisionCh:     make(chan []byte, 1000),
		notificationCh: make(chan []byte, 1000),
	}

	c := api.DefaultConfig(cf.Appid, cf.Secret)
	c.Pub.Endpoint = cf.PubEndpoint
	c.Sub.Endpoint = cf.SubEndpoint
	c.Admin.Endpoint = cf.AdminEndpoint
	this.client = api.NewClient(c)
	return this
}

func (this *Supervisor) Name() string {
	return "supervisor"
}

func (this *Supervisor) Start() error {
	go this.recvDecisions()
	go this.recvNotification()
	for i := 0; i < runtime.NumCPU(); i++ {
		go this.schedule()
	}
	return nil
}

func (this *Supervisor) Stop() {
	close(this.quit)
}
