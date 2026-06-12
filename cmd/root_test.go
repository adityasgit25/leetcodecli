package cmd

import (
	"strings"
	"testing"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func TestRootCommandMetadataUsesLeetcode(t *testing.T) {
	command := NewRootCommand()

	if command.Use != "leetcode" {
		t.Fatalf("Use = %q, want %q", command.Use, "leetcode")
	}

	for _, value := range []string{command.Use, command.Short, command.Long, command.Example} {
		if strings.Contains(value, "leetcodecli") {
			t.Fatalf("public metadata contains leetcodecli: %q", value)
		}
	}
}

func TestRootCommandHelpIsLocalAndHumanReadable(t *testing.T) {
	output, stderr, err := executeCommand(t, "--help")
	if err != nil {
		t.Fatalf("Execute help returned error: %v", err)
	}

	if !strings.Contains(output, "leetcode") {
		t.Fatalf("help output = %q, want command identity", output)
	}
	assertNoForbiddenHelpTerms(t, output)
	if stderr != "" {
		t.Fatalf("stderr = %q, want empty", stderr)
	}
}

func TestRootHelpListsStatsCommand(t *testing.T) {
	output, stderr, err := executeCommand(t, "help")
	if err != nil {
		t.Fatalf("Execute help returned error: %v", err)
	}

	if !strings.Contains(output, "stats") {
		t.Fatalf("root help output = %q, want stats command", output)
	}
	if strings.Contains(output, "leetcodecli") {
		t.Fatalf("root help output contains internal module identity: %q", output)
	}
	assertNoForbiddenHelpTerms(t, output)
	if stderr != "" {
		t.Fatalf("stderr = %q, want empty", stderr)
	}
}

func TestRootCommandTreeExposesOnlyStatsCommand(t *testing.T) {
	command := NewRootCommand()
	commands := command.Commands()

	if len(commands) != 1 {
		t.Fatalf("root command count = %d, want 1", len(commands))
	}
	if commands[0].Name() != "stats" {
		t.Fatalf("root command name = %q, want stats", commands[0].Name())
	}
	assertNoForbiddenCommandSurface(t, command)
}

func TestRootCommandFlagsDoNotExposeUnsupportedSurface(t *testing.T) {
	command := NewRootCommand()
	assertNoForbiddenFlags(t, command)
	for _, child := range command.Commands() {
		assertNoForbiddenFlags(t, child)
	}
}

func assertNoForbiddenCommandSurface(t *testing.T, command *cobra.Command) {
	t.Helper()

	values := []string{
		command.Use,
		command.Short,
		command.Long,
		command.Example,
		strings.Join(command.Aliases, " "),
	}
	for _, value := range values {
		assertNoForbiddenHelpTerms(t, value)
	}
	for _, child := range command.Commands() {
		assertNoForbiddenCommandSurface(t, child)
	}
}

func assertNoForbiddenFlags(t *testing.T, command *cobra.Command) {
	t.Helper()

	flagSets := []struct {
		name  string
		flags *pflag.FlagSet
	}{
		{name: "local", flags: command.LocalFlags()},
		{name: "persistent", flags: command.PersistentFlags()},
	}
	for _, flagSet := range flagSets {
		flagSet.flags.VisitAll(func(flag *pflag.Flag) {
			for _, value := range []string{flag.Name, flag.Shorthand, flag.Usage} {
				assertNoForbiddenHelpTerms(t, value)
			}
		})
	}
}
