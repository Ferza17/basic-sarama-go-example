package consume

import (
	"context"
	"os"

	"github.com/Shopify/sarama"
	"golang.org/x/sync/errgroup"

	"github.com/ferza17/kafka-basic/consumer/enum"
	"github.com/ferza17/kafka-basic/consumer/module/user/dto"
	"github.com/ferza17/kafka-basic/consumer/pkg/logger"
)

func (c *UserConsume) CreateUser(ctx context.Context, topic enum.KafkaTopicName, signals chan os.Signal) error {
	errGroup, ctx := errgroup.WithContext(ctx)
	userConsumer, userPartitionList, err := c.Kafka.GetConsumerTopicAndPartitionList(ctx, topic)
	if err != nil {
		logger.Infof(ctx, "error Get Consumer Topic And Partition List err : %v", err)
		return err
	}

	for _, partition := range userPartitionList {
		errGroup.Go(func() error {
			partitionConsumer, err := userConsumer.ConsumePartition(string(topic), partition, sarama.OffsetNewest)
			if err != nil {
				logger.Infof(ctx, "error initialize partition consumer err : %v", err)
				return err
			}
			defer func() {
				if err := partitionConsumer.Close(); err != nil {
					logger.Infof(ctx, "error Close Partition Consumer err : %v", err)
					return
				}
			}()
		ConsumerLoop:
			for {
				select {
				case msg := <-partitionConsumer.Messages():
					reqData := new(dto.CreateUser)

					if err := reqData.BindByteToStruct(msg.Value); err != nil {
						logger.Infof(ctx, "error Bind Byte To Struct : %v", err)
						return err
					}

					logger.Infof(ctx, "reqData : %v", reqData)

					//if expt := c.UserLogic.CreateUser(ctx, reqData); expt != nil {
					//	logger.WithField("RequestID", reqData.RequestID).Infof(ctx, "error: %v", expt.Message)
					//}
				case <-signals:
					logger.Infof(ctx, "closing user consumer")
					if err = partitionConsumer.Close(); err != nil {
						logger.Infof(ctx, "error closing consumer err : %v", err)
						return err
					}
					if err = userConsumer.Close(); err != nil {
						logger.Infof(ctx, "error closing consumer err : %v", err)
						return err
					}
					break ConsumerLoop
				}
			}
			return nil
		})
	}

	if err = errGroup.Wait(); err != nil {
		return err
	}

	return nil
}
