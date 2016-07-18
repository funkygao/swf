package mysql

import (
	"log"

	"github.com/funkygao/swf/services/manager"
	"github.com/go-ozzo/ozzo-dbx"
)

type Manager struct {
	dsn string
	db  *dbx.DB
}

func New(dsn string) manager.Service {
	this := &Manager{dsn: dsn}
	return this
}

func (this *Manager) Name() string {
	return "mysql"
}

func (this *Manager) Start() (err error) {
	this.db, err = dbx.Open("mysql", this.dsn)
	if err != nil {
		return
	}

	this.db.LogFunc = log.Printf

	return
}

func (this *Manager) Stop() {
	this.db.Close()
}
