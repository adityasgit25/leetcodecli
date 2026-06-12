---
baseline_commit: NO_VCS
---

# Story 1.2: Expose Root and Stats Help

Status: review

Completion Note: Ultimate context engine analysis completed - comprehensive developer guide created.

## Story

As a developer user,
I want clear help for the CLI and stats command,
so that I can discover how to run `leetcode stats <username>` without reading source code.

## Requirements Covered

- FR2: Users can discover the Stats Command and basic usage through CLI help.
- FR12: Users receive human-readable output only; v1 does not expose JSON, CSV, or machine-readable export flags.

## Acceptance Criteria

1. Given the CLI has a root command, when a user runs `leetcode help`, then help output lists `stats` as an available command, and it does not mention login, logout, session setup, token setup, config setup, JSON output, CSV output, dashboards, recommendations, goals, or reminders.
2. Given the stats command exists, when a user runs `leetcode help stats`, then help output explains `leetcode stats <username>`, and it states that v1 requires a username.
3. Given standard Cobra help behavior is available, when a user runs `leetcode --help`, then the CLI displays standard root help without crashing, and the visible command examples consistently use `leetcode`.

## Tasks / Subtasks

- [x] Build on the Story 1.1 Cobra scaffold. (AC: 1, 2, 3)
  - [x] Verify `go.mod`, `main.go`, and `cmd/root.go` exist before editing.
  - [x] If Story 1.1 did not add a `stats` command, add a thin `cmd/stats.go` with help metadata only.
  - [x] Keep `main.go` as a delegation-only entry point.

- [x] Implement root help content. (AC: 1, 3)
  - [x] Ensure root command `Use` is `leetcode`.
  - [x] Include `stats` as an available command.
  - [x] Use examples and descriptions that consistently show `leetcode`, not `leetcodecli`.
  - [x] Keep output human-readable and do not add structured output flags.

- [x] Implement stats help content. (AC: 2)
  - [x] Ensure stats command `Use` is `stats <username>`.
  - [x] Explain that v1 requires a username.
  - [x] Show `leetcode stats <username>` as the supported successful shape.
  - [x] Do not implement LeetCode retrieval yet; help text only is sufficient for this story.

- [x] Protect no-auth/no-export scope in help. (AC: 1, 2, 3)
  - [x] Do not mention login, logout, session setup, token setup, config setup, JSON, CSV, dashboards, recommendations, goals, reminders, topic-gap analysis, browser extensions, web apps, or TUI flows.
  - [x] Do not add Viper or config packages.
  - [x] Do not add flags for JSON, CSV, config, token, or session behavior.

- [x] Add command help tests. (AC: 1, 2, 3)
  - [x] Test `leetcode help` includes `stats`.
  - [x] Test `leetcode help stats` includes `leetcode stats <username>` and username-required language.
  - [x] Test `leetcode --help` succeeds and uses `leetcode`.
  - [x] Test forbidden words or command names are absent from help output.
  - [x] Use buffers and direct command execution; do not spawn network calls.

- [x] Verify. (AC: 1, 2, 3)
  - [x] Run `go test ./...`.
  - [x] Run `go build ./...`.

## Dev Notes

### Previous Story Context

- Story 1.1 creates the initial Go module and minimal Cobra scaffold.
- This story should update the command metadata and add the `stats` command help surface if it does not already exist.
- Preserve Story 1.1's thin-command boundary: `cmd` handles command construction, output routing, and argument/help behavior only.

### Architecture Guardrails

- Cobra provides root and subcommand help behavior.
- `cmd` owns help text and argument validation only.
- Keep public command identity as `leetcode`.
- `leetcode stats <username>` is the only successful v1 stats command shape.
- Bare `leetcode stats` behavior is completed in Story 1.3. Do not accidentally make bare stats successful here.
- No LeetCode GraphQL client, rendering package, config surface, auth surface, or persistence is needed for this story.

### Expected Files

- Update: `cmd/root.go`
- Add/update: `cmd/stats.go`
- Update/add tests: `cmd/root_test.go`, `cmd/stats_test.go`
- No changes expected under `internal/` for this story.

### Test Guidance

- Instantiate Cobra commands directly in tests.
- Set command output/error writers to buffers.
- Normalize line endings before comparing help output.
- Assert absence of forbidden surface terms in a case-insensitive way.
- Do not call `os.Exit` from tested code.

### References

- [Source: `_bmad-output/planning-artifacts/epics.md` - Story 1.2]
- [Source: `_bmad-output/planning-artifacts/architecture.md` - CLI Framework, Documentation Constraints, Enforcement Guidelines]
- [Source: `_bmad-output/planning-artifacts/prds/prd-LeetcodeCLI-2026-06-11/prd.md` - FR2, FR12, v1 command contract]
- [Previous: `_bmad-output/implementation-artifacts/1-1-set-up-initial-project-from-go-cobra-starter-template.md`]

## Dev Agent Record

### Agent Model Used

GPT-5 Codex

### Debug Log References

- 2026-06-12: Confirmed Story 1.1 scaffold files exist before editing (`go.mod`, `main.go`, `cmd/root.go`).
- 2026-06-12: Red-phase `go test ./...` failed because root help did not list `stats` and `leetcode help stats` did not expose stats usage.
- 2026-06-12: Initial green build passed, while stats help test caught that Cobra's help command only showed the long description; moved the supported usage shape into the long help text.
- 2026-06-12: Verification passed: `gofmt`, `go test ./...`, and `go build ./...`.

### Completion Notes List

- Added a thin `stats <username>` command with help metadata only.
- Registered `stats` on the root command and disabled Cobra's generated completion command to keep the v1 command surface narrow.
- Added command help tests for root help, stats help, username-required language, and forbidden no-auth/no-export help terms.
- Kept `main.go` unchanged as a delegation-only entry point.

### File List

- `_bmad-output/implementation-artifacts/1-2-expose-root-and-stats-help.md`
- `cmd/command_test.go`
- `cmd/root.go`
- `cmd/root_test.go`
- `cmd/stats.go`
- `cmd/stats_test.go`

## Change Log

- 2026-06-12: Added root/stats help discovery and marked Story 1.2 ready for review.
