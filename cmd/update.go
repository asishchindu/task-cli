package cmd

import (
	"task-cli/utils"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "to add task in to task list",
	Long:  "to add task in to task list",
	Run: func(cmd *cobra.Command, args []string) {
		utils.UpdateTask(args[0], args[1])
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
