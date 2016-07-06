package supervisor

import (
	//"github.com/funkygao/swf/models"
	"github.com/funkygao/swf/services"
)

type Service interface {
	services.Service

	NotifySupervisor()
}

var Default Service
