package cmd

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestReleaseWorkflowPublishesFromVersionTagsAfterValidation(t *testing.T) {
	workflow := readProjectFile(t, ".github/workflows/release.yml")

	for _, required := range []string{
		"push:",
		"tags:",
		"v*.*.*",
		"contents: write",
		"fetch-depth: 0",
		"actions/setup-go@v6",
		"go-version-file: go.mod",
		"go test ./...",
		"go build ./...",
		"goreleaser/goreleaser-action@v7",
		"distribution: goreleaser",
		"version: \"~> v2\"",
		"args: release --clean",
		"GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}",
	} {
		if !strings.Contains(workflow, required) {
			t.Fatalf("release workflow missing %q", required)
		}
	}

	testIndex := strings.Index(workflow, "go test ./...")
	buildIndex := strings.Index(workflow, "go build ./...")
	releaseIndex := strings.Index(workflow, "goreleaser/goreleaser-action@v7")
	if testIndex == -1 || buildIndex == -1 || releaseIndex == -1 {
		t.Fatal("release workflow missing validation or publish step")
	}
	if !(testIndex < releaseIndex && buildIndex < releaseIndex) {
		t.Fatal("release workflow must run go test and go build before GoReleaser publishes")
	}
}

func TestReleaseNotesTemplateDocumentsTrustContract(t *testing.T) {
	notes := readProjectFile(t, filepath.Join("docs", "release-notes-template.md"))
	config := readProjectFile(t, ".goreleaser.yaml")

	for _, required := range []string{
		"Version",
		"Windows",
		"macOS",
		"Linux",
		"leetcode",
		"checksums.txt",
		"checksum verification",
		"unofficial",
		"public LeetCode data",
		"stores no credentials, tokens, cookies, Session Data, or config files",
		"Breaking command behavior",
		"login",
		"private LeetCode data",
		"JSON",
		"CSV",
		"dashboards",
		"recommendations",
		"goals",
		"reminders",
	} {
		if !strings.Contains(notes, required) {
			t.Fatalf("release notes template missing %q", required)
		}
	}

	for _, required := range []string{
		"release:",
		"header: |",
		"footer: |",
		"checksums.txt",
		"v1 Trust Boundaries",
		"Breaking command behavior",
	} {
		if !strings.Contains(config, required) {
			t.Fatalf(".goreleaser.yaml release notes config missing %q", required)
		}
	}
}
