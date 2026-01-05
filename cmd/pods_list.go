package cmd

import (
	"log"

	"github.com/Vinaychinnu/kubex/pkg/pods"
	"github.com/spf13/cobra"
)

var podsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List pods in the cluster",
	Run: func(cmd *cobra.Command, args []string) {
		if err := pods.ListPods(); err != nil {
			log.Fatalf("failed to list pods: %v", err)
		}
	},
}

func init() {
	podsCmd.AddCommand(podsListCmd)
}
