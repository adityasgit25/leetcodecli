---
title: "LeetcodeCLI Product Brief"
status: "draft"
created: "2026-06-11"
updated: "2026-06-11"
---

# Product Brief: LeetcodeCLI

## Executive Summary

LeetcodeCLI is a public developer tool for viewing LeetCode profile and practice statistics directly from the terminal. Built with Go and the Cobra framework, it gives developers a fast command-line way to inspect solved-question counts, language breakdowns, and profile-level progress without opening the LeetCode web UI.

The first version is a simple stats reviewer, not a recommendation engine or full practice planner. A user logs in through a browser-based flow, then runs `leetcode stats` to view their own profile stats or `leetcode stats username` to view another user's stats.

The product is meant for anyone who practices on LeetCode: interview candidates, competitive programmers, students, and working developers. For v1, the practical center of gravity is developers who already live in the terminal and want quick visibility into their coding practice progress.

## The Problem

LeetCode users often want a quick sense of progress: total solved questions, profile information, and which programming languages they have used most. Today, that information is usually checked through the website, which is slower for terminal-native developers and awkward to incorporate into local workflows, scripts, dashboards, or routine progress checks.

For public developer-tool users, the friction is not only the data itself. The tool must be easy to install, trustworthy with login/session handling, predictable across platforms, and clear about what data it can and cannot show.

## The Solution

LeetcodeCLI provides a small, focused command-line interface for authenticated LeetCode stats review. After browser login, the user can run:

```text
leetcode stats
leetcode stats username
```

`leetcode stats` defaults to the logged-in user. Passing a username lets the user inspect any public LeetCode profile. The CLI returns a readable terminal table showing profile and practice statistics, starting with total solved questions and solved-question counts by programming language. The experience should feel fast, minimal, and familiar to developers who use tools like `gh`, `kubectl`, or `docker`.

## Known User Needs

- View LeetCode profile information from the CLI.
- See the total number of solved questions.
- See solved-question counts by programming language.
- Explore additional practice stats that help the user understand progress and patterns.
- Install and use the tool with minimal setup friction across common developer environments.
- Log in before viewing stats.
- Read output as pretty terminal tables.

## Who This Serves

LeetcodeCLI serves developers who practice coding problems and prefer terminal workflows. This includes interview candidates checking progress, students tracking practice habits, competitive programmers reviewing language usage, and working developers keeping their algorithm practice visible.

The audience can be broad, but v1 should optimize for the repeat individual user: someone who wants to run a fast command and understand their own or another user's LeetCode progress at a glance.

## Version 1 Scope

In scope:

- Go-based CLI using Cobra.
- Browser-based login flow.
- Browser login is the only supported authentication method.
- `leetcode stats` command defaulting to the logged-in user.
- `leetcode stats username` command for any public LeetCode profile.
- Profile summary.
- Total solved-question count.
- Solved-question counts by programming language.
- Pretty terminal table output.
- Public distribution as a developer tool.
- Windows, macOS, and Linux support.

Out of scope for v1:

- Practice recommendations.
- Goal setting.
- Reminders or streak nudges.
- Topic-gap analysis.
- JSON output.
- Team dashboards.
- Recruiter-focused reporting.
- Browser extension or web app.
- Non-browser login methods, including manual session-cookie paste/import.

## Product Principles

- Simple before smart: show trustworthy stats before adding planning features.
- Terminal-native: make output scannable, compact, and pleasant in a shell.
- Low setup friction: installation and login should be easy enough for casual use.
- Respect user trust: authentication/session behavior must be explicit, secure, and documented.
- Clear limitations: if a stat cannot be fetched reliably, the CLI should say so plainly.

## Success Criteria

- A new user can install the CLI, log in, and run `leetcode stats` without reading extensive documentation.
- A logged-in user can view any public LeetCode profile with `leetcode stats username`.
- Browser login works consistently enough across Windows, macOS, and Linux to be the default authentication path.
- The stats output is readable in a standard terminal without extra setup.
- The language breakdown and solved counts match the data users expect from LeetCode.
- Users understand where credentials or session data are stored.
- The v1 command set remains small enough to document and test thoroughly.

## Key Risks And Unknowns

- LeetCode data access and login behavior may depend on unofficial or unstable endpoints.
- Supporting login increases security expectations for credential/session handling.
- A broad target audience may dilute product choices unless v1 keeps a narrow stats-review focus.
- Pretty terminal tables are useful for humans but limit automation until structured output is added later.

## Early Product Shape

- Form factor: command-line application.
- Primary implementation stack: Go with Cobra.
- Primary data source: LeetCode user/profile data.
- Distribution intent: public developer tool usable by anyone.
- Initial interaction model: browser login first, then run stats commands.

## Open Questions

- What exact profile fields should appear in the v1 table?

## Vision

If v1 succeeds, LeetcodeCLI can become the fastest way for developers to understand their LeetCode progress from the terminal. Later versions could add structured output, topic breakdowns, local history snapshots, progress comparison over time, and lightweight practice planning while preserving the core promise: quick, trustworthy LeetCode stats without leaving the command line.
