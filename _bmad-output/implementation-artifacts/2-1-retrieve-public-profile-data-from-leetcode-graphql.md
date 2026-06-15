---
baseline_commit: NO_VCS
---

# Story 2.1: Retrieve Public Profile Data from LeetCode GraphQL

Status: review

Completion Note: Ultimate context engine analysis completed - comprehensive developer guide created.

## Story

As a developer user,
I want the stats command to retrieve public LeetCode profile data for a supplied username,
so that the CLI can build stats from the same public-data source every time.

## Requirements Covered

- FR3: Users can run `leetcode stats <username>` to view a Public Profile.
- FR5: LeetcodeCLI only displays stats it can retrieve for the requested Public Profile.

## Acceptance Criteria

1. Given a username is supplied, when the LeetCode client builds the request, then it sends a `POST` request to `https://leetcode.com/graphql`, and it passes the username as a GraphQL variable rather than interpolating it into the query string.
2. Given the LeetCode client sends a request, when request headers are inspected in tests, then the request includes `Content-Type: application/json`, and it includes a stable LeetcodeCLI user agent.
3. Given the client is tested, when unit tests run, then they use injected HTTP behavior, fake servers, fake transports, or fixtures, and they do not call the live LeetCode endpoint.
4. Given LeetCode returns transport, HTTP, or GraphQL-level failures, when the client handles the response, then it classifies not found, unavailable stats, endpoint failure, rate limit/access blocked, and malformed response states distinctly enough for command-level error mapping.

## Tasks / Subtasks

- [x] Add the LeetCode client package. (AC: 1, 2, 3, 4)
  - [x] Create `internal/leetcode`.
  - [x] Keep endpoint URL, static query string, request DTOs, response DTOs, HTTP behavior, and API-specific error classification inside this package.
  - [x] Do not pass raw GraphQL DTOs into `cmd` or `internal/render`.

- [x] Implement request construction. (AC: 1, 2)
  - [x] Use `POST https://leetcode.com/graphql`.
  - [x] Use Go standard `net/http`.
  - [x] Build requests with context.
  - [x] Pass `username` through GraphQL variables.
  - [x] Set `Content-Type: application/json`.
  - [x] Set a stable user agent such as `LeetcodeCLI/1.0`.
  - [x] Use a default timeout for the client.

- [x] Implement decode and client-level classification. (AC: 4)
  - [x] Decode GraphQL responses into private DTO structs.
  - [x] Classify `matchedUser == nil` as username not found.
  - [x] Classify HTTP 403/429 as rate limit/access blocked.
  - [x] Classify transport errors, timeouts, non-2xx responses, malformed JSON, and GraphQL errors distinctly enough for Story 2.5 command mapping.
  - [x] Do not implement automatic retries in v1.

- [x] Add isolated client tests. (AC: 1, 2, 3, 4)
  - [x] Use fake servers, fake transports, or injected `*http.Client`.
  - [x] Assert request method, URL/path, headers, and GraphQL variables.
  - [x] Assert username is not interpolated into the query string.
  - [x] Cover not found, access blocked/rate limited, endpoint failure, GraphQL errors, malformed response, and transport failure.
  - [x] Ensure tests never call the live LeetCode endpoint.

- [x] Verify. (AC: 1, 2, 3, 4)
  - [x] Run `go test ./...`.
  - [x] Run `go build ./...`.

## Dev Notes

### Previous Story Context

- Epic 1 established the CLI command surface and username-required behavior.
- Story 2.1 introduces external API retrieval but should not wire full successful stats rendering yet.
- If `cmd/stats.go` currently has a placeholder path, keep command-level changes minimal and testable.

### Architecture Guardrails

- External API boundary belongs to `internal/leetcode`.
- `cmd` must not know HTTP status codes, GraphQL query text, GraphQL field names, or raw DTO shapes.
- Use context-aware requests and injected HTTP behavior for deterministic tests.
- Unit tests must not make live requests to `https://leetcode.com/graphql`.
- No credentials, cookies, tokens, session files, config files, or browser-login behavior.
- No automatic retry in v1.

### GraphQL Query Shape

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

### Expected Files

- Add: `internal/leetcode/client.go`
- Add: `internal/leetcode/graphql.go`
- Add: `internal/leetcode/types.go`
- Add: `internal/leetcode/client_test.go`
- Add fixtures under `internal/leetcode/testdata/` as needed.
- Update `cmd` only if a test seam or interface is needed for later command wiring.

### Latest Technical Notes

- Use Go standard `net/http`; no third-party HTTP client is required.
- The query shape was verified in the PRD on 2026-06-11 for username `leetcode`.
- Keep dependency additions minimal. This story should not need `go-pretty`, `x/term`, Viper, or GoReleaser.

### References

- [Source: `_bmad-output/planning-artifacts/epics.md` - Story 2.1]
- [Source: `_bmad-output/planning-artifacts/architecture.md` - API & Communication Patterns, API Response Formats, Project Structure & Boundaries]
- [Source: `_bmad-output/planning-artifacts/prds/prd-LeetcodeCLI-2026-06-11/prd.md` - FR3, FR5, public API query contract]
- [Previous: `_bmad-output/implementation-artifacts/1-4-preserve-v1-scope-boundaries-in-the-cli-surface.md`]

## Dev Agent Record

### Agent Model Used

GPT-5 Codex

### Debug Log References

- 2026-06-12: Planned `internal/leetcode` as the API boundary with private GraphQL DTOs, injected HTTP behavior, default endpoint `https://leetcode.com/graphql`, and classified package errors for later command mapping.
- 2026-06-12: Red-phase `go test ./...` failed because `internal/leetcode` client APIs and error taxonomy were undefined.
- 2026-06-12: Verification passed: `gofmt`, `go test ./...`, and `go build ./...`.

### Completion Notes List

- Added `internal/leetcode` with context-aware GraphQL POST construction, default endpoint, stable user agent, and default HTTP timeout.
- Kept GraphQL request/response DTOs private to the package and did not pass raw API shapes into `cmd` or render code.
- Added classified package errors for not found, unavailable, endpoint failure, rate limited/access blocked, malformed response, and future missing-stats mapping.
- Added isolated fake-transport tests covering request shape, username variables, headers, and failure classification without live LeetCode calls.

### File List

- `_bmad-output/implementation-artifacts/2-1-retrieve-public-profile-data-from-leetcode-graphql.md`
- `internal/leetcode/client.go`
- `internal/leetcode/client_test.go`
- `internal/leetcode/errors.go`
- `internal/leetcode/graphql.go`
- `internal/leetcode/types.go`

## Change Log

- 2026-06-12: Added the LeetCode GraphQL client boundary and marked Story 2.1 ready for review.
