package mock

import (
	"context"
	"github.com/lucaskatayama/learn-hexagonal/go/app/service"
	"log"
)

type Adapter struct {
}

func New() *Adapter {
	return &Adapter{}
}
func (a *Adapter) Send(ctx context.Context, params service.Params) error {
	log.Printf("email sent: %v", params)
	return nil
}
