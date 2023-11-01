package rest

import (
	"github.com/lucaskatayama/learn-hexagonal/go/app/service"
	brokermock "github.com/lucaskatayama/learn-hexagonal/go/infra/broker/mock"
	dbmock "github.com/lucaskatayama/learn-hexagonal/go/infra/db/mock"
	emailmock "github.com/lucaskatayama/learn-hexagonal/go/infra/email/mock"
)

func Run() {
	db := dbmock.New()
	email := emailmock.New()
	broker := brokermock.New()
	svc := service.New(db, broker, email)
	rest := New(svc)

	rest.Start()
}
