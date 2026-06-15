package cmd

import (
	"bytes"
	"context"
	"errors"
	"io"
	"strings"
	"testing"

	"github.com/adityasgit25/leetcodecli/internal/leetcode"
)

const missingUsernameMessage = "Username required. Usage: leetcode stats <username>. Run \"leetcode help\" for help."

func TestStatsHelpExplainsUsernameRequiredUsage(t *testing.T) {
	output, stderr, err := executeCommand(t, "help", "stats")
	if err != nil {
		t.Fatalf("Execute stats help returned error: %v", err)
	}

	for _, expected := range []string{
		"stats <username>",
		"leetcode stats <username>",
		"requires a username",
	} {
		if !strings.Contains(output, expected) {
			t.Fatalf("stats help output missing %q:\n%s", expected, output)
		}
	}
	if strings.Contains(output, "leetcodecli") {
		t.Fatalf("stats help output contains internal module identity: %q", output)
	}
	assertNoForbiddenHelpTerms(t, output)
	if stderr != "" {
		t.Fatalf("stderr = %q, want empty", stderr)
	}
}

func TestStatsMissingUsernameReturnsUsageExitCode(t *testing.T) {
	exitCode, stdout, stderr := runCommand(t, "stats")

	if exitCode != 2 {
		t.Fatalf("exitCode = %d, want 2", exitCode)
	}
	if stdout != "" {
		t.Fatalf("stdout = %q, want empty", stdout)
	}
	if stderr != missingUsernameMessage+"\n" {
		t.Fatalf("stderr = %q, want %q", stderr, missingUsernameMessage+"\n")
	}
}

func TestStatsMissingUsernameDoesNotInvokeDependencies(t *testing.T) {
	fetchCalled := 0
	renderCalled := 0
	command := newStatsCommand(statsCommandConfig{
		fetch: func(context.Context, string) (leetcode.ProfileStats, error) {
			fetchCalled++
			return leetcode.ProfileStats{}, nil
		},
		render: func(leetcode.ProfileStats, io.Writer) (string, error) {
			renderCalled++
			return "", nil
		},
	})
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	command.SetOut(&stdout)
	command.SetErr(&stderr)
	command.SetArgs(nil)

	err := command.Execute()

	if err == nil {
		t.Fatal("Execute returned nil, want missing username error")
	}
	if fetchCalled != 0 {
		t.Fatalf("fetch called %d times, want 0", fetchCalled)
	}
	if renderCalled != 0 {
		t.Fatalf("render called %d times, want 0", renderCalled)
	}
	if stdout.Len() != 0 {
		t.Fatalf("stdout = %q, want empty", stdout.String())
	}
}

func TestStatsAcceptsExactlyOneUsername(t *testing.T) {
	fetchedUsername := ""
	renderCalled := false
	command := newStatsCommand(statsCommandConfig{
		fetch: func(_ context.Context, username string) (leetcode.ProfileStats, error) {
			fetchedUsername = username
			return commandStats(username), nil
		},
		render: func(stats leetcode.ProfileStats, _ io.Writer) (string, error) {
			renderCalled = true
			return "Profile Summary\nTotal Solved Count\nLanguage Breakdown\n" + stats.Summary.Username + "\n", nil
		},
	})
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	command.SetOut(&stdout)
	command.SetErr(&stderr)
	command.SetArgs([]string{"alice"})

	if err := command.Execute(); err != nil {
		t.Fatalf("Execute returned error: %v", err)
	}
	if fetchedUsername != "alice" {
		t.Fatalf("fetched username = %q, want alice", fetchedUsername)
	}
	if !renderCalled {
		t.Fatal("render was not called")
	}
	for _, expected := range []string{"Profile Summary", "Total Solved Count", "Language Breakdown", "alice"} {
		if !strings.Contains(stdout.String(), expected) {
			t.Fatalf("stdout missing %q:\n%s", expected, stdout.String())
		}
	}
	if stderr.Len() != 0 {
		t.Fatalf("stderr = %q, want empty", stderr.String())
	}
}

func TestStatsHappyPathReturnsExitCodeZeroWithInjectedDependencies(t *testing.T) {
	statsCommand := newStatsCommand(statsCommandConfig{
		fetch: func(_ context.Context, username string) (leetcode.ProfileStats, error) {
			return commandStats(username), nil
		},
		render: func(stats leetcode.ProfileStats, _ io.Writer) (string, error) {
			return "Profile Summary\nTotal Solved Count\nLanguage Breakdown\n" + stats.Summary.ProfileURL + "\n", nil
		},
	})
	root := newRootCommand(statsCommand)
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	exitCode := run(root, []string{"stats", "alice"}, &stdout, &stderr)

	if exitCode != 0 {
		t.Fatalf("exitCode = %d, want 0; stderr=%q", exitCode, stderr.String())
	}
	if stderr.Len() != 0 {
		t.Fatalf("stderr = %q, want empty", stderr.String())
	}
	for _, expected := range []string{"Profile Summary", "Total Solved Count", "Language Breakdown", "https://leetcode.com/u/alice/"} {
		if !strings.Contains(stdout.String(), expected) {
			t.Fatalf("stdout missing %q:\n%s", expected, stdout.String())
		}
	}
}

func TestStatsRejectsBlankUsername(t *testing.T) {
	exitCode, stdout, stderr := runCommand(t, "stats", "   ")

	if exitCode != 2 {
		t.Fatalf("exitCode = %d, want 2", exitCode)
	}
	if stdout != "" {
		t.Fatalf("stdout = %q, want empty", stdout)
	}
	if stderr != missingUsernameMessage+"\n" {
		t.Fatalf("stderr = %q, want %q", stderr, missingUsernameMessage+"\n")
	}
}

func TestStatsTrimsUsernameBeforeFetching(t *testing.T) {
	fetchedUsername := ""
	command := newStatsCommand(statsCommandConfig{
		fetch: func(_ context.Context, username string) (leetcode.ProfileStats, error) {
			fetchedUsername = username
			return commandStats(username), nil
		},
		render: func(leetcode.ProfileStats, io.Writer) (string, error) {
			return "ok", nil
		},
	})
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	command.SetOut(&stdout)
	command.SetErr(&stderr)
	command.SetArgs([]string{" alice "})

	if err := command.Execute(); err != nil {
		t.Fatalf("Execute returned error: %v", err)
	}
	if fetchedUsername != "alice" {
		t.Fatalf("fetched username = %q, want alice", fetchedUsername)
	}
}

func TestStatsRejectsExtraArgumentsAsUsageError(t *testing.T) {
	exitCode, stdout, stderr := runCommand(t, "stats", "alice", "bob")

	if exitCode != 2 {
		t.Fatalf("exitCode = %d, want 2", exitCode)
	}
	if stdout != "" {
		t.Fatalf("stdout = %q, want empty", stdout)
	}
	if !strings.Contains(stderr, "Usage: leetcode stats <username>.") {
		t.Fatalf("stderr = %q, want usage guidance", stderr)
	}
}

func TestStatsFailureMappings(t *testing.T) {
	tests := []struct {
		name       string
		fetchErr   error
		renderErr  error
		wantStderr string
	}{
		{
			name:       "not found",
			fetchErr:   &leetcode.Error{Kind: leetcode.ErrorKindNotFound, Err: errors.New("raw not found")},
			wantStderr: `No LeetCode profile found for "alice". Check the username and try again.`,
		},
		{
			name:       "unavailable",
			fetchErr:   &leetcode.Error{Kind: leetcode.ErrorKindUnavailable, Err: errors.New("raw graphql")},
			wantStderr: `Stats for "alice" are not available from LeetCode right now. Try again later.`,
		},
		{
			name:       "malformed response",
			fetchErr:   &leetcode.Error{Kind: leetcode.ErrorKindMalformedResponse, Err: errors.New("raw json")},
			wantStderr: `Stats for "alice" are not available from LeetCode right now. Try again later.`,
		},
		{
			name:       "endpoint failure",
			fetchErr:   &leetcode.Error{Kind: leetcode.ErrorKindEndpointFailure, Err: errors.New("raw network")},
			wantStderr: `Could not reach LeetCode. Check your connection and try again.`,
		},
		{
			name:       "rate limited",
			fetchErr:   &leetcode.Error{Kind: leetcode.ErrorKindRateLimited, Err: errors.New("raw blocked")},
			wantStderr: `LeetCode blocked or rate-limited the request. Try again later.`,
		},
		{
			name:       "missing stats",
			fetchErr:   &leetcode.Error{Kind: leetcode.ErrorKindMissingStats, Err: errors.New("raw missing")},
			wantStderr: `LeetCode did not return required stats for "alice". Try again later.`,
		},
		{
			name:       "render failure",
			renderErr:  errors.New("raw render"),
			wantStderr: `Could not render stats output. Try again later.`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			renderCalled := false
			statsCommand := newStatsCommand(statsCommandConfig{
				fetch: func(_ context.Context, username string) (leetcode.ProfileStats, error) {
					if tt.fetchErr != nil {
						return leetcode.ProfileStats{}, tt.fetchErr
					}
					return commandStats(username), nil
				},
				render: func(leetcode.ProfileStats, io.Writer) (string, error) {
					renderCalled = true
					if tt.renderErr != nil {
						return "", tt.renderErr
					}
					return "Profile Summary\nTotal Solved Count\nLanguage Breakdown\n", nil
				},
			})
			root := newRootCommand(statsCommand)
			var stdout bytes.Buffer
			var stderr bytes.Buffer

			exitCode := run(root, []string{"stats", "alice"}, &stdout, &stderr)

			if exitCode != 1 {
				t.Fatalf("exitCode = %d, want 1", exitCode)
			}
			if stdout.Len() != 0 {
				t.Fatalf("stdout = %q, want empty", stdout.String())
			}
			if got := normalizeLineEndings(stderr.String()); got != tt.wantStderr+"\n" {
				t.Fatalf("stderr = %q, want %q", got, tt.wantStderr+"\n")
			}
			if tt.fetchErr != nil && renderCalled {
				t.Fatal("render called after fetch failure")
			}
		})
	}
}

func TestStatsMissingUsernameStillUsesUsageExitCode(t *testing.T) {
	exitCode, stdout, stderr := runCommand(t, "stats")

	if exitCode != 2 {
		t.Fatalf("exitCode = %d, want 2", exitCode)
	}
	if stdout != "" {
		t.Fatalf("stdout = %q, want empty", stdout)
	}
	if stderr != missingUsernameMessage+"\n" {
		t.Fatalf("stderr = %q, want %q", stderr, missingUsernameMessage+"\n")
	}
}

func commandStats(username string) leetcode.ProfileStats {
	return leetcode.ProfileStats{
		Summary: leetcode.ProfileSummary{
			Username:    username,
			DisplayName: "Alice Example",
			Ranking:     "123",
			Reputation:  "7",
			ProfileURL:  "https://leetcode.com/u/" + username + "/",
		},
		TotalSolved: 42,
		LanguageCounts: []leetcode.LanguageCount{
			{Language: "Go", Solved: 10},
		},
	}
}
