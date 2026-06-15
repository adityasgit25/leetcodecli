# Usage

Supported v1 commands:

```text
leetcode help
leetcode help stats
leetcode stats <username>
```

Use `leetcode help` to see the root command help. Use `leetcode help stats` to see stats command help. Use `leetcode stats <username>` to fetch and render public profile stats for a specific LeetCode username.

The stats view is human-readable terminal output with these sections:

- Profile Summary
- Total Solved Count
- Language Breakdown

Bare `leetcode stats` is not a successful stats command. It returns usage guidance because v1 requires a username.

LeetcodeCLI is unofficial and depends on public LeetCode data. Because it depends on public data, public-data availability and LeetCode access behavior may change.

v1 stores no credentials, tokens, cookies, Session Data, or config files. It has no login, logout, session, config, or private-data workflow.
