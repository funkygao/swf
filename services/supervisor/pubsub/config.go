package pubsub

type config struct {
	PubEndpoint   string
	SubEndpoint   string
	AdminEndpoint string
	Appid, Secret string
}

func NewConfig() *config {
	return &config{
		PubEndpoint:   "localhost:9191",
		SubEndpoint:   "localhost:9192",
		AdminEndpoint: "localhost:9193",
	}
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
