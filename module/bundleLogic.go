package module

import (
	"go.uber.org/fx"

	"github.com/ferza17/kafka-basic/consumer/module/user/logic"
)

var BundleLogic = fx.Options(
	fx.Provide(logic.NewUserLogic),
)
