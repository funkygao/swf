package supervisor

import (
	"runtime"

	"github.com/funkygao/golib/idgen"
	"github.com/funkygao/swf/models"
)

type task struct {
	d *models.PollForDecisionTaskOutput
}

type Supervisor struct {
	idgen *idgen.IdGenerator

	notificationCh chan []byte
	tasks          map[string]task // FIXME

	quit chan struct{}
}

func New(idgen *idgen.IdGenerator) Service {
	return &Supervisor{
		quit:           make(chan struct{}),
		notificationCh: make(chan []byte, 1000),
		idgen:          idgen,
		tasks:          make(map[string]task, 1000),
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
