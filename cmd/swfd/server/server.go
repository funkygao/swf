package server

import (
	"fmt"
	"net/http"

	"github.com/funkygao/gafka/ctx"
	"github.com/funkygao/gafka/telementry"
	"github.com/funkygao/gafka/telementry/influxdb"
	"github.com/funkygao/gafka/zk"
	"github.com/funkygao/go-metrics"
	"github.com/funkygao/golib/idgen"
	log "github.com/funkygao/log4go"
	"github.com/funkygao/swf/services"
	"github.com/funkygao/swf/services/history"
	hm "github.com/funkygao/swf/services/history/mysql"
	"github.com/funkygao/swf/services/manager"
	mm "github.com/funkygao/swf/services/manager/mysql"
	"github.com/funkygao/swf/services/mom/pubsub"
	"github.com/funkygao/swf/services/supervisor"
)

// Server is the SimpleWorkFlow server engine.
type Server struct {
	apiServer *apiServer
	services  []services.Service

	zkzone     *zk.ZkZone
	idgen      *idgen.IdGenerator
	supervisor supervisor.Service

	shutdownChan chan struct{}
}

func init() {
	ctx.LoadFromHome()

	log.AddFilter("stdout", log.DEBUG, log.NewConsoleLogWriter())
}

func New() *Server {
	this := &Server{
		services:     make([]services.Service, 0),
		zkzone:       zk.NewZkZone(zk.DefaultConfig(Options.Zone, ctx.ZoneZkAddrs(Options.Zone))),
		shutdownChan: make(chan struct{}),
	}
	this.apiServer = newApiServer(this)
	var err error
	this.idgen, err = idgen.NewIdGenerator(1)
	if err != nil {
		panic(err)
	}

	return this
}

func (this *Server) setupApis() {
	m := this.Middleware

	if this.apiServer != nil {
		this.apiServer.Router().NotFound = http.HandlerFunc(this.apiServer.notFoundHandler)
		this.apiServer.Router().GET("/alive", m(this.apiServer.checkAliveHandler))
		this.apiServer.Router().POST("/v1", m(this.apiServer.handleApiV1))
	}

}

func (this *Server) setupServices() {
	manager.Default = mm.New("root:@/swf") // TODO
	this.addService(manager.Default)

	history.Default = hm.New("root:@/swf")
	this.addService(history.Default)

	this.supervisor = supervisor.New(pubsub.New(pubsub.NewConfig()), this.idgen)
	this.addService(this.supervisor)

	cf, err := influxdb.NewConfig(Options.InfluxServer, Options.InfluxDbName, "", "", Options.ReporterInterval)
	if err == nil {
		telementry.Default = influxdb.New(metrics.DefaultRegistry, cf)
		this.addService(telementry.Default)
	} else {
		log.Error("%v", err)
	}

}

func (this *Server) addService(svc services.Service) {
	this.services = append(this.services, svc)
}

func (this *Server) start() error {
	this.setupApis()
	this.setupServices()

	// start all the services before serving clients
	for _, svc := range this.services {
		go func(svc services.Service) {
			if err := svc.Start(); err != nil {
				panic(fmt.Sprintf("service[%s]: %v", svc.Name(), err))
			} else {
				log.Trace("service[%s] started", svc.Name())
			}
		}(svc)
	}

	if err := this.apiServer.start(); err != nil {
		return err
	}

	log.Info("server ready")

	return nil
}

func (this *Server) stop() {
	for _, svc := range this.services {
		svc.Stop()
		log.Trace("service[%s] stopped", svc.Name())
	}
}

func (this *Server) ServeForever() {
	if err := this.start(); err != nil {
		log.Error("fail to start: %v", err)
		return
	}

	<-this.shutdownChan
	this.stop()
}
