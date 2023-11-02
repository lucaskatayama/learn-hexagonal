package cli

import (
	"context"
	"github.com/lucaskatayama/learn-hexagonal/go/app/service"
	brokermock "github.com/lucaskatayama/learn-hexagonal/go/infra/broker/mock"
	dbmock "github.com/lucaskatayama/learn-hexagonal/go/infra/db/mock"
	emailmock "github.com/lucaskatayama/learn-hexagonal/go/infra/email/mock"
)

func DoSomething(ctx context.Context, name string, age int) error {
	db := dbmock.New()
	email := emailmock.New()
	broker := brokermock.New()
	svc := service.New(db, broker, email)
	cli := New(svc)
	return cli.DoSomething(ctx, name, age)
}
