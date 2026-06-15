---
baseline_commit: NO_VCS
---

# Story 1.1: Set Up Initial Project from Go/Cobra Starter Template

Status: review

Completion Note: Ultimate context engine analysis completed - comprehensive developer guide created.

## Story

As a developer user,
I want a runnable `leetcode` CLI entry point,
so that I can invoke the tool from my terminal and reach the v1 command surface.

## Requirements Covered

- FR1: Users can install and run LeetcodeCLI on Windows, macOS, and Linux.
- FR12: Users receive human-readable output only; v1 does not expose JSON, CSV, or machine-readable export flags.

## Acceptance Criteria

1. Given the repository has no existing Go CLI scaffold, when the implementation initializes the project, then it creates a Go module using module path `leetcodecli`, and it adds a minimal Cobra-based CLI entry point with `main.go` and thin command wiring under `cmd`.
2. Given the CLI has been initialized, when a developer runs the project locally through Go tooling, then the command can be invoked as the LeetcodeCLI executable surface, and the public command examples and command metadata consistently use `leetcode`.
3. Given the scaffold is complete, when `go test ./...` and `go build ./...` are run, then both commands complete successfully, and no auth, session, config, token, logout, cache, JSON, CSV, recommendation, dashboard, or TUI package is introduced.

## Tasks / Subtasks

- [x] Confirm the live workspace still has no application scaffold before creating files. (AC: 1)
  - [x] Check for existing `go.mod`, `main.go`, `cmd/`, `internal/`, and README/application files.
  - [x] Preserve all existing BMad planning artifacts; do not move or rewrite `_bmad-output` content.

- [x] Initialize the Go module at the project root. (AC: 1, 3)
  - [x] Run `go mod init leetcodecli` unless `go.mod` already exists.
  - [x] Use the local Go toolchain, targeting the current stable Go line where available.
  - [x] Do not add GoReleaser, CI, docs, release packaging, auth/session/config dependencies, or LeetCode API code in this story.

- [x] Add a minimal Cobra CLI entry point. (AC: 1, 2)
  - [x] Create `main.go` that only delegates to the command package execution path.
  - [x] Create `cmd/root.go` with a root command whose public `Use`/examples/metadata consistently use `leetcode`.
  - [x] Keep command files thin: command construction, output wiring, and execution only.
  - [x] Prefer hand-written minimal Cobra files if the `cobra-cli` generator creates unused boilerplate.
  - [x] If a generator is used, use `github.com/spf13/cobra-cli@v1.3.0`, not an unpinned `@latest` install.

- [x] Keep the v1 surface narrow. (AC: 2, 3)
  - [x] Do not introduce `login`, `logout`, session, token, config, cache, JSON, CSV, dashboard, recommendation, goal, reminder, topic-gap, browser extension, web app, or TUI commands or packages.
  - [x] Do not add `github.com/spf13/viper`; v1 has no config files or persistent preferences.
  - [x] Do not implement LeetCode GraphQL retrieval, table rendering, or stats error mapping in this story. Those belong to later Epic 2 stories.
  - [x] If a placeholder `stats` command is introduced for scaffold shape, it must not perform network work or pretend to render successful stats. Detailed help/discovery is Story 1.2, missing-username behavior is Story 1.3, and command surface auditing is Story 1.4.

- [x] Add focused scaffold tests. (AC: 2, 3)
  - [x] Add package-local tests under `cmd` that instantiate commands directly with buffers instead of calling `os.Exit`.
  - [x] Verify the root command metadata uses `leetcode`.
  - [x] Verify scaffold execution/help behavior does not require network, config files, credentials, or terminal-specific state.
  - [x] Keep tests deterministic on Windows, macOS, and Linux.

- [x] Verify the scaffold. (AC: 3)
  - [x] Run `go test ./...`.
  - [x] Run `go build ./...`.
  - [x] Optionally run `go build -o leetcode .` as a local binary-name sanity check, but do not commit the built binary.
  - [x] Run a final file/package check to confirm no forbidden v1 packages or command surfaces were introduced.

## Dev Notes

### Current Workspace State

- At story creation time, the workspace contains BMad configuration and planning artifacts only; no `go.mod`, `main.go`, `cmd/`, or application source files were present outside `_bmad` and `_bmad-output`.
- No previous implementation story exists for Epic 1, so there are no earlier dev notes, review findings, file patterns, or regressions to inherit.
- No Git repository was detected from the workspace root during story creation, so no commit-history intelligence is available.
- No `_bmad-output/implementation-artifacts/sprint-status.yaml` file exists to update. This story file is still ready for development.

### Architecture Guardrails

- Use a Go module plus minimal Cobra scaffold as the implementation foundation.
- The concrete local module path is `leetcodecli`.
- The public executable/command identity is `leetcode`. Do not use `leetcodecli` in visible command examples or Cobra `Use` strings.
- `main.go` should only call the command execution entry point, such as `cmd.Execute()`.
- `cmd` owns command wiring, argument validation, stdout/stderr routing, and exit-code mapping only.
- Keep generated or hand-written Cobra files thin. Do not put GraphQL parsing, HTTP calls, rendering logic, persistence, or product-domain transformation inside `cmd`.
- `internal/leetcode` and `internal/render` are planned architectural packages for later stories; do not create them in Story 1.1 unless a clear scaffold-only reason emerges.
- Add new internal packages only when they own a distinct responsibility. This story should normally need none.

### Starter and Dependency Guidance

- Architecture allows either:
  - `go mod init leetcodecli`, then hand-written minimal Cobra files, or
  - `go mod init leetcodecli`, `go install github.com/spf13/cobra-cli@v1.3.0`, `cobra-cli init`, and careful cleanup of unused boilerplate.
- Hand-written minimal files are preferred if they avoid generator noise.
- Use the Cobra runtime library `github.com/spf13/cobra`; let `go mod tidy` capture the exact resolved dependency versions in `go.mod` and `go.sum`.
- As of 2026-06-12, Go's featured stable download line is `go1.26.4`; use a compatible Go directive for the local toolchain and avoid setting a future version the local environment cannot build.
- As of 2026-06-12, the latest observed `github.com/spf13/cobra` release is `v1.10.1`; the separate `github.com/spf13/cobra-cli` bootstrapper latest release remains `v1.3.0`.

### Suggested Initial File Shape

```text
LeetcodeCLI/
|-- go.mod
|-- go.sum
|-- main.go
`-- cmd/
    |-- root.go
    `-- root_test.go
```

If the developer decides a scaffold-level `stats` command is necessary now, add:

```text
cmd/
|-- stats.go
`-- stats_test.go
```

Do not add future-oriented directories such as `internal/auth`, `internal/session`, `internal/config`, `internal/cache`, `internal/leetcode`, or `internal/render` during this story unless directly required by the scaffold acceptance criteria.

### Command Behavior Boundaries for This Story

- Required now:
  - The project builds as a Go module.
  - The root command is runnable through Go tooling.
  - Public command identity is `leetcode`.
  - Tests pass without network, config, credentials, or platform-specific state.
- Not required now:
  - Final root help content for `leetcode help`.
  - Final stats help content for `leetcode help stats`.
  - Exact bare `leetcode stats` usage error.
  - LeetCode GraphQL request/response handling.
  - Profile stats normalization.
  - Terminal table rendering.
  - README, docs, CI, release packaging, Homebrew, or GitHub Releases artifacts.

### Test Guidance

- Use Go's standard `testing` package.
- Test command constructors directly rather than spawning subprocesses for the core scaffold checks.
- Use `bytes.Buffer` or similar writers for command output in tests.
- Avoid `os.Exit` in testable command paths. If the public execution path exits, keep it at the outer boundary only.
- `go test ./...` and `go build ./...` are mandatory completion checks.
- Unit tests must not call `https://leetcode.com/graphql` or any live network endpoint.

### Anti-Patterns to Avoid

- Adding Viper or any config-file package.
- Adding auth, session, token, cookie, login, logout, cache, JSON, CSV, dashboard, recommendation, goal, reminder, topic-gap, browser extension, web app, or TUI scaffolding.
- Creating a fake success path for `leetcode stats <username>` before retrieval/rendering exists.
- Calling LeetCode or any HTTP endpoint from `cmd`.
- Hiding unsupported future work behind placeholder packages.
- Committing generated binaries or local build output.

### References

- [Source: `_bmad-output/planning-artifacts/epics.md` - Story 1.1, Epic 1 implementation notes, requirements inventory]
- [Source: `_bmad-output/planning-artifacts/architecture.md` - Starter Template Evaluation, Selected Starter, Project Structure & Boundaries, Implementation Patterns & Consistency Rules]
- [Source: `_bmad-output/planning-artifacts/prds/prd-LeetcodeCLI-2026-06-11/prd.md` - FR1, FR12, v1 scope boundaries, public command contract]
- [Source: `_bmad-output/planning-artifacts/prds/prd-LeetcodeCLI-2026-06-11/addendum.md` - preserved implementation context, no-auth v1 scope, standard command boundary]
- [External: `https://go.dev/dl/` - current Go stable downloads checked on 2026-06-12]
- [External: `https://github.com/spf13/cobra/releases` - Cobra runtime releases checked on 2026-06-12]
- [External: `https://github.com/spf13/cobra-cli/releases` - Cobra CLI bootstrapper releases checked on 2026-06-12]

## Project Structure Notes

- This story creates the first application source files, so there are no existing application files marked UPDATE.
- Preserve BMad artifacts in place.
- The initial source tree should be intentionally small. Future architecture target folders exist in the architecture document, but Story 1.1 should only create the subset needed for a runnable Cobra scaffold.
- There is no detected project-context file for additional repository rules.

## Dev Agent Record

### Agent Model Used

GPT-5 Codex

### Debug Log References

- 2026-06-12: Confirmed no existing application scaffold (`go.mod`, `main.go`, `cmd/`, `internal/`, README application file) before creating files.
- 2026-06-12: `git rev-parse HEAD` failed because the workspace is not a Git repository; baseline recorded as `NO_VCS`.
- 2026-06-12: Local Go toolchain detected as `go1.26.4 windows/amd64`.
- 2026-06-12: Red-phase `go test ./...` failed on missing `NewRootCommand` after adding command tests.
- 2026-06-12: `go mod tidy` resolved `github.com/spf13/cobra v1.10.2`.
- 2026-06-12: Verification passed: `go test ./...`, `go build ./...`, `go list ./...`, and forbidden v1 surface scan over Go/module files.

### Completion Notes List

- Initialized module `leetcodecli` with a minimal hand-written Cobra root command whose public command identity is `leetcode`.
- Added `main.go` as a thin delegation point to `cmd.Execute()`.
- Added deterministic command tests for root metadata and local help behavior.
- Kept the scaffold narrow: no stats behavior, network calls, config, auth/session/cache, machine-readable output flags, CI, docs, or release packaging were added.

### File List

- `_bmad-output/implementation-artifacts/1-1-set-up-initial-project-from-go-cobra-starter-template.md`
- `cmd/root.go`
- `cmd/root_test.go`
- `go.mod`
- `go.sum`
- `main.go`

## Change Log

- 2026-06-12: Created the initial Go/Cobra CLI scaffold and marked Story 1.1 ready for review.
