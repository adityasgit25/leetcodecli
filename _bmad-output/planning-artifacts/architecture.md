---
stepsCompleted: [1, 2, 3, 4, 5, 6, 7, 8]
inputDocuments:
  - "_bmad-output/planning-artifacts/prds/prd-LeetcodeCLI-2026-06-11/prd.md"
  - "_bmad-output/planning-artifacts/prds/prd-LeetcodeCLI-2026-06-11/addendum.md"
  - "_bmad-output/planning-artifacts/briefs/brief-LeetcodeCLI-2026-06-11/brief.md"
  - "_bmad-output/planning-artifacts/prds/prd-LeetcodeCLI-2026-06-11/review-rubric.md"
  - "_bmad-output/planning-artifacts/prds/prd-LeetcodeCLI-2026-06-11/reconcile-brief.md"
workflowType: 'architecture'
project_name: 'LeetcodeCLI'
user_name: 'Adi'
date: '2026-06-11'
lastStep: 8
status: 'complete'
completedAt: '2026-06-11'
---

# Architecture Decision Document

_This document builds collaboratively through step-by-step discovery. Sections are appended as we work through each architectural decision together._

## Project Context Analysis

### Requirements Overview

LeetcodeCLI v1 is now a small public-profile stats CLI. The core command is `leetcode stats <username>`, which fetches and renders public LeetCode profile statistics. Bare `leetcode stats` should not attempt authenticated self-resolution; it should show clear usage/help explaining that a username is required.

The architecture should exclude Browser Login, Session Data persistence, and `leetcode logout` from v1. This removes the local credential/session trust surface and shifts the main risk to unauthenticated LeetCode GraphQL availability, response shape stability, rate limiting, and clear terminal failures.

Mandatory successful output remains:

- Profile Summary
- Total Solved Count
- Language Breakdown

Command parsing must remain separate from data retrieval, data retrieval must remain separate from terminal rendering, and raw GraphQL responses should be normalized before validation and rendering.

### Non-Functional Requirements

Reliability is the main driver. LeetcodeCLI depends on LeetCode's unofficial GraphQL endpoint, so architecture must isolate the LeetCode client behind a narrow boundary with timeouts, clear error mapping, and tests for missing mandatory fields.

Portability still matters across Windows, macOS, and Linux for executable behavior, terminal rendering, path-independent operation, line endings, and release validation.

Security scope is reduced because v1 stores no Session Data and handles no user credentials. The tool must still avoid misleading users about data source reliability or unofficial status.

### Scale & Complexity

- Primary domain: cross-platform CLI plus external public API integration
- Product complexity: low
- Operational/integration complexity: medium
- Main complexity drivers: LeetCode GraphQL drift, public-profile availability, terminal rendering determinism, release portability, and clear error behavior

### Technical Constraints & Dependencies

- Go-based CLI using Cobra.
- `leetcode stats <username>` is the v1 happy path.
- Bare `leetcode stats` shows usage/help because no authenticated default profile exists.
- Data retrieval uses `POST https://leetcode.com/graphql` where available without Browser Login.
- Output is human-readable Terminal Table only.
- JSON, CSV, dashboards, recommendations, goals, reminders, topic-gap analysis, browser extensions, web apps, Browser Login, Session Data, and logout are out of scope.
- macOS distribution uses Homebrew.
- Windows and Linux distribution use checksummed GitHub Releases binaries with documented install and PATH setup.
- LeetcodeCLI is an unofficial third-party tool.

### Cross-Cutting Concerns Identified

- LeetCode GraphQL isolation behind a gateway/client interface.
- Normalization of GraphQL response data before mandatory-field validation.
- Explicit error taxonomy for username missing, username not found, stats unavailable, mandatory stats missing, network/API failure, rate limiting or access blocking, rendering failure, and unsupported command usage.
- stdout vs stderr discipline and stable exit-code behavior.
- Deterministic rendering tests with sample payloads, golden output, 80-column checks, and CRLF/LF normalization.
- Documentation/help alignment with actual v1 behavior: username-required stats, no login/logout, unofficial data source, supported OSes, limitations, and failure modes.
- Scope discipline: no auth/session architecture should be prebuilt for v1 unless reintroduced intentionally.

## Starter Template Evaluation

### Primary Technology Domain

CLI tool, based on the product requirements and preserved implementation context.

The selected implementation foundation is Go with Cobra. The repo currently has no `go.mod`, Go source files, or existing CLI scaffold, so architecture should initialize a new Go module rather than adapt an existing application.

### Starter Options Considered

**Option 1: Standard Go module plus Cobra**

Use Go modules and Cobra for command structure. Cobra provides conventional command wiring, argument validation, help behavior, and subcommand organization while keeping v1 close to the brief's Go implementation context.

`cobra-cli` may be used as a one-time scaffold helper, but generated code must be committed and reviewed. The generator is not a runtime dependency, and implementation should pin or record the generator version rather than relying on `@latest` in reproducible setup notes.

**Option 2: Hand-written minimal Cobra scaffold**

Manually create `main.go`, `cmd/root.go`, and `cmd/stats.go` using Cobra directly.

This is also acceptable for v1 because the command surface is tiny. If the generator produces unused boilerplate, hand-written files are preferable to keeping unnecessary generated structure.

**Option 3: Cobra plus Viper**

Rejected for v1. LeetcodeCLI has no config file, no Session Data, no login state, no token storage, and no persistent user preferences after the architecture scope override. Viper would add configuration surface area before there is a real product need.

**Option 4: TUI-style starter**

Rejected for v1. LeetcodeCLI output is a human-readable terminal table, not a stateful interactive terminal UI.

### Selected Starter: Go Module plus Minimal Cobra Scaffold

**Rationale for Selection:**

Cobra matches the brief's implementation context and the v1 command model. It gives `leetcode`, `leetcode stats <username>`, help text, argument validation, and command organization without introducing config, TUI, or web-app concerns.

The starter should be used only for command scaffolding. Domain logic, LeetCode GraphQL access, response normalization, mandatory-stat validation, rendering, and error mapping should live outside Cobra command files.

**Initialization Command:**

```bash
go mod init leetcodecli
go install github.com/spf13/cobra-cli@v1.3.0
cobra-cli init
cobra-cli add stats
```

The concrete local module path is `leetcodecli` because this workspace does not yet have a Git remote. If the project is later published as a public source repository, the module path may be intentionally changed before release packaging. If the implementation team hand-writes the minimal Cobra files instead of using `cobra-cli`, it should preserve the same command contract and thin command boundaries.

**Architectural Decisions Provided by Starter:**

**Language & Runtime:**

- Go module-based project.
- Target current stable Go unless release compatibility requires a lower supported version.
- Public binary name remains `leetcode`.

**CLI Framework:**

- Cobra provides root command and subcommand structure.
- `stats` is a Cobra subcommand with exactly one required username argument for v1.
- Bare `leetcode stats` returns `Username required. Usage: leetcode stats <username>. Run "leetcode help" for help.`, exits with code `2`, and makes no network call.

**Build Tooling:**

- Native Go build tooling.
- Initial verification commands are `go test ./...` and `go build`.
- Release automation should be decided separately; do not add GoReleaser or equivalent tooling in the starter unless the release architecture explicitly adopts it.

**Testing Framework:**

- Use Go's standard `testing` package as the default.
- Unit tests must not make live LeetCode calls.
- LeetCode client tests should use fixtures, fake servers, and injected `*http.Client` or transport behavior.
- Rendering tests should inject terminal width and include golden output, 80-column behavior, CRLF/LF normalization, long usernames, long language names, and empty/non-empty Language Breakdown cases.

**Code Organization:**

- Keep generated Cobra files thin.
- Recommended initial responsibility split:
  - `cmd`: Cobra command wiring and argument validation only
  - `internal/leetcode`: GraphQL client, response DTOs, response normalization, timeout-aware HTTP behavior
  - `internal/render`: terminal table rendering from normalized profile stats
- Add `internal/profile` only if domain transformation grows beyond the LeetCode client boundary.
- Add `internal/apperr` or equivalent only if a dedicated application error taxonomy becomes clearer than ordinary Go errors plus command-level mapping.
- Do not create config, auth, session, token, browser-login, or logout packages for v1.

**Development Experience:**

- `go test ./...` is the primary verification command.
- `go build` verifies local binary construction.
- The first implementation test matrix should cover:
  - `leetcode stats` with no args: username-required error, exit code `2`, no network call
  - `leetcode stats alice`: happy-path fixture renders Profile Summary, Total Solved Count, and Language Breakdown
  - unknown, private, unavailable, malformed, and mandatory-field-missing responses
  - empty and non-empty Language Breakdown
  - long username and language names at 80 columns
  - HTTP timeout and network/path errors mapped to user-safe stderr

**Documentation Constraints:**

- Document only `leetcode stats <username>` as the successful stats command.
- Bare `leetcode stats` may appear only as a usage-error example that points to `leetcode help`.
- Do not add README or help sections named Authentication, Session, Logout, Config, Token, or similar v1 placeholders.
- Help text must state that output includes Profile Summary, Total Solved Count, and Language Breakdown.
- Docs must state that LeetcodeCLI is an unofficial third-party tool and that public-profile data availability can differ from LeetCode UI behavior.

**Note:** Project initialization using this foundation should be the first implementation story.

## Core Architectural Decisions

### Decision Priority Analysis

**Critical Decisions (Block Implementation):**

- Use `github.com/jedib0t/go-pretty/v6/table` for terminal table rendering.
- Use Go standard `net/http` for LeetCode GraphQL calls with injected `*http.Client` or transport.
- Normalize GraphQL responses before validation and rendering.
- Treat `leetcode stats <username>` as the only successful stats command.
- Return the standard username-required usage error and no network call for bare `leetcode stats`.

**Important Decisions (Shape Architecture):**

- Use `golang.org/x/term` for terminal width detection with an 80-column fallback.
- Use explicit stdout/stderr and exit-code policy.
- Use GitHub Actions OS matrix for Windows, macOS, and Linux.
- Use GoReleaser for v1 release packaging when distribution stories begin.

**Deferred Decisions (Post-MVP):**

- Authentication, Session Data, config, logout, JSON/CSV output, dashboards, and recommendations remain out of scope.
- Caching is deferred because v1 makes one public-profile request per command.
- Observability beyond concise user-facing errors is deferred because this is a local CLI without a backend service.

### Data Architecture

LeetcodeCLI stores no application data in v1.

Data flow is request-only:
`username argument -> LeetCode GraphQL client -> normalized profile stats -> mandatory-field validation -> terminal renderer`.

GraphQL DTOs stay inside `internal/leetcode`. The renderer receives only normalized stats, never raw GraphQL payloads.

Mandatory validation checks:

- Profile Summary exists.
- Total Solved Count exists from the `All` row.
- Language Breakdown exists; an empty valid list renders as none found.
- Missing mandatory fields fail the command rather than rendering partial success.

### Authentication & Security

No Browser Login, cookies, tokens, Session Data, config file, credential storage, or logout command exists in v1.

Security posture:

- Username input is sent as a GraphQL variable, not interpolated into a query string.
- No secrets are read, stored, printed, or logged.
- Error output must not imply official LeetCode affiliation or private-data access.
- Public-profile availability is treated as an external dependency, not a guarantee.

### API & Communication Patterns

LeetCode access uses `POST https://leetcode.com/graphql`.

The client uses:

- Go standard `net/http`.
- Context-aware requests.
- A default request timeout.
- Injected HTTP client or transport for tests.
- Static GraphQL query plus variables.
- Minimal headers: `Content-Type: application/json` and a stable CLI user agent.

No automatic retry in v1. A failed request should produce a clear error instead of increasing load on an unofficial endpoint.

Error mapping:

- Missing username: usage error with exact copy `Username required. Usage: leetcode stats <username>. Run "leetcode help" for help.`
- `matchedUser == nil`: username not found.
- GraphQL errors or missing mandatory fields: stats unavailable or required stats missing.
- HTTP timeout/network failure: could not reach LeetCode.
- HTTP 403/429: access blocked or rate limited.
- Rendering failure: command failure with concise stderr.

Exit policy:

- `0`: successful stats render or explicit help through `leetcode help` or Cobra help flags.
- `2`: usage/argument error.
- `1`: network, API, data validation, or rendering failure.

### Frontend Architecture

Not applicable. LeetcodeCLI has no web, desktop, mobile, or TUI frontend in v1.

### Infrastructure & Deployment

CI uses GitHub Actions with Windows, macOS, and Linux runners.

Required CI checks:

- `go test ./...`
- `go build ./...`
- No live LeetCode calls in unit tests.

Release packaging:

- Use GoReleaser for GitHub Releases binaries, checksums, and Homebrew packaging when release stories begin.
- Keep release config separate from the initial scaffold story if it slows first implementation.
- Release artifacts must preserve the public executable name `leetcode`.

### Decision Impact Analysis

**Implementation Sequence:**

1. Initialize Go module and minimal Cobra scaffold.
2. Add `stats <username>` command behavior and usage errors.
3. Add normalized profile stats model and LeetCode client with fake-server tests.
4. Add mandatory-field validation and error mapping.
5. Add `go-pretty` renderer with width-injected golden tests.
6. Add CI matrix.
7. Add GoReleaser configuration during release packaging work.

**Cross-Component Dependencies:**

- `cmd` depends on application interfaces, not raw HTTP or raw GraphQL.
- `internal/leetcode` owns external API details.
- `internal/render` owns table layout and width behavior.
- Error mapping sits at the command boundary so domain/client code remains testable.

## Implementation Patterns & Consistency Rules

### Pattern Categories Defined

**Critical Conflict Points Identified:**

There are 12 areas where AI agents could make incompatible choices if not specified:

1. Cobra command behavior and argument validation.
2. Package ownership between `cmd`, `internal/leetcode`, and `internal/render`.
3. Raw GraphQL DTO naming versus normalized domain naming.
4. Mandatory stats validation timing.
5. User-facing error copy and exit codes.
6. stdout versus stderr usage.
7. HTTP timeout and client injection.
8. Terminal width detection and fallback.
9. Table rendering style and golden output stability.
10. Fixture location and live-network test avoidance.
11. Documentation examples for `leetcode stats <username>`.
12. No-auth/no-session/no-config scope enforcement.

### Naming Patterns

**Database Naming Conventions:**

Not applicable. LeetcodeCLI v1 stores no application data and has no database, migrations, tables, indexes, or local cache.

**API Naming Conventions:**

- LeetcodeCLI exposes no HTTP API.
- External GraphQL request code lives under `internal/leetcode`.
- GraphQL DTO names should describe the external response shape and stay unexported unless another package truly needs them.
- Normalized stats types should describe product concepts, not GraphQL field names.

Examples:

- Good DTO name inside `internal/leetcode`: `matchedUserResponse`.
- Good normalized type: `ProfileStats`.
- Avoid leaking GraphQL names into render code, such as `SubmitStatsGlobal` or `AcSubmissionNum`.

**Code Naming Conventions:**

- Go package names use lowercase, short, noun-style names: `leetcode`, `render`, `cmd`.
- Go file names use lowercase snake case when multiple words are needed: `stats_command.go`, `profile_stats.go`.
- Exported types use product/domain language: `ProfileStats`, `ProfileSummary`, `LanguageCount`.
- Functions that perform I/O should make that clear: `FetchProfileStats`, `RenderStats`, `DetectWidth`.
- Test fixtures use descriptive names: `profile_success.json`, `profile_missing_total.json`, `profile_not_found.json`.

### Structure Patterns

**Project Organization:**

- `cmd` owns Cobra command wiring, argument validation, stdout/stderr wiring, and exit-code mapping.
- `internal/leetcode` owns LeetCode GraphQL endpoint details, request/response DTOs, HTTP behavior, response normalization, and API-specific error classification.
- `internal/render` owns table layout, terminal width handling, and output formatting.
- Do not add `internal/auth`, `internal/session`, `internal/config`, `internal/cache`, or `internal/logout` in v1.
- Add `internal/profile` only if normalized domain logic becomes too large to live cleanly at the `internal/leetcode` boundary.
- Add `internal/apperr` only if error mapping becomes clearer as a shared application error taxonomy than command-local mapping.

**File Structure Patterns:**

- Unit tests live beside the package they test using `*_test.go`.
- Test fixtures live in package-local `testdata` directories.
- Golden files for rendering live under `internal/render/testdata`.
- Cobra command tests should execute commands with injected stdout/stderr buffers and injected service/client behavior.
- No test should depend on the live `https://leetcode.com/graphql` endpoint.

### Format Patterns

**API Response Formats:**

There is no LeetcodeCLI public API response format.

For the external LeetCode GraphQL call:

- Use a static GraphQL query string.
- Pass username through GraphQL variables.
- Decode into private DTO structs in `internal/leetcode`.
- Convert DTOs into normalized stats before validation/rendering.
- Treat GraphQL `errors`, `matchedUser == nil`, missing mandatory stats, HTTP failures, and decode failures as distinct internal cases that can map to stable user-facing errors.

**Data Exchange Formats:**

The internal render input is normalized Go structs, not JSON maps and not raw GraphQL DTOs.

Canonical normalized shape:

```go
type ProfileStats struct {
    Summary        ProfileSummary
    TotalSolved    int
    LanguageCounts []LanguageCount
}
```

Rules:

- Missing nullable display fields render as `N/A`.
- Missing mandatory stats fail before rendering.
- Empty valid `LanguageCounts` renders as an explicit none-found state.
- Profile URL is derived from username in normalized/domain code, not inside table layout code.

### Communication Patterns

**Event System Patterns:**

Not applicable. LeetcodeCLI v1 has no event bus, background jobs, async workflow, or service-to-service communication.

**State Management Patterns:**

Not applicable beyond request-local command state.

Rules:

- Do not introduce global mutable state for current username, HTTP clients, terminal width, or command output.
- Pass dependencies explicitly into command constructors or package functions.
- Keep command execution deterministic under tests by injecting stdout, stderr, HTTP behavior, and terminal width.

### Process Patterns

**Error Handling Patterns:**

- Library/internal packages return errors with enough classification for the command boundary to map them.
- User-facing copy is emitted at the command boundary, not deep inside `internal/leetcode`.
- Successful stats output writes to stdout.
- Help and usage may write through Cobra's configured output.
- Failures write concise messages to stderr.
- Exit code `0` means successful stats render or explicit help.
- Exit code `2` means usage/argument error.
- Exit code `1` means network, API, data validation, or rendering failure.

Stable error categories:

- Missing username.
- Username not found.
- Stats unavailable.
- Required stats missing.
- Could not reach LeetCode.
- Access blocked or rate limited.
- Rendering failed.

**Loading State Patterns:**

Not applicable. The CLI is a synchronous command and should not introduce spinners, progress bars, background polling, or retry loops in v1.

### Enforcement Guidelines

**All AI Agents MUST:**

- Keep Cobra command files thin and free of direct GraphQL parsing.
- Keep LeetCode HTTP and DTO details inside `internal/leetcode`.
- Render only normalized stats through `internal/render`.
- Preserve `leetcode stats <username>` as the only successful stats command.
- Ensure bare `leetcode stats` emits the standard username-required usage error, exits `2`, and does not make a network call.
- Avoid auth, session, config, token, logout, cache, JSON, CSV, dashboard, recommendation, or planning features in v1.
- Use fixtures or fake servers for tests instead of live LeetCode calls.
- Inject HTTP behavior and terminal width in tests.
- Keep docs and help examples aligned with the username-required command contract.

**Pattern Enforcement:**

- `go test ./...` must pass before implementation work is considered complete.
- Command tests must verify stdout/stderr routing and exit behavior.
- Renderer golden tests must normalize line endings before comparison.
- A change that adds a new package must state which architecture responsibility it owns.
- A change that touches command help or README examples must preserve the no-auth/no-session v1 scope.

### Pattern Examples

**Good Examples:**

```go
// cmd/stats.go
func newStatsCommand(fetcher StatsFetcher, widthDetector WidthDetector, out, errOut io.Writer) *cobra.Command
```

```go
// internal/leetcode/client.go
func (c *Client) FetchProfileStats(ctx context.Context, username string) (ProfileStats, error)
```

```go
// internal/render/stats.go
func RenderStats(stats leetcode.ProfileStats, width int) (string, error)
```

Good command behavior:

```text
leetcode stats alice
```

Fetches public profile stats for `alice`, validates mandatory sections, and writes tables to stdout.

```text
leetcode stats
```

Emits `Username required. Usage: leetcode stats <username>. Run "leetcode help" for help.`, exits with code `2`, and makes no network call.

**Anti-Patterns:**

- Calling `http.Post` directly from `cmd/stats.go`.
- Passing raw GraphQL JSON or DTO structs into `internal/render`.
- Reading terminal width directly inside golden tests.
- Adding `leetcode login`, `leetcode logout`, config files, cookie storage, or token handling.
- Adding `--json`, `--csv`, recommendations, reminders, or dashboards.
- Writing errors to stdout.
- Returning partial success when Total Solved Count or Language Breakdown is missing.

## Project Structure & Boundaries

### Complete Project Directory Structure

```text
LeetcodeCLI/
|-- README.md
|-- go.mod
|-- go.sum
|-- main.go
|-- .gitignore
|-- .github/
|   `-- workflows/
|       `-- ci.yml
|-- cmd/
|   |-- root.go
|   |-- root_test.go
|   |-- stats.go
|   `-- stats_test.go
|-- internal/
|   |-- leetcode/
|   |   |-- client.go
|   |   |-- client_test.go
|   |   |-- graphql.go
|   |   |-- normalize.go
|   |   |-- normalize_test.go
|   |   |-- types.go
|   |   `-- testdata/
|   |       |-- profile_success.json
|   |       |-- profile_not_found.json
|   |       |-- profile_missing_total.json
|   |       |-- profile_missing_languages.json
|   |       |-- profile_empty_languages.json
|   |       `-- graphql_error.json
|   `-- render/
|       |-- stats.go
|       |-- stats_test.go
|       |-- width.go
|       |-- width_test.go
|       `-- testdata/
|           |-- stats_success_80.golden
|           |-- stats_long_values_80.golden
|           `-- stats_empty_languages_80.golden
|-- docs/
|   |-- usage.md
|   `-- limitations.md
`-- _bmad-output/
    `-- planning-artifacts/
        `-- architecture.md
```

Release files such as `.goreleaser.yaml` should be added when release packaging stories begin, not during the initial scaffold unless implementation explicitly includes release setup.

### Architectural Boundaries

**API Boundaries:**

- LeetcodeCLI exposes no HTTP API.
- The only external API boundary is `internal/leetcode` calling `POST https://leetcode.com/graphql`.
- `cmd` must not know GraphQL query shape, HTTP headers, or response DTO fields.
- `internal/render` must not know HTTP status codes, GraphQL errors, or external response shapes.

**Component Boundaries:**

- `main.go` calls `cmd.Execute()` or equivalent root command execution only.
- `cmd` wires dependencies, validates arguments, maps errors to user-facing messages, and routes stdout/stderr.
- `internal/leetcode` fetches, decodes, normalizes, and validates profile stats from LeetCode.
- `internal/render` turns normalized stats into deterministic terminal tables.

**Service Boundaries:**

- There are no backend services, background workers, or service-to-service calls in v1.
- Internal communication is direct Go function/interface calls.
- Tests should replace external communication with fake servers, fake transports, or injected fetcher interfaces.

**Data Boundaries:**

- Raw GraphQL response structs stay private to `internal/leetcode`.
- Normalized profile stats are the only data passed from `internal/leetcode` to `cmd` and `internal/render`.
- No application data is persisted to disk.
- No cache directory, config directory, or session directory exists in v1.

### Requirements To Structure Mapping

**CLI Entry Point and Command Discovery:**

- `main.go`
- `cmd/root.go`
- `cmd/root_test.go`
- `README.md`
- `docs/usage.md`

**Stats Command and Username Targeting:**

- `cmd/stats.go`
- `cmd/stats_test.go`
- `internal/leetcode/client.go`
- `internal/leetcode/client_test.go`

`leetcode stats <username>` is implemented here. Bare `leetcode stats` is tested as the standard username-required usage error with no network call.

**Stats Data Retrieval and Interpretation:**

- `internal/leetcode/graphql.go`
- `internal/leetcode/types.go`
- `internal/leetcode/normalize.go`
- `internal/leetcode/normalize_test.go`
- `internal/leetcode/testdata/*.json`

Profile Summary, Total Solved Count, and Language Breakdown validation lives at this boundary.

**Terminal Output:**

- `internal/render/stats.go`
- `internal/render/stats_test.go`
- `internal/render/width.go`
- `internal/render/width_test.go`
- `internal/render/testdata/*.golden`

Table rendering, width fallback, long-value handling, empty languages, and line-ending normalization are tested here.

**Documentation and Trust:**

- `README.md`
- `docs/usage.md`
- `docs/limitations.md`
- Cobra help text in `cmd/root.go` and `cmd/stats.go`

Docs must state that the tool is unofficial, username-based, no-login, no-session, and dependent on public LeetCode data availability.

**Cross-Platform Build and Validation:**

- `.github/workflows/ci.yml`
- `go.mod`
- `go.sum`

CI validates `go test ./...` and `go build ./...` on Windows, macOS, and Linux.

### Integration Points

**Internal Communication:**

The command layer depends on interfaces or function types for fetching stats and detecting width. This lets command tests verify behavior without live network calls or real terminal width.

Recommended flow:

```text
cmd/stats.go
  -> internal/leetcode.Client.FetchProfileStats(ctx, username)
  -> internal/render.RenderStats(stats, width)
  -> stdout or stderr
```

**External Integrations:**

- LeetCode GraphQL endpoint: `https://leetcode.com/graphql`
- GitHub Actions for CI.
- GoReleaser for future release packaging.

**Data Flow:**

```text
username
  -> GraphQL variables
  -> LeetCode HTTP response
  -> private DTO decode
  -> normalized ProfileStats
  -> mandatory validation
  -> go-pretty table rendering
  -> stdout
```

Failures break the flow at the boundary where they are detected and return to `cmd` for stderr and exit-code mapping.

### File Organization Patterns

**Configuration Files:**

- `go.mod` and `go.sum` define Go module dependencies.
- `.github/workflows/ci.yml` defines CI checks.
- No `.env`, config YAML, local settings file, token file, or session file exists in v1.

**Source Organization:**

- `cmd` for command surface only.
- `internal/leetcode` for external API and normalized stats retrieval.
- `internal/render` for terminal presentation.
- Add new internal packages only when a distinct responsibility emerges and cannot fit cleanly in these boundaries.

**Test Organization:**

- Use co-located `*_test.go`.
- Use package-local `testdata`.
- Use fake server or injected transport for HTTP tests.
- Use golden files only for deterministic renderer output.
- Normalize line endings before golden comparisons.

**Asset Organization:**

There are no static assets in v1.

### Development Workflow Integration

**Development Server Structure:**

Not applicable. LeetcodeCLI has no development server.

**Build Process Structure:**

- Local test: `go test ./...`
- Local build: `go build ./...`
- CLI run during development: `go run . stats <username>`

**Deployment Structure:**

- Initial implementation produces a local `leetcode` binary.
- CI validates source across supported OSes.
- Release stories add GitHub Releases binaries, checksums, and Homebrew packaging through GoReleaser.

## Architecture Validation Results

### Coherence Validation

**Decision Compatibility:**

The architecture is internally coherent for the revised v1 scope. Go, Cobra, `net/http`, `go-pretty/v6/table`, `x/term`, GitHub Actions, and later GoReleaser fit together without conflicting responsibilities.

The PRD and addendum have been updated to match the no-login v1 scope, removing the previous drift around Browser Login, Session Data, own-profile default lookup, and logout.

**Pattern Consistency:**

Patterns consistently enforce thin Cobra commands, isolated LeetCode GraphQL access, normalized stats before rendering, deterministic tests, stdout/stderr discipline, and no auth/session/config scope.

**Structure Alignment:**

The project tree supports all architectural decisions. `cmd`, `internal/leetcode`, and `internal/render` have clear ownership. Tests and fixtures are placed where implementation agents can follow them consistently.

### Requirements Coverage Validation

**Feature Coverage:**

The revised v1 feature set is covered:

- `leetcode stats <username>`
- public-profile stats retrieval
- bare `leetcode stats` usage error with no network call
- `leetcode help` for help
- Profile Summary
- Total Solved Count
- Language Breakdown
- human-readable table output
- cross-platform validation
- documentation of unofficial/public-data limitations

**Functional Requirements Coverage:**

All revised FR categories are architecturally supported:

- CLI entry point and command discovery map to `main.go` and `cmd`.
- Public-profile targeting maps to `cmd/stats.go` and `internal/leetcode`.
- Stats retrieval and validation map to `internal/leetcode`.
- Terminal rendering maps to `internal/render`.
- Documentation and trust requirements map to `README.md`, `docs/usage.md`, `docs/limitations.md`, and Cobra help text.

**Non-Functional Requirements Coverage:**

Reliability, portability, maintainability, security/privacy, and terminal readability are covered. Security scope is reduced because v1 stores no credentials, tokens, cookies, Session Data, or config files.

### Implementation Readiness Validation

**Decision Completeness:**

Critical decisions are documented with current technology choices and versions where relevant. The concrete local Go module path is `leetcodecli`.

**Structure Completeness:**

The project structure is complete enough for implementation and avoids unused auth/config/session directories.

**Pattern Completeness:**

Major conflict points are addressed: package ownership, DTO boundaries, validation timing, error mapping, exit codes, stdout/stderr, width injection, fixtures, and no live-network unit tests.

### Gap Analysis Results

No critical implementation-blocking gaps remain.

**Deferred, Non-Blocking Items:**

- Public repository module path can be changed intentionally before release packaging if the project is published as source.
- GoReleaser configuration is deferred until release packaging stories begin.
- Browser Login, Session Data, logout, config/default username, JSON/CSV output, caching, and richer analytics remain future-scope decisions.

### Validation Issues Addressed

- PRD and addendum were reconciled with the revised no-login architecture scope.
- The local Go module path was set to `leetcodecli`.
- Bare `leetcode stats` now has exact behavior: username-required usage error, exit code `2`, no network call, and help direction via `leetcode help`.
- Standard v1 user-facing error copy was finalized.
- Live LeetCode calls remain excluded from unit tests.
- Rendering and HTTP behavior remain injectable for deterministic tests.

### Architecture Completeness Checklist

**Requirements Analysis**

- [x] Project context thoroughly analyzed
- [x] Scale and complexity assessed
- [x] Technical constraints identified
- [x] Cross-cutting concerns mapped

**Architectural Decisions**

- [x] Critical decisions documented with versions
- [x] Technology stack fully specified
- [x] Integration patterns defined
- [x] Performance considerations addressed

**Implementation Patterns**

- [x] Naming conventions established
- [x] Structure patterns defined
- [x] Communication patterns specified
- [x] Process patterns documented

**Project Structure**

- [x] Complete directory structure defined
- [x] Component boundaries established
- [x] Integration points mapped
- [x] Requirements to structure mapping complete

### Architecture Readiness Assessment

**Overall Status:** READY FOR IMPLEMENTATION

**Confidence Level:** High for the revised no-login v1 scope.

**Key Strengths:**

- Small command surface.
- Clear package boundaries.
- No credential/session handling.
- Deterministic test strategy.
- Explicit external API isolation.
- Cross-platform CI path.
- PRD and architecture now describe the same v1 product.

**Areas for Future Enhancement:**

- Add release configuration when distribution work starts.
- Change module path before public source publishing if needed.
- Revisit auth/session only if deliberately reintroduced in a future version.

### Implementation Handoff

**AI Agent Guidelines:**

- Follow all architectural decisions exactly as documented.
- Do not reintroduce auth, session, logout, config, or JSON output.
- Keep Cobra command files thin.
- Keep LeetCode API details inside `internal/leetcode`.
- Keep rendering deterministic and fixture-tested.
- Use exact standard error copy from the PRD/addendum.

**First Implementation Priority:**

Initialize the Go module and minimal Cobra scaffold for `leetcode stats <username>`:

```bash
go mod init leetcodecli
```
