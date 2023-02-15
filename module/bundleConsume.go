package module

import (
	"go.uber.org/fx"

	"github.com/ferza17/kafka-basic/consumer/module/user/consume"
)

var BundleConsume = fx.Options(
	fx.Provide(consume.NewUserConsume),
)
