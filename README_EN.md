# FastGox API Starter

English | [中文](README.md)

A minimal Go API starter template with layered architecture, ready to use out of the box.

## Project Structure

```
main.go                    # Entry point
config/                    # Configuration files (YAML, per environment)
deploy/docker/             # Dockerfile + docker-compose
docs/
  db/                      # SQL migration files
  swagger/                 # Swagger docs
src/
  core/
    config/                # Config loader
    database/              # DB connection, migration, utilities
    i18n/                  # Internationalization (zh/en)
    session/               # Session management
    tcp/                   # TCP long-connection server
  models/
    config/                # Config structs
    dto/                   # Data transfer objects (request/response/error codes)
    entity/                # Database entities
  repository/              # Data access layer
  services/                # Business logic layer
  router/
    handle/                # HTTP route handlers
    middleware/             # Middleware (auth, CORS, i18n, recovery)
  utils/                   # Utility functions (JWT, AES, Hash, etc.)
  pkg/                     # External service wrappers
test/                      # Unit tests
templates/                 # HTML templates
```

## Features

- **Layered architecture**: Repository → Service → Router with clear separation of concerns
- **HTTP + TCP dual protocol**: HTTP API and TCP long-connection server running in parallel
- **Internationalization (i18n)**: Built-in zh/en support, auto-switching by request header
- **Unified response format**: Standardized success/failure/pagination response structures
- **SQL migration version control**: Auto-execution based on `schema_migrations` table
- **Graceful shutdown**: Both HTTP and TCP services support graceful shutdown
- **Auto-generated Swagger docs**
- **JWT authentication + CORS + Recovery middleware**
- **Multi-environment config**: dev / test / prod config isolation
- **One-command Docker deployment**
- **Task-based project management** (dev, test, build, deploy in one place)
- **Unit test coverage**: config, database, entity, i18n, tcp, utils, and more

## Quick Start

### Prerequisites

- Go 1.24+
- [Task](https://taskfile.dev) (recommended)
- Docker / Docker Compose (for deployment)

### Development

```bash
# Initialize project (dependencies + Swagger)
task setup

# Start dev server
task dev

# Format code
task fmt

# Generate Swagger docs
task swagger
```

### Testing

```bash
# Run all unit tests
task test

# Run tests with coverage report
task test:cover

# Quick tests (skip integration tests)
task test:short
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
APP_ENV=dev    # loads config/dev.yaml
APP_ENV=test   # loads config/test.yaml
APP_ENV=prod   # loads config/prod.yaml
```

Main configuration sections:

| Section | Description |
| --- | --- |
| `app` | App name, port, debug mode, Swagger, shutdown timeout |
| `database` | DB driver, connection params, auto-migration |
| `jwt` | JWT secret key |
| `sms-code` | SMS code length, expiry, rate limiting, whitelist |
| `tcp` | TCP server port, max connections, timeouts, heartbeat |
| `file` | File upload size limit, count limit, allowed extensions |

## Database Migrations

Add numbered SQL files under `docs/db/`:

```
docs/db/000001_create_users.sql
docs/db/000002_create_sms_codes.sql
docs/db/000003_your_migration.sql
```

Migrations run automatically on startup when `auto_migrate: true` is set.

## License

Apache 2.0 - See [LICENSE](LICENSE)
