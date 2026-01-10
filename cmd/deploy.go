package cmd

import "github.com/spf13/cobra"

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Manage Kubernetes deployments",
}

func init() {
	rootCmd.AddCommand(deployCmd)
}
