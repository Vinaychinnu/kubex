package cmd

import (
	"log"

	"github.com/Vinaychinnu/kubex/pkg/deploy"
	"github.com/spf13/cobra"
)

var deleteNamespace string

var deployDeleteCmd = &cobra.Command{
	Use:   "delete <name>",
	Short: "Delete a deployment",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := deploy.Delete(args[0], deleteNamespace); err != nil {
			log.Fatalf("failed to delete deployment: %v", err)
		}
	},
}

func init() {

	deployDeleteCmd.Flags().StringVarP(
		&deleteNamespace,
		"namespace",
		"n",
		"default",
		"Namespace to delete deployment from",
	)

	deployCmd.AddCommand(deployDeleteCmd)
}
