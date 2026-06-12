# Addendum: LeetcodeCLI PRD

## Source Inputs

- `_bmad-output/planning-artifacts/briefs/brief-LeetcodeCLI-2026-06-11/brief.md`

## Architecture Scope Correction

During architecture planning on 2026-06-11, the v1 product scope was simplified:

- Browser Login is deferred.
- Session Data persistence is deferred.
- `leetcode logout` is deferred.
- Own-profile default lookup through `leetcode stats` is deferred.
- Config files, token storage, cookie storage, and credential handling are out of scope for v1.

The revised v1 happy path is:

```text
leetcode stats <username>
```

Bare `leetcode stats` is a usage error, performs no network call, and points users to `leetcode help`.

## Preserved Implementation Context

The product brief states that LeetcodeCLI is a Go-based CLI using Cobra. The PRD preserves that as implementation context instead of making it a user-facing product behavior requirement.

The architecture selected a minimal Go module plus Cobra scaffold. The concrete local module path for implementation is `leetcodecli` because this workspace does not yet have a Git remote. If the project is later published as a public source repository, the module path may be intentionally changed before release packaging.

## Decisions Preserved From Brief

- v1 is a stats reviewer, not a recommendation engine or full practice planner.
- The primary command remains `leetcode stats`, now requiring `<username>` for successful execution.
- `leetcode stats <username>` targets a public LeetCode profile.
- Profile Summary, Total Solved Count, and Language Breakdown are all mandatory for v1.
- Profile Summary fields are username, display name, ranking, reputation, and derived profile URL.
- Output is a pretty human-readable terminal table.
- macOS distribution uses Homebrew.
- Windows and Linux distribution use checksummed GitHub Releases binaries with documented install and PATH setup.
- JSON output, dashboards, browser extensions, web apps, reminders, goals, and topic-gap analysis are out of scope for v1.
- Supported operating systems are Windows, macOS, and Linux.

## Deferred From Brief

- Browser Login.
- Session Data persistence.
- `leetcode logout`.
- Authenticated own-profile default lookup.
- Non-browser authentication and manual cookie import.
- Any private or authenticated LeetCode data access.

## Downstream Architecture Notes

- LeetCode data access uses `POST https://leetcode.com/graphql` with `matchedUser`, `submitStatsGlobal.acSubmissionNum`, and `languageProblemCount`.
- Architecture should isolate LeetCode data retrieval from terminal rendering so endpoint changes do not spread through the CLI.
- Stats rendering should be tested separately from data retrieval with deterministic sample payloads.
- Terminal table rendering should be tested at common widths, especially 80 columns.
- Error states should be modeled explicitly: missing username, username not found, profile unavailable, endpoint failure, rate limit/access blocked, mandatory stats missing, and rendering failure.
- Unit tests should not call the live LeetCode endpoint.

## Standard v1 Error Copy

- Missing username: `Username required. Usage: leetcode stats <username>. Run "leetcode help" for help.`
- Username not found: `No LeetCode profile found for "<username>". Check the username and try again.`
- Profile or stats unavailable: `Stats for "<username>" are not available from LeetCode right now. Try again later.`
- Public API or network failure: `Could not reach LeetCode. Check your connection and try again.`
- Rate limited or access blocked: `LeetCode blocked or rate-limited the request. Try again later.`
- Mandatory stats missing: `LeetCode did not return required stats for "<username>". Try again later.`
- Rendering failure: `Could not render stats output. Try again later.`

## Research Note

A lightweight external landscape check did not introduce requirements beyond the source brief. A direct read-only request to LeetCode GraphQL on 2026-06-11 verified that the selected query shape returns username/profile fields, solved counts, and language counts for username `leetcode`.
