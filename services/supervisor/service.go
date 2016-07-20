package supervisor

import (
	"github.com/funkygao/swf/services"
)

type Service interface {
	services.Service

	Fire(input interface{}) (output interface{}, err error)
	AddTopic(cluster, domain, topic, ver string) error
}
