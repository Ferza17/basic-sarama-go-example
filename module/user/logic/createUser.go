package logic

import (
	"context"

	"github.com/ferza17/kafka-basic/consumer/exception"
	"github.com/ferza17/kafka-basic/consumer/module/user/dto"
	"github.com/ferza17/kafka-basic/consumer/pkg/logger"
)

func (u *UserLogic) CreateUser(ctx context.Context, reqData *dto.CreateUser) *exception.Exception {

	message, err := reqData.ParseToProducerMessage(ctx)
	if err != nil {
		logger.Infof(ctx, "error ParseToProducerMessage err : %v", err)
		return exception.GetException(exception.GENERAL_INTERNAL_SERVER_ERROR)
	}

	if err = u.Kafka.SendMessage(ctx, message); err != nil {
		logger.Infof(ctx, "error AsyncProduce err : %v", err)
		return exception.GetException(exception.GENERAL_INTERNAL_SERVER_ERROR)
	}

	return nil
}
