# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Build & Run

```bash
# Run the server
go run cmd/server/main.go

# Run all tests
go test ./...

# Run tests for a specific package
go test ./internal/domain/game/...
```

## Architecture

This is a Go backend for **"Sandwich"** — a 3-player card game with 5 rounds. The project follows **Domain-Driven Design (DDD)** with clear layer separation.

### DDD Layers

| Layer | Directory | Purpose |
|-------|-----------|---------|
| **Domain** | `internal/domain/` | Core business logic, entities, value objects, repository interfaces |
| **Application** | `internal/application/` | Use cases, service orchestration |
| **Infrastructure** | `internal/infrastructure/` | External dependencies (DB, Redis, config) |
| **Interface** | `internal/interface/` | HTTP handlers, WebSocket/Socket.IO servers |

### Key Components

**Domain Layer:**
- `domain/game/` — GameState, PlayerState, Card, Suit, Rank, HandType, deck, hand evaluation, scoring rules
- `domain/user/` — User entity with GORM tags
- `domain/matchmaking/` — Matchmaking queue entity

**Application Layer:**
- `application/game/service.go` — CreateMatch, PlayCards, AdvanceRound
- `application/auth/service.go` — Register, Login, ValidateToken (JWT + bcrypt)
- `application/matchmaking/service.go` — JoinQueue, TryCreateMatch

**Infrastructure Layer:**
- `infrastructure/persistence/mysql/` — MySQL driver support
- `infrastructure/persistence/sqlite/` — SQLite driver for testing
- `infrastructure/cache/` — Redis client, GameRepository, MatchmakingRepository implementations
- `infrastructure/wire.go` — DI container with `InitializeApp(cfg AppConfig)`

**Interface Layer:**
- `interface/api/` — Gin router, HTTP handlers for /api/register and /api/login
- `interface/ws/` — Socket.IO server via go-socket.io

### Game Rules

- **3 players** per match, **5 rounds**, 54-card deck (52 standard + 2 jokers)
- Each round: players play cards, combined with 4 public cards (one per round)
- **Scoring (Sandwich rule)**: 1st place wins points equal to round number; 2nd place pays double to 1st and gives points to 3rd
- Hand rankings (low to high): HighCard < Pair < Straight < Flush < StraightFlush < Leopard
- Jokers act as wildcards in hand evaluation

### Configuration

All config via `config.yaml` using Viper:
- `database.driver` — "mysql" or "sqlite"
- `database.dsn` — connection string
- `redis.addr/password/db` — Redis connection
- `auth.jwt_secret` — JWT signing secret
- `app.host/app.port` — server binding

### Dependencies

- **gin-gonic/gin** — HTTP framework
- **googollee/go-socket.io** — Real-time communication
- **gorm** — ORM (MySQL/SQLite drivers)
- **go-redis** — Redis client
- **golang-jwt** — JWT authentication
- **golang.org/x/crypto/bcrypt** — Password hashing
- **spf13/viper** — Configuration management
