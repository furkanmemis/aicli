package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Generate AI commit message",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Generating commit message...")
	},
}

func init() {
	rootCmd.AddCommand(commitCmd)
}