package cmd

import (
	"fmt"

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

		fmt.Println("Staged diff:")
		fmt.Println()
		fmt.Println(diff)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(commitCmd)
}
