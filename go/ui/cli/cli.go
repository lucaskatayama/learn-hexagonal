package cli

import (
	"context"
	"github.com/lucaskatayama/learn-hexagonal/go/app/service"
	"log"
)

type CLI struct {
	svc *service.Service
}

func New(svc *service.Service) *CLI {
	return &CLI{
		svc: svc,
	}
}

func (cli *CLI) DoSomething(ctx context.Context, arg1 string, arg2 int) error {
	p := service.Params{
		Name: arg1,
		Age:  arg2,
	}
	res, err := cli.svc.ServiceMethod(ctx, p)
	if err != nil {
		log.Printf("something went wrong: %v\n", err)
		return err
	}
	log.Printf("sucessful execution: %v\n", res)
	return nil
}
