---
baseline_commit: NO_VCS
---

# Story 1.3: Enforce Username-Required Stats Usage

Status: review

Completion Note: Ultimate context engine analysis completed - comprehensive developer guide created.

## Story

As a developer user,
I want `leetcode stats` to fail locally with clear usage guidance when I omit the username,
so that I understand how to correct the command without triggering an unnecessary network request.

## Requirements Covered

- FR4: Users who run `leetcode stats` without a username receive usage guidance instead of a network request.

## Acceptance Criteria

1. Given the stats command is available, when a user runs `leetcode stats` without a username, then the CLI exits with usage exit code `2`, and it writes exactly `Username required. Usage: leetcode stats <username>. Run "leetcode help" for help.` to stderr.
2. Given a username is missing, when `leetcode stats` handles the usage error, then it performs no LeetCode GraphQL request, and command tests prove no fetcher, HTTP client, or network dependency is invoked.
3. Given the usage error is emitted, when stdout and stderr are inspected in tests, then no successful stats output appears on stdout, and the failure message remains concise and human-readable.

## Tasks / Subtasks

- [x] Build on Story 1.2's root and stats command help. (AC: 1, 2, 3)
  - [x] Verify `cmd/stats.go` exists.
  - [x] Keep existing help output and command names intact.

- [x] Add exact missing-username behavior. (AC: 1, 3)
  - [x] Configure stats command argument validation for exactly one required username.
  - [x] Ensure missing username returns/propagates a usage error that maps to exit code `2`.
  - [x] Write exactly `Username required. Usage: leetcode stats <username>. Run "leetcode help" for help.` to stderr.
  - [x] Do not include extra Cobra usage text unless tests account for it and the exact required line remains isolated as the failure message.

- [x] Preserve no-network behavior. (AC: 2)
  - [x] If a fetcher interface or placeholder dependency exists, ensure it is not invoked when username is missing.
  - [x] If no fetcher exists yet, keep it that way. Do not add LeetCode client code for this story.
  - [x] Avoid any HTTP package dependency in `cmd/stats.go` except test fakes if already introduced.

- [x] Add command tests. (AC: 1, 2, 3)
  - [x] Test missing username returns usage exit code `2` through the command execution path.
  - [x] Test stderr equals the exact required copy, with only expected line ending normalization.
  - [x] Test stdout is empty for the missing-username path.
  - [x] Test fake fetcher or network dependency call count remains zero.

- [x] Verify. (AC: 1, 2, 3)
  - [x] Run `go test ./...`.
  - [x] Run `go build ./...`.

## Dev Notes

### Previous Story Context

- Story 1.1 provides the Go/Cobra scaffold.
- Story 1.2 exposes help for root and stats commands.
- This story completes the local usage-error path before any public LeetCode retrieval is implemented.

### Architecture Guardrails

- Missing username is a local command validation error, not an API/data/rendering error.
- Exit code policy: usage/argument failures exit `2`.
- Failure messages write to stderr. Successful stats output writes to stdout, but this story must not produce any success output.
- Do not call `https://leetcode.com/graphql` or introduce network work.
- Keep user-facing copy at the command boundary.
- Keep command files thin; do not add raw GraphQL DTOs or rendering logic.

### Expected Files

- Update: `cmd/stats.go`
- Update/add tests: `cmd/stats_test.go`
- Possible update: `cmd/root.go` only if command construction needs a testable execution helper.
- No `internal/leetcode` or `internal/render` changes expected.

### Test Guidance

- Prefer a command execution helper that returns an application exit code instead of calling `os.Exit`.
- In tests, pass buffers for stdout and stderr.
- If an application-level `Execute()` function exits, separate it from a testable `ExecuteCommand()` or `Run()` helper.
- Add a fake fetcher only if it helps prove no network path is invoked. Do not build the real client in this story.

### References

- [Source: `_bmad-output/planning-artifacts/epics.md` - Story 1.3]
- [Source: `_bmad-output/planning-artifacts/architecture.md` - API & Communication Patterns, Error Handling Patterns, Good command behavior]
- [Source: `_bmad-output/planning-artifacts/prds/prd-LeetcodeCLI-2026-06-11/prd.md` - FR4, standard v1 error copy, exit behavior]
- [Source: `_bmad-output/planning-artifacts/prds/prd-LeetcodeCLI-2026-06-11/addendum.md` - standard v1 error copy]
- [Previous: `_bmad-output/implementation-artifacts/1-2-expose-root-and-stats-help.md`]

## Dev Agent Record

### Agent Model Used

GPT-5 Codex

### Debug Log References

- 2026-06-12: Confirmed `cmd/stats.go` exists and Story 1.2 help behavior is present before editing.
- 2026-06-12: Red-phase `go test ./...` failed because `Run` and testable `newStatsCommand` did not exist.
- 2026-06-12: Verification passed: `gofmt`, `go test ./...`, and `go build ./...`.

### Completion Notes List

- Added a testable `Run(args, stdout, stderr)` execution path that maps command usage errors to exit code `2`.
- Preserved `main.go` by keeping `cmd.Execute()` as the outer process-exit boundary.
- Added exact missing-username validation for `leetcode stats` with the required stderr copy and empty stdout.
- Added a stats runner hook and tests proving the runner is not invoked when username validation fails; no LeetCode client or HTTP dependency was introduced.

### File List

- `_bmad-output/implementation-artifacts/1-3-enforce-username-required-stats-usage.md`
- `cmd/command_test.go`
- `cmd/root.go`
- `cmd/stats.go`
- `cmd/stats_test.go`

## Change Log

- 2026-06-12: Added username-required stats usage handling and marked Story 1.3 ready for review.
