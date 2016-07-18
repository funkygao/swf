// Package mom provides message oriented middleware implementation.
package mom

import (
	"github.com/funkygao/swf/services"
)

type Service interface {
	services.Service

	AddTopic(cluster, appid, topic, ver string) error
	Pub(appid, topic, ver string, msg []byte) error
	Sub(appid, topic, ver string)
}
