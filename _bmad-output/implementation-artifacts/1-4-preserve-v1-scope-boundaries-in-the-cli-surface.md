---
baseline_commit: NO_VCS
---

# Story 1.4: Preserve v1 Scope Boundaries in the CLI Surface

Status: review

Completion Note: Ultimate context engine analysis completed - comprehensive developer guide created.

## Story

As a developer user,
I want the CLI surface to show only the supported v1 stats workflow,
so that I am not misled into expecting authentication, configuration, structured exports, or planning features.

## Requirements Covered

- FR12: Users receive human-readable output only; v1 does not expose JSON, CSV, or machine-readable export flags.

## Acceptance Criteria

1. Given the root and stats commands are implemented, when help text, command descriptions, examples, and available flags are reviewed, then they describe only the username-based public stats workflow, and they expose no login, logout, session, token, config, JSON, CSV, dashboard, recommendation, goal, reminder, topic-gap, browser extension, web app, or TUI workflow.
2. Given v1 scope is intentionally public-profile-only, when command tests inspect the available command tree, then unsupported auth/session/config/export commands are absent, and the stats command accepts exactly one required username argument for successful execution.
3. Given the CLI command surface is complete for Epic 1, when `go test ./...` is run, then command tests verify help output, missing-username behavior, stdout/stderr routing, and v1 scope boundaries.

## Tasks / Subtasks

- [x] Audit the command tree. (AC: 1, 2)
  - [x] Inspect root command, subcommands, flags, descriptions, examples, and aliases.
  - [x] Confirm only the root/help surface and `stats <username>` are exposed for v1.
  - [x] Confirm no unsupported command or flag names exist.

- [x] Tighten stats argument behavior. (AC: 2)
  - [x] Ensure stats requires exactly one username for successful execution.
  - [x] Ensure missing username remains the Story 1.3 exact stderr and exit code `2`.
  - [x] Ensure extra positional args are rejected as usage errors.

- [x] Harden tests against scope drift. (AC: 1, 2, 3)
  - [x] Add tests that list root subcommands and assert forbidden commands are absent.
  - [x] Add tests that inspect persistent/local flags and assert no JSON/CSV/config/token/session flags exist.
  - [x] Add tests for forbidden text in help output.
  - [x] Add tests for stats accepting exactly one username.

- [x] Preserve stdout/stderr rules. (AC: 3)
  - [x] Ensure help/usage behavior remains deterministic in tests.
  - [x] Ensure usage failures do not emit successful stats output to stdout.
  - [x] Keep failure messages on stderr.

- [x] Verify Epic 1 completion. (AC: 3)
  - [x] Run `go test ./...`.
  - [x] Run `go build ./...`.

## Dev Notes

### Previous Story Context

- Story 1.1 initializes the Go/Cobra scaffold.
- Story 1.2 exposes root and stats help.
- Story 1.3 enforces missing-username behavior with exact stderr, exit code `2`, and no network call.
- This story is an Epic 1 hardening pass. It should not introduce Epic 2 retrieval/rendering behavior.

### Architecture Guardrails

- Preserve `leetcode stats <username>` as the only successful stats command.
- Do not add auth, session, config, token, browser-login, logout, cache, JSON, CSV, dashboard, recommendation, goal, reminder, topic-gap, browser extension, web app, or TUI features.
- Do not add Viper.
- Do not introduce global mutable state for username, HTTP clients, terminal width, or command output.
- Command tests should instantiate commands with injected writers and dependencies.

### Expected Files

- Update: `cmd/root.go`, `cmd/stats.go`
- Update tests: `cmd/root_test.go`, `cmd/stats_test.go`
- No new `internal/` packages expected unless implementation already introduced a small testability helper.

### Test Guidance

- Build a forbidden terms list once in tests and apply it to root help, stats help, commands, and flags.
- Keep assertions precise enough to prevent accidental future `--json`, `--csv`, `login`, `logout`, or `config` surfaces.
- Continue to normalize CRLF/LF in command output comparisons.

### References

- [Source: `_bmad-output/planning-artifacts/epics.md` - Story 1.4]
- [Source: `_bmad-output/planning-artifacts/architecture.md` - Enforcement Guidelines, Anti-Patterns, Documentation Constraints]
- [Source: `_bmad-output/planning-artifacts/prds/prd-LeetcodeCLI-2026-06-11/prd.md` - FR12, Non-Goals, command contract]
- [Previous: `_bmad-output/implementation-artifacts/1-3-enforce-username-required-stats-usage.md`]

## Dev Agent Record

### Agent Model Used

GPT-5 Codex

### Debug Log References

- 2026-06-12: Audited current command tree after Stories 1.1-1.3; root exposes `stats` with Cobra help behavior and no completion command.
- 2026-06-12: Scope hardening tests passed immediately against the current command implementation.
- 2026-06-12: Verification passed: `gofmt`, `go test ./...`, `go build ./...`, and `go mod tidy`.

### Completion Notes List

- Added command tree tests that assert the root exposes only the supported `stats` command.
- Added flag and help surface tests to prevent unsupported auth/session/config/export/planning terminology from appearing.
- Added stats argument boundary tests for exactly one username and extra positional arguments.
- No production command changes were required because Stories 1.2 and 1.3 already kept the v1 surface narrow.

### File List

- `_bmad-output/implementation-artifacts/1-4-preserve-v1-scope-boundaries-in-the-cli-surface.md`
- `cmd/root_test.go`
- `cmd/stats_test.go`
- `go.mod`

## Change Log

- 2026-06-12: Added Epic 1 command surface guardrail tests and marked Story 1.4 ready for review.
