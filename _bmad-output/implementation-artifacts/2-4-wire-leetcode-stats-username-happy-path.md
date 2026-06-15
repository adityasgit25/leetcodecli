---
baseline_commit: NO_VCS
---

# Story 2.4: Wire `leetcode stats <username>` Happy Path

Status: review

Completion Note: Ultimate context engine analysis completed - comprehensive developer guide created.

## Story

As a developer user,
I want `leetcode stats <username>` to render a complete Stats View for a valid public profile,
so that I can check public LeetCode progress without opening the website.

## Requirements Covered

- FR3: Users can run `leetcode stats <username>` to view a Public Profile.
- FR5: LeetcodeCLI only displays stats it can retrieve.
- FR6: Profile Summary.
- FR7: Total Solved Count.
- FR8: Language Breakdown.
- FR10: Pretty Terminal Table output.
- FR12: Human-output-only v1.

## Acceptance Criteria

1. Given the CLI receives `leetcode stats <username>`, when the supplied username resolves to valid public stats, then the command fetches public profile data, validates mandatory stats, renders the Stats View, writes it to stdout, and exits `0`.
2. Given a successful Stats View is rendered, when the output is inspected, then it identifies the requested or resolved LeetCode profile, and it includes Profile Summary, Total Solved Count, and Language Breakdown.
3. Given the stats command is wired to data retrieval and rendering, when command tests run, then they use injected fetcher and width behavior, and they do not require a live LeetCode request or a real terminal.
4. Given the command has a successful username-based workflow, when help examples and command metadata are inspected, then they still describe `leetcode stats <username>` as the supported v1 happy path, and they do not imply own-profile default lookup or authenticated private-data access.

## Tasks / Subtasks

- [x] Define command-level seams for dependencies. (AC: 1, 3)
  - [x] Add interfaces or function types for fetching normalized profile stats.
  - [x] Add an injectable width detector or renderer function.
  - [x] Keep real dependency wiring at the command construction boundary.
  - [x] Avoid global mutable state for current username, HTTP clients, width, stdout, or stderr.

- [x] Wire successful stats flow. (AC: 1, 2)
  - [x] Validate exactly one username using the existing Epic 1 behavior.
  - [x] Call `internal/leetcode` to fetch/normalize/validate public profile stats.
  - [x] Call `internal/render` to produce the Stats View.
  - [x] Write successful output to stdout.
  - [x] Return exit code `0` for successful rendering.

- [x] Preserve command help and scope. (AC: 4)
  - [x] Keep examples as `leetcode stats <username>`.
  - [x] Do not imply bare `leetcode stats` succeeds.
  - [x] Do not imply own-profile defaults, login, private data, session use, tokens, config, JSON, CSV, dashboards, or recommendations.

- [x] Add command happy-path tests. (AC: 1, 2, 3, 4)
  - [x] Use fake fetcher returning normalized stats.
  - [x] Use fake width behavior or fake renderer so no real terminal is required.
  - [x] Assert stdout contains Profile Summary, Total Solved Count, and Language Breakdown.
  - [x] Assert stderr is empty on success.
  - [x] Assert exit code `0` on success.
  - [x] Assert no live LeetCode request is made in command tests.

- [x] Verify. (AC: 1, 2, 3, 4)
  - [x] Run `go test ./...`.
  - [x] Run `go build ./...`.

## Dev Notes

### Previous Story Context

- Story 2.1 implemented LeetCode GraphQL retrieval and API error classification.
- Story 2.2 normalized and validated mandatory stats.
- Story 2.3 rendered normalized stats into terminal tables.
- This story connects those pieces at the command boundary for the success path only. Failure copy is completed in Story 2.5.

### Architecture Guardrails

- `cmd` wires dependencies, validates args, maps errors, and routes stdout/stderr.
- `cmd` must not parse raw GraphQL JSON or build HTTP requests directly.
- `internal/leetcode` owns endpoint details and normalized stats retrieval.
- `internal/render` owns table layout and width behavior.
- Successful output goes to stdout.
- No JSON/CSV/machine-readable output flags.
- Unit tests do not call live LeetCode or depend on real terminal width.

### Expected Files

- Update: `cmd/stats.go`
- Update: `cmd/stats_test.go`
- Possible update: `cmd/root.go` if dependency injection requires command construction changes.
- Reuse: `internal/leetcode`, `internal/render`

### Test Guidance

- Command tests should use fakes at command seams.
- Keep client and renderer behavior covered in their own package tests.
- Include at least one command-level test proving requested/resolved username is visible in success output.
- Keep missing-username tests from Story 1.3 passing unchanged.

### References

- [Source: `_bmad-output/planning-artifacts/epics.md` - Story 2.4]
- [Source: `_bmad-output/planning-artifacts/architecture.md` - Integration Points, Error Handling Patterns, Component Boundaries]
- [Source: `_bmad-output/planning-artifacts/prds/prd-LeetcodeCLI-2026-06-11/prd.md` - FR3, FR5, FR6, FR7, FR8, FR10, FR12]
- [Previous: `_bmad-output/implementation-artifacts/2-3-render-human-readable-stats-tables.md`]

## Dev Agent Record

### Agent Model Used

GPT-5 Codex

### Debug Log References

- 2026-06-12: Planned command-level seams for fetching normalized stats and rendering output, with real dependencies wired only in `NewStatsCommand`.
- 2026-06-12: Red-phase `go test ./...` failed because `statsCommandConfig`, `newRootCommand`, and injectable run helpers were undefined.
- 2026-06-12: Verification passed: `gofmt`, `go test ./...`, and `go build ./...`.

### Completion Notes List

- Added command-level fetch/render seams using function types with real dependency wiring in `NewStatsCommand`.
- Wired the production stats command to `internal/leetcode.Client.FetchProfileStats` and `internal/render.RenderStatsWithWidthDetector`.
- Preserved username validation, help text, and no-auth/no-export scope.
- Added happy-path command tests with fake fetcher/renderer dependencies, asserting stdout sections, empty stderr, and exit code `0`.

### File List

- `_bmad-output/implementation-artifacts/2-4-wire-leetcode-stats-username-happy-path.md`
- `cmd/root.go`
- `cmd/stats.go`
- `cmd/stats_test.go`

## Change Log

- 2026-06-12: Wired the successful `leetcode stats <username>` command path and marked Story 2.4 ready for review.
