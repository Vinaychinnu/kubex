package cmd

import (
	"log"

	"github.com/Vinaychinnu/kubex/pkg/deploy"
	"github.com/spf13/cobra"
)

var filePath string
var deployNamespace string
var dryRun bool

var deployApplyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Apply a deployment manifest",
	Run: func(cmd *cobra.Command, args []string) {
		if err := deploy.Apply(filePath, deployNamespace, dryRun); err != nil {
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

	deployApplyCmd.Flags().StringVarP(
		&deployNamespace,
		"namespace",
		"n",
		"default",
		"Namespace to apply deployment to",
	)

	deployApplyCmd.Flags().BoolVar(
		&dryRun,
		"dry-run",
		false,
		"Show what would be applied without making changes",
	)

	deployApplyCmd.MarkFlagRequired("file")
	deployCmd.AddCommand(deployApplyCmd)

}
