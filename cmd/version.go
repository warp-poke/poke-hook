package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	version = "0.0.1"
	githash = "HEAD"
)

func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Poke-me %s (%s)\n", version, githash)
	},
}
