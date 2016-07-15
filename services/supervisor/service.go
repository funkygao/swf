package supervisor

import (
	"github.com/funkygao/swf/services"
)

type Service interface {
	services.Service

	Fire(interface{})
	AddTopic(cluster, topic, ver string) error
}

var Default Service
