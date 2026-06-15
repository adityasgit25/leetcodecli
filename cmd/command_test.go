package cmd

import (
	"bytes"
	"strings"
	"testing"
)

var forbiddenHelpTerms = []string{
	"login",
	"logout",
	"session",
	"token",
	"config",
	"json",
	"csv",
	"dashboard",
	"recommendation",
	"goal",
	"reminder",
	"topic-gap",
	"browser extension",
	"web app",
	"tui",
}

func executeCommand(t *testing.T, args ...string) (string, string, error) {
	t.Helper()

	command := NewRootCommand()
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	command.SetOut(&stdout)
	command.SetErr(&stderr)
	command.SetArgs(args)

	err := command.Execute()

	return normalizeLineEndings(stdout.String()), normalizeLineEndings(stderr.String()), err
}

func runCommand(t *testing.T, args ...string) (int, string, string) {
	t.Helper()

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	exitCode := Run(args, &stdout, &stderr)

	return exitCode, normalizeLineEndings(stdout.String()), normalizeLineEndings(stderr.String())
}

func normalizeLineEndings(value string) string {
	return strings.ReplaceAll(value, "\r\n", "\n")
}

func assertNoForbiddenHelpTerms(t *testing.T, output string) {
	t.Helper()

	lower := strings.ToLower(output)
	for _, term := range forbiddenHelpTerms {
		if strings.Contains(lower, term) {
			t.Fatalf("help output contains forbidden term %q:\n%s", term, output)
		}
	}
}
