---
stepsCompleted: [1, 2, 3, 4]
inputDocuments:
  - "_bmad-output/planning-artifacts/prds/prd-LeetcodeCLI-2026-06-11/prd.md"
  - "_bmad-output/planning-artifacts/prds/prd-LeetcodeCLI-2026-06-11/addendum.md"
  - "_bmad-output/planning-artifacts/architecture.md"
workflowType: 'epics-and-stories'
project_name: 'LeetcodeCLI'
user_name: 'Adi'
date: '2026-06-11'
lastStep: 4
status: 'complete'
completedAt: '2026-06-12'
---

# LeetcodeCLI - Epic Breakdown

## Overview

This document provides the complete epic and story breakdown for LeetcodeCLI, decomposing the requirements from the PRD, UX Design if it exists, and Architecture requirements into implementable stories.

## Requirements Inventory

### Functional Requirements

FR1: Users can install and run LeetcodeCLI on Windows, macOS, and Linux.

FR2: Users can discover the Stats Command and basic usage through CLI help.

FR3: Users can run `leetcode stats <username>` to view a Public Profile.

FR4: Users who run `leetcode stats` without a username receive usage guidance instead of a network request.

FR5: LeetcodeCLI only displays stats it can retrieve for the requested Public Profile.

FR6: Users can see a Profile Summary in the Stats View, including username, display name, ranking, reputation, and derived profile URL, with missing nullable values shown as `N/A`.

FR7: Users can see the Total Solved Count in the Stats View, sourced from the `All` row in `matchedUser.submitStatsGlobal.acSubmissionNum`.

FR8: Users can see solved-question counts by programming language in the Stats View, sourced from `matchedUser.languageProblemCount`, with empty valid language data shown as none found.

FR9: LeetcodeCLI may show additional practice stats only when they are reliable, understandable, documented, and do not turn v1 into a planning product.

FR10: Users can read the Stats View as one or more pretty Terminal Tables with visible section labels `Profile Summary`, `Total Solved Count`, and `Language Breakdown`.

FR11: Users can read essential stats in standard terminal widths, including 80 columns, without long values corrupting adjacent columns.

FR12: Users receive human-readable output only; v1 does not expose JSON, CSV, or machine-readable export flags.

FR13: Users can install, run stats commands, and understand limitations without reading source code.

FR14: Users can evaluate whether LeetcodeCLI is safe and appropriate for their local environment through documentation of data source, supported OSes, unofficial status, and no credential/session storage.

### NonFunctional Requirements

NFR1: LeetcodeCLI v1 must not read, store, print, log, or expose credentials, tokens, cookies, Session Data, or config files.

NFR2: LeetcodeCLI must fail with actionable messages when LeetCode data retrieval, profile targeting, mandatory stats validation, rate limiting/access blocking, or rendering fails; it must not panic or render fabricated stats.

NFR3: LeetcodeCLI should avoid unnecessary local delay after LeetCode responds; table rendering should complete quickly enough that network latency dominates normal Stats Command runtime, with a practical v1 target of under 250 ms local processing time after successful data retrieval.

NFR4: LeetcodeCLI must work in common shells on Windows, macOS, and Linux.

NFR5: Public command behavior must remain stable across v1 patch releases unless a breaking change is documented.

### Additional Requirements

- Use a Go module plus minimal Cobra scaffold as the implementation foundation; the concrete local module path is `leetcodecli`.
- Preserve `leetcode` as the public executable name and `leetcode stats <username>` as the v1 happy path.
- Implement bare `leetcode stats` as a usage error that emits exactly `Username required. Usage: leetcode stats <username>. Run "leetcode help" for help.`, exits with code `2`, and makes no network call.
- Keep generated or hand-written Cobra command files thin; command code owns argument validation, dependency wiring, stdout/stderr routing, and exit-code mapping only.
- Organize implementation around `cmd`, `internal/leetcode`, and `internal/render`; add other internal packages only when a distinct responsibility emerges.
- Keep LeetCode GraphQL endpoint details, static query, request/response DTOs, HTTP behavior, normalization, and API-specific error classification inside `internal/leetcode`.
- Retrieve data with `POST https://leetcode.com/graphql` using Go standard `net/http`, context-aware requests, a default timeout, injected HTTP client or transport for tests, GraphQL variables for username, `Content-Type: application/json`, and a stable CLI user agent.
- Use this v1 GraphQL shape: `matchedUser.username`, `matchedUser.profile.realName`, `matchedUser.profile.ranking`, `matchedUser.profile.reputation`, `matchedUser.submitStatsGlobal.acSubmissionNum`, and `matchedUser.languageProblemCount`.
- Normalize GraphQL responses into product/domain stats before validation and rendering; renderer code must not receive raw GraphQL payloads or DTOs.
- Validate mandatory stats before rendering: Profile Summary exists, Total Solved Count exists from the `All` row, and Language Breakdown exists; missing mandatory stats fail the command.
- Treat empty but valid Language Breakdown as a renderable none-found state.
- Derive profile URL from username in normalized/domain code, not table layout code.
- Use `github.com/jedib0t/go-pretty/v6/table` for terminal table rendering.
- Use `golang.org/x/term` for terminal width detection with an 80-column fallback.
- Keep successful stats output on stdout, failure messages on stderr, explicit help exit code `0`, usage/argument failures exit code `2`, and network/API/data/rendering failures exit code `1`.
- Model or classify these failure states explicitly: missing username, username not found, profile unavailable, endpoint failure, rate limit/access blocked, mandatory stats missing, and rendering failure.
- Use the PRD/addendum standard v1 error copy for user-facing failures.
- Do not implement automatic retries in v1.
- Do not create config, auth, session, token, browser-login, logout, cache, JSON, CSV, dashboard, recommendation, goal, reminder, topic-gap, browser extension, web app, or TUI features in v1.
- Unit tests must not call the live LeetCode endpoint; use fixtures, fake servers, fake transports, and injected fetcher behavior.
- Rendering tests must inject terminal width and cover golden output, 80-column behavior, CRLF/LF normalization, long usernames, long language names, empty Language Breakdown, and non-empty Language Breakdown.
- Command tests must verify stdout/stderr routing, exit behavior, the no-network bare stats path, and happy/error-path command behavior.
- Use package-local `testdata` directories for LeetCode JSON fixtures and renderer golden files.
- Add GitHub Actions CI with Windows, macOS, and Linux runners that run `go test ./...` and `go build ./...`.
- Add GoReleaser configuration only when release packaging stories begin; release artifacts must preserve the public executable name `leetcode`.
- Documentation and Cobra help must align with v1 scope: username-required stats, Profile Summary, Total Solved Count, Language Breakdown, unofficial third-party status, public-data dependency, supported OSes, and no credential/session/config storage.
- macOS distribution uses Homebrew; Windows and Linux distribution use checksummed GitHub Releases binaries with documented install and PATH setup.

### UX Design Requirements

No UX Design document was provided or found for this workflow run, so no separate UX-DR requirements were extracted.

### FR Coverage Map

FR1: Epic 1, Epic 3, and Epic 4 - Epic 1 covers local runnable executable behavior; Epic 3 covers supported install/run guidance and CI validation; Epic 4 covers packaged release artifacts and install validation across Windows, macOS, and Linux.

FR2: Epic 1 - CLI help and stats command discovery.

FR3: Epic 2 - `leetcode stats <username>` public profile lookup.

FR4: Epic 1 - bare `leetcode stats` usage error with no network call.

FR5: Epic 2 - only retrieved public-profile stats are displayed.

FR6: Epic 2 - Profile Summary.

FR7: Epic 2 - Total Solved Count.

FR8: Epic 2 - Language Breakdown.

FR9: Epic 2 - additional practice stats guardrail.

FR10: Epic 2 - pretty terminal table output.

FR11: Epic 2 - terminal width resilience.

FR12: Epic 1 and Epic 2 - Epic 1 covers no structured-output/config/auth surface in help or usage behavior; Epic 2 covers human-readable stats output with no JSON, CSV, or export flags.

FR13: Epic 3 and Epic 4 - Epic 3 covers installation, usage, and limitations documentation; Epic 4 aligns that documentation with real release artifacts, checksums, Homebrew, and GitHub Releases.

FR14: Epic 3 and Epic 4 - Epic 3 covers trust, unofficial status, no credential/session storage, public-data dependency, and supported operating systems; Epic 4 adds release provenance, checksum verification, and distribution trust signals.

## Epic List

### Epic 1: Safe CLI Entry and Stats Command Discovery

Users can run the `leetcode` command, discover `leetcode stats <username>`, and receive safe usage guidance when they invoke the command incorrectly.

**FRs covered:** FR1, FR2, FR4, FR12

**Implementation notes:** Establish the Go module and minimal Cobra command shell, preserve `leetcode` as the public executable name, keep command files thin, expose help for root and stats usage, prevent any auth/session/config or structured-output surface from appearing in help, and ensure bare `leetcode stats` emits the exact username-required usage error, exits `2`, and performs no network call.

### Epic 2: Public Profile Stats Retrieval and Display

Users can run `leetcode stats <username>` and receive a complete, validated, human-readable Stats View for a public LeetCode profile.

**FRs covered:** FR3, FR5, FR6, FR7, FR8, FR9, FR10, FR11, FR12

**Implementation notes:** Implement the public LeetCode GraphQL retrieval path, normalize responses before rendering, validate mandatory stats, map profile/API/data/rendering failures to stable user-facing errors, render Profile Summary, Total Solved Count, and Language Breakdown as readable terminal tables, keep output human-readable only, and test with fixtures/fakes rather than live LeetCode calls.

### Epic 3: Public Release Trust and Cross-Platform Readiness

Users can understand how to install and use LeetcodeCLI safely across supported platforms, including its unofficial status, public-data dependency, and no-credential/no-session boundary.

**FRs covered:** FR1, FR13, FR14

**Implementation notes:** Add or finalize README and usage/limitations documentation, align docs with Cobra help and v1 scope, document supported OSes and public-data failure modes, add cross-platform CI for Windows, macOS, and Linux, and defer GoReleaser packaging until distribution stories intentionally begin.

### Epic 4: Release Packaging and Public Distribution

Users can install verified LeetcodeCLI release artifacts through the intended public distribution channels: Homebrew on macOS and checksummed GitHub Releases binaries on Windows and Linux.

**FRs covered:** FR1, FR13, FR14

**Implementation notes:** Establish release provenance before publishing, decide whether the public module path must change before release packaging, add GoReleaser configuration, produce checksummed Windows/macOS/Linux artifacts that preserve the `leetcode` executable name, publish through GitHub Releases and Homebrew, keep v1 trust boundaries intact, and update installation guidance from intended flow to actual release flow.

## Epic 1: Safe CLI Entry and Stats Command Discovery

Users can run the `leetcode` command, discover `leetcode stats <username>`, and receive safe usage guidance when they invoke the command incorrectly.

### Story 1.1: Set Up Initial Project from Go/Cobra Starter Template

**Requirements covered:** FR1, FR12

As a developer user,
I want a runnable `leetcode` CLI entry point,
So that I can invoke the tool from my terminal and reach the v1 command surface.

**Acceptance Criteria:**

**Given** the repository has no existing Go CLI scaffold
**When** the implementation initializes the project
**Then** it creates a Go module using module path `leetcodecli`
**And** it adds a minimal Cobra-based CLI entry point with `main.go` and thin command wiring under `cmd`.

**Given** the CLI has been initialized
**When** a developer runs the project locally through Go tooling
**Then** the command can be invoked as the LeetcodeCLI executable surface
**And** the public command examples and command metadata consistently use `leetcode`.

**Given** the scaffold is complete
**When** `go test ./...` and `go build ./...` are run
**Then** both commands complete successfully
**And** no auth, session, config, token, logout, cache, JSON, CSV, recommendation, dashboard, or TUI package is introduced.

### Story 1.2: Expose Root and Stats Help

**Requirements covered:** FR2, FR12

As a developer user,
I want clear help for the CLI and stats command,
So that I can discover how to run `leetcode stats <username>` without reading source code.

**Acceptance Criteria:**

**Given** the CLI has a root command
**When** a user runs `leetcode help`
**Then** help output lists `stats` as an available command
**And** it does not mention login, logout, session setup, token setup, config setup, JSON output, CSV output, dashboards, recommendations, goals, or reminders.

**Given** the stats command exists
**When** a user runs `leetcode help stats`
**Then** help output explains `leetcode stats <username>`
**And** it states that v1 requires a username.

**Given** standard Cobra help behavior is available
**When** a user runs `leetcode --help`
**Then** the CLI displays standard root help without crashing
**And** the visible command examples consistently use `leetcode`.

### Story 1.3: Enforce Username-Required Stats Usage

**Requirements covered:** FR4

As a developer user,
I want `leetcode stats` to fail locally with clear usage guidance when I omit the username,
So that I understand how to correct the command without triggering an unnecessary network request.

**Acceptance Criteria:**

**Given** the stats command is available
**When** a user runs `leetcode stats` without a username
**Then** the CLI exits with usage exit code `2`
**And** it writes exactly `Username required. Usage: leetcode stats <username>. Run "leetcode help" for help.` to stderr.

**Given** a username is missing
**When** `leetcode stats` handles the usage error
**Then** it performs no LeetCode GraphQL request
**And** command tests prove no fetcher, HTTP client, or network dependency is invoked.

**Given** the usage error is emitted
**When** stdout and stderr are inspected in tests
**Then** no successful stats output appears on stdout
**And** the failure message remains concise and human-readable.

### Story 1.4: Preserve v1 Scope Boundaries in the CLI Surface

**Requirements covered:** FR12

As a developer user,
I want the CLI surface to show only the supported v1 stats workflow,
So that I am not misled into expecting authentication, configuration, structured exports, or planning features.

**Acceptance Criteria:**

**Given** the root and stats commands are implemented
**When** help text, command descriptions, examples, and available flags are reviewed
**Then** they describe only the username-based public stats workflow
**And** they expose no login, logout, session, token, config, JSON, CSV, dashboard, recommendation, goal, reminder, topic-gap, browser extension, web app, or TUI workflow.

**Given** v1 scope is intentionally public-profile-only
**When** command tests inspect the available command tree
**Then** unsupported auth/session/config/export commands are absent
**And** the stats command accepts exactly one required username argument for successful execution.

**Given** the CLI command surface is complete for Epic 1
**When** `go test ./...` is run
**Then** command tests verify help output, missing-username behavior, stdout/stderr routing, and v1 scope boundaries.

## Epic 2: Public Profile Stats Retrieval and Display

Users can run `leetcode stats <username>` and receive a complete, validated, human-readable Stats View for a public LeetCode profile.

### Story 2.1: Retrieve Public Profile Data from LeetCode GraphQL

**Requirements covered:** FR3, FR5

As a developer user,
I want the stats command to retrieve public LeetCode profile data for a supplied username,
So that the CLI can build stats from the same public-data source every time.

**Acceptance Criteria:**

**Given** a username is supplied
**When** the LeetCode client builds the request
**Then** it sends a `POST` request to `https://leetcode.com/graphql`
**And** it passes the username as a GraphQL variable rather than interpolating it into the query string.

**Given** the LeetCode client sends a request
**When** request headers are inspected in tests
**Then** the request includes `Content-Type: application/json`
**And** it includes a stable LeetcodeCLI user agent.

**Given** the client is tested
**When** unit tests run
**Then** they use injected HTTP behavior, fake servers, fake transports, or fixtures
**And** they do not call the live LeetCode endpoint.

**Given** LeetCode returns transport, HTTP, or GraphQL-level failures
**When** the client handles the response
**Then** it classifies not found, unavailable stats, endpoint failure, rate limit/access blocked, and malformed response states distinctly enough for command-level error mapping.

### Story 2.2: Normalize and Validate Mandatory Profile Stats

**Requirements covered:** FR5, FR6, FR7, FR8

As a developer user,
I want public LeetCode data normalized into trustworthy profile stats,
So that the CLI never renders fabricated or misleading partial results.

**Acceptance Criteria:**

**Given** LeetCode returns a successful GraphQL response
**When** the response is decoded
**Then** raw GraphQL DTOs remain inside `internal/leetcode`
**And** the rest of the application receives normalized profile stats rather than raw response payloads.

**Given** a public profile response contains profile fields
**When** normalization runs
**Then** it produces username, display name, ranking, reputation, and a profile URL derived from username
**And** nullable display fields can be represented as `N/A` without breaking rendering.

**Given** `matchedUser.submitStatsGlobal.acSubmissionNum` contains an `All` row
**When** normalization runs
**Then** Total Solved Count is set from the `All` row `count`
**And** missing Total Solved Count is classified as a mandatory-stats failure.

**Given** `matchedUser.languageProblemCount` is present
**When** normalization runs
**Then** it produces language-name and solved-count rows
**And** an empty but valid language list remains a successful none-found state.

**Given** mandatory profile summary, total solved, or language breakdown data is missing
**When** validation runs
**Then** the stats command cannot proceed to successful rendering
**And** tests cover success, profile not found, missing total, missing languages, empty languages, and GraphQL error fixtures.

### Story 2.3: Render Human-Readable Stats Tables

**Requirements covered:** FR6, FR7, FR8, FR9, FR10, FR11, FR12

As a developer user,
I want public profile stats rendered as readable terminal tables,
So that I can quickly scan progress from a standard shell.

**Acceptance Criteria:**

**Given** normalized profile stats are available
**When** rendering runs
**Then** the output includes visible sections labeled `Profile Summary`, `Total Solved Count`, and `Language Breakdown`
**And** the renderer does not require raw GraphQL DTOs or HTTP details.

**Given** Profile Summary contains nullable display values
**When** rendering runs
**Then** missing nullable values display as `N/A`
**And** table layout remains coherent.

**Given** Language Breakdown is empty but valid
**When** rendering runs
**Then** the CLI shows an explicit none-found state
**And** it does not crash or treat the empty list as a mandatory-stats failure.

**Given** terminal width is available
**When** rendering detects width
**Then** it uses `golang.org/x/term` where appropriate
**And** falls back to a safe 80-column layout when width cannot be detected.

**Given** long usernames or language names are rendered at 80 columns
**When** renderer tests compare output
**Then** adjacent columns remain readable
**And** golden tests normalize CRLF and LF line endings.

**Given** v1 output is human-readable only
**When** the Stats View is rendered
**Then** it uses pretty terminal table output through `github.com/jedib0t/go-pretty/v6/table`
**And** it does not produce JSON, CSV, or machine-readable export output.

### Story 2.4: Wire `leetcode stats <username>` Happy Path

**Requirements covered:** FR3, FR5, FR6, FR7, FR8, FR10, FR12

As a developer user,
I want `leetcode stats <username>` to render a complete Stats View for a valid public profile,
So that I can check public LeetCode progress without opening the website.

**Acceptance Criteria:**

**Given** the CLI receives `leetcode stats <username>`
**When** the supplied username resolves to valid public stats
**Then** the command fetches public profile data, validates mandatory stats, renders the Stats View, writes it to stdout, and exits `0`.

**Given** a successful Stats View is rendered
**When** the output is inspected
**Then** it identifies the requested or resolved LeetCode profile
**And** it includes Profile Summary, Total Solved Count, and Language Breakdown.

**Given** the stats command is wired to data retrieval and rendering
**When** command tests run
**Then** they use injected fetcher and width behavior
**And** they do not require a live LeetCode request or a real terminal.

**Given** the command has a successful username-based workflow
**When** help examples and command metadata are inspected
**Then** they still describe `leetcode stats <username>` as the supported v1 happy path
**And** they do not imply own-profile default lookup or authenticated private-data access.

### Story 2.5: Map Stats Failures to User-Safe Errors

**Requirements covered:** FR5, FR9, FR12

As a developer user,
I want profile, API, data, and rendering failures to produce clear messages,
So that I can understand what went wrong without seeing misleading partial stats.

**Acceptance Criteria:**

**Given** LeetCode returns no matching user
**When** the stats command handles the failure
**Then** it writes `No LeetCode profile found for "<username>". Check the username and try again.` to stderr
**And** it exits `1`.

**Given** public profile or stats data is unavailable
**When** the stats command handles the failure
**Then** it writes `Stats for "<username>" are not available from LeetCode right now. Try again later.` to stderr
**And** it exits `1`.

**Given** the public API or network cannot be reached
**When** the stats command handles the failure
**Then** it writes `Could not reach LeetCode. Check your connection and try again.` to stderr
**And** it exits `1`.

**Given** LeetCode rate-limits or blocks the request
**When** the stats command handles the failure
**Then** it writes `LeetCode blocked or rate-limited the request. Try again later.` to stderr
**And** it exits `1`.

**Given** mandatory stats are missing
**When** the stats command handles the failure
**Then** it writes `LeetCode did not return required stats for "<username>". Try again later.` to stderr
**And** it exits `1`.

**Given** stats output cannot be rendered
**When** the stats command handles the failure
**Then** it writes `Could not render stats output. Try again later.` to stderr
**And** it exits `1`.

**Given** any stats failure occurs after a username is supplied
**When** stdout and stderr are inspected
**Then** failure messages are written to stderr
**And** no fabricated, partial, JSON, CSV, or machine-readable success output is written to stdout.

## Epic 3: Public Release Trust and Cross-Platform Readiness

Users can understand how to install and use LeetcodeCLI safely across supported platforms, including its unofficial status, public-data dependency, and no-credential/no-session boundary.

### Story 3.1: Document Usage, Scope, and Trust Boundaries

**Requirements covered:** FR13, FR14

As a developer user,
I want clear usage and limitations documentation,
So that I can install and use LeetcodeCLI with accurate expectations.

**Acceptance Criteria:**

**Given** v1 command behavior is implemented
**When** README and usage documentation are reviewed
**Then** they include examples for `leetcode help`, `leetcode help stats`, and `leetcode stats <username>`
**And** they do not present bare `leetcode stats` as a successful stats command.

**Given** LeetcodeCLI depends on public LeetCode data
**When** limitations documentation is reviewed
**Then** it states that LeetcodeCLI is unofficial unless that changes
**And** it explains that public-data availability and LeetCode access behavior may change.

**Given** v1 has no authentication or persistence scope
**When** documentation is reviewed
**Then** it states that v1 stores no credentials, tokens, cookies, Session Data, or config files
**And** it states that login, logout, own-profile default lookup, recommendations, goals, reminders, topic-gap analysis, JSON output, dashboards, browser extensions, desktop UI, and web apps are out of scope.

**Given** command help and documentation both describe v1
**When** they are compared
**Then** terminology, examples, and limitations are consistent across README, docs, and Cobra help text.

### Story 3.2: Add Cross-Platform CI Validation

**Requirements covered:** FR1, FR14

As a developer user,
I want automated validation on supported operating systems,
So that LeetcodeCLI has release confidence beyond a single local machine.

**Acceptance Criteria:**

**Given** the project supports Windows, macOS, and Linux
**When** CI runs
**Then** GitHub Actions validates the project on Windows, macOS, and Linux runners
**And** each runner executes `go test ./...` and `go build ./...`.

**Given** unit tests run in CI
**When** the test suite executes
**Then** tests do not call the live LeetCode endpoint
**And** GraphQL behavior remains covered by fixtures, fake servers, fake transports, or injected fetchers.

**Given** terminal output tests run across platforms
**When** golden output is compared
**Then** line endings are normalized
**And** rendering remains deterministic across Windows, macOS, and Linux.

### Story 3.3: Prepare Public Release Installation Guidance

**Requirements covered:** FR1, FR13, FR14

As a developer user,
I want platform-specific installation guidance,
So that I can run LeetcodeCLI on my operating system with minimal setup friction.

**Acceptance Criteria:**

**Given** supported operating systems are documented
**When** installation guidance is reviewed
**Then** it includes a macOS Homebrew path
**And** it includes Windows and Linux GitHub Releases binary paths with checksum and PATH setup expectations.

**Given** v1 release packaging may be completed after core implementation
**When** release-readiness documentation is reviewed
**Then** GoReleaser is identified as the intended packaging path when distribution stories begin
**And** no release configuration is required before it is intentionally scoped.

**Given** users evaluate whether to install the tool
**When** release documentation is reviewed
**Then** it repeats the public-data dependency, unofficial status, supported OSes, and no credential/session/config storage boundary
**And** it keeps the public executable name `leetcode` consistent.

## Epic 4: Release Packaging and Public Distribution

Users can install verified LeetcodeCLI release artifacts through the intended public distribution channels: Homebrew on macOS and checksummed GitHub Releases binaries on Windows and Linux.

### Story 4.1: Establish Release Provenance and Versioning

**Requirements covered:** FR1, FR13, FR14

As a developer user,
I want LeetcodeCLI releases to come from a traceable source revision and version tag,
So that I can trust the artifacts I install and understand which v1 behavior they contain.

**Acceptance Criteria:**

**Given** release packaging work begins
**When** source-control and release metadata are reviewed
**Then** the project has a documented release provenance approach that identifies the source revision used for each public artifact
**And** public release artifacts are not published from an untracked or ambiguous source state.

**Given** the project may be published as a public source repository
**When** release readiness is reviewed
**Then** the team has explicitly decided whether the Go module path remains `leetcodecli` or changes before release packaging
**And** any module-path change is completed before GoReleaser artifacts are produced.

**Given** v1 command behavior must remain stable across patch releases
**When** release versioning is documented
**Then** the project uses semantic versioning or an equivalent public release convention
**And** release tags, artifact versions, release notes, and documentation refer to the same version.

**Given** release provenance is established
**When** `go test ./...` and `go build ./...` are run before packaging
**Then** both commands complete successfully
**And** the command surface still exposes only the supported v1 public-profile stats workflow.

### Story 4.2: Configure GoReleaser for Checksummed Cross-Platform Artifacts

**Requirements covered:** FR1, FR13, FR14

As a developer user,
I want GoReleaser to build LeetcodeCLI artifacts for supported operating systems,
So that Windows, macOS, and Linux users can install the same tested CLI release.

**Acceptance Criteria:**

**Given** GoReleaser configuration is added
**When** the release build matrix is inspected
**Then** it builds supported Windows, macOS, and Linux artifacts
**And** each artifact preserves the public executable name `leetcode`.

**Given** GitHub Releases binaries are the Windows and Linux distribution path
**When** GoReleaser packaging runs
**Then** it produces downloadable archives for Windows and Linux
**And** it produces checksums that users can verify before installation.

**Given** macOS distribution uses Homebrew
**When** GoReleaser packaging runs
**Then** it also produces the macOS artifact inputs required by the Homebrew release path
**And** it does not introduce unsupported Windows or Linux package-manager distribution for v1.

**Given** release packaging is configured
**When** a snapshot or dry-run release is executed locally or in CI
**Then** GoReleaser completes without publishing public artifacts
**And** the generated artifact names, checksums, and executable names match the documented install flow.

**Given** v1 has strict scope boundaries
**When** release packaging files are reviewed
**Then** they do not add config, auth, session, token, browser-login, logout, JSON, CSV, dashboard, recommendation, goal, reminder, topic-gap, browser extension, web app, or TUI features.

### Story 4.3: Publish GitHub Releases from Tagged Builds

**Requirements covered:** FR1, FR13, FR14

As a developer user,
I want tagged releases to publish checksummed binaries through GitHub Releases,
So that Windows and Linux installation instructions point to real, verifiable artifacts.

**Acceptance Criteria:**

**Given** a release tag is pushed
**When** the release workflow runs
**Then** it executes the required tests and build validation before publishing artifacts
**And** it fails without publishing if validation does not pass.

**Given** validation passes for a release tag
**When** GoReleaser publishes to GitHub Releases
**Then** the release includes Windows, macOS, and Linux artifacts as configured
**And** it includes checksum files for artifact verification.

**Given** users inspect the GitHub Release
**When** release notes are reviewed
**Then** they identify the release version, supported operating systems, executable name `leetcode`, checksum verification expectation, and v1 trust boundaries
**And** they do not promise authentication, private LeetCode data, JSON/CSV output, dashboards, recommendations, goals, reminders, or other deferred features.

**Given** public command behavior must remain stable
**When** a release is prepared
**Then** release notes call out any breaking command behavior changes before publication
**And** patch releases do not silently change command names, argument meaning, output category names, or no-auth behavior.

### Story 4.4: Publish the macOS Homebrew Distribution Path

**Requirements covered:** FR1, FR13, FR14

As a macOS developer user,
I want to install LeetcodeCLI through Homebrew,
So that I can use the expected macOS package-management flow instead of manually unpacking binaries.

**Acceptance Criteria:**

**Given** Homebrew is the intended macOS distribution channel
**When** release packaging is configured
**Then** GoReleaser updates the agreed Homebrew tap or formula path for LeetcodeCLI
**And** the formula installs the executable as `leetcode`.

**Given** a macOS user follows installation guidance
**When** they run the documented `brew install` command after release publication
**Then** Homebrew installs LeetcodeCLI successfully
**And** running `leetcode help` from the installed binary does not crash.

**Given** the Homebrew formula is reviewed
**When** formula metadata and tests are inspected
**Then** they reference the correct release version, artifact URL, and checksum
**And** they do not require credentials, tokens, cookies, Session Data, or config files.

**Given** Homebrew distribution is live
**When** README and installation docs are reviewed
**Then** placeholder tap or release commands are replaced with actual published Homebrew instructions
**And** Windows and Linux guidance continues to use checksummed GitHub Releases binaries.

### Story 4.5: Validate Public Installation and Release Documentation

**Requirements covered:** FR1, FR13, FR14

As a developer user,
I want release installation instructions to match the published artifacts,
So that I can install LeetcodeCLI without reading source code or guessing which file to download.

**Acceptance Criteria:**

**Given** a release candidate has been packaged
**When** install validation runs for Windows, macOS, and Linux
**Then** each supported operating system has at least one documented install/run path validated against the packaged artifacts
**And** the installed executable can run help commands without crashing.

**Given** Windows and Linux users install from GitHub Releases
**When** installation validation follows the documented flow
**Then** checksum verification is possible before placing the binary on `PATH`
**And** the docs explain archive extraction and `PATH` setup clearly enough to install without reading source code.

**Given** macOS users install from Homebrew
**When** installation validation follows the documented flow
**Then** the `brew install` path works for the published release
**And** the installed command name remains `leetcode`.

**Given** public release docs are updated
**When** README, usage, limitations, and installation docs are compared with Cobra help and release notes
**Then** command examples, supported OSes, unofficial third-party status, public-data dependency, no credential/session/config storage, and no-auth v1 scope remain consistent
**And** the docs no longer describe release artifacts as merely future or intended if they have been published.

**Given** release validation is complete
**When** `go test ./...` and `go build ./...` are run
**Then** both commands pass
**And** documentation tests, if present, still protect installation guidance, trust boundaries, and release artifact expectations.
