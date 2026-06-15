# Installation

LeetcodeCLI supports Windows, macOS, and Linux. The public executable name is `leetcode`.

Public releases are built by GoReleaser from tagged source revisions. GoReleaser produces GitHub Releases archives, `checksums.txt`, release notes, and the macOS Homebrew cask entry.

Release provenance, versioning, and module-path decisions are documented in [release.md](release.md).

## macOS

Use Homebrew:

```text
brew install --cask adityasgit25/leetcodecli/leetcode
leetcode help
leetcode help stats
leetcode stats <username>
```

## Windows

Use a checksummed binary from GitHub Releases:

1. Download `leetcode_<version>_windows_<arch>.zip` and `checksums.txt` from GitHub Releases for the same version.
2. Verify the archive against `checksums.txt` before running the binary.
3. Extract `leetcode.exe`.
4. Place `leetcode.exe` in a directory on PATH.
5. Run `leetcode help`, `leetcode help stats`, or `leetcode stats <username>`.

PowerShell checksum and extraction example:

```powershell
Get-FileHash -Algorithm SHA256 .\leetcode_<version>_windows_<arch>.zip
Select-String -Path .\checksums.txt -Pattern "leetcode_<version>_windows_<arch>.zip"
Expand-Archive .\leetcode_<version>_windows_<arch>.zip -DestinationPath .\leetcode
```

## Linux

Use a checksummed binary from GitHub Releases:

1. Download `leetcode_<version>_linux_<arch>.tar.gz` and `checksums.txt` from GitHub Releases for the same version.
2. Verify the archive against `checksums.txt` before running the binary.
3. Extract `leetcode`.
4. Place `leetcode` in a directory on PATH.
5. Run `leetcode help`, `leetcode help stats`, or `leetcode stats <username>`.

Shell checksum and extraction example:

```sh
sha256sum -c checksums.txt --ignore-missing
tar -xzf leetcode_<version>_linux_<arch>.tar.gz
chmod +x leetcode
```

## Output

`leetcode stats <username>` renders human-readable output with:

- Profile Summary
- Total Solved Count
- Language Breakdown

Bare `leetcode stats` is only a usage-error path.

## Trust Boundaries

LeetcodeCLI is unofficial and depends on public LeetCode data. LeetCode access behavior and public-data availability may change.

v1 stores no credentials, tokens, cookies, Session Data, or config files. v1 has no login, logout, session, config, or private-data workflow.
