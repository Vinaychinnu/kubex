package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kubex",
	Short: "kubex is a CLI tool for Kubernetes automation",
	Long: `kubex is a Go-based CLI tool that interacts directly
with the Kubernetes API to automate common cluster operations.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
