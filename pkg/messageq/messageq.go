package messageq

import (
	"context"
	"github.com/cookienyancloud/testrpckafkapsqlclick/internal/config"
	"github.com/segmentio/kafka-go"
)

func NewKafkaBroker(ctx context.Context, cfg *config.Config) (*kafka.Conn, error) {
	return kafka.DialLeader(ctx,
		cfg.Kafka.Net,
		cfg.Kafka.Addr,
		cfg.Kafka.Topic,
		cfg.Kafka.Partition)
}

func NewKafkaConsumer(cfg *config.Config) *kafka.Reader {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{cfg.Kafka.Addr},
		Partition: cfg.Kafka.Partition,
		Topic:     cfg.Kafka.Topic,
	})
	return r
}
