package pubsub

type config struct {
	PubsubEndpoint string
	Appid, Secret  string
}

func NewConfig() *config {
	return &config{}
}

func (this *config) decisionQueue() string {
	return "__decision"
}

func (this *config) supervisorQueue() string {
	return "__supervisor"
}

func (this *config) decisionGroup() string {
	return "_sp_d_"
}

func (this *config) supervisorGroup() string {
	return "sx"
}

func (this *config) version() string {
	return "v1"
}
