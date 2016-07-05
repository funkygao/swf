package swf

type Client struct {
	cf *config
}

func New(cf *config) *Client {
	return &Client{cf: cf}
}

func Default() *Client {
	return &Client{cf: NewConfig()}
}
