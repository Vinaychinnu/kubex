package cmd

import "github.com/spf13/cobra"

var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "Manage Kubernetes services",
}

func init() {
	rootCmd.AddCommand(serviceCmd)
}
