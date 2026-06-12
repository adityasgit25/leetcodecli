---
baseline_commit: NO_VCS
---

# Story 3.3: Prepare Public Release Installation Guidance

Status: review

Completion Note: Ultimate context engine analysis completed - comprehensive developer guide created.

## Story

As a developer user,
I want platform-specific installation guidance,
so that I can run LeetcodeCLI on my operating system with minimal setup friction.

## Requirements Covered

- FR1: Users can install and run LeetcodeCLI on Windows, macOS, and Linux.
- FR13: Users can install, run stats commands, and understand limitations without reading source code.
- FR14: Users can evaluate whether LeetcodeCLI is safe and appropriate for their local environment.

## Acceptance Criteria

1. Given supported operating systems are documented, when installation guidance is reviewed, then it includes a macOS Homebrew path, and it includes Windows and Linux GitHub Releases binary paths with checksum and PATH setup expectations.
2. Given v1 release packaging may be completed after core implementation, when release-readiness documentation is reviewed, then GoReleaser is identified as the intended packaging path when distribution stories begin, and no release configuration is required before it is intentionally scoped.
3. Given users evaluate whether to install the tool, when release documentation is reviewed, then it repeats the public-data dependency, unofficial status, supported OSes, and no credential/session/config storage boundary, and it keeps the public executable name `leetcode` consistent.

## Tasks / Subtasks

- [x] Update installation guidance. (AC: 1, 3)
  - [x] Document supported OSes: Windows, macOS, and Linux.
  - [x] Include a macOS Homebrew path.
  - [x] Include Windows GitHub Releases binary install path with checksum verification and PATH setup expectations.
  - [x] Include Linux GitHub Releases binary install path with checksum verification and PATH setup expectations.
  - [x] Keep executable name `leetcode`.

- [x] Document release packaging intent without adding premature config. (AC: 2)
  - [x] State GoReleaser is the intended packaging path when distribution stories begin.
  - [x] Do not add `.goreleaser.yaml` unless release packaging has been explicitly scoped.
  - [x] Keep release docs honest if no public release artifacts exist yet.

- [x] Repeat trust boundaries in release docs. (AC: 3)
  - [x] State LeetcodeCLI is unofficial unless project status has intentionally changed.
  - [x] State it depends on public LeetCode data and access behavior may change.
  - [x] State v1 stores no credentials, tokens, cookies, Session Data, or config files.
  - [x] State no login/logout/session/private-data workflow exists in v1.

- [x] Align docs. (AC: 1, 2, 3)
  - [x] Update README install section.
  - [x] Update `docs/usage.md` or add a release/install doc if useful.
  - [x] Update `docs/limitations.md` if trust statements have drifted.
  - [x] Ensure all examples use `leetcode`.

- [x] Verify. (AC: 1, 2, 3)
  - [x] Run `go test ./...`.
  - [x] Run `go build ./...`.
  - [x] If docs tests exist, run them too.

## Dev Notes

### Previous Story Context

- Story 3.1 documents usage, scope, and trust boundaries.
- Story 3.2 adds CI validation across supported OSes.
- Story 3.3 completes v1 release-readiness documentation, not release automation.

### Architecture Guardrails

- macOS distribution uses Homebrew.
- Windows and Linux distribution use checksummed GitHub Releases binaries with documented install and PATH setup.
- GoReleaser is deferred until distribution stories intentionally begin.
- Release artifacts must preserve the public executable name `leetcode`.
- Do not create config, auth, session, token, logout, browser-login, JSON, CSV, dashboard, recommendation, or TUI docs.

### Latest Technical Notes

- GoReleaser supports package managers including Homebrew and works with GitHub Actions, but this story should document intent rather than add release configuration unless the release packaging story is explicitly in scope.
- Keep install docs practical even before release artifacts exist: mark commands as intended release flow if binaries/formula are not yet published.

### Expected Files

- Update: `README.md`
- Update: `docs/usage.md`
- Update: `docs/limitations.md`
- Optional add: `docs/installation.md` or `docs/release.md` if the docs need a separate install page.
- No `.goreleaser.yaml` expected unless scope changes.

### Test Guidance

- If docs are tested, assert command examples use `leetcode`, not `leetcodecli`.
- Assert install docs mention checksums for Windows/Linux binaries.
- Assert trust boundary statements remain present.

### References

- [Source: `_bmad-output/planning-artifacts/epics.md` - Story 3.3]
- [Source: `_bmad-output/planning-artifacts/architecture.md` - Infrastructure & Deployment, Development Workflow Integration]
- [Source: `_bmad-output/planning-artifacts/prds/prd-LeetcodeCLI-2026-06-11/prd.md` - FR1, FR13, FR14, Distribution]
- [External: `https://goreleaser.com/` - GoReleaser package-manager and GitHub Actions support]
- [Previous: `_bmad-output/implementation-artifacts/3-2-add-cross-platform-ci-validation.md`]

## Dev Agent Record

### Agent Model Used

GPT-5 Codex

### Debug Log References

- 2026-06-12: Planned release-readiness documentation only, with no GoReleaser configuration until release packaging is explicitly scoped.
- 2026-06-12: Red-phase `go test ./...` failed because `docs/installation.md` did not exist.
- 2026-06-12: Docs test caught missing literal `Homebrew` wording in README; install heading was aligned.
- 2026-06-12: Verification passed: `go test ./...` and `go build ./...`.

### Completion Notes List

- Added platform-specific installation guidance for Windows, macOS, and Linux.
- Documented macOS Homebrew intent and Windows/Linux GitHub Releases binary flows with checksum verification and PATH setup expectations.
- Documented GoReleaser as the intended future packaging path without adding `.goreleaser.yaml`.
- Repeated release trust boundaries: unofficial status, public LeetCode data dependency, changing access behavior, and no credential/session/config storage.
- Extended docs tests to assert install/release guidance and absence of premature GoReleaser config.

### File List

- `_bmad-output/implementation-artifacts/3-3-prepare-public-release-installation-guidance.md`
- `README.md`
- `cmd/docs_test.go`
- `docs/installation.md`

## Change Log

- 2026-06-12: Added release installation guidance and marked Story 3.3 ready for review.
