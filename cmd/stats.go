package cmd

import (
	"context"
	"fmt"
	"io"

	"github.com/spf13/cobra"

	"leetcodecli/internal/leetcode"
	statsrender "leetcodecli/internal/render"
)

type statsFetcher func(context.Context, string) (leetcode.ProfileStats, error)

type statsRenderer func(leetcode.ProfileStats) (string, error)

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
		render: func(stats leetcode.ProfileStats) (string, error) {
			return statsrender.RenderStatsWithWidthDetector(stats, statsrender.DetectTerminalWidth)
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

			output, err := config.render(stats)
			if err != nil {
				return renderFailureError()
			}

			_, err = io.WriteString(command.OutOrStdout(), output)
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
		config.render = func(leetcode.ProfileStats) (string, error) {
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
	return nil
}
