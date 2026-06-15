package cmd

import (
	"context"
	"fmt"
	"io"
<<<<<<< HEAD
=======
	"strings"
>>>>>>> release-code

	"github.com/spf13/cobra"

	"github.com/adityasgit25/leetcodecli/internal/leetcode"
	statsrender "github.com/adityasgit25/leetcodecli/internal/render"
)

type statsFetcher func(context.Context, string) (leetcode.ProfileStats, error)

<<<<<<< HEAD
type statsRenderer func(leetcode.ProfileStats) (string, error)
=======
type statsRenderer func(leetcode.ProfileStats, io.Writer) (string, error)
>>>>>>> release-code

type statsCommandConfig struct {
	fetch  statsFetcher
	render statsRenderer
}

type userSafeError struct {
	message string
}

func (err userSafeError) Error() string {
	return err.message
}

func NewStatsCommand() *cobra.Command {
	client := leetcode.NewClient()
	return newStatsCommand(statsCommandConfig{
		fetch: client.FetchProfileStats,
<<<<<<< HEAD
		render: func(stats leetcode.ProfileStats) (string, error) {
			return statsrender.RenderStatsWithWidthDetector(stats, statsrender.DetectTerminalWidth)
=======
		render: func(stats leetcode.ProfileStats, output io.Writer) (string, error) {
			return statsrender.RenderStatsWithWidthDetector(stats, statsrender.DetectWriterWidth(output))
>>>>>>> release-code
		},
	})
}

func newStatsCommand(config statsCommandConfig) *cobra.Command {
	config = config.withDefaults()

	return &cobra.Command{
		Use:   "stats <username>",
		Short: "Show public LeetCode profile statistics for a username.",
		Long:  "Show public LeetCode profile statistics with leetcode stats <username>. In v1, this command requires a username.",
		Example: `  leetcode stats <username>
  leetcode stats alice`,
		Args:          validateStatsUsername,
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(command *cobra.Command, args []string) error {
			username := args[0]
			stats, err := config.fetch(command.Context(), username)
			if err != nil {
				return mapStatsFetchError(username, err)
			}

<<<<<<< HEAD
			output, err := config.render(stats)
=======
			outputWriter := command.OutOrStdout()
			output, err := config.render(stats, outputWriter)
>>>>>>> release-code
			if err != nil {
				return renderFailureError()
			}

<<<<<<< HEAD
			_, err = io.WriteString(command.OutOrStdout(), output)
=======
			_, err = io.WriteString(outputWriter, output)
>>>>>>> release-code
			if err != nil {
				return renderFailureError()
			}
			return nil
		},
	}
}

func (config statsCommandConfig) withDefaults() statsCommandConfig {
	if config.fetch == nil {
		config.fetch = func(context.Context, string) (leetcode.ProfileStats, error) {
			return leetcode.ProfileStats{}, nil
		}
	}
	if config.render == nil {
<<<<<<< HEAD
		config.render = func(leetcode.ProfileStats) (string, error) {
=======
		config.render = func(leetcode.ProfileStats, io.Writer) (string, error) {
>>>>>>> release-code
			return "", nil
		}
	}
	return config
}

func mapStatsFetchError(username string, err error) error {
	switch {
	case leetcode.IsErrorKind(err, leetcode.ErrorKindNotFound):
		return userSafeError{message: fmt.Sprintf(`No LeetCode profile found for "%s". Check the username and try again.`, username)}
	case leetcode.IsErrorKind(err, leetcode.ErrorKindUnavailable),
		leetcode.IsErrorKind(err, leetcode.ErrorKindMalformedResponse):
		return userSafeError{message: fmt.Sprintf(`Stats for "%s" are not available from LeetCode right now. Try again later.`, username)}
	case leetcode.IsErrorKind(err, leetcode.ErrorKindEndpointFailure):
		return userSafeError{message: "Could not reach LeetCode. Check your connection and try again."}
	case leetcode.IsErrorKind(err, leetcode.ErrorKindRateLimited):
		return userSafeError{message: "LeetCode blocked or rate-limited the request. Try again later."}
	case leetcode.IsErrorKind(err, leetcode.ErrorKindMissingStats):
		return userSafeError{message: fmt.Sprintf(`LeetCode did not return required stats for "%s". Try again later.`, username)}
	default:
		return userSafeError{message: fmt.Sprintf(`Stats for "%s" are not available from LeetCode right now. Try again later.`, username)}
	}
}

func renderFailureError() error {
	return userSafeError{message: "Could not render stats output. Try again later."}
}

func validateStatsUsername(_ *cobra.Command, args []string) error {
	if len(args) == 0 {
		return usageError{message: `Username required. Usage: leetcode stats <username>. Run "leetcode help" for help.`}
	}
	if len(args) != 1 {
		return usageError{message: `Usage: leetcode stats <username>. Run "leetcode help" for help.`}
	}
<<<<<<< HEAD
=======
	args[0] = strings.TrimSpace(args[0])
	if args[0] == "" {
		return usageError{message: `Username required. Usage: leetcode stats <username>. Run "leetcode help" for help.`}
	}
>>>>>>> release-code
	return nil
}
