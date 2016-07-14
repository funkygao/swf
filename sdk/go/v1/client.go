package swfapi

import (
	"github.com/funkygao/gafka/ctx"
)

type Client struct {
	cf *config
}

func New(cf *config) *Client {
	return &Client{cf: cf}
}

func WithZone(zone string) *Client {
	return &Client{cf: NewConfig(zone)}
}

func Default() *Client {
	return &Client{cf: NewConfig(ctx.DefaultZone())}
}
