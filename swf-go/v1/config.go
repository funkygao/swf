package swfapi

type config struct {
	Endpoint string
}

func NewConfig() *config {
	return &config{}
}
