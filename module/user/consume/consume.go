package consume

import (
	"context"
	"os"

	"go.uber.org/fx"

	"github.com/ferza17/kafka-basic/consumer/enum"
	"github.com/ferza17/kafka-basic/consumer/external/kafka"
	userLogic "github.com/ferza17/kafka-basic/consumer/module/user/logic"
)

type (
	IUserConsume interface {
		CreateUser(ctx context.Context, topic enum.KafkaTopicName, signal chan os.Signal) error
	}
	UserConsume struct {
		fx.In
		Kafka     kafka.IKafka
		UserLogic userLogic.UserLogic
	}
)

func NewUserConsume(u UserConsume) IUserConsume {
	return &u
}
