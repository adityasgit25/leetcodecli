# LeetcodeCLI

LeetcodeCLI is a small terminal tool for viewing public LeetCode profile statistics.

## Usage

```text
leetcode help
leetcode help stats
leetcode stats <username>
```

`leetcode stats <username>` fetches public profile data and renders a human-readable stats view with:

- Profile Summary
- Total Solved Count
- Language Breakdown

Bare `leetcode stats` is a usage error because v1 requires an explicit username.

## Installation

Supported OSes are Windows, macOS, and Linux. Public releases are built by GoReleaser from tagged source revisions, and the public executable name is `leetcode`.

macOS Homebrew release path:

```text
brew install --cask adityasgit25/leetcodecli/leetcode
leetcode help
```

Windows release path:

1. Download `leetcode_<version>_windows_<arch>.zip` and `checksums.txt` from GitHub Releases for the same version.
2. Verify the archive against `checksums.txt` before running the binary.
3. Extract `leetcode.exe`.
4. Place `leetcode.exe` in a directory on PATH.
5. Run `leetcode help`, `leetcode help stats`, or `leetcode stats <username>`.

Linux release path:

1. Download `leetcode_<version>_linux_<arch>.tar.gz` and `checksums.txt` from GitHub Releases for the same version.
2. Verify the archive against `checksums.txt` before running the binary.
3. Extract `leetcode`.
4. Place `leetcode` in a directory on PATH.
5. Run `leetcode help`, `leetcode help stats`, or `leetcode stats <username>`.

GoReleaser produces GitHub Releases archives, checksums, release notes, and the macOS Homebrew cask entry.

Release provenance, versioning, and module-path decisions are documented in [docs/release.md](docs/release.md).

## Scope

LeetcodeCLI v1 is unofficial and uses public LeetCode data. Because it depends on public data, public-data availability and LeetCode access behavior may change.

v1 stores no credentials, tokens, cookies, Session Data, or config files. See [docs/usage.md](docs/usage.md) and [docs/limitations.md](docs/limitations.md) for details.
