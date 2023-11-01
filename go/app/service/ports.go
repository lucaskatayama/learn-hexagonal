package service

import "context"

type DBPort interface {
	Save(ctx context.Context, p Params) error
}

type BrokerPort interface {
	Publish(ctx context.Context, p Params) error
}

type EmailPort interface {
	Send(ctx context.Context, p Params) error
}
