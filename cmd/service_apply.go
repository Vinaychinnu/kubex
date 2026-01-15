package cmd

import (
	"log"

	"github.com/Vinaychinnu/kubex/pkg/service"
	"github.com/spf13/cobra"
)

var serviceFile string
var serviceNamespace string

var serviceApplyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Apply a service manifest",
	Run: func(cmd *cobra.Command, args []string) {
		if err := service.Apply(serviceFile, serviceNamespace); err != nil {
			log.Fatalf("failed to apply service: %v", err)
		}
	},
}

func init() {
	serviceApplyCmd.Flags().StringVarP(&serviceFile, "file", "f", "", "Path to service YAML file")
	serviceApplyCmd.Flags().StringVarP(&serviceNamespace, "namespace", "n", "default", "Namespace to apply service to")
	_ = serviceApplyCmd.MarkFlagRequired("file")

	serviceCmd.AddCommand(serviceApplyCmd)
}
