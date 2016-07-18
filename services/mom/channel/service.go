package channel

import (
	"github.com/funkygao/swf/services/mom"
)

type Channel struct {
	chs map[string]chan []byte
}

func New() mom.Service {
	return &Channel{
		chs: make(map[string]chan []byte),
	}
}

func (this *Channel) Name() string {
	return "mom.channel"
}

func (this *Channel) Start() error {
	return nil
}

func (this *Channel) Stop() {

}
