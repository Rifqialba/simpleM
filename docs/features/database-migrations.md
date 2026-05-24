# Database Migration System

## Tool
golang-migrate

## Purpose
Schema version management.

## Migration Workflow

Create migration:
```bash
migrate create -ext sql -dir database/migrations migration_name
```
Run migration:
```bash
migrate -path database/migrations \
-database "postgres://..." up
```
Rollback:
```bash
migrate -path database/migrations \
-database "postgres://..." down
```
