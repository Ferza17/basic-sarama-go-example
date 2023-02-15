package module

import (
	"go.uber.org/fx"

	userRoute "github.com/ferza17/kafka-basic/consumer/module/user/route"
)

// Register Routes
var BundleRoutes = fx.Options(
	fx.Invoke(userRoute.NewRoutes),
)
