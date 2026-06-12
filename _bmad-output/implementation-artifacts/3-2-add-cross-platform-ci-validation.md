---
baseline_commit: NO_VCS
---

# Story 3.2: Add Cross-Platform CI Validation

Status: review

Completion Note: Ultimate context engine analysis completed - comprehensive developer guide created.

## Story

As a developer user,
I want automated validation on supported operating systems,
so that LeetcodeCLI has release confidence beyond a single local machine.

## Requirements Covered

- FR1: Users can install and run LeetcodeCLI on Windows, macOS, and Linux.
- FR14: Users can evaluate whether LeetcodeCLI is safe and appropriate for their local environment.

## Acceptance Criteria

1. Given the project supports Windows, macOS, and Linux, when CI runs, then GitHub Actions validates the project on Windows, macOS, and Linux runners, and each runner executes `go test ./...` and `go build ./...`.
2. Given unit tests run in CI, when the test suite executes, then tests do not call the live LeetCode endpoint, and GraphQL behavior remains covered by fixtures, fake servers, fake transports, or injected fetchers.
3. Given terminal output tests run across platforms, when golden output is compared, then line endings are normalized, and rendering remains deterministic across Windows, macOS, and Linux.

## Tasks / Subtasks

- [x] Add GitHub Actions workflow. (AC: 1)
  - [x] Create `.github/workflows/ci.yml`.
  - [x] Trigger on push and pull request unless project conventions say otherwise.
  - [x] Use a matrix for `ubuntu-latest`, `macos-latest`, and `windows-latest`.
  - [x] Use `actions/checkout` and `actions/setup-go`.
  - [x] Run `go test ./...` on every runner.
  - [x] Run `go build ./...` on every runner.

- [x] Ensure tests are CI-safe. (AC: 2, 3)
  - [x] Confirm unit tests do not call `https://leetcode.com/graphql`.
  - [x] Confirm LeetCode behavior is covered by fixtures, fake servers, fake transports, or injected fetchers.
  - [x] Confirm renderer golden tests normalize CRLF and LF.
  - [x] Confirm tests do not require a real terminal width.

- [x] Add minimal workflow permissions and caching. (AC: 1)
  - [x] Set `permissions: contents: read`.
  - [x] Use setup-go caching if compatible with the project.
  - [x] Use `go-version-file: go.mod` or a quoted explicit Go version consistent with `go.mod`.

- [x] Run local verification. (AC: 1, 2, 3)
  - [x] Run `go test ./...`.
  - [x] Run `go build ./...`.
  - [x] If possible, inspect workflow YAML for syntax issues.

## Dev Notes

### Previous Story Context

- Story 3.1 documents usage and trust boundaries.
- By this story, the codebase should already have command, client, normalization, rendering, and docs tests.
- CI should validate the already-implemented test suite without adding live network dependencies.

### Architecture Guardrails

- CI validates Windows, macOS, and Linux.
- CI runs `go test ./...` and `go build ./...`.
- No live LeetCode calls in unit tests.
- Rendering tests must normalize line endings to be stable across OSes.
- Release packaging remains deferred to Story 3.3 and later distribution work; do not add GoReleaser config in this story unless deliberately scoped.

### Latest Technical Notes

- GitHub Actions workflow matrix supports multiple OS runners using `jobs.<job_id>.strategy.matrix`.
- As of 2026-06-12, `actions/setup-go` documents `actions/setup-go@v6`, `go-version-file: go.mod`, caching, and `permissions: contents: read`.

### Expected Files

- Add: `.github/workflows/ci.yml`
- Possible updates: tests that are not CI-safe across Windows/macOS/Linux.
- No `.goreleaser.yaml` expected in this story.

### Example Workflow Shape

```yaml
name: CI

on:
  push:
  pull_request:

permissions:
  contents: read

jobs:
  test:
    strategy:
      matrix:
        os: [ubuntu-latest, macos-latest, windows-latest]
    runs-on: ${{ matrix.os }}
    steps:
      - uses: actions/checkout@v6
      - uses: actions/setup-go@v6
        with:
          go-version-file: go.mod
      - run: go test ./...
      - run: go build ./...
```

Adjust action major versions if the project standardizes on a different supported version before implementation.

### References

- [Source: `_bmad-output/planning-artifacts/epics.md` - Story 3.2]
- [Source: `_bmad-output/planning-artifacts/architecture.md` - Infrastructure & Deployment, Cross-Platform Build and Validation]
- [Source: `_bmad-output/planning-artifacts/prds/prd-LeetcodeCLI-2026-06-11/prd.md` - FR1, FR14]
- [External: `https://docs.github.com/en/actions/reference/workflows-and-actions/workflow-syntax` - matrix syntax]
- [External: `https://github.com/actions/setup-go` - setup-go usage]
- [Previous: `_bmad-output/implementation-artifacts/3-1-document-usage-scope-and-trust-boundaries.md`]

## Dev Agent Record

### Agent Model Used

GPT-5 Codex

### Debug Log References

- 2026-06-12: Planned a GitHub Actions workflow with OS matrix, setup-go from `go.mod`, read-only contents permission, and test/build steps.
- 2026-06-12: Red-phase `go test ./...` failed because `.github/workflows/ci.yml` did not exist.
- 2026-06-12: Test-source scan found no hard-coded `https://leetcode.com/graphql` live endpoint in `*_test.go` files.
- 2026-06-12: Verification passed: `go test ./...` and `go build ./...`.

### Completion Notes List

- Added GitHub Actions CI workflow for Ubuntu, macOS, and Windows runners.
- Workflow uses read-only contents permission, `actions/checkout@v6`, `actions/setup-go@v6`, `go-version-file: go.mod`, setup-go caching, `go test ./...`, and `go build ./...`.
- Added a workflow-content test to guard required CI matrix, permissions, setup, test, and build steps.
- Confirmed tests use fixtures/fakes/injected dependencies and renderer tests normalize line endings without real terminal width.

### File List

- `_bmad-output/implementation-artifacts/3-2-add-cross-platform-ci-validation.md`
- `.github/workflows/ci.yml`
- `cmd/ci_test.go`

## Change Log

- 2026-06-12: Added cross-platform GitHub Actions CI and marked Story 3.2 ready for review.
