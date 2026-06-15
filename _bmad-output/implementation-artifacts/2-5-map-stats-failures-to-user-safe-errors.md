---
baseline_commit: NO_VCS
---

# Story 2.5: Map Stats Failures to User-Safe Errors

Status: review

Completion Note: Ultimate context engine analysis completed - comprehensive developer guide created.

## Story

As a developer user,
I want profile, API, data, and rendering failures to produce clear messages,
so that I can understand what went wrong without seeing misleading partial stats.

## Requirements Covered

- FR5: LeetcodeCLI only displays stats it can retrieve.
- FR9: Additional practice stats guardrail.
- FR12: Human-output-only v1.

## Acceptance Criteria

1. Given LeetCode returns no matching user, when the stats command handles the failure, then it writes `No LeetCode profile found for "<username>". Check the username and try again.` to stderr, and it exits `1`.
2. Given public profile or stats data is unavailable, when the stats command handles the failure, then it writes `Stats for "<username>" are not available from LeetCode right now. Try again later.` to stderr, and it exits `1`.
3. Given the public API or network cannot be reached, when the stats command handles the failure, then it writes `Could not reach LeetCode. Check your connection and try again.` to stderr, and it exits `1`.
4. Given LeetCode rate-limits or blocks the request, when the stats command handles the failure, then it writes `LeetCode blocked or rate-limited the request. Try again later.` to stderr, and it exits `1`.
5. Given mandatory stats are missing, when the stats command handles the failure, then it writes `LeetCode did not return required stats for "<username>". Try again later.` to stderr, and it exits `1`.
6. Given stats output cannot be rendered, when the stats command handles the failure, then it writes `Could not render stats output. Try again later.` to stderr, and it exits `1`.
7. Given any stats failure occurs after a username is supplied, when stdout and stderr are inspected, then failure messages are written to stderr, and no fabricated, partial, JSON, CSV, or machine-readable success output is written to stdout.

## Tasks / Subtasks

- [x] Finalize error taxonomy. (AC: 1, 2, 3, 4, 5, 6)
  - [x] Ensure `internal/leetcode` exposes classified errors without leaking raw HTTP/GraphQL details into `cmd`.
  - [x] Ensure mandatory stats failures are distinguishable from unavailable profile/stats failures where required.
  - [x] Ensure renderer failures are distinguishable at the command boundary.

- [x] Map failures to exact stderr copy. (AC: 1, 2, 3, 4, 5, 6, 7)
  - [x] Map username not found to `No LeetCode profile found for "<username>". Check the username and try again.`
  - [x] Map unavailable public profile/stats to `Stats for "<username>" are not available from LeetCode right now. Try again later.`
  - [x] Map network/API reachability failures to `Could not reach LeetCode. Check your connection and try again.`
  - [x] Map 403/429 access/rate limiting to `LeetCode blocked or rate-limited the request. Try again later.`
  - [x] Map missing mandatory stats to `LeetCode did not return required stats for "<username>". Try again later.`
  - [x] Map render failures to `Could not render stats output. Try again later.`
  - [x] Return exit code `1` for these runtime failures.

- [x] Preserve output boundaries. (AC: 7)
  - [x] On failure, write no successful stats table to stdout.
  - [x] Do not print raw GraphQL errors, stack traces, tokens, cookies, or debug details.
  - [x] Do not produce JSON/CSV/machine-readable failure output.
  - [x] Keep missing username behavior from Story 1.3 as exit code `2`, not `1`.

- [x] Add command failure tests. (AC: 1, 2, 3, 4, 5, 6, 7)
  - [x] Use fake fetchers/renderers to trigger each error category.
  - [x] Assert exact stderr copy and exit code.
  - [x] Assert stdout is empty for each failure.
  - [x] Assert no live network call is required.
  - [x] Keep happy-path tests from Story 2.4 passing.

- [x] Verify. (AC: 1, 2, 3, 4, 5, 6, 7)
  - [x] Run `go test ./...`.
  - [x] Run `go build ./...`.

## Dev Notes

### Previous Story Context

- Story 2.4 wires the success path for `leetcode stats <username>`.
- Story 2.5 completes the runtime failure behavior and should leave Epic 2 ready for docs and release-readiness work.

### Architecture Guardrails

- User-facing copy is emitted at the command boundary.
- Internal packages should return classified errors with enough detail for mapping, not preformatted CLI messages.
- Never render partial success when mandatory data is missing.
- Avoid exposing GraphQL response details, HTTP internals, or implementation stack traces to users.
- Successful output writes to stdout; failures write to stderr.
- Runtime/API/data/rendering failures exit `1`; usage failures exit `2`; explicit help exits `0`.

### Expected Files

- Update: `cmd/stats.go`
- Update: `cmd/stats_test.go`
- Possible update: `internal/leetcode/errors.go` or equivalent if a dedicated taxonomy is clearer.
- Possible update: `internal/render` only to return a classified rendering error.

### Test Guidance

- Prefer table-driven command tests for error mapping.
- Keep exact message strings centralized in tests or constants to prevent drift.
- Verify missing username still uses the Story 1.3 exact copy and exit code `2`.
- Verify every failure case keeps stdout empty.

### References

- [Source: `_bmad-output/planning-artifacts/epics.md` - Story 2.5]
- [Source: `_bmad-output/planning-artifacts/architecture.md` - Error Handling Patterns, API & Communication Patterns]
- [Source: `_bmad-output/planning-artifacts/prds/prd-LeetcodeCLI-2026-06-11/prd.md` - standard v1 error copy and exit behavior]
- [Source: `_bmad-output/planning-artifacts/prds/prd-LeetcodeCLI-2026-06-11/addendum.md` - downstream architecture notes and standard error copy]
- [Previous: `_bmad-output/implementation-artifacts/2-4-wire-leetcode-stats-username-happy-path.md`]

## Dev Agent Record

### Agent Model Used

GPT-5 Codex

### Debug Log References

- 2026-06-12: Planned command-boundary mapping from `internal/leetcode` classified errors and render-call failures to exact user-safe stderr messages.
- 2026-06-12: Red-phase `go test ./...` failed because command failures printed raw internal errors.
- 2026-06-12: Verification passed: `gofmt`, `go test ./...`, and `go build ./...`.

### Completion Notes List

- Added command-boundary safe error mapping for profile not found, unavailable stats, endpoint/network failure, rate limiting/access blocked, missing mandatory stats, and render/write failures.
- Preserved Story 1.3 usage behavior: missing username remains exit code `2` with exact stderr copy.
- Added table-driven failure tests using fake fetch/render functions, asserting exact stderr, exit code `1`, empty stdout, and no render call after fetch failures.
- Kept happy-path command tests passing.

### File List

- `_bmad-output/implementation-artifacts/2-5-map-stats-failures-to-user-safe-errors.md`
- `cmd/stats.go`
- `cmd/stats_test.go`

## Change Log

- 2026-06-12: Added user-safe stats failure mapping and marked Story 2.5 ready for review.
