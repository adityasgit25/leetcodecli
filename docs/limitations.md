# Limitations

LeetcodeCLI v1 supports:

```text
leetcode help
leetcode help stats
leetcode stats <username>
```

The successful stats view includes Profile Summary, Total Solved Count, and Language Breakdown.

LeetcodeCLI is unofficial. Because it depends on public data, public-data availability and LeetCode access behavior may change, so v1 cannot promise permanent compatibility with LeetCode's public GraphQL behavior.

v1 stores no credentials, tokens, cookies, Session Data, or config files.

Out of scope for v1:

- login
- logout
- own-profile default lookup
- recommendations
- goals
- reminders
- topic-gap analysis
- JSON output
- dashboards
- browser extensions
- desktop UI
- web apps

Bare `leetcode stats` is only a usage-error path. The supported stats command shape is `leetcode stats <username>`.
