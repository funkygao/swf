package engine

import (
	"net/http"
	"time"
)

type apiServer struct {
	*webServer
}

func newApiServer() *apiServer {
	this := &apiServer{
		webServer: newWebServer("api", Options.ApiHttpAddr, Options.ApiHttpsAddr),
	}

	return this
}

func (this *apiServer) punishClient(r *http.Request) {
	time.Sleep(time.Second * 2)
}
