package history

import (
	"github.com/funkygao/swf/services"
)

type Service interface {
	services.Service
}

var Default Service
