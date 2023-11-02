package service

import "context"

type DBPort interface {
	Save(ctx context.Context, params Params) error
}

type BrokerPort interface {
	Publish(ctx context.Context, params Params) error
}

type EmailPort interface {
	Send(ctx context.Context, params Params) error
}
