package dummy

import (
	"github.com/funkygao/swf/services/manager"
)

type Manager struct {
}

func New() manager.Service {
	return &Manager{}
}

func (this *Manager) Start() error {
	return nil
}

func (this *Manager) Stop() {

}
