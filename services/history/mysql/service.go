package mysql

import (
	"log"

	"github.com/funkygao/swf/services/history"
	"github.com/go-ozzo/ozzo-dbx"
)

type service struct {
	dsn string
	db  *dbx.DB
}

func New(dsn string) history.Service {
	this := &service{dsn: dsn}
	return this
}

func (this *service) Name() string {
	return "mysql"
}

func (this *service) Start() (err error) {
	this.db, err = dbx.Open("mysql", this.dsn)
	if err != nil {
		return
	}

	this.db.LogFunc = log.Printf

	return
}

func (this *service) Stop() {
	this.db.Close()
}
