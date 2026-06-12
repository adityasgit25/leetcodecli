package cmd

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

const (
	exitCodeOK    = 0
	exitCodeError = 1
	exitCodeUsage = 2
)

type usageError struct {
	message string
}

func (e usageError) Error() string {
	return e.message
}

func NewRootCommand() *cobra.Command {
	return newRootCommand(NewStatsCommand())
}

func newRootCommand(statsCommand *cobra.Command) *cobra.Command {
	command := &cobra.Command{
		Use:   "leetcode",
		Short: "Inspect public LeetCode profile statistics from the terminal.",
		Long:  "leetcode is a terminal tool for viewing public LeetCode profile statistics.",
		Example: `  leetcode --help
  leetcode help`,
		SilenceUsage:      true,
		SilenceErrors:     true,
		CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
	}
	if statsCommand != nil {
		command.AddCommand(statsCommand)
	}

	return command
}

func Execute() {
	os.Exit(Run(os.Args[1:], os.Stdout, os.Stderr))
}

func Run(args []string, stdout io.Writer, stderr io.Writer) int {
	return run(NewRootCommand(), args, stdout, stderr)
}

func run(command *cobra.Command, args []string, stdout io.Writer, stderr io.Writer) int {
	command.SetOut(stdout)
	command.SetErr(stderr)
	command.SetArgs(args)

	if err := command.Execute(); err != nil {
		var usage usageError
		if errors.As(err, &usage) {
			_, _ = fmt.Fprintln(stderr, usage.Error())
			return exitCodeUsage
		}

		_, _ = fmt.Fprintln(stderr, err)
		return exitCodeError
	}

	return exitCodeOK
}
