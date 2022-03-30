package service

const (
	Name = "Service"
)

type Service interface {
	Echo(request string, response *string) error
}