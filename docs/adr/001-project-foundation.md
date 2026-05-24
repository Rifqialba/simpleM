# ADR 001 — Project Foundation

## Status
Accepted

## Context
simpleM membutuhkan arsitektur yang scalable, maintainable, dan cocok untuk realtime collaborative environment.

## Decision
Menggunakan monorepo structure dengan:
- Next.js frontend
- Go backend
- PostgreSQL
- Redis
- Docker-based local development

## Consequences
- Struktur project lebih konsisten
- Shared contracts lebih mudah
- Development workflow lebih stabil