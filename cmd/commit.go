package cmd

import (
	"fmt"

	"github.com/furkanmemis/aicli/internal/ai"
	gitinternal "github.com/furkanmemis/aicli/internal/git"
	"github.com/spf13/cobra"
)

var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Generate AI commit message",
	RunE: func(cmd *cobra.Command, args []string) error {

		diff, err := gitinternal.GetStagedDiff()
		if err != nil {
			return err
		}

		if diff == "" {
			fmt.Println("No staged changes found.")
			return nil
		}

		fmt.Println("Generating commit message...")

		message, err := ai.GenerateCommitMessage(diff)
		if err != nil {
			return err
		}

		fmt.Println()
		fmt.Println("Suggested commit:")
		fmt.Println(message)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(commitCmd)
}
