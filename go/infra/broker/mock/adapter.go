package mock

import (
	"context"
	"github.com/lucaskatayama/learn-hexagonal/go/app/service"
	"log"
)

type ProducerAdapter struct {
}

func New() *ProducerAdapter {
	return &ProducerAdapter{}
}

func (p *ProducerAdapter) Publish(ctx context.Context, params service.Params) error {
	log.Printf("excuted broker mock: %v\n", params)
	return nil
}
