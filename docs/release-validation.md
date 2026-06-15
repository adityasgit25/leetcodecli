# Release Validation

Validation target: `v1.0.0` release candidate.

No public release was published from this workspace; the release is not published. No tag was pushed, no GitHub Release was created, no Homebrew tap push was attempted, and no release credentials were used.

## Local Evidence

- `go test ./...` passed with a workspace-local Go build cache.
- `go build ./...` passed with a workspace-local Go build cache.
- Static tests verify GoReleaser archive targets, `checksums.txt`, the release workflow, release notes, Homebrew cask configuration, installation docs, and v1 trust boundaries.
- GoReleaser unavailable: `goreleaser --version` is not available on PATH in this environment.
- Homebrew unavailable: `brew --version` is not available in this Windows environment.
- Git provenance unavailable: this workspace is not a local Git repository to the `git` command, so live source-revision validation must happen in GitHub Actions.

## Expected Release Assets

For tag `v1.0.0`, expected artifact names include:

- `leetcode_1.0.0_windows_amd64.zip`
- `leetcode_1.0.0_windows_arm64.zip`
- `leetcode_1.0.0_linux_amd64.tar.gz`
- `leetcode_1.0.0_linux_arm64.tar.gz`
- `leetcode_1.0.0_darwin_amd64.tar.gz`
- `leetcode_1.0.0_darwin_arm64.tar.gz`
- `checksums.txt`

Each archive must contain `leetcode` or `leetcode.exe` as appropriate.

## Manual Checks Required After Publishing Approval

1. Install GoReleaser in the release environment.
2. Configure `HOMEBREW_TAP_GITHUB_TOKEN` with content-write access to `adityasgit25/homebrew-leetcodecli`.
3. Push the SemVer tag, for example `v1.0.0`, only from a clean tracked source revision.
4. Let `.github/workflows/release.yml` run `go test ./...`, `go build ./...`, and GoReleaser before publishing.
5. Verify GitHub Releases contains Windows, macOS, Linux archives and `checksums.txt`.
6. On Windows, verify the archive checksum with `Get-FileHash -Algorithm SHA256`, extract with `Expand-Archive`, place `leetcode.exe` on PATH, then run `leetcode help` and `leetcode help stats`.
7. On Linux, verify with `sha256sum -c checksums.txt`, extract with `tar -xzf`, place `leetcode` on PATH, then run `leetcode help` and `leetcode help stats`.
8. On macOS, run `brew install --cask adityasgit25/leetcodecli/leetcode`, then run `leetcode help` and `leetcode help stats`.

## Trust Boundaries

LeetcodeCLI is unofficial and depends on public LeetCode data.

v1 stores no credentials, tokens, cookies, Session Data, or config files. It does not support login, logout, own-profile default lookup, private LeetCode data, JSON output, CSV output, dashboards, recommendations, goals, reminders, or topic-gap analysis.
