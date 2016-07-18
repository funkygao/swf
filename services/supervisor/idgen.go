package supervisor

import (
	"crypto/rand"
	"fmt"
	"time"

	"github.com/funkygao/golib/idgen"
)

func (this *Supervisor) nextId() int64 {
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

func (this *Supervisor) nextTaskToken() string {
	b := make([]byte, 8)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
