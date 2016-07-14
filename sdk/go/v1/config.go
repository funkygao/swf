package swfapi

import (
	"github.com/funkygao/gafka/ctx"
)

type config struct {
	zone string
}

func NewConfig(zone string) *config {
	return &config{
		zone: zone,
	}
}

func (this config) Endpoint() string {
	return ctx.Zone(this.zone).SwfEndpoint
}
