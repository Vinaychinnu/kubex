package cmd

import (
	"log"

	"github.com/Vinaychinnu/kubex/pkg/pods"
	"github.com/spf13/cobra"
)

var namespaceName string

var podsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List pods in a namespace",
	Run: func(cmd *cobra.Command, args []string) {
		if err := pods.ListPods(namespaceName); err != nil {
			log.Fatalf("failed to list pods: %v", err)
		}
	},
}

func init() {
	podsListCmd.Flags().StringVarP(&namespaceName, "namespace", "n", "default", "Namespace to list pods from")
	podsCmd.AddCommand(podsListCmd)
}
