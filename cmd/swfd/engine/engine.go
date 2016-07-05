package engine

import (
	"net/http"
)

// Engine is the SimpleWorkFlow server engine.
type Engine struct {
	apiServer *apiServer
}

func New() *Engine {
	this := &Engine{}
	return this
}

func (this *Engine) setupApis() {
	m := this.Middleware

	if this.apiServer != nil {
		this.apiServer.Router().GET("/alive", m(this.apiServer.checkAliveHandler))
		this.apiServer.Router().NotFound = http.HandlerFunc(this.apiServer.notFoundHandler)
		this.apiServer.Router().POST("/v1", m(this.apiServer.handleApiV1))
	}

}

func (this *Engine) ServeForever() {
	this.setupApis()

	select {}
}
