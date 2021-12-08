package messageq

import "github.com/Shopify/sarama"

func NewConsumer(brokers []string) (sarama.SyncProducer, error) {
	config := sarama.NewClient()
	config.Consumer
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(brokers, config)
	return producer, err
}

