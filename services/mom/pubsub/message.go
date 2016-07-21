package pubsub

import (
	"fmt"
	"net/http"

	"github.com/funkygao/gafka/cmd/kateway/api/v1"
	log "github.com/funkygao/log4go"
)

func (this *PubSub) Pub(appid, topic, ver string, msg []byte) error {
	log.Debug("PUB(%s, %s/%s) %s", appid, topic, ver, string(msg))

	return this.client.Pub("key", msg, api.PubOption{
		Topic: topic,
		Ver:   ver,
	})
}

func (this *PubSub) AddTopic(cluster, appid, topic, ver string) error {
	return this.client.AddTopic(cluster, appid, topic, ver)
}

func (this *PubSub) Sub(appid, topic, ver string) (payload []byte, err error) {
	err = this.client.Sub(api.SubOption{
		AppId: appid,
		Topic: topic,
		Ver:   ver,
		Group: fmt.Sprintf("%s.%s.%s", appid, topic, ver), // FIXME
	}, func(statusCode int, msg []byte) error {
		if statusCode != http.StatusOK {
			if statusCode == http.StatusNoContent {
				return nil
			}

			return fmt.Errorf("%v", http.StatusText(statusCode))
		}

		payload = msg
		return api.ErrSubStop
	})

	return
}
