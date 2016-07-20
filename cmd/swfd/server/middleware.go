package server

import (
	"net/http"
	"time"

	log "github.com/funkygao/log4go"
	"github.com/julienschmidt/httprouter"
)

func (this *Server) Middleware(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

		w.Header().Set("Server", "swf")
		w.Header().Set("Content-Type", "application/json; charset=utf8")

		t0 := time.Now()
		h(w, r, params)
		log.Debug("%s[%s] %s", r.Method, r.RequestURI, time.Since(t0))

	}
}
