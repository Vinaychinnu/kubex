package cmd

import (
	"log"

	"github.com/Vinaychinnu/kubex/pkg/namespace"
	"github.com/spf13/cobra"
)

var namespaceDeleteCmd = &cobra.Command{
	Use:   "delete <name>",
	Short: "Delete a namespace",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := namespace.Delete(args[0]); err != nil {
			log.Fatalf("failed to delete namespace: %v", err)
		}
	},
}

func init() {
	namespaceCmd.AddCommand(namespaceDeleteCmd)
}
