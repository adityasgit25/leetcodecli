---
title: "LeetcodeCLI PRD"
status: "draft"
created: "2026-06-11"
updated: "2026-06-11"
---

# PRD: LeetcodeCLI

## 0. Document Purpose

This PRD defines the revised v1 product requirements for LeetcodeCLI, a public command-line tool for viewing public LeetCode profile and practice statistics from the terminal.

This revision supersedes the earlier Browser Login scope. LeetcodeCLI v1 does not include Browser Login, Session Data, own-profile default lookup, config files, token handling, or `leetcode logout`. The v1 happy path is `leetcode stats <username>`.

## 1. Vision

LeetcodeCLI gives developers a fast, terminal-native way to inspect public LeetCode practice progress without opening the LeetCode website. A user runs `leetcode stats <username>` and sees a readable summary of profile and solved-question statistics directly in the shell.

The v1 product is intentionally small: it is a public-profile stats reviewer, not a recommendation engine, practice planner, dashboard, authentication tool, or analytics product. It should feel familiar to developers who already use CLIs such as `gh`, `kubectl`, or `docker`: predictable commands, compact output, useful errors, and minimal setup ceremony.

## 2. Target User

### 2.1 Jobs To Be Done

- Check public LeetCode progress quickly while staying in a terminal workflow.
- See total solved-question progress for a supplied LeetCode username.
- Understand which programming languages appear most in accepted solutions.
- Inspect another public LeetCode profile from the same command surface.
- Install and run the tool across common developer operating systems with minimal setup friction.
- Understand that the tool is unofficial and depends on public LeetCode data availability.

### 2.2 Non-Users (v1)

- Users who need private/authenticated LeetCode account data.
- Users seeking practice recommendations, topic-gap analysis, reminders, or streak coaching.
- Automation-heavy users who require JSON, CSV, or structured output in v1.
- Recruiters or teams needing reports, exports, comparisons, or dashboards.

### 2.3 Key User Journeys

- **UJ-1. Anika checks public practice progress before a study session.**
  Anika runs `leetcode stats anikaUsername` from her terminal. The CLI fetches the requested public LeetCode profile and renders a Stats View containing a Profile Summary, Total Solved Count, and Language Breakdown.

- **UJ-2. Marco reviews a friend's public profile from the terminal.**
  Marco runs `leetcode stats friendUsername`. If the username is valid and public stats are available, the CLI renders the same Stats View shape. If the username is invalid or stats cannot be read, the CLI explains the problem and exits without showing misleading partial data.

- **UJ-3. Priya installs the CLI and reaches first stats.**
  Priya installs LeetcodeCLI, runs `leetcode help`, then runs `leetcode stats leetcode`. The value lands when she sees the first Terminal Table and understands the tool's public-data limitations.

## 3. Glossary

- **Language Breakdown** - Solved-question counts grouped by programming language, based on data available from the LeetCode Profile.
- **LeetCode Profile** - A LeetCode user's profile data, including public profile fields and practice statistics that LeetcodeCLI can retrieve.
- **LeetcodeCLI** - The public command-line application defined by this PRD.
- **Profile Summary** - The top-level identity and progress fields displayed before detailed practice statistics.
- **Public API** - LeetCode's public GraphQL endpoint at `https://leetcode.com/graphql`, used by LeetcodeCLI to retrieve mandatory v1 stats.
- **Public Profile** - A LeetCode Profile addressable by username and visible enough for LeetcodeCLI to retrieve supported v1 stats.
- **Stats Command** - The `leetcode stats <username>` command.
- **Stats View** - The complete human-readable output rendered by the Stats Command.
- **Terminal Table** - A pretty, aligned terminal output format optimized for humans reading in a shell.
- **Total Solved Count** - The total number of solved LeetCode questions shown in the Stats View.

## 4. Features

### 4.1 Installation and CLI Entry Point

**Description:** LeetcodeCLI must be installable and runnable as a normal developer CLI on Windows, macOS, and Linux. The command surface should be small, discoverable, and consistent with common CLI conventions.

#### FR-1: Cross-platform executable

Users can install and run LeetcodeCLI on Windows, macOS, and Linux.

**Consequences (testable):**
- Release validation includes at least one supported install/run path for Windows, macOS, and Linux.
- Running the installed executable with no arguments or help commands does not crash.
- The executable name and documented command examples use `leetcode` consistently.

#### FR-2: Command discovery

Users can discover the Stats Command and basic usage through CLI help.

**Consequences (testable):**
- `leetcode help` lists `stats` as an available command.
- `leetcode help stats` explains `leetcode stats <username>`.
- `leetcode --help` may also expose standard Cobra help behavior.
- Help text states that v1 requires a username and does not include login, logout, session, token, or config setup.

### 4.2 Stats Command and Profile Targeting

**Description:** The Stats Command is the center of v1. It requires a username and targets that public LeetCode profile.

#### FR-3: Public-profile stats by username

Users can run `leetcode stats <username>` to view a Public Profile.

**Consequences (testable):**
- `leetcode stats <username>` targets the supplied username.
- The output identifies the requested or resolved LeetCode Profile.
- Invalid, missing, blocked, or unreadable usernames produce a clear failure message.

#### FR-4: Bare stats help behavior

Users who run `leetcode stats` without a username receive usage guidance instead of a network request.

**Consequences (testable):**
- `leetcode stats` does not call LeetCode.
- `leetcode stats` exits with a usage error.
- The message includes `Username required. Usage: leetcode stats <username>. Run "leetcode help" for help.`

#### FR-5: Targeting boundary

LeetcodeCLI only displays stats it can retrieve for the requested Public Profile.

**Consequences (testable):**
- The CLI does not imply access to private or unavailable profile data.
- If LeetCode returns partial or restricted data, the CLI fails clearly rather than fabricating values.
- Error copy distinguishes not found, unavailable stats, network/API failure, and missing mandatory stats when those states can be reliably detected.

### 4.3 Stats Data Retrieval and Interpretation

**Description:** LeetcodeCLI retrieves LeetCode profile and practice statistics, then presents a trustworthy v1 subset: Profile Summary, Total Solved Count, and Language Breakdown.

#### FR-6: Profile Summary

Users can see a Profile Summary in the Stats View.

**Consequences (testable):**
- The Profile Summary includes username, display name, ranking, reputation, and profile URL.
- The Profile Summary uses the LeetCode GraphQL fields `matchedUser.username`, `matchedUser.profile.realName`, `matchedUser.profile.ranking`, and `matchedUser.profile.reputation`; profile URL is derived from username.
- Missing nullable profile values are shown as `N/A` without breaking table rendering.
- Avatar URL is not shown in the v1 Terminal Table because it is not useful in a text-only terminal output.
- Profile Summary is mandatory for a successful v1 Stats View.

#### FR-7: Total Solved Count

Users can see the Total Solved Count in the Stats View.

**Consequences (testable):**
- The Stats View includes a Total Solved Count field.
- The Total Solved Count matches the `All` row from `matchedUser.submitStatsGlobal.acSubmissionNum`.
- If the Total Solved Count cannot be retrieved, the CLI exits unsuccessfully rather than showing a successful incomplete Stats View.

#### FR-8: Language Breakdown

Users can see solved-question counts by programming language in the Stats View.

**Consequences (testable):**
- The Stats View includes a Language Breakdown table based on `matchedUser.languageProblemCount`.
- Each Language Breakdown row includes a programming language name and solved-question count.
- If the Language Breakdown cannot be retrieved, the CLI exits unsuccessfully rather than showing a successful incomplete Stats View.
- Empty but valid language data is displayed as none found, not as a crash.

#### FR-9: Additional practice stats guardrail

LeetcodeCLI may show additional practice stats only when they are reliable, understandable, and do not turn v1 into a planning product.

**Consequences (testable):**
- Any additional stat has a documented source and display label.
- Additional stats do not introduce recommendations, goals, streak nudges, or topic-gap analysis in v1.
- Additional stats can be removed from the Stats View without breaking FR-6, FR-7, or FR-8.

### 4.4 Terminal Output

**Description:** The Stats View should be compact, readable, and pleasant in a standard terminal. It is optimized for human scanning, not machine parsing in v1.

#### FR-10: Pretty Terminal Table rendering

Users can read the Stats View as one or more pretty Terminal Tables.

**Consequences (testable):**
- Profile Summary and Total Solved Count render in stable human-readable sections.
- Language Breakdown renders with aligned language names and counts.
- The visible section labels are `Profile Summary`, `Total Solved Count`, and `Language Breakdown`.
- Table output remains readable in common light and dark terminal themes without requiring custom fonts or shell plugins.

#### FR-11: Terminal width resilience

Users can read essential stats in standard terminal widths.

**Consequences (testable):**
- The Stats View remains coherent at 80 columns.
- Long usernames or language labels do not corrupt adjacent columns.
- If terminal width cannot be detected, the CLI uses a safe default layout.

#### FR-12: Human-output-only v1

Users receive human-readable output only; structured machine output is deferred.

**Consequences (testable):**
- v1 does not expose JSON output, CSV output, or machine-readable export flags.
- Documentation identifies structured output as out of scope for v1.
- The absence of JSON output does not prevent the Stats Command from succeeding.

### 4.5 Documentation and Trust

**Description:** Because LeetcodeCLI is a public developer tool depending on unofficial public LeetCode data access, docs must make setup, commands, limitations, and data-source risk explicit.

#### FR-13: Minimum user documentation

Users can install, run stats commands, and understand limitations without reading source code.

**Consequences (testable):**
- Documentation includes installation instructions for supported platforms.
- Documentation includes examples for `leetcode help`, `leetcode help stats`, and `leetcode stats <username>`.
- Documentation states that v1 does not include login, logout, Session Data, config files, token storage, recommendations, goal setting, reminders, topic-gap analysis, JSON output, dashboards, browser extensions, or web apps.
- Documentation describes LeetCode public-data dependency and likely failure modes.

#### FR-14: Public release readiness

Users can evaluate whether LeetcodeCLI is safe and appropriate for their local environment.

**Consequences (testable):**
- The project describes its data source and the risk that LeetCode access behavior may change.
- The project documents supported operating systems.
- The project states that LeetcodeCLI is not an official LeetCode product unless that becomes true. [ASSUMPTION: LeetcodeCLI is an unofficial third-party tool.]
- The project states that v1 stores no credentials, tokens, cookies, Session Data, or config files.

## 5. Non-Goals (Explicit)

- LeetcodeCLI v1 will not authenticate users.
- LeetcodeCLI v1 will not include Browser Login.
- LeetcodeCLI v1 will not store Session Data, cookies, tokens, credentials, or config files.
- LeetcodeCLI v1 will not provide `leetcode logout`.
- LeetcodeCLI v1 will not default to the current user's profile.
- LeetcodeCLI v1 will not recommend what problem to solve next.
- LeetcodeCLI v1 will not set goals, reminders, streak nudges, or practice plans.
- LeetcodeCLI v1 will not perform topic-gap analysis.
- LeetcodeCLI v1 will not provide JSON, CSV, or other structured output.
- LeetcodeCLI v1 will not provide team, recruiter, dashboard, browser extension, desktop UI, or web app workflows.
- LeetcodeCLI v1 will not guarantee access to private or unavailable LeetCode data.

## 6. MVP Scope

### 6.1 In Scope

- Public command-line application named LeetcodeCLI.
- `leetcode help` and `leetcode help stats`.
- `leetcode stats <username>` for a Public Profile.
- Bare `leetcode stats` usage error with no network call.
- Profile Summary.
- Total Solved Count.
- Language Breakdown.
- Pretty Terminal Table output.
- Windows, macOS, and Linux support.
- User documentation for install, command usage, limitations, unofficial status, and public-data dependency.

### 6.2 Out of Scope for MVP

- Browser Login, Session Data, logout, config files, tokens, cookies, credentials, and own-profile default lookup.
- Practice recommendations, goal setting, reminders, topic-gap analysis, structured output, dashboards, recruiter reporting, browser extension, desktop UI, and web app surfaces.

## 7. Cross-Cutting Non-Functional Requirements

- **NFR-1: Security and privacy.** LeetcodeCLI v1 must not read, store, print, log, or expose credentials, tokens, cookies, Session Data, or config files.
- **NFR-2: Reliability.** LeetcodeCLI must fail with actionable messages when LeetCode data retrieval, profile targeting, mandatory stats validation, rate limiting/access blocking, or rendering fails. It must not panic or render fabricated stats.
- **NFR-3: Performance.** LeetcodeCLI should avoid unnecessary local delay after LeetCode responds; table rendering should complete quickly enough that network latency dominates normal Stats Command runtime. [ASSUMPTION: A practical v1 target is under 250 ms local processing time after successful data retrieval.]
- **NFR-4: Portability.** LeetcodeCLI must work in common shells on Windows, macOS, and Linux.
- **NFR-5: Maintainability.** Public command behavior must remain stable across v1 patch releases unless a breaking change is documented.

## 8. Constraints and Guardrails

- **LeetCode access stability.** LeetcodeCLI depends on LeetCode's public GraphQL endpoint, which is not a formally versioned public API contract. The product must document this risk and avoid promising permanent compatibility.
- **No-auth boundary.** LeetcodeCLI v1 must not introduce Browser Login, credentials, manual cookie import, token storage, Session Data, config files, own-profile default lookup, or logout.
- **Data boundary.** LeetcodeCLI must display only data available through LeetCode's public GraphQL endpoint for the requested Public Profile.
- **Mandatory stats boundary.** Profile Summary, Total Solved Count, and Language Breakdown are all mandatory for a successful v1 Stats View.
- **Output boundary.** LeetcodeCLI must keep v1 output human-readable and avoid structured export flags until structured output is intentionally scoped.
- **Implementation context.** The source brief specifies Go with Cobra. That implementation choice is preserved in `addendum.md` rather than treated as a user-facing product behavior requirement.

## 9. Public Surface and Versioning

### 9.1 Command Contract

- `leetcode stats <username>` displays the requested Public Profile's Stats View.
- `leetcode stats` exits with a usage error, makes no network call, and points users to `leetcode help`.
- `leetcode help` and `leetcode help stats` document supported v1 usage.

### 9.2 Exit Behavior

- Successful Stats Command runs exit `0` after rendering the Stats View to stdout.
- Explicit help exits `0`.
- Usage or argument failures exit `2`.
- Network, API, mandatory-data, access-blocking, rate-limit, and rendering failures exit `1`.
- Failure messages write concise human-readable copy to stderr.

**Standard v1 error copy:**
- Missing username: `Username required. Usage: leetcode stats <username>. Run "leetcode help" for help.`
- Username not found: `No LeetCode profile found for "<username>". Check the username and try again.`
- Profile or stats unavailable: `Stats for "<username>" are not available from LeetCode right now. Try again later.`
- Public API or network failure: `Could not reach LeetCode. Check your connection and try again.`
- Rate limited or access blocked: `LeetCode blocked or rate-limited the request. Try again later.`
- Mandatory stats missing: `LeetCode did not return required stats for "<username>". Try again later.`
- Rendering failure: `Could not render stats output. Try again later.`

### 9.3 Versioning

- v1 should keep the `leetcode stats <username>` command stable.
- Breaking changes to command names, argument meaning, output category names, or no-auth behavior require a documented version change. [ASSUMPTION: The project will use semantic versioning or an equivalent public release convention.]

### 9.4 Distribution

- macOS distribution uses Homebrew.
- Windows distribution uses checksummed GitHub Releases binaries with documented install and PATH setup.
- Linux distribution uses checksummed GitHub Releases binaries with documented install and PATH setup.
- Additional package-manager distribution for Windows or Linux is optional after v1.

### 9.5 Public API Query Contract

LeetcodeCLI v1 uses `POST https://leetcode.com/graphql` with this query shape:

```graphql
query userProfile($username: String!) {
  matchedUser(username: $username) {
    username
    profile {
      realName
      ranking
      reputation
    }
    submitStatsGlobal {
      acSubmissionNum {
        difficulty
        count
        submissions
      }
    }
    languageProblemCount {
      languageName
      problemsSolved
    }
  }
}
```

Field mapping:
- Profile Summary: `username`, `profile.realName`, `profile.ranking`, `profile.reputation`, and derived profile URL.
- Total Solved Count: `submitStatsGlobal.acSubmissionNum` row where `difficulty` is `All`, using `count`.
- Language Breakdown: `languageProblemCount.languageName` and `languageProblemCount.problemsSolved`.

The query was verified against `https://leetcode.com/graphql` on 2026-06-11 with username `leetcode`.

## 10. Success Metrics

**Primary**

- **SM-1: Time to first stats** - A new user can install LeetcodeCLI, discover usage with `leetcode help`, and run `leetcode stats <username>` without reading source code. Validates FR-1, FR-2, FR-3, FR-4, and FR-13.
- **SM-2: Stats correctness** - Profile Summary, Total Solved Count, and Language Breakdown match LeetCode GraphQL responses for tested public profiles. Validates FR-6, FR-7, and FR-8.
- **SM-3: Command reliability** - The Stats Command succeeds for valid public-profile scenarios in release validation across Windows, macOS, and Linux. Validates FR-3, FR-10, and FR-11.

**Secondary**

- **SM-4: Trust clarity** - Users can find that LeetcodeCLI is unofficial, public-data dependent, and stores no credentials or Session Data. Validates FR-13 and FR-14.
- **SM-5: Scope discipline** - v1 ships without authentication, logout, config, recommendations, planning features, dashboards, or structured output. Validates FR-9 and FR-12.

**Counter-metrics (do not optimize)**

- **SM-C1: Command count growth** - Do not optimize for more commands in v1; adding commands may dilute the narrow stats-review promise.
- **SM-C2: Output density** - Do not optimize for showing every available LeetCode field; too much output can reduce terminal readability.
- **SM-C3: Auth convenience** - Do not optimize for login shortcuts, cookie import, or token persistence in v1.

## 11. Open Questions

No product-scope open questions remain after the 2026-06-11 architecture scope correction. Implementation should still validate LeetCode GraphQL schema stability and failure-state behavior.

## 12. Assumptions Index

- Section 4.5, FR-14 - LeetcodeCLI is an unofficial third-party tool.
- Section 7, NFR-3 - A practical v1 target is under 250 ms local processing time after successful data retrieval.
- Section 9.3 - The project will use semantic versioning or an equivalent public release convention.
