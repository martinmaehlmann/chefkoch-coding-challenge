package command

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"gitlab.com/m.maehlmann/chefkoch-coding-challenge/internal/wire"
	"go.uber.org/zap"
)

// configCmd represents the config command.
// nolint:gochecknoglobals // global variable needed by cobra
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "pretty prints the application configuration",
	Long:  `This command pretty prints the application configuration.`,
	Run:   Config,
}

// nolint:gochecknoinits // needed by cobra
func init() {
	rootCmd.AddCommand(configCmd)

	// local flags for this command only
	serveCmd.Flags().StringP("indent", "i", "  ", "sets the indentation for the json pretty print")
}

// Config config prints the application configuration in a pretty json format.
func Config(_ *cobra.Command, _ []string) {
	// Initialize a new logger
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(fmt.Sprintf("could not initialize logger: %v", err))
	}

	// initialize the configuration from viper
	configRegistry := wire.InitializeConfig(logger)

	// get the indentation from the viper flag
	indentation, err := serveCmd.Flags().GetString("indent")
	if err != nil {
		logger.Fatal("required flag 'indent' not set. This should not be possible, as it has a default value")
	}

	// get the pretty printed json
	jsonString, err := configRegistry.PrettyString(indentation)
	if err != nil {
		logger.Fatal("could not pretty print the configuration.")
	}

	// print the pretty printed json
	fmt.Println(jsonString)
}
