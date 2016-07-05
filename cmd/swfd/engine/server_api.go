package engine

import (
	"net/http"
	"time"
)

type apiServer struct {
	*webServer
}

func (this *apiServer) punishClient(r *http.Request) {
	time.Sleep(time.Second * 2)
}
