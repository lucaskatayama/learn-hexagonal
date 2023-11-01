package mock

import (
	"context"
	"github.com/lucaskatayama/learn-hexagonal/go/app/service"
)

type Adapter struct {
}

func New() *Adapter {
	return &Adapter{}
}

func (a *Adapter) Save(ctx context.Context, params service.Params) error {
	//TODO implement me
	panic("implement me")
}
