# BitBook IPFS Fork Testing Strategy

This policy applies SQLite's reliability-oriented testing strategy, as adopted by Keel,
to BitBook's legacy IPFS/libp2p fork. Because the upstream tree is large and old,
evidence is concentrated on the packages and protocol boundaries BitBook actually
changes and maintains. SQLite's scale and 100% branch coverage are inspirations, not
claims about this repository's current maturity.

## Standing Rules

1. **Tests lead implementation.** Every behavior change starts with a focused test that
   fails for the intended reason. The active ticket names the authorized red and green
   commands. Production code follows only after the red result is understood.
2. **Falsify important tests before trusting them.** Temporarily suppress or break the
   mechanism under test and prove the test fails. Do not commit the falsification change;
   report the method and result.
3. **One regression test per bug.** A bug fix is incomplete until a test reproduces the
   bug before the fix and passes after it.
4. **Assert non-vacuous outcomes.** Zero peers, empty provider results, skipped transfers,
   ignored errors, and timeout-only loops are not successful distributed tests unless
   emptiness is the required behavior.
5. **Exercise the real path.** Discovery, routing, transfer, persistence, and restart
   tests may not bypass the mechanism they claim to prove with manual peer injection or
   pre-seeded answers.

## Required Techniques

Use these where a BitBook-specific ticket touches the corresponding boundary:

- native Go fuzzing for multiformats, frames, records, protocol messages, and other
  attacker-controlled parsers;
- boundary tests immediately below, at, and above defined message, block, peer, timeout,
  retry, and resource limits;
- property tests for round trips, determinism, canonical encoding, idempotence,
  monotonicity, bounded resources, and compatibility invariants;
- failure injection for disconnects, DHT timeouts, cancellation, partial reads/writes,
  storage faults, corrupted records, and restart recovery;
- compound-failure tests when another fault occurs during cleanup or recovery; and
- in-process multi-node tests for discovery, provider lookup, routing, transfer,
  partition, node loss, and convergence without manual shortcuts.

Tests must be offline, deterministic, credential-free, and independent of public peers,
wall-clock luck, or mutable third-party services. A ticket that changes a fork-specific
protocol must include compatibility fixtures for the previous accepted BitBook behavior.
Protocol- and storage-critical results should have an independent oracle where practical:
a canonical fixture, a maintained upstream behavior, or an invariant checked through a
separate path. Tests must detect leaked goroutines, file descriptors, sockets, and
temporary state when the touched boundary owns those resources.

## Ticket and Review Contract

Each implementation ticket must state:

- why the fork must change instead of consuming a maintained upstream component;
- the invariant or observable protocol behavior being proved;
- authorized test paths before production paths;
- the exact targeted red command and expected failure;
- the exact targeted green command and broader acceptance commands;
- how at least one high-value test will be falsified; and
- relevant fuzz, property, failure, race, compatibility, and multi-node coverage.

If dependencies, build inputs, protocol parsing, cryptography, storage, or downstream
release content can change, the ticket must also name the applicable security scans,
their exact commands, and the finding threshold that blocks acceptance.

The ticket-authorized implementation developer authors bounded test source before
production source, but does not execute tests. Jr Dev — Codex Luna integrates the test-only
drop first, records the expected red result, integrates the production drop, runs the
targeted and broader acceptance commands, and publishes the implementation evidence.
The reviewer independently inspects the tests, rejects tautological or shortcut proofs,
and accepts or rejects the integrated evidence.

## CI and Coverage

- Run changed and maintained concurrent packages under the race detector.
- Keep fuzz seeds as ordinary regression inputs and run affected fuzz targets for a
  bounded time.
- Coverage instrumentation tests the test suite. Re-run release-critical behavior in the
  configuration actually delivered downstream.
- Coverage is a ratchet for BitBook-maintained packages, not a global target imposed on
  the entire inherited upstream tree. Raise package floors as tests land; never lower a
  floor without a ticketed rationale.
- Do not rebuild downstream desktop release artifacts for an isolated substrate change
  unless the ticket identifies a compatibility risk requiring that proof.

Human or model review remains necessary for architecture and protocol compatibility, but
is not a substitute for executable tests.

## Security and Supply-Chain Evidence

- Scan Go source and the resolved legacy dependency graph with pinned,
  ecosystem-appropriate tools. Scan committed content for secrets before acceptance.
- Record inherited findings as an explicit reviewed baseline rather than silently
  normalizing them. New findings fail the ratchet; critical or plausibly exploitable
  findings require immediate review regardless of the baseline.
- A suppression requires a ticketed rationale, owner, affected version or path, and
  expiry or removal condition. Severity labels alone do not replace exploitability
  review.
- When this repository directly supplies a release artifact, generate a machine-readable
  SPDX or CycloneDX SBOM from its exact resolved inputs, attach it to the release, and
  scan the artifact or SBOM before publication. Downstream-only releases own their own
  final SBOM; routine pushes do not regenerate one.
- Pin scanner and SBOM-generator versions. Preserve their reports as CI or release
  evidence. Network access is limited to fetching tools and signed advisory data; tests
  must not depend on a mutable service response.

Lineage: [How SQLite Is Tested](https://www.sqlite.org/testing.html).
