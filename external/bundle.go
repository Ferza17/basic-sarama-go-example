package external

import (
	"go.uber.org/fx"

	"github.com/ferza17/kafka-basic/consumer/external/kafka"
)

var BundleExternal = fx.Options(
	fx.Provide(kafka.NewKafka),
)
