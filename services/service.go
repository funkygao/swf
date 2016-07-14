package services

type Service interface {
	Start() error
	Stop()

	Name() string
}
