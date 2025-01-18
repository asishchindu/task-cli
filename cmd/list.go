package cmd

import (
	"task-cli/utils"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "to add task in to task list",
	Long:  "to add task in to task list",
	Run: func(cmd *cobra.Command, args []string) {
		utils.ListTasks()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
