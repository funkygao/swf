package pubsub

import (
	"time"

	"github.com/funkygao/golib/idgen"
)

func (this *Supervisor) nextEventId() int64 {
	for {
		eventId, err := this.idgen.Next()
		if err != nil {
			if err == idgen.ErrorClockBackwards {
				time.Sleep(time.Millisecond * 50)
				continue
			} else {
				panic(err)
			}
		}

		return eventId

	}
}
