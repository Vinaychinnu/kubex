package cmd

import "github.com/spf13/cobra"

var podsCmd = &cobra.Command{
	Use:   "pods",
	Short: "Manage Kubernetes pods",
	Long:  "Commands to interact with Kubernetes pods",
}

func init() {
	rootCmd.AddCommand(podsCmd)
}
