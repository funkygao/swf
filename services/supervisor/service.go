package supervisor

import (
	"github.com/funkygao/swf/services"
)

type Service interface {
	services.Service

	NotifySupervisor(interface{})
	AddTopic(cluster, topic, ver string) error
}

var Default Service
