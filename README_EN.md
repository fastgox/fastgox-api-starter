# FastGox API Starter

English | [中文](README.md)

A minimal Go API starter template with layered architecture, ready to use out of the box.

## Project Structure

```
main.go                    # Entry point
config/                    # Configuration files (YAML)
deploy/docker/             # Dockerfile + docker-compose
docs/
  db/                      # SQL migration files
  swagger/                 # Swagger docs
src/
  core/                    # Core: config loader, database, session
  models/                  # Data models: entity, DTO, config
  repository/              # Data access layer
  services/                # Business logic layer
  router/                  # Routes + middleware
  utils/                   # Utility functions
  pkg/                     # External service wrappers
```

## Features

- Layered architecture: Repository -> Service -> Router
- Auto route registration (via init())
- SQL migration version control (schema_migrations)
- Graceful shutdown (production ready)
- Auto-generated Swagger docs
- JWT authentication + CORS middleware
- One-command Docker deployment
- Task-based project management

## Quick Start

### Prerequisites

- Go 1.24+
- [Task](https://taskfile.dev) (recommended)
- Docker / Docker Compose (for deployment)

### Development

```bash
# Initialize project
task setup

# Start dev server
task dev

# Format code
task fmt

# Generate Swagger docs
task swagger
```

### Build and Release

```bash
# Local build
task build

# Release build (clean + fmt + tidy + swagger + build)
task release
```

### Docker Deployment

```bash
# Build image
task docker-build

# Build and start (with PostgreSQL)
task deploy

# View logs
task docker-logs

# Stop services
task docker-down
```

### All Available Tasks

```bash
task --list
```

## Configuration

Config files are in the `config/` directory, selected by `APP_ENV`:

```bash
APP_ENV=dev   # loads config/dev.yaml
APP_ENV=prod  # loads config/prod.yaml
```

## Database Migrations

Add numbered SQL files under `docs/db/`:

```
docs/db/000001_create_users.sql
docs/db/000002_create_sms_codes.sql
docs/db/000003_your_migration.sql
```

Migrations run automatically on startup when `auto_migrate: true` is set.

## License

MIT - See [LICENSE](LICENSE)
