package kafka

import (
	"context"
	"strings"
	"time"

	"github.com/Shopify/sarama"

	"github.com/ferza17/kafka-basic/consumer/config"
	"github.com/ferza17/kafka-basic/consumer/enum"
	"github.com/ferza17/kafka-basic/consumer/pkg/logger"
)

type (
	Kafka struct {
		cfg          *sarama.Config
		syncProducer sarama.SyncProducer
	}

	IKafka interface {
		GetConsumerTopicAndPartitionList(ctx context.Context, topic enum.KafkaTopicName) (sarama.Consumer, []int32, error)
		SendMessage(ctx context.Context, msg *sarama.ProducerMessage) error
		Close() error
	}
)

func NewKafka() IKafka {
	var (
		k   = new(Kafka)
		err error
	)
	k.cfg = sarama.NewConfig()
	// Client Config
	k.cfg.ClientID = "kafka-basic"

	// Consumer Config
	k.cfg.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.BalanceStrategyRoundRobin}
	k.cfg.Consumer.Offsets.Initial = sarama.OffsetOldest
	k.cfg.Consumer.Offsets.AutoCommit.Enable = true
	k.cfg.Consumer.Offsets.AutoCommit.Interval = 1 * time.Minute

	// Producer Config
	k.cfg.Producer.Partitioner = sarama.NewRandomPartitioner
	k.cfg.Producer.RequiredAcks = sarama.WaitForAll
	k.cfg.Producer.Return.Successes = true

	k.syncProducer, err = sarama.NewSyncProducer(strings.Split(config.Get().KafkaBrokers, ","), k.cfg)
	if err != nil {
		logger.Errorf(context.TODO(), "error initialize sync producer err : %v", err)
	}

	return k
}

func (k *Kafka) GetConsumerTopicAndPartitionList(ctx context.Context, topic enum.KafkaTopicName) (sarama.Consumer, []int32, error) {

	consumer, err := sarama.NewConsumer(strings.Split(config.Get().KafkaBrokers, ","), k.cfg)
	if err != nil {
		logger.Infof(ctx, "error initialize consumer err : %v", err)
		return nil, nil, err
	}

	partitionList, err := consumer.Partitions(string(topic))
	if err != nil {
		logger.Infof(ctx, "error find consumer partition list err : %v", err)
		return nil, nil, err
	}

	return consumer, partitionList, nil
}

func (k *Kafka) SendMessage(ctx context.Context, msg *sarama.ProducerMessage) error {
	if _, _, err := k.syncProducer.SendMessage(msg); err != nil {
		logger.Infof(ctx, "error SendMessage err : %v", err)
	}
	return nil
}

func (k *Kafka) Close() error {
	return k.syncProducer.Close()
}
