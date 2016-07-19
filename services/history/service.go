package history

import (
	"github.com/funkygao/swf/models"
	"github.com/funkygao/swf/services"
)

type Service interface {
	services.Service

	SaveHistoryEvents(runId int64, evts models.HistoryEvents) (err error)
	LoadHistoryEvents(runId int64) (evts models.HistoryEvents, err error)
}

var Default Service
