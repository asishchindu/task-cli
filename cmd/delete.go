package cmd

import (
	"task-cli/utils"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "to add task in to task list",
	Long:  "to add task in to task list",
	Run: func(cmd *cobra.Command, args []string) {
		utils.DeleteTask(args[0])
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
