package cmd

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestGoReleaserConfigBuildsChecksummedCrossPlatformArtifacts(t *testing.T) {
	config := readProjectFile(t, ".goreleaser.yaml")

	for _, required := range []string{
		"version: 2",
		"project_name: leetcode",
		"builds:",
		"main: .",
		"binary: leetcode",
		"goos:",
		"windows",
		"darwin",
		"linux",
		"goarch:",
		"amd64",
		"arm64",
		"archives:",
		"name_template:",
		"format_overrides:",
		"goos: windows",
		"format: zip",
		"checksum:",
		"name_template: checksums.txt",
		"algorithm: sha256",
	} {
		if !strings.Contains(config, required) {
			t.Fatalf(".goreleaser.yaml missing %q", required)
		}
	}

	for _, forbidden := range []string{
		"brews:",
		"nfpms:",
		"chocolateys:",
		"scoops:",
		"winget:",
		"aurs:",
		"nix:",
		"docker",
	} {
		if strings.Contains(strings.ToLower(config), strings.ToLower(forbidden)) {
			t.Fatalf(".goreleaser.yaml contains unsupported v1 release surface %q", forbidden)
		}
	}

	for _, forbiddenSecret := range []string{
		"ghp_",
		"github_pat_",
		"glpat-",
	} {
		if strings.Contains(strings.ToLower(config), strings.ToLower(forbiddenSecret)) {
			t.Fatalf(".goreleaser.yaml appears to contain a plaintext secret marker %q", forbiddenSecret)
		}
	}
}

func TestHomebrewCaskDistributionPathIsConfigured(t *testing.T) {
	config := readProjectFile(t, ".goreleaser.yaml")
	workflow := readProjectFile(t, ".github/workflows/release.yml")
	readme := readProjectFile(t, "README.md")
	install := readProjectFile(t, filepath.Join("docs", "installation.md"))

	for _, required := range []string{
		"homebrew_casks:",
		"name: leetcode",
		"ids:",
		"- release-archives",
		"binaries:",
		"- leetcode",
		"owner: adityasgit25",
		"name: homebrew-leetcodecli",
		"branch: main",
		"token: \"{{ .Env.HOMEBREW_TAP_GITHUB_TOKEN }}\"",
		"directory: Casks",
		"homepage: \"https://github.com/adityasgit25/leetcodecli\"",
		"description:",
		"system_command \"#{staged_path}/leetcode\", args: [\"help\"]",
	} {
		if !strings.Contains(config, required) {
			t.Fatalf(".goreleaser.yaml missing Homebrew cask config %q", required)
		}
	}

	if !strings.Contains(workflow, "HOMEBREW_TAP_GITHUB_TOKEN: ${{ secrets.HOMEBREW_TAP_GITHUB_TOKEN }}") {
		t.Fatal("release workflow does not pass Homebrew tap token from secrets")
	}

	for _, content := range []string{readme, install} {
		for _, required := range []string{
			"brew install --cask adityasgit25/leetcodecli/leetcode",
			"leetcode help",
			"GitHub Releases",
			"checksum",
			"PATH",
			"stores no credentials, tokens, cookies, Session Data, or config files",
		} {
			if !strings.Contains(content, required) {
				t.Fatalf("Homebrew install docs missing %q", required)
			}
		}
		if strings.Contains(content, "<tap>") {
			t.Fatal("Homebrew install docs still contain placeholder <tap>")
		}
	}
}
