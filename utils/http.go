package utils

import (
	"net/http"
	"strings"
)

// HttpRemoteIp returns ip only, without remote port.
func HttpRemoteIp(r *http.Request) string {
	forwardedFor := r.Header.Get("X-Forwarded-For") // client_ip,proxy_ip,proxy_ip,...
	if forwardedFor == "" {
		// not behind haproxy, directly connected
		p := strings.SplitN(r.RemoteAddr, ":", 2)
		return p[0]
	}

	return forwardedFor // FIXME forwardedFor might be comma seperated ip list, but here for performance ignore it
}
