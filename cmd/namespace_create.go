package cmd

import (
	"log"

	"github.com/Vinaychinnu/kubex/pkg/namespace"
	"github.com/spf13/cobra"
)

var namespaceCreateCmd = &cobra.Command{
	Use:   "create <name>",
	Short: "Create a namespace",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := namespace.Create(args[0]); err != nil {
			log.Fatalf("failed to create namespace: %v", err)
		}
	},
}

func init() {
	namespaceCmd.AddCommand(namespaceCreateCmd)
}
