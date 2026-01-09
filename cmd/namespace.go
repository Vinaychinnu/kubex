package cmd

import "github.com/spf13/cobra"

var namespaceCmd = &cobra.Command{
	Use:   "namespace",
	Short: "Manage Kubernetes namespaces",
}

func init() {
	rootCmd.AddCommand(namespaceCmd)
}
