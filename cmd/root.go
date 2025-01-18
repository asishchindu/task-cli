package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "task-cli",
	Short: "task-cli is a cli todo tool",
	Long:  "task-cli is a cli todo tool, it can perform multiple operations",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func Execute() {
	// completion this is a default command. that im disabling.
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Oops. An error while executing task-cli '%s'\n", err)
		os.Exit(1)
	}
}
