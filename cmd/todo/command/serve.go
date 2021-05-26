package command

import (
	"fmt"
	"log"

	"gitlab.com/m.maehlmann/chefkoch-coding-challenge/internal/config"
	"gitlab.com/m.maehlmann/chefkoch-coding-challenge/internal/repository"

	"github.com/spf13/cobra"
	"gitlab.com/m.maehlmann/chefkoch-coding-challenge/internal/wire"
	"go.uber.org/zap"
)

// automigrate boolean flag weather to automigrate the database or not
// nolint:gochecknoglobals // cobra syntax
var automigrate bool

// serveCmd represents the serve command.
// nolint:gochecknoglobals // global variable needed by cobra
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "serves the gin-gonic server.",
	Long: `Serves the gin-gonic server.

Reads the configuration from viper. The default configuration file is in under $HOME/.todo`,
	Run: Serve,
}

// nolint:gochecknoinits // needed by cobra
func init() {
	rootCmd.AddCommand(serveCmd)

	serveCmd.Flags().BoolVar(&automigrate, "automigrate", false,
		"automatically migrate your schema, to keep your schema up to date.")
}

// Serve serves the gin-gonic server.
func Serve(_ *cobra.Command, _ []string) {
	// get a new logger instance
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(fmt.Sprintf("could not initialize logger: %v", err))
	}

	// initialize the repository and connect to it
	todoRepository := repository.NewTodoRepository(config.NewPostgresConfig(), logger)
	defer todoRepository.Close()
	todoRepository.Connect()

	// if automigration is needed, migrate / create the schema
	if automigrate {
		migrateDatabase(todoRepository)
	}

	// start the server and wait for connections
	ginServer := wire.InitializeServer(todoRepository, logger)
	ginServer.Run()
}

// migrateDatabase automigrates the repository.
func migrateDatabase(todoRepository repository.TodoRepository) {
	err := todoRepository.AutoMigrate()
	if err != nil {
		log.Fatal(fmt.Sprintf("could not automigrate database: %v", err))
	}
}
