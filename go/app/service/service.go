package service

import (
	"context"
	"fmt"
)

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
	err := svc.email.Send(ctx, p)
	res := Result{Msg: fmt.Sprintf("%s is %d years old", p.Name, p.Age)}
	return res, err
}
