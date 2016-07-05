package engine

import (
	"net/http"

	log "github.com/funkygao/log4go"
)

// Engine is the SimpleWorkFlow server engine.
type Engine struct {
	apiServer *apiServer

	shutdownChan chan struct{}
}

func New() *Engine {
	this := &Engine{
		apiServer:    newApiServer(),
		shutdownChan: make(chan struct{}),
	}
	return this
}

func (this *Engine) setupApis() {
	m := this.Middleware

	if this.apiServer != nil {
		this.apiServer.Router().NotFound = http.HandlerFunc(this.apiServer.notFoundHandler)
		this.apiServer.Router().GET("/alive", m(this.apiServer.checkAliveHandler))
		this.apiServer.Router().POST("/v1", m(this.apiServer.handleApiV1))
	}

}

func (this *Engine) start() error {
	this.setupApis()
	if err := this.apiServer.start(); err != nil {
		return err
	}

	log.Info("engine ready")

	return nil
}

func (this *Engine) ServeForever() {
	if err := this.start(); err != nil {
		log.Error("fail to start: %v", err)
		return
	}

	<-this.shutdownChan
}
