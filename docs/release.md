# Release

LeetcodeCLI public artifacts must be traceable to a source revision, a version tag, release notes, and checksums. The public executable name is `leetcode`.

## Provenance Policy

Every public release artifact must be built from a tracked source revision and a version tag. The release source revision is the commit checked out by the tagged GitHub Actions release workflow, and the version tag is the Git tag that triggered that workflow.

Release artifacts must not be published from `NO_VCS`, an uncommitted source state, a dirty worktree, or any ambiguous source snapshot. If the release job cannot identify a source revision and version tag, the job must fail before publishing artifacts.

Release metadata is captured in these places:

- GitHub Release: version tag, release notes, checksummed artifact uploads, and source revision associated with the tag.
- GoReleaser output: artifact names, artifact versions, and checksum file entries.
- Project documentation: installation commands and validation notes for the same version.

## Module Path Decision

The public Go module path is `github.com/adityasgit25/leetcodecli`. Earlier planning mentioned the local module path `leetcodecli` because the workspace started before a public repository path was known. The current module path is final for release packaging unless a future public repository move is explicitly completed before artifacts are built.

The executable name remains `leetcode`; documentation and release artifacts must not rename the installed command to the module path.

## Versioning Rules

LeetcodeCLI uses SemVer for public releases. Version tags use the `vMAJOR.MINOR.PATCH` format, such as `v1.0.0`.

The version tag, artifact version, release notes heading, installation documentation, and validation notes must refer to the same version. Patch releases must not silently change command names, argument meaning, output section names, or no-auth behavior.

Release notes must call out any breaking command behavior change before publication. Stable v1 behavior is:

```text
leetcode help
leetcode help stats
leetcode stats <username>
```

## Trust Boundaries

LeetcodeCLI v1 is unofficial and depends on public LeetCode data. Public-data availability and LeetCode access behavior may change.

v1 stores no credentials, tokens, cookies, Session Data, or config files. It does not support login, logout, own-profile default lookup, private LeetCode data, JSON output, CSV output, dashboards, recommendations, goals, reminders, topic-gap analysis, browser extensions, desktop UI, web apps, or TUI features.

