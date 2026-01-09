package cmd

import (
	"log"

	"github.com/Vinaychinnu/kubex/pkg/namespace"
	"github.com/spf13/cobra"
)

var namespaceListCmd = &cobra.Command{
	Use:   "list",
	Short: "List namespaces",
	Run: func(cmd *cobra.Command, args []string) {
		if err := namespace.List(); err != nil {
			log.Fatalf("failed to list namespaces: %v", err)
		}
	},
}

func init() {
	namespaceCmd.AddCommand(namespaceListCmd)
}
