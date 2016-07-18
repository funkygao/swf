package supervisor

import (
	"runtime"

	"github.com/funkygao/golib/idgen"
	"github.com/funkygao/swf/services/mom"
)

type Supervisor struct {
	m     mom.Service
	idgen *idgen.IdGenerator

	notificationCh chan []byte

	quit chan struct{}
}

func New(m mom.Service, idgen *idgen.IdGenerator) Service {
	return &Supervisor{
		quit:           make(chan struct{}),
		notificationCh: make(chan []byte, 1000),
		m:              m,
		idgen:          idgen,
	}
}

func (this *Supervisor) Name() string {
	return "supervisor"
}

func (this *Supervisor) Start() error {
	go this.recvNotification()

	for i := 0; i < runtime.NumCPU(); i++ {
		go this.schedule()
	}
	return nil
}

func (this *Supervisor) Stop() {
	close(this.quit)
}
