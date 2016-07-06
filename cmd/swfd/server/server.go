package server

import (
	"net/http"

	"github.com/funkygao/gafka/ctx"
	"github.com/funkygao/gafka/zk"
	log "github.com/funkygao/log4go"
	"github.com/funkygao/swf/services"
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

}

func (this *Server) addService(svc services.Service) {
	this.services = append(this.services, svc)
}

func (this *Server) start() error {
	this.setupApis()
	this.setupServices()

	if err := this.apiServer.start(); err != nil {
		return err
	}

	log.Info("engine ready")

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
