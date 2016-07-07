package pubsub

type config struct {
	PubsubEndpoint string
}

func NewConfig() *config {
	return &config{}
}

func (this *config) queue() string {
	return "__supervisor"
}
