# IPFS-GOV-002 — Adopt SQLite/Keel Test-First Strategy

Status: ACCEPTED

Reviewer: Lead Engineer/Reviewer — Codex

Implementation actor: Reviewer (governance-only publication)

Source baseline: `46e45f2a3cb26e11d755c20d1e2d04fabf71fc7e`

## Objective

Make test-first development, test falsification, regression coverage, hostile-boundary
testing, real multi-node proofs, compatibility fixtures, security scanning, release SBOM
evidence, and maintained-package coverage and vulnerability ratchets standing
requirements for future fork changes.

## Authorized Paths

- `AGENTS.md`
- `TESTING.md`
- `docs/handoff/CURRENT_TASK.md`
- `tickets/IPFS-GOV-002.md`

## Acceptance

- The policy requires red-before-green evidence and falsification of important tests.
- It covers fuzz, property, failure-injection, race, regression, compatibility, and real
  multi-node testing without imposing a global target on inherited upstream code.
- It requires pinned source, dependency, and secret scanning plus release-time SBOM and
  artifact evidence while explicitly baselining inherited legacy findings.
- No production or test source, dependency, generated state, or network behavior changes.
- `git diff --check` passes for the authorized paths.

## Reviewer Acceptance

Accepted as a governance-only testing baseline. No implementation task is authorized.
