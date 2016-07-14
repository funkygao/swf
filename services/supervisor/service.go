package supervisor

import (
	//"github.com/funkygao/swf/models"
	"github.com/funkygao/swf/services"
)

type Service interface {
	services.Service

	NotifySupervisor()
	AddTopic(cluster, topic, ver string) error
}

var Default Service
