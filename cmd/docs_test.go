package cmd

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestUserDocsMatchV1CommandScope(t *testing.T) {
	docs := map[string]string{
		"README.md":            readProjectFile(t, "README.md"),
		"docs/usage.md":        readProjectFile(t, filepath.Join("docs", "usage.md")),
		"docs/limitations.md":  readProjectFile(t, filepath.Join("docs", "limitations.md")),
		"docs/installation.md": readProjectFile(t, filepath.Join("docs", "installation.md")),
	}

	for path, content := range docs {
		for _, required := range []string{
			"leetcode help",
			"leetcode help stats",
			"leetcode stats <username>",
			"Profile Summary",
			"Total Solved Count",
			"Language Breakdown",
		} {
			if !strings.Contains(content, required) {
				t.Fatalf("%s missing required text %q", path, required)
			}
		}
		if strings.Contains(content, "Bare `leetcode stats` succeeds") {
			t.Fatalf("%s presents bare stats as successful", path)
		}
	}

	limitations := docs["docs/limitations.md"]
	for _, required := range []string{
		"unofficial",
		"public-data availability and LeetCode access behavior may change",
		"stores no credentials, tokens, cookies, Session Data, or config files",
		"login",
		"logout",
		"own-profile default lookup",
		"recommendations",
		"goals",
		"reminders",
		"topic-gap analysis",
		"JSON output",
		"dashboards",
		"browser extensions",
		"desktop UI",
		"web apps",
	} {
		if !strings.Contains(limitations, required) {
			t.Fatalf("docs/limitations.md missing required text %q", required)
		}
	}
}

func TestInstallationDocsDescribeReleaseReadiness(t *testing.T) {
	install := readProjectFile(t, filepath.Join("docs", "installation.md"))
	readme := readProjectFile(t, "README.md")

	for _, content := range []string{install, readme} {
		for _, required := range []string{
			"Windows",
			"macOS",
			"Linux",
			"Homebrew",
			"GitHub Releases",
			"checksum",
			"PATH",
			"leetcode",
			"GoReleaser",
			"unofficial",
			"public LeetCode data",
			"stores no credentials, tokens, cookies, Session Data, or config files",
		} {
			if !strings.Contains(content, required) {
				t.Fatalf("installation docs missing %q", required)
			}
		}
		if strings.Contains(content, "leetcodecli") {
			t.Fatalf("installation docs use internal module name: %s", content)
		}
	}

	if _, err := os.Stat(filepath.Join("..", ".goreleaser.yaml")); err == nil {
		t.Fatal(".goreleaser.yaml exists, but release packaging is not scoped")
	} else if !os.IsNotExist(err) {
		t.Fatalf("stat .goreleaser.yaml: %v", err)
	}
}

func TestDocsExamplesMatchHelpCommands(t *testing.T) {
	rootHelp, _, err := executeCommand(t, "help")
	if err != nil {
		t.Fatalf("root help error: %v", err)
	}
	statsHelp, _, err := executeCommand(t, "help", "stats")
	if err != nil {
		t.Fatalf("stats help error: %v", err)
	}

	readme := readProjectFile(t, "README.md")
	for _, expected := range []string{"leetcode help", "leetcode help stats", "leetcode stats <username>"} {
		if !strings.Contains(readme, expected) {
			t.Fatalf("README missing help-aligned example %q", expected)
		}
	}
	if !strings.Contains(rootHelp, "stats") {
		t.Fatal("root help does not list stats")
	}
	if !strings.Contains(statsHelp, "leetcode stats <username>") {
		t.Fatal("stats help does not include leetcode stats <username>")
	}
}

func readProjectFile(t *testing.T, path string) string {
	t.Helper()

	data, err := os.ReadFile(filepath.Join("..", path))
	if err != nil {
		t.Fatalf("read %s: %v", path, err)
	}

	return string(data)
}
