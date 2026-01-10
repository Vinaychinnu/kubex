package cmd

import (
	"log"

	"github.com/Vinaychinnu/kubex/pkg/deploy"
	"github.com/spf13/cobra"
)

var filePath string

var deployApplyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Apply a deployment manifest",
	Run: func(cmd *cobra.Command, args []string) {
		if err := deploy.Apply(filePath); err != nil {
			log.Fatalf("failed to apply deployment: %v", err)
		}
	},
}

func init() {
	deployApplyCmd.Flags().StringVarP(
		&filePath,
		"file",
		"f",
		"",
		"Path to deployment YAML file",
	)
	deployApplyCmd.MarkFlagRequired("file")

	deployCmd.AddCommand(deployApplyCmd)
}
