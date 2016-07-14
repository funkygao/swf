package server

import (
	"net/http"
	"time"
)

type apiServer struct {
	*webServer
}

func newApiServer(ctx *Server) *apiServer {
	this := &apiServer{
		webServer: newWebServer(ctx, "api", Options.ApiHttpAddr, Options.ApiHttpsAddr),
	}

	return this
}

func (this *apiServer) punishClient(r *http.Request) {
	time.Sleep(time.Second * 2)
}
