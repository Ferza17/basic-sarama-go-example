package cmd

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Shopify/sarama"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"

	"github.com/ferza17/kafka-basic/consumer/config"
	"github.com/ferza17/kafka-basic/consumer/external"
	"github.com/ferza17/kafka-basic/consumer/module"
	"github.com/ferza17/kafka-basic/consumer/pkg/logger"
	"github.com/ferza17/kafka-basic/consumer/restapi"
	"github.com/ferza17/kafka-basic/consumer/stream"
)

func init() {
	rootCmd.AddCommand(consumerCMD)
}

var consumerCMD = &cobra.Command{
	Use:   "serve-api",
	Short: "Start Serving Consumer",
	Long:  `Start Serving Consumer`,
	Run: func(cmd *cobra.Command, args []string) {
		fx.New(
			fx.Provide(stream.NewStreaming),
			fx.Provide(restapi.NewRouter),
			module.BundleConsume,
			module.BundleLogic,
			module.BundleRoutes,
			external.BundleExternal,
			fx.Invoke(registerHooks),
			fx.WithLogger(fxLogger),
		).Run()
	},
}

func fxLogger() fxevent.Logger {
	return &fxevent.NopLogger
}

func registerHooks(lifecycle fx.Lifecycle, s *stream.Streaming, e *restapi.Router) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {

				// Consumer
				go func() {
					sarama.Logger = log.New(os.Stdout, "[sarama] ", log.LstdFlags)
					if err := s.StartConsumer(ctx); err != nil {
						logger.Errorf(ctx, "error StartConsumer err : %v", err)
						return
					}
					return
				}()

				// Rest API
				go func() {
					e.Start(fmt.Sprintf(":%s", config.Get().Port))
					return
				}()

				return nil
			},
			OnStop: func(ctx context.Context) error {
				logger.Infof(context.TODO(), "Shutting down Server")

				// Closing producer
				if err := s.Kafka.Close(); err != nil {
					logger.Errorf(ctx, "error Closing producer err : %v", err)
					return err
				}

				// Closing HTTP Server
				e.Close()

				return nil
			},
		},
	)
}
