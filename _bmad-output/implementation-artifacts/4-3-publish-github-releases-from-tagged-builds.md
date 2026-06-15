---
baseline_commit: NO_VCS
---

# Story 4.3: Publish GitHub Releases from Tagged Builds

Status: review

Completion Note: Ultimate context engine analysis completed - comprehensive developer guide created.

## Story

As a developer user,
I want tagged releases to publish checksummed binaries through GitHub Releases,
so that Windows and Linux installation instructions point to real, verifiable artifacts.

## Requirements Covered

- FR1: Users can install and run LeetcodeCLI on Windows, macOS, and Linux.
- FR13: Users can install, run stats commands, and understand limitations without reading source code.
- FR14: Users can evaluate whether LeetcodeCLI is safe and appropriate for their local environment.

## Acceptance Criteria

1. Given a release tag is pushed, when the release workflow runs, then it executes the required tests and build validation before publishing artifacts, and it fails without publishing if validation does not pass.
2. Given validation passes for a release tag, when GoReleaser publishes to GitHub Releases, then the release includes Windows, macOS, and Linux artifacts as configured, and it includes checksum files for artifact verification.
3. Given users inspect the GitHub Release, when release notes are reviewed, then they identify the release version, supported operating systems, executable name `leetcode`, checksum verification expectation, and v1 trust boundaries, and they do not promise authentication, private LeetCode data, JSON/CSV output, dashboards, recommendations, goals, reminders, or other deferred features.
4. Given public command behavior must remain stable, when a release is prepared, then release notes call out any breaking command behavior changes before publication, and patch releases do not silently change command names, argument meaning, output category names, or no-auth behavior.

## Tasks / Subtasks

- [x] Add a release workflow. (AC: 1, 2)
  - [x] Add `.github/workflows/release.yml`.
  - [x] Trigger release publishing from version tags using the project tag convention from Story 4.1.
  - [x] Run `go test ./...` before publishing.
  - [x] Run `go build ./...` before publishing.
  - [x] Use checkout with full history for GoReleaser.
  - [x] Use the Go version from `go.mod`.
  - [x] Use `goreleaser/goreleaser-action` to run GoReleaser.

- [x] Configure minimal release permissions. (AC: 1, 2)
  - [x] Grant only the permissions required to create GitHub Releases and upload release assets.
  - [x] Use `GITHUB_TOKEN` unless Homebrew cross-repository publishing later requires a dedicated token in Story 4.4.
  - [x] Do not store plaintext tokens in repository files.

- [x] Ensure release notes carry the user trust contract. (AC: 3, 4)
  - [x] Add or document a release-notes template/source that names the version, supported OSes, executable `leetcode`, checksum verification, unofficial status, public-data dependency, and no credential/session/config storage.
  - [x] State that deferred features remain out of scope.
  - [x] State any breaking command behavior changes before publication.

- [x] Add workflow tests. (AC: 1, 2, 3, 4)
  - [x] Add or extend tests that inspect `.github/workflows/release.yml`.
  - [x] Assert tag trigger, test/build steps, GoReleaser action, full checkout history, and release permissions.
  - [x] Assert release docs/notes mention checksums and v1 trust boundaries.

- [x] Verify. (AC: 1, 2)
  - [x] Run `go test ./...`.
  - [x] Run `go build ./...`.
  - [x] If GoReleaser is available, run a non-publishing local validation.

## Dev Notes

### Previous Story Context

- Story 4.1 establishes provenance, versioning, and module-path decision.
- Story 4.2 adds `.goreleaser.yaml` and confirms non-publishing artifact/checksum generation.
- This story wires tagged releases to GitHub Releases; it should not add Homebrew tap publishing unless Story 4.4 scope is being implemented at the same time.

### Architecture Guardrails

- Publishing must happen only after test/build validation succeeds.
- Release artifacts must preserve executable name `leetcode`.
- Windows and Linux users install checksummed binaries from GitHub Releases.
- macOS artifacts may also be present for Homebrew consumption, but Homebrew tap publishing is Story 4.4.
- Keep no-auth, no-session, no-config, no-private-data, and human-readable-only scope intact.

### Latest Technical Notes

- GoReleaser's GitHub Actions docs show `goreleaser/goreleaser-action@v7`.
- GoReleaser requires checkout with `fetch-depth: 0` so it can read full Git history and tags.
- GitHub recommends least-privilege `GITHUB_TOKEN` permissions.
- GitHub release events can be filtered by activity type, but this story should publish from version tags unless Story 4.1 chooses a different convention.

### Expected Files

- Add: `.github/workflows/release.yml`
- Add or update: workflow-content tests under `cmd/`
- Add or update: release notes template/doc if needed, such as `docs/release.md`
- Update: `README.md` or `docs/installation.md` only if release asset naming becomes concrete in this story

### Test Guidance

- Static workflow tests should follow the existing `cmd/ci_test.go` style.
- Do not publish a real release during local tests.
- Do not require live LeetCode requests.
- Run `go test ./...` and `go build ./...`.

### References

- [Source: `_bmad-output/planning-artifacts/epics.md` - Story 4.3]
- [Source: `_bmad-output/planning-artifacts/prds/prd-LeetcodeCLI-2026-06-11/prd.md` - Versioning, Distribution, Trust]
- [Source: `_bmad-output/planning-artifacts/architecture.md` - release packaging and CI guidance]
- [External: `https://goreleaser.com/customization/ci/actions/` - GoReleaser GitHub Actions]
- [External: `https://docs.github.com/en/actions/tutorials/authenticate-with-github_token` - GITHUB_TOKEN permissions]
- [External: `https://docs.github.com/en/actions/reference/workflows-and-actions/events-that-trigger-workflows` - release events]

## Dev Agent Record

### Agent Model Used

GPT-5 Codex

### Debug Log References

- `go test ./cmd` initially failed as expected because `.github/workflows/release.yml` and `docs/release-notes-template.md` did not exist.
- `gofmt -w cmd\release_config_test.go cmd\release_workflow_test.go` completed.
- `go test ./...` passed with workspace-local `GOCACHE`.
- `go build ./...` passed with workspace-local `GOCACHE`.
- `goreleaser --version` failed because GoReleaser is not installed on PATH.

### Completion Notes List

- Added `.github/workflows/release.yml` to publish only from version tag pushes matching `v*.*.*`, with full checkout history, Go version from `go.mod`, `go test ./...`, `go build ./...`, and `goreleaser/goreleaser-action@v7`.
- Scoped workflow permissions to `contents: write` and used `GITHUB_TOKEN` without committing plaintext credentials.
- Added `docs/release-notes-template.md` and GoReleaser release header/footer content covering version, supported OSes, executable `leetcode`, checksum verification, breaking command behavior disclosure, unofficial status, public-data dependency, and v1 deferred features.
- Added static workflow/release-notes tests that enforce tag trigger, validation-before-publish ordering, GoReleaser action configuration, release permissions, checksums, and trust boundaries.
- No GitHub Release was published, no tag was pushed, and no live credentials were used. Local GoReleaser validation remains deferred because GoReleaser is unavailable on PATH.

### File List

- .github/workflows/release.yml
- .goreleaser.yaml
- docs/release-notes-template.md
- cmd/release_config_test.go
- cmd/release_workflow_test.go
- _bmad-output/implementation-artifacts/sprint-status.yaml
- _bmad-output/implementation-artifacts/4-3-publish-github-releases-from-tagged-builds.md

## Change Log

- 2026-06-12: Created Story 4.3 ready for development.
- 2026-06-12: Added tag-driven GitHub Release workflow, release-notes trust contract, GoReleaser release notes config, and static workflow tests.
