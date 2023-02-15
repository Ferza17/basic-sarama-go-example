package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/ferza17/kafka-basic/consumer/config"
	"github.com/ferza17/kafka-basic/consumer/pkg/logger"
)

var rootCmd = &cobra.Command{
	Use:   "github.com/ferza17/kafka-basic",
	Short: "kafka-basic Service is Running...",
	Long:  `kafka-basic Service is Running...`,
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	config.SetConfig(".")

	// Init Logger
	err := logger.NewLogger()
	if err != nil {
		log.Fatalf("rootCMD| initConfig | Could not instantiate log %s", err.Error())
	}

}
