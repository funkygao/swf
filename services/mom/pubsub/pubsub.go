package pubsub

import (
	"github.com/funkygao/gafka/cmd/kateway/api/v1"
	"github.com/funkygao/swf/services/mom"
)

type PubSub struct {
	cf     *config
	client *api.Client
}

func New(cf *config) mom.Service {
	this := &PubSub{
		cf: cf,
	}
	c := api.DefaultConfig(cf.Appid, cf.Secret)
	c.Pub.Endpoint = cf.PubEndpoint
	c.Sub.Endpoint = cf.SubEndpoint
	c.Admin.Endpoint = cf.AdminEndpoint
	this.client = api.NewClient(c)
	return this
}

func (this *PubSub) Name() string {
	return "pubsub"
}

func (this *PubSub) Start() error {
	return nil
}

func (this *PubSub) Stop() {

}
