package server

import (
	"net/http"

	log "github.com/funkygao/log4go"
)

// Server is the SimpleWorkFlow server engine.
type Server struct {
	apiServer *apiServer

	shutdownChan chan struct{}
}

func New() *Server {
	this := &Server{
		apiServer:    newApiServer(),
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

func (this *Server) start() error {
	this.setupApis()
	if err := this.apiServer.start(); err != nil {
		return err
	}

	log.Info("engine ready")

	return nil
}

func (this *Server) ServeForever() {
	if err := this.start(); err != nil {
		log.Error("fail to start: %v", err)
		return
	}

	<-this.shutdownChan
}
