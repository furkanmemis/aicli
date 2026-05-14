package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "aicli",
	Short: "AI-powered git commit generator",
	Long:  "Generate conventional commit messages using local AI models.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("aicli running...")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}