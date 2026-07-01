# OMNIAGENT.md — theburrowhub/minion-tamagotchi

Guidance for Omniagent when working in this repository.

## Project context

A small Go CLI: a tamagotchi-style pet that reacts to AI-agent activity. Kept
deliberately simple — it is the end-to-end test bed for Omniagent itself.

## Conventions

- Go 1.22+. Standard library only where possible; avoid new dependencies unless
  an issue explicitly calls for one.
- Package layout: `cmd/minion` (CLI), `internal/minion` (model), `internal/activity`
  (agent-activity readers).
- Keep functions small and pure where you can; the model must be unit-testable
  without touching the filesystem.
- Match the existing style; run `gofmt`.

## Build / test / verify

- Build: `go build ./...`
- Test: `go test ./...`
- Both must pass before a change is complete. CI runs the same on every PR.

## Per-stage guidance

- **Triage**: label with `area:cli`, `area:model`, or `area:activity`; set a
  priority. Keep the classification one line.
- **Refinement**: express acceptance criteria as concrete, testable assertions
  (e.g. "hunger never exceeds 100").
- **Development**: implement the smallest change that satisfies the criteria,
  with a unit test. Do not add dependencies gratuitously. Keep the CLI's output
  stable unless the issue changes it.
- **Verify**: `go build ./...` and `go test ./...` must be green.
- **Follow-up**: address each review comment; if a comment is out of scope, say
  so briefly rather than expanding the PR.

## Guardrails

- Never delete existing commands or break their output contract without an issue
  asking for it.
- Stats are clamped to 0–100; preserve that invariant.
