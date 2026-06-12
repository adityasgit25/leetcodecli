# Reconciliation: Product Brief to Revised PRD

Source: `_bmad-output/planning-artifacts/briefs/brief-LeetcodeCLI-2026-06-11/brief.md`

## Covered Directly

- Public developer CLI for LeetCode stats: PRD sections 1, 2, 4, and 6.
- `leetcode stats <username>` for public LeetCode profiles: FR-3 and Public Surface.
- Profile summary: FR-6.
- Total solved count: FR-7.
- Solved counts by programming language: FR-8.
- Pretty terminal tables: FR-10 and FR-11.
- Windows, macOS, and Linux support: FR-1, MVP Scope, and NFR-4.
- Trust and clear limitations: FR-13, FR-14, NFR-1, NFR-2, and Constraints.
- v1 out-of-scope items: Non-Goals and MVP Scope.

## Revised From Brief

Architecture planning simplified the v1 scope after user confirmation:

- Browser Login is deferred.
- Session Data persistence is deferred.
- `leetcode logout` is deferred.
- Authenticated own-profile default lookup is deferred.
- Config, token, cookie, and credential storage are out of scope.

The revised v1 happy path is:

```text
leetcode stats <username>
```

Bare `leetcode stats` is a usage error, performs no network call, and points to `leetcode help`.

## Expanded Into Requirements

- Missing username behavior became FR-4.
- Username/profile targeting boundaries became FR-5.
- Additional practice stats were bounded by FR-9 so v1 does not silently expand into recommendations or planning.
- Human-output-only scope became FR-12.
- Public release trust requirements became FR-13 and FR-14.
- Command contract, exit behavior, and versioning were added as developer-product public-surface requirements.

## Open Gaps Resolved By Review

- Exact v1 Profile Summary fields are username, display name, ranking, reputation, and derived profile URL.
- LeetCode data source is `POST https://leetcode.com/graphql`.
- v1 excludes login, logout, and Session Data.
- macOS distribution is Homebrew; Windows and Linux use checksummed GitHub Releases binaries with documented install and PATH setup.
- Standard unavailable-data, network, rate-limit/access-blocked, rendering, and missing-username copy is captured in the PRD.
