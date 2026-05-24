# simpleM

Persistent collaborative workspace with realtime voice communication.

simpleM adalah platform collaboration-first yang menggabungkan:

* persistent workspace
* realtime collaboration
* ambient voice communication
* multi-tab collaborative environment

Inspirasi utama:

* Excalidraw
* Discord
* FigJam
* Notion collaborative

---

# Current Phase

Saat ini project berada pada:

```text
Phase 1 — Foundation & System Design
```

Focus utama phase ini:

* infrastructure setup
* backend foundation
* database connection
* redis integration
* clean architecture foundation
* testing setup
* documentation

---

# Tech Stack

## Frontend

| Technology  | Purpose                    |
| ----------- | -------------------------- |
| Next.js     | Main frontend framework    |
| TailwindCSS | Styling                    |
| Zustand     | State management           |
| Yjs         | Realtime collaboration     |
| Excalidraw  | Whiteboard engine          |
| Tiptap      | Collaborative notes editor |

---

## Backend

| Technology | Purpose            |
| ---------- | ------------------ |
| Go         | Backend language   |
| Fiber      | HTTP framework     |
| PostgreSQL | Main database      |
| Redis      | Presence & pub/sub |
| Zerolog    | Structured logging |

---

## Infrastructure

| Technology           | Purpose               |
| -------------------- | --------------------- |
| Docker               | Local infrastructure  |
| Docker Compose       | Service orchestration |
| PostgreSQL 17 Alpine | Database container    |
| Redis 8 Alpine       | Cache container       |

---

# Project Structure

```text
simplem/
│
├── apps/
│   ├── web/                 # frontend app
│   └── server/              # backend app
│
├── docs/
│   ├── adr/                 # architecture decisions
│   └── features/            # feature documentation
│
├── infrastructure/
│   └── docker/
│
├── packages/
│   └── shared/
│
├── scripts/
│
├── docker-compose.yml
├── README.md
└── .env.example
```

---

# Current Backend Architecture

```text
apps/server
├── cmd/api
├── internal
│   ├── app
│   ├── cache
│   ├── config
│   ├── database
│   ├── handler
│   ├── logger
│   ├── middleware
│   ├── routes
│   └── server
├── tests
├── go.mod
└── go.sum
```

---

# Current Features

## Infrastructure

* Docker Compose setup
* PostgreSQL integration
* Redis integration
* Healthcheck support
* Persistent docker volume
* ARM64 / Mac M1 compatible

---

## Backend Foundation

* Fiber HTTP server
* Structured logger
* Environment config loader
* Graceful shutdown
* Health endpoint
* PostgreSQL connection pool
* Redis connection
* Dependency container foundation
* Hot reload development
* Testing setup

---

# Local Development

## Start Infrastructure

```bash
docker compose up -d
```

---

## Stop Infrastructure

```bash
docker compose down
```

---

## Backend Development

Masuk ke backend:

```bash
cd apps/server
```

Run server:

```bash
go run cmd/api/main.go
```

Run hot reload:

```bash
air
```

Run tests:

```bash
go test ./...
```

---

# Environment Variables

## Backend `.env`

```env
APP_PORT=8080
APP_ENV=development

POSTGRES_HOST=localhost
POSTGRES_PORT=5432
POSTGRES_USER={username}
POSTGRES_PASSWORD={password}
POSTGRES_DB=simplem_db
POSTGRES_SSLMODE=disable

REDIS_HOST=localhost
REDIS_PORT=6379
```

---

# Health Endpoint

Endpoint:

```http
GET /health
```

Expected response:

```json
{
  "status": "ok",
  "services": {
    "database": "ok",
    "redis": "ok"
  }
}
```

---

# Development Principles

Project ini dikembangkan dengan pendekatan:

* incremental development
* validation-driven development
* documentation-first
* maintainable architecture
* testing-first mindset
* scalable realtime architecture

---

# Documentation Structure

## ADR

Architecture Decision Record:

```text
docs/adr/
```

Digunakan untuk:

* architecture decisions
* tradeoff analysis
* long-term maintenance

---

## Feature Documentation

```text
docs/features/
```

Digunakan untuk:

* feature explanation
* architecture notes
* maintenance notes
* future improvements

---

# Current Validation Status

| Component             | Status |
| --------------------- | ------ |
| Docker Infrastructure | ✅      |
| PostgreSQL            | ✅      |
| Redis                 | ✅      |
| Fiber Server          | ✅      |
| Graceful Shutdown     | ✅      |
| Health Endpoint       | ✅      |
| Hot Reload            | ✅      |
| Testing Setup         | ✅      |
| Mac M1 Compatibility  | ✅      |

---

# Known Issues

## PostgreSQL Role Error

Jika muncul error:

```text
FATAL: role "simplem" does not exist
```

biasanya disebabkan karena:

* volume PostgreSQL lama masih tersimpan
* environment PostgreSQL berubah setelah container pertama dibuat

Solusi:

```bash
docker compose down -v
```

Lalu:

```bash
docker compose up -d
```

Karena PostgreSQL hanya membuat user/database saat volume pertama kali dibuat.

---

# Next Planned Phase

Phase berikutnya:

```text
Database Migration System
```

Akan mencakup:

* migration management
* schema versioning
* UUID strategy
* repository foundation
* user entity foundation
* audit fields

---

# Long-Term Vision

simpleM bukan:

* Zoom clone
* Discord clone
* Excalidraw clone

simpleM adalah:

> collaborative presence workspace.