package command

import (
	// noling:golint // this is used to easier embed the version information into the code. This can be changed with a
	// simple sed command by any automatic process.
	_ "embed"
	"fmt"

	"github.com/spf13/cobra"
)

//go:embed version.txt
var version string

// versionCmd represents the serve command.
// nolint:gochecknoglobals // global variable needed by cobra
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "prints version information",
	Long:  `prints version information`,
	Run:   Version,
}

// nolint:gochecknoinits // needed by cobra
func init() {
	rootCmd.AddCommand(versionCmd)
}

// Version prints version information.
func Version(_ *cobra.Command, _ []string) {
	fmt.Println(version)
}
