package messageq

import (
	"bytes"
	"github.com/Shopify/sarama"
	"github.com/cookienyancloud/testrpckafkapsqlclick/internal/domain"
	"github.com/cookienyancloud/testrpckafkapsqlclick/internal/repo/clickLog"
	lg "github.com/cookienyancloud/testrpckafkapsqlclick/pkg/logger/logger"
)

type mq struct {
	lgdb  clickLog.IClickLog
	con   sarama.Consumer
	prod  sarama.SyncProducer
	topic string
}

func Newmq(
	lgdb clickLog.IClickLog,
	prod sarama.SyncProducer,
	topic string,
) Imq {
	return &mq{
		lgdb:  lgdb,
		prod:  prod,
		topic: topic,
	}
}

type Imq interface {
	Subscribe()
	MessageReceived(message *sarama.ConsumerMessage) (*domain.UserLog, error)
	MessageToQueue(message domain.UserLog) *sarama.ProducerMessage
}

func (m *mq) Subscribe() {
	partitionList, err := m.con.Partitions(m.topic)
	if err != nil {
		lg.Errorf("err in partitions:%v\n", err)
		return
	}
	initialOffset := sarama.OffsetOldest

	for _, partition := range partitionList {
		pc, _ := m.con.ConsumePartition(m.topic, partition, initialOffset)

		go func(pc sarama.PartitionConsumer) {
			for message := range pc.Messages() {
				userLog, err := m.MessageReceived(message)
				if err != nil {
					lg.Errorf("err in receiving message:%v\n", err)
				}
				err = m.lgdb.LogNewUser(userLog)
				if err != nil {
					lg.Errorf("err logging:%v\n", err)
				}
			}
		}(pc)
	}
}

func (m *mq) MessageReceived(message *sarama.ConsumerMessage) (*domain.UserLog, error) {
	var msg domain.UserLog
	var msgBytes bytes.Buffer
	msgBytes.Write(message.Value)
	err := domain.Decode(&msgBytes, &msg)
	if err != nil {
		return &domain.UserLog{}, nil
	}
	return &msg, nil
}

func (m *mq) MessageToQueue(message domain.UserLog) *sarama.ProducerMessage {
	msg := &sarama.ProducerMessage{
		Topic:     m.topic,
		Partition: -1,
		Value:     message,
	}

	return msg
}
