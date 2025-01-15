package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
    "os"
)

var rootCmd = &cobra.Command{
    Use:   "task-cli",
    Short: "task-cli is a cli todo tool",
    Long:  "task-cli is a cli todo tool, it can perform multiple operations",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Hello")
    },
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Fprintf(os.Stderr, "Oops. An error while executing Zero '%s'\n", err)
        os.Exit(1)
    }
}