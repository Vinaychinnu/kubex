package cmd

import (
	"log"

	"github.com/Vinaychinnu/kubex/pkg/service"
	"github.com/spf13/cobra"
)

var serviceDeleteNamespace string

var serviceDeleteCmd = &cobra.Command{
	Use:   "delete <name>",
	Short: "Delete a service",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := service.Delete(args[0], serviceDeleteNamespace); err != nil {
			log.Fatalf("failed to delete service: %v", err)
		}
	},
}

func init() {
	serviceDeleteCmd.Flags().StringVarP(&serviceDeleteNamespace, "namespace", "n", "default", "Namespace to delete service from")
	serviceCmd.AddCommand(serviceDeleteCmd)
}
