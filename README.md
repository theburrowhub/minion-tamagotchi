# Minion Tamagotchi 🥚

A tiny CLI pet that lives in your terminal and **reacts to your AI-agent
activity**. The more you pair with coding agents (Claude Code, Codex, Gemini),
the happier and more energetic your minion gets. Neglect them and it sulks.

> This repository is a **test bed for [Omniagent](https://github.com/theburrowhub/omniagent)** —
> its issues drive Omniagent's autonomous pipeline end-to-end. Specs, limits and
> features here are intentionally small and self-contained.

## What it does

- Reads recent activity from AI coding agents on your machine (session/log
  counts under `~/.claude`, `~/.codex`, `~/.gemini`).
- Maps that activity onto the minion's **stats**: `energy`, `mood`, `hunger`.
- Renders the minion (ASCII) and lets you interact with it from the CLI.

## Commands

| Command | Description |
|---|---|
| `minion status` | Show the minion (ASCII) + current stats. |
| `minion feed`   | Feed the minion (raises energy, lowers hunger). |
| `minion tick`   | Advance time one step: activity feeds mood, hunger rises. |

## Stats model

- `energy` 0–100 — rises with agent activity and feeding, decays over time.
- `hunger` 0–100 — rises over time, lowered by feeding.
- `mood` — derived from energy & hunger: `happy | content | grumpy | sad`.

State persists at `~/.minion/state.json`.

## Build & test

```sh
go build ./...
go test ./...
```

## Layout

- `cmd/minion` — CLI entrypoint.
- `internal/minion` — the minion model (stats, mood, rendering).
- `internal/activity` — reads agent activity (to be implemented).
