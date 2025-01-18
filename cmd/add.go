package cmd

import (
	"task-cli/utils"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "to add task in to task list",
	Long:  "to add task in to task list",
	Run: func(cmd *cobra.Command, args []string) {
		utils.AddTask(args[0])
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
