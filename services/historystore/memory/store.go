package memory

import (
	"github.com/funkygao/swf/services/historystore"
)

type HistoryStore struct {
}

func New() historystore.Service {
	return &HistoryStore{}
}

func (this *HistoryStore) Start() error {
	return nil
}

func (this *HistoryStore) Stop() {

}
