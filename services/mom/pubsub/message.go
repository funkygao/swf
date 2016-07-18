package pubsub

import (
	"github.com/funkygao/gafka/cmd/kateway/api/v1"
)

func (this *PubSub) Pub(appid, topic, ver string, msg []byte) error {
	return this.client.Pub("key", msg, api.PubOption{
		Topic: topic,
		Ver:   ver,
	})
}

func (this *PubSub) AddTopic(cluster, appid, topic, ver string) error {
	return this.client.AddTopic(cluster, appid, topic, ver)
}

func (this *PubSub) Sub(appid, topic, ver string) []byte {
	return nil
}
