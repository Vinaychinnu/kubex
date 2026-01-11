package cmd

import (
	"log"

	"github.com/Vinaychinnu/kubex/pkg/deploy"
	"github.com/spf13/cobra"
)

var deployDeleteCmd = &cobra.Command{
	Use:   "delete <name>",
	Short: "Delete a deployment",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := deploy.Delete(args[0]); err != nil {
			log.Fatalf("failed to delete deployment: %v", err)
		}
	},
}

func init() {
	deployCmd.AddCommand(deployDeleteCmd)
}
