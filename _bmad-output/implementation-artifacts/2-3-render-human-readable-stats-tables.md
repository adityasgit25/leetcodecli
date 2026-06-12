---
baseline_commit: NO_VCS
---

# Story 2.3: Render Human-Readable Stats Tables

Status: review

Completion Note: Ultimate context engine analysis completed - comprehensive developer guide created.

## Story

As a developer user,
I want public profile stats rendered as readable terminal tables,
so that I can quickly scan progress from a standard shell.

## Requirements Covered

- FR6: Users can see a Profile Summary in the Stats View.
- FR7: Users can see the Total Solved Count in the Stats View.
- FR8: Users can see solved-question counts by programming language in the Stats View.
- FR9: Additional practice stats guardrail.
- FR10: Users can read the Stats View as pretty Terminal Tables.
- FR11: Users can read essential stats in standard terminal widths, including 80 columns.
- FR12: Users receive human-readable output only.

## Acceptance Criteria

1. Given normalized profile stats are available, when rendering runs, then the output includes visible sections labeled `Profile Summary`, `Total Solved Count`, and `Language Breakdown`, and the renderer does not require raw GraphQL DTOs or HTTP details.
2. Given Profile Summary contains nullable display values, when rendering runs, then missing nullable values display as `N/A`, and table layout remains coherent.
3. Given Language Breakdown is empty but valid, when rendering runs, then the CLI shows an explicit none-found state, and it does not crash or treat the empty list as a mandatory-stats failure.
4. Given terminal width is available, when rendering detects width, then it uses `golang.org/x/term` where appropriate, and falls back to a safe 80-column layout when width cannot be detected.
5. Given long usernames or language names are rendered at 80 columns, when renderer tests compare output, then adjacent columns remain readable, and golden tests normalize CRLF and LF line endings.
6. Given v1 output is human-readable only, when the Stats View is rendered, then it uses pretty terminal table output through `github.com/jedib0t/go-pretty/v6/table`, and it does not produce JSON, CSV, or machine-readable export output.

## Tasks / Subtasks

- [x] Add the render package. (AC: 1, 6)
  - [x] Create `internal/render`.
  - [x] Render from normalized stats only.
  - [x] Do not import raw GraphQL DTO types or HTTP code.
  - [x] Use `github.com/jedib0t/go-pretty/v6/table`.

- [x] Render mandatory sections. (AC: 1, 2, 3)
  - [x] Include visible labels `Profile Summary`, `Total Solved Count`, and `Language Breakdown`.
  - [x] Render username, display name, ranking, reputation, and profile URL in Profile Summary.
  - [x] Render Total Solved Count from normalized stats.
  - [x] Render language rows with language name and solved count.
  - [x] Render missing nullable display values as `N/A`.
  - [x] Render an explicit none-found state for present empty language counts.

- [x] Implement width detection and fallback. (AC: 4, 5)
  - [x] Use `golang.org/x/term` for terminal width detection.
  - [x] Fall back to 80 columns when width cannot be detected or is invalid.
  - [x] Keep width detection injectable/testable.
  - [x] Use go-pretty width controls or deterministic wrapping/trimming to preserve readability at 80 columns.

- [x] Keep output human-readable only. (AC: 6)
  - [x] Do not add JSON, CSV, TSV, Markdown, or HTML output modes.
  - [x] Do not expose renderer output flags.
  - [x] Even though go-pretty supports other formats, use table rendering only.

- [x] Add renderer tests and golden files. (AC: 1, 2, 3, 4, 5, 6)
  - [x] Add golden output for normal stats at 80 columns.
  - [x] Add golden output for long username and language names at 80 columns.
  - [x] Add golden output for empty Language Breakdown.
  - [x] Normalize CRLF/LF before comparison.
  - [x] Test width fallback when detection fails.
  - [x] Test renderer does not depend on raw DTOs or network behavior.

- [x] Verify. (AC: 1, 2, 3, 4, 5, 6)
  - [x] Run `go test ./...`.
  - [x] Run `go build ./...`.

## Dev Notes

### Previous Story Context

- Story 2.2 produces normalized profile stats and validates mandatory data.
- This story builds table rendering only. The final command happy path is wired in Story 2.4.

### Architecture Guardrails

- `internal/render` owns table layout, terminal width handling, and output formatting.
- Renderer must not know HTTP status codes, GraphQL errors, endpoint URLs, or external response shapes.
- Render only normalized stats.
- Missing mandatory data should have failed before renderer; renderer should still handle empty valid languages.
- Output must remain human-readable; no machine-readable export modes.

### Latest Technical Notes

- `github.com/jedib0t/go-pretty/v6/table` provides a table writer via `NewWriter()`, supports `AppendHeader`, `AppendRow`, `SetTitle`, `SetStyle`, and width controls such as row/column width limits.
- `golang.org/x/term.GetSize(fd)` returns visible terminal dimensions and should be wrapped behind a testable width detector.
- As of 2026-06-12, `pkg.go.dev` shows go-pretty v6 package docs and `golang.org/x/term` exposes `GetSize(fd)`.

### Expected Files

- Add: `internal/render/stats.go`
- Add: `internal/render/width.go`
- Add: `internal/render/stats_test.go`
- Add: `internal/render/width_test.go`
- Add: `internal/render/testdata/stats_success_80.golden`
- Add: `internal/render/testdata/stats_long_values_80.golden`
- Add: `internal/render/testdata/stats_empty_languages_80.golden`
- Update `go.mod` and `go.sum` for `go-pretty` and `x/term`.

### Test Guidance

- Keep renderer tests pure: pass normalized stats and explicit widths.
- Normalize line endings before golden comparisons.
- Avoid terminal-dependent tests by injecting width.
- Include long values that would reveal column corruption at 80 columns.
- Do not snapshot ANSI color unless intentionally used; plain stable output is easier to maintain.

### References

- [Source: `_bmad-output/planning-artifacts/epics.md` - Story 2.3]
- [Source: `_bmad-output/planning-artifacts/architecture.md` - Terminal Output mapping, Format Patterns, Testing Framework]
- [Source: `_bmad-output/planning-artifacts/prds/prd-LeetcodeCLI-2026-06-11/prd.md` - FR6-FR12]
- [External: `https://pkg.go.dev/github.com/jedib0t/go-pretty/v6/table` - table writer and width controls]
- [External: `https://pkg.go.dev/golang.org/x/term` - terminal size detection]
- [Previous: `_bmad-output/implementation-artifacts/2-2-normalize-and-validate-mandatory-profile-stats.md`]

## Dev Agent Record

### Agent Model Used

GPT-5 Codex

### Debug Log References

- 2026-06-12: Planned `internal/render` to accept only normalized `leetcode.ProfileStats`, use `go-pretty` table rendering, and wrap terminal width detection behind injectable functions.
- 2026-06-12: Red-phase `go test ./...` failed because renderer APIs and width constants were undefined.
- 2026-06-12: Initial golden run exposed wrapped go-pretty table titles; moved section labels outside table titles for stable readability.
- 2026-06-12: `go mod tidy` resolved `github.com/jedib0t/go-pretty/v6 v6.8.1` and `golang.org/x/term v0.44.0`.
- 2026-06-12: Verification passed: `gofmt`, `go test ./...`, `go build ./...`, and `go mod tidy`.

### Completion Notes List

- Added `internal/render` with normalized-stats-only table rendering through `github.com/jedib0t/go-pretty/v6/table`.
- Rendered `Profile Summary`, `Total Solved Count`, and `Language Breakdown`, including `N/A` values and an explicit empty-language state.
- Added injectable terminal width detection using `golang.org/x/term`, with a safe 80-column fallback.
- Added golden tests for normal, long-value, and empty-language output, plus width fallback tests.
- Kept output human-readable only; no renderer modes or output flags were added.

### File List

- `_bmad-output/implementation-artifacts/2-3-render-human-readable-stats-tables.md`
- `go.mod`
- `go.sum`
- `internal/render/stats.go`
- `internal/render/stats_test.go`
- `internal/render/testdata/stats_empty_languages_80.golden`
- `internal/render/testdata/stats_long_values_80.golden`
- `internal/render/testdata/stats_success_80.golden`
- `internal/render/width.go`

## Change Log

- 2026-06-12: Added human-readable stats table rendering and marked Story 2.3 ready for review.
