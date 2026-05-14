package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"time"

	"github.com/charmbracelet/lipgloss"
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

		generatingStyle := lipgloss.NewStyle().
			Foreground(lipgloss.Color("#ff0000")).
			Bold(true)

		fmt.Println(
			generatingStyle.Render(
				"Generating commit message...",
			),
		)

		start := time.Now()

		message, err := ai.GenerateCommitMessage(diff)
		if err != nil {
			return err
		}

		elapsed := time.Since(start).Seconds()

		fmt.Println(
			generatingStyle.Render(
				fmt.Sprintf(
					"Generated commit message in %.2fs",
					elapsed,
				),
			),
		)

		commitStyle := lipgloss.NewStyle().
			Foreground(lipgloss.Color("#00FF87"))

		titleStyle := lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFA500")).
			Bold(true)

		fmt.Println()
		fmt.Println(titleStyle.Render("Suggested commit:"))
		fmt.Println(commitStyle.Render(message))
		fmt.Println()

		reader := bufio.NewReader(os.Stdin)

		commitApprove := lipgloss.NewStyle().
			Foreground(lipgloss.Color("#0400ff"))

		fmt.Print(commitApprove.Render("Commit? (y/n): "))

		input, err := reader.ReadString('\n')
		if err != nil {
			return err
		}

		input = strings.TrimSpace(strings.ToLower(input))

		if input != "y" {
			fmt.Println(titleStyle.Render("Commit cancelled."))
			return nil
		}

		err = gitinternal.Commit(message)
		if err != nil {
			return err
		}

		fmt.Println("Commit created successfully.")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(commitCmd)
}
