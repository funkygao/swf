package mysql

import (
	"errors"

	"github.com/funkygao/swf/models"
	"github.com/go-ozzo/ozzo-dbx"
)

func (this *service) SaveHistoryEvents(runId int64, evts models.HistoryEvents) (err error) {
	_, err = this.db.Insert("history", dbx.Params{
		"run_id": runId,
		"events": evts.Bytes(),
	}).Execute()
	return
}

func (this *service) LoadHistoryEvents(runId int64) (evts models.HistoryEvents, err error) {
	q := this.db.NewQuery("SELECT run_id, events FROM history WHERE run_id={:id}")
	q.Bind(dbx.Params{"id": runId})

	var row dbx.NullStringMap
	err = q.One(&row)
	if err != nil {
		return
	}

	if row["events"].Valid {
		evts.From([]byte(row["events"].String))
	} else {
		err = errors.New("invalid events in history table") // TODO
	}

	return
}
