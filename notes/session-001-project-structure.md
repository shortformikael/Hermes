# Session 001 — Project Structure

Date: 2026-03-07

## Project Overview

Hermes is a lightweight Go daemon that receives API calls and forwards notifications to Telegram. Intended for Linux servers. End goal: TUI + installable via curl/bash to `/usr/local/bin/`.

## Architecture

The system has two layers:

- **Engine:** daemon + REST API + Telegram sender + logger
- **Chassis:** TUI layered on top showing config, status, health

## File Structure

```
cmd/
  hermes/
    main.go        ← entry point, wires everything together

internal/
  server/          ← starts HTTP server, applies middleware
  api/             ← defines routes, wires controllers
  controllers/     ← handle request logic, call dependencies
  clients/
    telegram/      ← stateless Telegram client
  config/          ← app configuration
  logger/          ← logging
  tui/             ← terminal UI (later)
```

## Key Concepts Covered

### Why `cmd/` sits outside `internal/`

`internal/` is compiler-enforced — nothing outside the module can import from it. It holds all business logic.

`cmd/hermes/main.go` is the entry point (executable, not a library). It imports from `internal/`, initializes all components, injects dependencies, and starts the server. It is the composer; `internal/` packages are the instruments.

### Dependency Injection over Singletons

The Telegram client is initialized once at startup in `main.go` and passed down to controllers that need it. This is preferable to a true global singleton because:

- Dependencies are explicit
- Components are testable (swap real client for a fake)
- No hidden global state

### Concurrency & Mutexes

Go's HTTP server handles each request in its own goroutine automatically. Mutexes are needed to protect **shared mutable state** — e.g. a contact list loaded into memory.

- `sync.Mutex` — exclusive access for reads and writes
- `sync.RWMutex` — allows concurrent reads, exclusive writes (better for data read often, written rarely)

A stateless component (like the Telegram client, which just holds a token and HTTP client) does not need a mutex — there is nothing to corrupt.

### The `-race` flag

Run `go test -race ./...` or `go run -race .` to detect data races at runtime.

### `clients/` directory

Telegram lives under `internal/clients/telegram/` so other future clients (email, Slack, webhooks) have a natural home alongside it.
