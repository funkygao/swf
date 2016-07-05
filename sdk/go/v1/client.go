package swfapi

import (
	"github.com/funkygao/swf"
	"github.com/parnurzeal/gorequest"
)

type Client struct {
	cf *config
}

func New(cf *config) *Client {
	return &Client{cf: cf}
}

func Default() *Client {
	return &Client{cf: NewConfig()}
}

func (this *Client) call(op string, payload interface{}) (gorequest.Response, []byte, []error) {
	agent := gorequest.New()
	return agent.Post(this.cf.Endpoint).
		Set("User-Agent", "swf-go:"+swf.Version).
		Set("X-Swf-Api", op).
		SendStruct(payload).
		EndBytes()
}
