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

func (this *config) Queue() string {
	return "__supervisor"
}

func (this *config) Group() string {
	return "_sp_d_"
}

func (this *config) version() string {
	return "v1"
}
