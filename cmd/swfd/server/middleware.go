package server

import (
	"net/http"

	log "github.com/funkygao/log4go"
	"github.com/julienschmidt/httprouter"
)

func (this *Server) Middleware(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		log.Debug("%s %s", r.Method, r.RequestURI)

		w.Header().Set("Server", "swf")
		w.Header().Set("Content-Type", "application/json; charset=utf8")

		h(w, r, params)
	}
}
