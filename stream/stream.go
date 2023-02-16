package stream

import (
	"context"
	"os"
	"os/signal"

	"golang.org/x/sync/errgroup"

	"github.com/ferza17/kafka-basic/consumer/enum"
	"github.com/ferza17/kafka-basic/consumer/external/kafka"
	"github.com/ferza17/kafka-basic/consumer/module/user/consume"
)

type Streaming struct {
	Kafka        kafka.IKafka
	UserConsumer consume.IUserConsume
}

func NewStreaming(k kafka.IKafka, userConsumer consume.IUserConsume) *Streaming {
	return &Streaming{
		Kafka:        k,
		UserConsumer: userConsumer,
	}
}

func (e *Streaming) StartConsumer(ctx context.Context) error {
	// register Consumer
	errG, ctx := errgroup.WithContext(ctx)
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	// User Consumer
	errG.Go(func() error {
		if err := e.UserConsumer.CreateUser(ctx, enum.StreamingUserCreateUserJson, signals); err != nil {
			return err
		}
		return nil
	})

	err := errG.Wait()
	if err != nil {
		return err
	}
	return err
}
