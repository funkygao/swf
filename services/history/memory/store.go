package memory

import (
	"github.com/funkygao/swf/services/historystore"
)

type History struct {
}

func New() historystore.Service {
	return &History{}
}

func (this *History) Start() error {
	return nil
}

func (this *History) Stop() {

}
