package swfapi

type config struct {
	Endpoint string
}

func NewConfig() *config {
	return &config{
		Endpoint: "http://192.168.10.134:9191/v1",
	}
}
