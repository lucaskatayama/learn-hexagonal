package kafka

import (
	"github.com/IBM/sarama"
)

type Config struct {
	BootstrapServers []string
}

func NewConsumer(cfg Config) (sarama.Consumer, error) {
	return sarama.NewConsumer(cfg.BootstrapServers, nil)
}

func NewSyncProducer(cfg Config) (sarama.SyncProducer, error) {
	return sarama.NewSyncProducer(cfg.BootstrapServers, nil)
}
