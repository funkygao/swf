package supervisor

import (
	"github.com/funkygao/swf/services"
)

type Service interface {
	services.Service

	Push(topic string, key, val []byte) error
	Pop(topic string) (key, val []byte, err error)
}

var Default Service
