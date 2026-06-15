---
baseline_commit: NO_VCS
---

# Story 4.4: Publish the macOS Homebrew Distribution Path

Status: review

Completion Note: Ultimate context engine analysis completed - comprehensive developer guide created.

## Story

As a macOS developer user,
I want to install LeetcodeCLI through Homebrew,
so that I can use the expected macOS package-management flow instead of manually unpacking binaries.

## Requirements Covered

- FR1: Users can install and run LeetcodeCLI on Windows, macOS, and Linux.
- FR13: Users can install, run stats commands, and understand limitations without reading source code.
- FR14: Users can evaluate whether LeetcodeCLI is safe and appropriate for their local environment.

## Acceptance Criteria

1. Given Homebrew is the intended macOS distribution channel, when release packaging is configured, then GoReleaser updates the agreed Homebrew tap or formula path for LeetcodeCLI, and the formula or tap entry installs the executable as `leetcode`.
2. Given a macOS user follows installation guidance, when they run the documented `brew install` command after release publication, then Homebrew installs LeetcodeCLI successfully, and running `leetcode help` from the installed binary does not crash.
3. Given the Homebrew formula or tap entry is reviewed, when metadata and tests are inspected, then they reference the correct release version, artifact URL, and checksum, and they do not require credentials, tokens, cookies, Session Data, or config files.
4. Given Homebrew distribution is live, when README and installation docs are reviewed, then placeholder tap or release commands are replaced with actual published Homebrew instructions, and Windows and Linux guidance continues to use checksummed GitHub Releases binaries.

## Tasks / Subtasks

- [x] Confirm Homebrew tap ownership and publishing path. (AC: 1, 3)
  - [x] Identify the actual Homebrew tap repository and branch.
  - [x] Decide whether the project uses the current GoReleaser Homebrew cask mechanism or another supported Homebrew formula/tap mechanism.
  - [x] Record the token/permission expectation without committing secret values.

- [x] Configure GoReleaser Homebrew publishing. (AC: 1, 3)
  - [x] Update `.goreleaser.yaml` to publish the macOS Homebrew tap entry.
  - [x] Ensure the installed command is `leetcode`.
  - [x] Ensure the tap entry references the correct release version, artifact URL, and checksum.
  - [x] Do not add Windows or Linux package-manager publishing.

- [x] Add Homebrew validation. (AC: 2, 3)
  - [x] Add tests that inspect the GoReleaser Homebrew configuration for tap owner/name, binary `leetcode`, artifact/checksum linkage, and no plaintext tokens.
  - [x] If the generated tap entry is available locally, validate it includes a Homebrew test or equivalent install check that runs `leetcode help`.
  - [x] If macOS/Homebrew is unavailable locally, document the validation gap and rely on static tests plus CI/release workflow validation.

- [x] Replace placeholder installation docs. (AC: 4)
  - [x] Update README and `docs/installation.md` from `brew install <tap>/leetcode` to the actual published Homebrew command.
  - [x] Keep Windows and Linux as checksummed GitHub Releases binary flows.
  - [x] Keep unofficial/public-data/no credential-session-config trust boundaries visible.

- [x] Verify. (AC: 1, 2, 3, 4)
  - [x] Run `go test ./...`.
  - [x] Run `go build ./...`.
  - [x] Run non-publishing GoReleaser validation if possible.
  - [x] On macOS with Homebrew available, validate the documented install path after release publication or document why live validation is deferred.

## Dev Notes

### Previous Story Context

- Story 4.2 should have configured GoReleaser artifacts and checksums.
- Story 4.3 should have configured tagged GitHub Releases.
- Story 3.3 added placeholder Homebrew guidance that must become real in this story.

### Architecture Guardrails

- macOS distribution uses Homebrew.
- Windows and Linux remain checksummed GitHub Releases binary installs.
- Public executable name stays `leetcode`.
- No credentials, tokens, cookies, Session Data, or config files are required by the installed CLI.
- Do not introduce login, logout, private-data access, JSON/CSV output, dashboards, recommendations, goals, reminders, topic-gap analysis, browser extensions, desktop UI, web apps, or TUI behavior.

### Latest Technical Notes

- As of 2026-06-12, current GoReleaser docs emphasize `homebrew_casks` for generated Homebrew tap entries and include a `binaries` field for installed binary names.
- GoReleaser Homebrew publishing can push generated files to a configured tap repository and may require a token for cross-repository publishing.
- Homebrew formula/cask validation should include a test or install check that runs the installed executable, such as `leetcode help`.

### Expected Files

- Update: `.goreleaser.yaml`
- Update: `README.md`
- Update: `docs/installation.md`
- Add or update: Homebrew/release config tests under `cmd/`
- Optional: generated tap output under `dist/` only during local validation; do not commit generated release artifacts unless project conventions explicitly require it.

### Test Guidance

- Static tests should verify no placeholder `<tap>` remains after Homebrew distribution is live.
- Static tests should verify docs still mention Windows, Linux, GitHub Releases, checksums, and PATH.
- Static tests should verify `.goreleaser.yaml` does not contain plaintext secret values.
- Run `go test ./...` and `go build ./...`.

### References

- [Source: `_bmad-output/planning-artifacts/epics.md` - Story 4.4]
- [Source: `_bmad-output/planning-artifacts/prds/prd-LeetcodeCLI-2026-06-11/prd.md` - Distribution]
- [Source: `docs/installation.md` - current placeholder Homebrew guidance]
- [External: `https://goreleaser.com/customization/publish/homebrew_casks/` - GoReleaser Homebrew publishing]
- [External: `https://docs.brew.sh/Formula-Cookbook` - Homebrew tests and formula guidance]

## Dev Agent Record

### Agent Model Used

GPT-5 Codex

### Debug Log References

- `go test ./cmd` initially failed as expected because `.goreleaser.yaml` did not have `homebrew_casks`.
- `gofmt -w cmd\docs_test.go cmd\release_config_test.go` completed.
- `go test ./...` passed with workspace-local `GOCACHE`.
- `go build ./...` passed with workspace-local `GOCACHE`.
- `brew --version` failed because Homebrew is not installed on this Windows environment.
- `goreleaser --version` failed because GoReleaser is not installed on PATH.

### Completion Notes List

- Configured GoReleaser `homebrew_casks` for tap repository `adityasgit25/homebrew-leetcodecli` on branch `main`, with cask name `leetcode`, installed binary `leetcode`, and cask output under `Casks`.
- Recorded cross-repository tap publishing token expectation through `HOMEBREW_TAP_GITHUB_TOKEN`, passed from GitHub Actions secrets and referenced in GoReleaser without plaintext secret values.
- Added a cask post-install hook that runs `leetcode help` from the staged binary as the generated tap entry's install-time check.
- Replaced Homebrew placeholder docs with `brew install --cask adityasgit25/leetcodecli/leetcode` and kept Windows/Linux on checksummed GitHub Releases archives.
- Added static Homebrew release tests for tap owner/name/branch, binary name, token-from-env, install check, no placeholder `<tap>`, GitHub Releases/checksum/PATH guidance, and trust boundaries.
- Live Homebrew install validation and tap push are deferred because this environment is Windows without Homebrew, GoReleaser is unavailable, and no release credentials or publishing approval were provided.

### File List

- .github/workflows/release.yml
- .goreleaser.yaml
- README.md
- docs/installation.md
- cmd/docs_test.go
- cmd/release_config_test.go
- _bmad-output/implementation-artifacts/sprint-status.yaml
- _bmad-output/implementation-artifacts/4-4-publish-the-macos-homebrew-distribution-path.md

## Change Log

- 2026-06-12: Created Story 4.4 ready for development.
- 2026-06-12: Configured Homebrew cask publishing path, tap token expectation, install-time help check, real Homebrew command docs, and static Homebrew validation tests.
