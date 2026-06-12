# Installation

LeetcodeCLI supports Windows, macOS, and Linux. The public executable name is `leetcode`.

Public release artifacts may not exist yet. When distribution stories begin, GoReleaser is the intended packaging path for GitHub Releases binaries and Homebrew distribution. No release configuration is required before that work is intentionally scoped.

## macOS

The intended macOS path is Homebrew:

```text
brew install <tap>/leetcode
leetcode help
leetcode help stats
leetcode stats <username>
```

## Windows

The intended Windows path is a checksummed binary from GitHub Releases:

1. Download the Windows archive and checksum from GitHub Releases.
2. Verify the checksum before running the binary.
3. Extract `leetcode.exe`.
4. Place `leetcode.exe` in a directory on PATH.
5. Run `leetcode help`, `leetcode help stats`, or `leetcode stats <username>`.

## Linux

The intended Linux path is a checksummed binary from GitHub Releases:

1. Download the Linux archive and checksum from GitHub Releases.
2. Verify the checksum before running the binary.
3. Extract `leetcode`.
4. Place `leetcode` in a directory on PATH.
5. Run `leetcode help`, `leetcode help stats`, or `leetcode stats <username>`.

## Output

`leetcode stats <username>` renders human-readable output with:

- Profile Summary
- Total Solved Count
- Language Breakdown

Bare `leetcode stats` is only a usage-error path.

## Trust Boundaries

LeetcodeCLI is unofficial and depends on public LeetCode data. LeetCode access behavior and public-data availability may change.

v1 stores no credentials, tokens, cookies, Session Data, or config files. v1 has no login, logout, session, config, or private-data workflow.
