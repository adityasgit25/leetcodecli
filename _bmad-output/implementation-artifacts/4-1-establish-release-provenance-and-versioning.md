---
baseline_commit: NO_VCS
---

# Story 4.1: Establish Release Provenance and Versioning

Status: review

Completion Note: Ultimate context engine analysis completed - comprehensive developer guide created.

## Story

As a developer user,
I want LeetcodeCLI releases to come from a traceable source revision and version tag,
so that I can trust the artifacts I install and understand which v1 behavior they contain.

## Requirements Covered

- FR1: Users can install and run LeetcodeCLI on Windows, macOS, and Linux.
- FR13: Users can install, run stats commands, and understand limitations without reading source code.
- FR14: Users can evaluate whether LeetcodeCLI is safe and appropriate for their local environment.

## Acceptance Criteria

1. Given release packaging work begins, when source-control and release metadata are reviewed, then the project has a documented release provenance approach that identifies the source revision used for each public artifact, and public release artifacts are not published from an untracked or ambiguous source state.
2. Given the project may be published as a public source repository, when release readiness is reviewed, then the team has explicitly decided whether the Go module path remains `leetcodecli` or changes before release packaging, and any module-path change is completed before GoReleaser artifacts are produced.
3. Given v1 command behavior must remain stable across patch releases, when release versioning is documented, then the project uses semantic versioning or an equivalent public release convention, and release tags, artifact versions, release notes, and documentation refer to the same version.
4. Given release provenance is established, when `go test ./...` and `go build ./...` are run before packaging, then both commands complete successfully, and the command surface still exposes only the supported v1 public-profile stats workflow.

## Tasks / Subtasks

- [x] Document release provenance policy. (AC: 1)
  - [x] Add or update release documentation that states public artifacts must be built from a tracked source revision and version tag.
  - [x] State that artifacts must not be published from `NO_VCS`, an uncommitted source state, or an ambiguous source snapshot.
  - [x] Record where the release source revision, tag, and release notes will be captured.

- [x] Resolve the public module path decision. (AC: 2)
  - [x] Compare planning history that mentions local module path `leetcodecli` with current `go.mod`, which uses `github.com/adityasgit25/leetcodecli`.
  - [x] Decide whether the current public module path is final for release.
  - [x] If the module path changes, update imports, docs, tests, and any release docs before packaging stories begin.
  - [x] If the module path remains, document that decision so release packaging does not reopen it.

- [x] Define versioning and tag rules. (AC: 3)
  - [x] Document the versioning convention, expected tag format, and release-note version format.
  - [x] Preserve stable v1 command behavior unless a breaking change is explicitly documented.
  - [x] State that patch releases must not silently change command names, argument meaning, output section names, or no-auth behavior.

- [x] Keep v1 trust and scope boundaries visible. (AC: 1, 3)
  - [x] Repeat that v1 is unofficial and depends on public LeetCode data.
  - [x] Repeat that v1 stores no credentials, tokens, cookies, Session Data, or config files.
  - [x] Repeat that login, logout, own-profile default lookup, private-data access, JSON output, dashboards, recommendations, goals, reminders, and topic-gap analysis remain out of scope.

- [x] Add documentation tests for release provenance. (AC: 1, 2, 3)
  - [x] Extend existing docs tests or add a focused test that verifies release docs mention source revision, version tag, executable name `leetcode`, module-path decision, SemVer or equivalent versioning, and no-auth trust boundaries.
  - [x] Keep existing docs tests that guard command examples and trust boundaries.

- [x] Verify. (AC: 4)
  - [x] Run `go test ./...`.
  - [x] Run `go build ./...`.
  - [x] Confirm no GoReleaser configuration or public release publishing is added in this story.

## Dev Notes

### Previous Story Context

- Epic 3 completed release-readiness documentation without adding `.goreleaser.yaml`.
- Epic 3 retrospective identified missing sprint tracking and absent Git metadata as process/release-provenance risks.
- Implementation readiness found the PRD/addendum/architecture/epics aligned, but `.decision-log.md` needed superseded auth/session/logout decisions annotated. That annotation is now present.

### Current Codebase Context

- `go.mod` currently declares `module github.com/adityasgit25/leetcodecli`.
- Earlier planning artifacts mention `leetcodecli` as a local module path because the workspace originally had no Git remote.
- The public executable name remains `leetcode`; do not rename commands or examples to `leetcodecli`.
- Existing docs tests live in `cmd/docs_test.go` and already check install guidance, command examples, and no premature GoReleaser config.

### Architecture Guardrails

- Release provenance must be established before public artifacts are published.
- Release artifacts must preserve executable name `leetcode`.
- v1 command surface remains `leetcode help`, `leetcode help stats`, and `leetcode stats <username>`.
- Do not add config, auth, session, token, browser-login, logout, cache, JSON, CSV, dashboard, recommendation, goal, reminder, topic-gap, browser extension, web app, or TUI behavior.

### Expected Files

- Add or update: `docs/release.md` or `docs/release-provenance.md`
- Update: `README.md` and/or `docs/installation.md` only if needed to link release provenance guidance
- Update: `cmd/docs_test.go` or add a focused release-docs test
- Optional update: `go.mod` and imports only if the module-path decision changes
- Do not add: `.goreleaser.yaml` in this story
- Do not add: `.github/workflows/release.yml` in this story

### Test Guidance

- Prefer content tests that read docs files, matching the existing `cmd/docs_test.go` pattern.
- Tests should ensure release docs say public artifacts trace to a source revision and version tag.
- Tests should ensure docs continue to say `leetcode`, not `leetcodecli`, for the executable.
- Run `go test ./...` and `go build ./...`.

### References

- [Source: `_bmad-output/planning-artifacts/epics.md` - Story 4.1]
- [Source: `_bmad-output/planning-artifacts/prds/prd-LeetcodeCLI-2026-06-11/prd.md` - Versioning, Distribution, FR1, FR13, FR14]
- [Source: `_bmad-output/planning-artifacts/prds/prd-LeetcodeCLI-2026-06-11/addendum.md` - module path and release packaging note]
- [Source: `_bmad-output/planning-artifacts/architecture.md` - release packaging and module-path guidance]
- [Source: `_bmad-output/implementation-artifacts/epic-3-retro-2026-06-12.md` - provenance and sprint-tracking action items]
- [Current code: `go.mod` - current module path]

## Dev Agent Record

### Agent Model Used

GPT-5 Codex

### Debug Log References

- `go test ./cmd` initially failed as expected because `docs/release.md` did not exist.
- `go test ./...` passed with workspace-local `GOCACHE`.
- `go build ./...` passed with workspace-local `GOCACHE`.
- `Test-Path .goreleaser.yaml` returned `False`.

### Completion Notes List

- Added `docs/release.md` documenting release provenance, tracked source revision and version tag requirements, forbidden `NO_VCS`/dirty/ambiguous publishing states, release metadata locations, SemVer tag rules, and v1 trust boundaries.
- Recorded the module-path decision: `github.com/adityasgit25/leetcodecli` remains final for release packaging, while the public executable remains `leetcode`.
- Linked release provenance guidance from README and installation docs.
- Added a documentation test that guards source revision, version tag, module path, SemVer, executable name, and no-auth trust-boundary wording.

### File List

- README.md
- docs/installation.md
- docs/release.md
- cmd/docs_test.go
- _bmad-output/implementation-artifacts/sprint-status.yaml
- _bmad-output/implementation-artifacts/4-1-establish-release-provenance-and-versioning.md

## Change Log

- 2026-06-12: Created Story 4.1 ready for development.
- 2026-06-12: Implemented release provenance, versioning, module-path decision docs, and provenance documentation tests.
