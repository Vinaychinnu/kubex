package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var podsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List pods in the cluster",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Listing pods")
	},
}

func init() {
	podsCmd.AddCommand(podsListCmd)
}
