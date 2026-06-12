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

Supported OSes are Windows, macOS, and Linux. Public release artifacts may not exist yet; the intended public executable name is `leetcode`.

macOS Homebrew release path:

```text
brew install <tap>/leetcode
leetcode help
```

Windows release path:

1. Download the Windows binary archive and checksum from GitHub Releases.
2. Verify the checksum before running the binary.
3. Place `leetcode.exe` in a directory on PATH.
4. Run `leetcode help`, `leetcode help stats`, or `leetcode stats <username>`.

Linux release path:

1. Download the Linux binary archive and checksum from GitHub Releases.
2. Verify the checksum before running the binary.
3. Place `leetcode` in a directory on PATH.
4. Run `leetcode help`, `leetcode help stats`, or `leetcode stats <username>`.

GoReleaser is the intended packaging path when distribution stories begin. This repository does not require release configuration before packaging is intentionally scoped.

## Scope

LeetcodeCLI v1 is unofficial and uses public LeetCode data. Because it depends on public data, public-data availability and LeetCode access behavior may change.

v1 stores no credentials, tokens, cookies, Session Data, or config files. See [docs/usage.md](docs/usage.md) and [docs/limitations.md](docs/limitations.md) for details.
