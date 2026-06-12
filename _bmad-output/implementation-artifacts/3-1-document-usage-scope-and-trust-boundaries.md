---
baseline_commit: NO_VCS
---

# Story 3.1: Document Usage, Scope, and Trust Boundaries

Status: review

Completion Note: Ultimate context engine analysis completed - comprehensive developer guide created.

## Story

As a developer user,
I want clear usage and limitations documentation,
so that I can install and use LeetcodeCLI with accurate expectations.

## Requirements Covered

- FR13: Users can install, run stats commands, and understand limitations without reading source code.
- FR14: Users can evaluate whether LeetcodeCLI is safe and appropriate for their local environment.

## Acceptance Criteria

1. Given v1 command behavior is implemented, when README and usage documentation are reviewed, then they include examples for `leetcode help`, `leetcode help stats`, and `leetcode stats <username>`, and they do not present bare `leetcode stats` as a successful stats command.
2. Given LeetcodeCLI depends on public LeetCode data, when limitations documentation is reviewed, then it states that LeetcodeCLI is unofficial unless that changes, and it explains that public-data availability and LeetCode access behavior may change.
3. Given v1 has no authentication or persistence scope, when documentation is reviewed, then it states that v1 stores no credentials, tokens, cookies, Session Data, or config files, and it states that login, logout, own-profile default lookup, recommendations, goals, reminders, topic-gap analysis, JSON output, dashboards, browser extensions, desktop UI, and web apps are out of scope.
4. Given command help and documentation both describe v1, when they are compared, then terminology, examples, and limitations are consistent across README, docs, and Cobra help text.

## Tasks / Subtasks

- [x] Create or update user-facing docs. (AC: 1, 2, 3, 4)
  - [x] Add/update `README.md`.
  - [x] Add/update `docs/usage.md`.
  - [x] Add/update `docs/limitations.md`.
  - [x] Keep docs concise and accurate to implemented v1 behavior.

- [x] Document supported usage. (AC: 1, 4)
  - [x] Include `leetcode help`.
  - [x] Include `leetcode help stats`.
  - [x] Include `leetcode stats <username>`.
  - [x] Present bare `leetcode stats` only as a usage-error example, if mentioned.
  - [x] Describe output sections: Profile Summary, Total Solved Count, and Language Breakdown.

- [x] Document trust boundaries. (AC: 2, 3)
  - [x] State LeetcodeCLI is unofficial unless project status has changed intentionally.
  - [x] State public-data availability and LeetCode access behavior may change.
  - [x] State v1 stores no credentials, tokens, cookies, Session Data, or config files.
  - [x] State no login/logout/session/config/private-data workflow exists in v1.

- [x] Document out-of-scope items. (AC: 3)
  - [x] List own-profile default lookup, recommendations, goals, reminders, topic-gap analysis, JSON output, dashboards, browser extensions, desktop UI, and web apps as out of scope.
  - [x] Avoid sections that imply planned auth/session setup.

- [x] Align docs with Cobra help. (AC: 4)
  - [x] Compare README/docs examples against `leetcode help` and `leetcode help stats`.
  - [x] Update either docs or help if terminology drift exists.
  - [x] Add docs/help consistency tests if the codebase has a practical command-output test pattern.

- [x] Verify. (AC: 1, 2, 3, 4)
  - [x] Run `go test ./...`.
  - [x] Run `go build ./...`.

## Dev Notes

### Previous Story Context

- Epic 1 establishes the command surface.
- Epic 2 completes stats retrieval, rendering, and failure behavior.
- This story should document what actually exists after implementation, not future aspirations.

### Architecture Guardrails

- Documentation and Cobra help must align with v1 scope.
- Do not add auth/session/config docs placeholders.
- Do not present bare `leetcode stats` as successful.
- Do not promise permanent LeetCode API compatibility.
- Do not document JSON/CSV output or private/authenticated data access.

### Expected Files

- Add/update: `README.md`
- Add/update: `docs/usage.md`
- Add/update: `docs/limitations.md`
- Possible update: `cmd/root.go`, `cmd/stats.go` if help text has drifted.
- Possible update: command/docs tests if existing patterns support it.

### Test Guidance

- If tests inspect docs, keep them focused on high-risk drift: required examples, forbidden terms, and exact command shape.
- Prefer command tests for help output and markdown lint-style checks only if already present or easy to maintain.

### References

- [Source: `_bmad-output/planning-artifacts/epics.md` - Story 3.1]
- [Source: `_bmad-output/planning-artifacts/architecture.md` - Documentation Constraints, Documentation and Trust mapping]
- [Source: `_bmad-output/planning-artifacts/prds/prd-LeetcodeCLI-2026-06-11/prd.md` - FR13, FR14, Non-Goals, Distribution]
- [Previous: `_bmad-output/implementation-artifacts/2-5-map-stats-failures-to-user-safe-errors.md`]

## Dev Agent Record

### Agent Model Used

GPT-5 Codex

### Debug Log References

- 2026-06-12: Confirmed no existing `README.md`, `docs/usage.md`, or `docs/limitations.md` content before adding user-facing documentation.
- 2026-06-12: Red-phase `go test ./...` failed because README and docs files were missing.
- 2026-06-12: Docs consistency test caught a case-sensitive trust-boundary phrase mismatch; wording was aligned across README, usage, and limitations docs.
- 2026-06-12: Verification passed: `go test ./...` and `go build ./...`.

### Completion Notes List

- Added README, usage docs, and limitations docs covering supported commands, stats output sections, unofficial/public-data trust boundaries, and v1 no-auth/no-persistence scope.
- Documented bare `leetcode stats` only as a usage-error path and kept the supported command shape as `leetcode stats <username>`.
- Added command/docs consistency tests for required examples, limitations wording, and help alignment.

### File List

- `_bmad-output/implementation-artifacts/3-1-document-usage-scope-and-trust-boundaries.md`
- `README.md`
- `cmd/docs_test.go`
- `docs/limitations.md`
- `docs/usage.md`

## Change Log

- 2026-06-12: Added usage, scope, and trust-boundary documentation and marked Story 3.1 ready for review.
