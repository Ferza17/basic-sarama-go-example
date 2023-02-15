package dto

import (
	"context"
	"encoding/json"
	"time"

	"github.com/Shopify/sarama"

	"github.com/ferza17/kafka-basic/consumer/enum"
	"github.com/ferza17/kafka-basic/consumer/pkg/logger"
)

type CreateUser struct {
	// From Header
	RequestID string `swaggerignore:"true"`

	Name string `json:"name"`
}

func (d *CreateUser) Validate() error {
	return nil
}

func (d *CreateUser) BindByteToStruct(b []byte) error {
	return json.Unmarshal(b, d)
}

func (d *CreateUser) ParseToProducerMessage(ctx context.Context) (*sarama.ProducerMessage, error) {

	value, err := json.Marshal(d)
	if err != nil {
		logger.Infof(ctx, "error marshall err : %v", err)
	}

	msg := &sarama.ProducerMessage{
		Topic:     string(enum.CreateUserTopic),
		Key:       sarama.StringEncoder(d.RequestID),
		Value:     sarama.StringEncoder(value),
		Offset:    sarama.OffsetNewest,
		Partition: 1,
		Timestamp: time.Now(),
	}
	return msg, nil
}
