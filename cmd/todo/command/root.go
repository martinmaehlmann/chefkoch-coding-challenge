package command

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "An application, that manages todos.",
	Long: `An application, that manages todos.

This application was created as the result of a coding challenge for chefkoch.de. Its parameters were as follows:

	- We want you to create a REST-service using go (https://golang.org/)
	- The service shall manage ToDo's
	- A ToDo consists of an arbitrary list of Subtasks and is structured as follows:
		{
			id [mandatory]
			name [mandatory]
			description
			tasks: [
				{
					id [mandatory]
					name [mandatory]
					description
				}
			]
		}

	- The service shall serve the following endpoints:
		- GET /todos → Returns a list of all Todos
		- POST /todos → Expects a Todo (without id) and returns a Todo with id
		- GET /todos/{id} → Returns a Todo
		- PUT /todos/{id} → Overwrites an existing Todo
		- DELETE /todos/{id} → Deletes a Todo
	- All ToDo's have to be persisted, the means are up to the applicant.
`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.chefkoch-coding-challenge.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".todo" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".todo")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
