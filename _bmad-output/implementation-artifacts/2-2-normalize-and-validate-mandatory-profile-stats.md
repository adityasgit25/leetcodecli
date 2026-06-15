---
baseline_commit: NO_VCS
---

# Story 2.2: Normalize and Validate Mandatory Profile Stats

Status: review

Completion Note: Ultimate context engine analysis completed - comprehensive developer guide created.

## Story

As a developer user,
I want public LeetCode data normalized into trustworthy profile stats,
so that the CLI never renders fabricated or misleading partial results.

## Requirements Covered

- FR5: LeetcodeCLI only displays stats it can retrieve for the requested Public Profile.
- FR6: Users can see a Profile Summary in the Stats View.
- FR7: Users can see the Total Solved Count in the Stats View.
- FR8: Users can see solved-question counts by programming language in the Stats View.

## Acceptance Criteria

1. Given LeetCode returns a successful GraphQL response, when the response is decoded, then raw GraphQL DTOs remain inside `internal/leetcode`, and the rest of the application receives normalized profile stats rather than raw response payloads.
2. Given a public profile response contains profile fields, when normalization runs, then it produces username, display name, ranking, reputation, and a profile URL derived from username, and nullable display fields can be represented as `N/A` without breaking rendering.
3. Given `matchedUser.submitStatsGlobal.acSubmissionNum` contains an `All` row, when normalization runs, then Total Solved Count is set from the `All` row `count`, and missing Total Solved Count is classified as a mandatory-stats failure.
4. Given `matchedUser.languageProblemCount` is present, when normalization runs, then it produces language-name and solved-count rows, and an empty but valid language list remains a successful none-found state.
5. Given mandatory profile summary, total solved, or language breakdown data is missing, when validation runs, then the stats command cannot proceed to successful rendering, and tests cover success, profile not found, missing total, missing languages, empty languages, and GraphQL error fixtures.

## Tasks / Subtasks

- [x] Define normalized stats types. (AC: 1, 2, 3, 4)
  - [x] Export product/domain types from `internal/leetcode`, such as `ProfileStats`, `ProfileSummary`, and `LanguageCount`.
  - [x] Use product names, not GraphQL names, in exported types.
  - [x] Keep raw GraphQL DTO structs private to `internal/leetcode`.

- [x] Implement normalization. (AC: 1, 2, 3, 4)
  - [x] Convert `matchedUser.username` into normalized username.
  - [x] Convert `profile.realName`, `profile.ranking`, and `profile.reputation` into normalized summary fields.
  - [x] Derive profile URL from username in normalized/domain code, not renderer code.
  - [x] Find the `All` row in `submitStatsGlobal.acSubmissionNum` and use its `count` as Total Solved Count.
  - [x] Convert `languageProblemCount` entries into language-name and solved-count rows.

- [x] Implement mandatory validation. (AC: 3, 4, 5)
  - [x] Fail if profile summary/matched user is missing.
  - [x] Fail if Total Solved Count is missing because no `All` row exists.
  - [x] Fail if `languageProblemCount` is missing/null.
  - [x] Treat present but empty language list as successful.
  - [x] Use typed/classified errors that Story 2.5 can map to exact user-facing stderr.

- [x] Add fixtures and tests. (AC: 1, 2, 3, 4, 5)
  - [x] Add `profile_success.json`.
  - [x] Add `profile_not_found.json`.
  - [x] Add `profile_missing_total.json`.
  - [x] Add `profile_missing_languages.json`.
  - [x] Add `profile_empty_languages.json`.
  - [x] Add `graphql_error.json` if not already present from Story 2.1.
  - [x] Test nullable display fields normalize to values renderable as `N/A`.
  - [x] Test no raw GraphQL DTO type is needed outside `internal/leetcode`.

- [x] Verify. (AC: 1, 2, 3, 4, 5)
  - [x] Run `go test ./...`.
  - [x] Run `go build ./...`.

## Dev Notes

### Previous Story Context

- Story 2.1 owns request construction, GraphQL DTO decode, and client-level error classification.
- This story completes the product/domain normalization boundary before rendering exists.

### Architecture Guardrails

- Data flow: `username argument -> LeetCode GraphQL client -> normalized profile stats -> mandatory-field validation -> terminal renderer`.
- Renderer must receive normalized stats only.
- Missing mandatory stats fail before rendering.
- Empty but valid language data renders later as none found.
- Do not fabricate values. Nullable display fields can become `N/A`; mandatory sections cannot.
- Do not add renderer table layout in this story.

### Suggested Normalized Shape

```go
type ProfileStats struct {
    Summary        ProfileSummary
    TotalSolved    int
    LanguageCounts []LanguageCount
}
```

Use a type or flag that can distinguish missing `LanguageCounts` from present empty `LanguageCounts`.

### Expected Files

- Update/add: `internal/leetcode/types.go`
- Add/update: `internal/leetcode/normalize.go`
- Add/update: `internal/leetcode/normalize_test.go`
- Add/update fixtures: `internal/leetcode/testdata/*.json`
- Update client tests if the public client now returns normalized stats.

### Test Guidance

- Test normalization as pure logic where possible.
- Keep fixture JSON close to the actual GraphQL shape.
- Avoid live LeetCode requests.
- Assert profile URL derivation is deterministic, for example `https://leetcode.com/u/<username>/` if that format is chosen.

### References

- [Source: `_bmad-output/planning-artifacts/epics.md` - Story 2.2]
- [Source: `_bmad-output/planning-artifacts/architecture.md` - Data Architecture, Data Exchange Formats, API Response Formats]
- [Source: `_bmad-output/planning-artifacts/prds/prd-LeetcodeCLI-2026-06-11/prd.md` - FR5, FR6, FR7, FR8, public API field mapping]
- [Previous: `_bmad-output/implementation-artifacts/2-1-retrieve-public-profile-data-from-leetcode-graphql.md`]

## Dev Agent Record

### Agent Model Used

GPT-5 Codex

### Debug Log References

- 2026-06-12: Planned normalized `ProfileStats`, `ProfileSummary`, and `LanguageCount` types in `internal/leetcode`, with `FetchProfileStats` as the public client method for later command wiring.
- 2026-06-12: Red-phase `go test ./...` failed because normalized types, `normalizeProfileStats`, and `FetchProfileStats` were undefined.
- 2026-06-12: Boundary scan passed: no raw GraphQL DTO names were found outside `internal/leetcode`.
- 2026-06-12: Verification passed: `gofmt`, `go test ./...`, and `go build ./...`.

### Completion Notes List

- Added exported normalized product types: `ProfileStats`, `ProfileSummary`, and `LanguageCount`.
- Implemented normalization from private GraphQL DTOs, including profile URL derivation and `All` row extraction for total solved.
- Added mandatory validation for missing profile summary, missing total solved, and missing/null language stats while allowing present empty languages.
- Added fixture-backed tests for success, not found, missing total, missing languages, empty languages, GraphQL errors, and nullable display fields.
- Added `Client.FetchProfileStats` so later command wiring can fetch validated normalized stats without raw GraphQL DTOs.

### File List

- `_bmad-output/implementation-artifacts/2-2-normalize-and-validate-mandatory-profile-stats.md`
- `internal/leetcode/client.go`
- `internal/leetcode/normalize.go`
- `internal/leetcode/normalize_test.go`
- `internal/leetcode/testdata/graphql_error.json`
- `internal/leetcode/testdata/profile_empty_languages.json`
- `internal/leetcode/testdata/profile_missing_languages.json`
- `internal/leetcode/testdata/profile_missing_total.json`
- `internal/leetcode/testdata/profile_not_found.json`
- `internal/leetcode/testdata/profile_success.json`
- `internal/leetcode/types.go`

## Change Log

- 2026-06-12: Added normalized profile stats and mandatory validation, then marked Story 2.2 ready for review.
