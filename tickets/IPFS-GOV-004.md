# IPFS-GOV-004 — Correct Jr Dev Role to Codex Luna

Status: ACCEPTED

Reviewer: Lead Engineer/Reviewer — Codex

Implementation actor: Reviewer (governance-only publication)

Source baseline: `4c937e97c43d4664e3a9eadeba0341be6f606599`

## Objective

Correct the obsolete Hermes Jr Dev label to Codex Luna using `gpt-5.6-luna` while
preserving the established integration and acceptance boundaries.

## Authorized Paths

- `AGENTS.md`
- `TESTING.md`
- `docs/engineering/DEVELOPMENT_ROLES.md`
- `docs/handoff/CURRENT_TASK.md`
- `tickets/IPFS-GOV-004.md`

## Acceptance

- Codex Luna owns integration, command execution, evidence, and developer-drop Git.
- It does not design or author tests.
- No implementation, dependency, generated state, CI, or protocol behavior changes.
- `git diff --check` passes.

## Reviewer Acceptance

Accepted as a governance-only role correction which supersedes the Jr Dev naming in
`IPFS-GOV-003` without rewriting that historical record.
