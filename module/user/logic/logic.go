package logic

import (
	"context"

	"go.uber.org/fx"

	"github.com/ferza17/kafka-basic/consumer/exception"
	"github.com/ferza17/kafka-basic/consumer/external/kafka"
	"github.com/ferza17/kafka-basic/consumer/module/user/dto"
)

type (
	IUserLogic interface {
		CreateUser(ctx context.Context, reqData *dto.CreateUser) *exception.Exception
	}
	UserLogic struct {
		fx.In
		Kafka kafka.IKafka
	}
)

func NewUserLogic(u UserLogic) IUserLogic {
	return &u
}
