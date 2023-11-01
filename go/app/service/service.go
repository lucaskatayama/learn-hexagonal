package service

import "context"

type Service struct {
	db     DBPort
	broker BrokerPort
	email  EmailPort
}

func New(db DBPort, broker BrokerPort, email EmailPort) *Service {
	return &Service{
		db:     db,
		broker: broker,
		email:  email,
	}
}

func (svc *Service) ServiceMethod(ctx context.Context, p Params) (Result, error) {
	panic("not implemented")
}
