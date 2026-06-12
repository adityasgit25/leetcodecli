package cmd

import (
	"strings"
	"testing"
)

func TestGitHubActionsCIWorkflow(t *testing.T) {
	workflow := readProjectFile(t, ".github/workflows/ci.yml")

	for _, required := range []string{
		"push:",
		"pull_request:",
		"contents: read",
		"ubuntu-latest",
		"macos-latest",
		"windows-latest",
		"actions/checkout@v6",
		"actions/setup-go@v6",
		"go-version-file: go.mod",
		"go test ./...",
		"go build ./...",
	} {
		if !strings.Contains(workflow, required) {
			t.Fatalf("ci workflow missing %q", required)
		}
	}
}
