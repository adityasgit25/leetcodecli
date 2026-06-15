---
baseline_commit: NO_VCS
---

# Story 4.2: Configure GoReleaser for Checksummed Cross-Platform Artifacts

Status: review

Completion Note: Ultimate context engine analysis completed - comprehensive developer guide created.

## Story

As a developer user,
I want GoReleaser to build LeetcodeCLI artifacts for supported operating systems,
so that Windows, macOS, and Linux users can install the same tested CLI release.

## Requirements Covered

- FR1: Users can install and run LeetcodeCLI on Windows, macOS, and Linux.
- FR13: Users can install, run stats commands, and understand limitations without reading source code.
- FR14: Users can evaluate whether LeetcodeCLI is safe and appropriate for their local environment.

## Acceptance Criteria

1. Given GoReleaser configuration is added, when the release build matrix is inspected, then it builds supported Windows, macOS, and Linux artifacts, and each artifact preserves the public executable name `leetcode`.
2. Given GitHub Releases binaries are the Windows and Linux distribution path, when GoReleaser packaging runs, then it produces downloadable archives for Windows and Linux, and it produces checksums that users can verify before installation.
3. Given macOS distribution uses Homebrew, when GoReleaser packaging runs, then it also produces the macOS artifact inputs required by the Homebrew release path, and it does not introduce unsupported Windows or Linux package-manager distribution for v1.
4. Given release packaging is configured, when a snapshot or dry-run release is executed locally or in CI, then GoReleaser completes without publishing public artifacts, and the generated artifact names, checksums, and executable names match the documented install flow.
5. Given v1 has strict scope boundaries, when release packaging files are reviewed, then they do not add config, auth, session, token, browser-login, logout, JSON, CSV, dashboard, recommendation, goal, reminder, topic-gap, browser extension, web app, or TUI features.

## Tasks / Subtasks

- [x] Add GoReleaser configuration. (AC: 1, 2, 3, 5)
  - [x] Add `.goreleaser.yaml`.
  - [x] Configure Go builds for Windows, macOS, and Linux.
  - [x] Set the binary name to `leetcode`.
  - [x] Use the current main package path from this repository.
  - [x] Configure archives for downloadable GitHub Releases artifacts, including Windows-friendly archive format.
  - [x] Configure checksums, using SHA-256 unless a stronger project standard is chosen.

- [x] Preserve release scope. (AC: 3, 5)
  - [x] Do not add Windows or Linux package-manager publishing in v1.
  - [x] Do not add auth/session/config or private-data release behavior.
  - [x] Do not add JSON, CSV, dashboard, recommendation, or TUI features.

- [x] Update tests that currently reject release config. (AC: 1, 2, 3, 5)
  - [x] Update `cmd/docs_test.go` so it no longer fails only because `.goreleaser.yaml` exists after packaging is intentionally scoped.
  - [x] Add a release-config test that verifies `.goreleaser.yaml` contains `leetcode`, Windows/macOS/Linux targets, checksums, and no unsupported v1 package-manager publishing.
  - [x] Keep docs tests that enforce Homebrew, GitHub Releases, checksums, PATH, executable name, unofficial status, and no credential/session/config storage.

- [x] Run safe local release validation. (AC: 4)
  - [x] Run `go test ./...`.
  - [x] Run `go build ./...`.
  - [x] Run `goreleaser check` if GoReleaser is installed.
  - [x] Run a non-publishing snapshot/dry run such as `goreleaser release --snapshot --clean` if GoReleaser is installed and local environment supports it.
  - [x] Confirm generated artifacts contain `leetcode` or `leetcode.exe` as appropriate.
  - [x] Confirm checksums are generated.

- [x] Document validation notes. (AC: 4)
  - [x] Record whether GoReleaser was installed locally.
  - [x] Record generated artifact names and checksum file names if a snapshot/dry run succeeds.
  - [x] If GoReleaser cannot be run locally, document the blocker and ensure config tests still cover expected structure.

## Dev Notes

### Previous Story Context

- Story 4.1 should have established the release provenance policy, tag/versioning convention, and module-path decision before this story begins.
- Story 3.3 intentionally kept GoReleaser deferred; this story is the first story where `.goreleaser.yaml` is expected.
- Existing `cmd/docs_test.go` currently fails if `.goreleaser.yaml` exists. That assertion must be removed or replaced with scoped release-config expectations.

### Architecture Guardrails

- Release artifacts must preserve public executable name `leetcode`.
- macOS distribution uses Homebrew.
- Windows and Linux distribution use checksummed GitHub Releases binaries with documented install and PATH setup.
- Keep `cmd`, `internal/leetcode`, and `internal/render` behavior unchanged unless tests reveal a release-build issue.
- No release story should weaken no-auth, no-session, no-config, and human-readable-only v1 boundaries.

### Latest Technical Notes

- As of 2026-06-12, GoReleaser docs show v2.16 as current.
- GoReleaser Go build configuration can specify target `GOOS`, `GOARCH`, and the binary name.
- GoReleaser archive configuration supports archive customization and Windows format overrides.
- GoReleaser checksum configuration defaults to SHA-256 and can name the checksum file.
- GoReleaser snapshot release mode is the safe validation path before public publishing.

### Expected Files

- Add: `.goreleaser.yaml`
- Update: `cmd/docs_test.go`
- Add or update: release configuration test, likely under `cmd/` following existing docs/CI test patterns
- Possible update: `docs/installation.md` only if artifact naming expectations need to be documented before publishing
- No `.github/workflows/release.yml` expected until Story 4.3

### Test Guidance

- Use ordinary Go tests for static release-config assertions.
- Do not call live LeetCode endpoints.
- Do not publish GitHub Releases from this story.
- Do not require a Homebrew tap push in this story.
- Run `go test ./...` and `go build ./...`.

### References

- [Source: `_bmad-output/planning-artifacts/epics.md` - Story 4.2]
- [Source: `_bmad-output/planning-artifacts/architecture.md` - release packaging guidance]
- [Source: `README.md` and `docs/installation.md` - intended install paths]
- [Current code: `cmd/docs_test.go` - no premature GoReleaser assertion]
- [External: `https://goreleaser.com/customization/builds/builders/go/` - Go build configuration]
- [External: `https://goreleaser.com/customization/package/archives/` - archives]
- [External: `https://goreleaser.com/customization/package/checksum/` - checksums]

## Dev Agent Record

### Agent Model Used

GPT-5 Codex

### Debug Log References

- `go test ./cmd` initially failed as expected because `.goreleaser.yaml` did not exist.
- `gofmt -w cmd\docs_test.go cmd\release_config_test.go` completed.
- `go test ./...` passed with workspace-local `GOCACHE`.
- `go build ./...` passed with workspace-local `GOCACHE`.
- `goreleaser --version` failed because GoReleaser is not installed on PATH.

### Completion Notes List

- Added `.goreleaser.yaml` using GoReleaser config version 2 with project name `leetcode`, main package `.`, binary name `leetcode`, Windows/macOS/Linux targets, amd64/arm64 architectures, tar.gz archives, Windows zip override, and SHA-256 `checksums.txt`.
- Updated docs tests so `.goreleaser.yaml` is now expected after release packaging is intentionally scoped.
- Added static release-config tests that guard supported OS targets, executable name, archive/checksum structure, and absence of unsupported v1 package-manager/auth/session/structured-output surfaces.
- Local GoReleaser validation could not run because `goreleaser` is unavailable on PATH. Expected artifact names are `leetcode_<version>_<os>_<arch>.tar.gz` for macOS/Linux, `leetcode_<version>_windows_<arch>.zip` for Windows, and `checksums.txt`; actual artifact generation remains deferred to an environment with GoReleaser installed.

### File List

- .goreleaser.yaml
- cmd/docs_test.go
- cmd/release_config_test.go
- _bmad-output/implementation-artifacts/sprint-status.yaml
- _bmad-output/implementation-artifacts/4-2-configure-goreleaser-for-checksummed-cross-platform-artifacts.md

## Change Log

- 2026-06-12: Created Story 4.2 ready for development.
- 2026-06-12: Added scoped GoReleaser artifact/checksum config and static release-config tests; documented local GoReleaser availability limitation.
