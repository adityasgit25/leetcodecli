# LeetcodeCLI Release Notes Template

## Version

Version: `vMAJOR.MINOR.PATCH`

## Supported Platforms

- Windows
- macOS
- Linux

## Installation Assets

- Public executable: `leetcode`
- Windows and Linux users should download the matching GitHub Releases archive and verify it with `checksums.txt` before placing the binary on PATH.
- macOS users should install through the documented Homebrew command once the tap release is published.

## Checksum Verification

Release notes must remind users that checksum verification is expected before installing archives from GitHub Releases.

## Breaking Command Behavior

Breaking command behavior: document any change here before publication.

List any breaking command behavior before publication, including command names, argument meaning, output section names, or no-auth behavior. Use "None" for patch releases with no breaking command behavior.

## v1 Trust Boundaries

LeetcodeCLI is unofficial and depends on public LeetCode data.

v1 stores no credentials, tokens, cookies, Session Data, or config files.

Out of scope for v1: login, logout, own-profile default lookup, private LeetCode data, JSON output, CSV output, dashboards, recommendations, goals, reminders, topic-gap analysis, browser extensions, desktop UI, web apps, and TUI features.
