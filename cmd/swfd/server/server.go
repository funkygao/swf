package server

import (
	"fmt"
	"net/http"

	"github.com/funkygao/gafka/ctx"
	"github.com/funkygao/gafka/zk"
	log "github.com/funkygao/log4go"
	"github.com/funkygao/swf/services"
	"github.com/funkygao/swf/services/history"
	hm "github.com/funkygao/swf/services/history/memory"
	"github.com/funkygao/swf/services/manager"
	mm "github.com/funkygao/swf/services/manager/memory"
	"github.com/funkygao/swf/services/supervisor"
	ps "github.com/funkygao/swf/services/supervisor/pubsub"
)

// Server is the SimpleWorkFlow server engine.
type Server struct {
	apiServer *apiServer
	services  []services.Service

	zkzone *zk.ZkZone

	shutdownChan chan struct{}
}

func init() {
	ctx.LoadFromHome()

	log.AddFilter("stdout", log.DEBUG, log.NewConsoleLogWriter())
}

func New() *Server {
	this := &Server{
		apiServer:    newApiServer(),
		services:     make([]services.Service, 0),
		zkzone:       zk.NewZkZone(zk.DefaultConfig(Options.Zone, ctx.ZoneZkAddrs(Options.Zone))),
		shutdownChan: make(chan struct{}),
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
	manager.Default = mm.New()
	this.addService(manager.Default)

	history.Default = hm.New()
	this.addService(history.Default)

	supervisor.Default = ps.New(ps.NewConfig())
	this.addService(supervisor.Default)
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
	}
}

func (this *Server) ServeForever() {
	if err := this.start(); err != nil {
		log.Error("fail to start: %v", err)
		return
	}

	<-this.shutdownChan
}
