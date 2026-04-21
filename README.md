# Homestack

> Production-ready template for self-hosted applications.

Built with Go + React. Single binary deployment, embedded frontend, pre-configured CI/CD, and systemd service setup.

## Quick Start

### Development

```bash
git clone https://github.com/smalex-z/homestack.git my-app
cd my-app
make dev
```

- Frontend: <http://localhost:5173>
- Backend API: <http://localhost:8080>

> **Requirements:** Go 1.22+, Node.js 18+

### Production

```bash
# Build single binary
./scripts/build.sh

# Install as systemd service
sudo ./scripts/install.sh
```

Or download a pre-built binary from [Releases](https://github.com/smalex-z/homestack/releases):

```bash
wget https://github.com/smalex-z/homestack/releases/latest/download/homestack-linux-amd64
chmod +x homestack-linux-amd64
./homestack-linux-amd64
```

## Project Structure

```
homestack/
├── cmd/server/          # Go entry point (embeds frontend/dist)
├── frontend/            # React 18 + TypeScript + Vite + Tailwind CSS
├── internal/
│   ├── api/             # Chi router, middleware, handlers
│   ├── db/              # SQLite + GORM models
│   ├── service/         # Business logic layer
│   └── config/          # Environment-based configuration
├── scripts/             # build.sh, dev.sh, install.sh
└── .github/workflows/   # test.yml, lint.yml, release.yml
```

## Tech Stack

| Layer      | Technology                           |
|------------|--------------------------------------|
| Backend    | Go 1.22+ · Chi router                |
| Frontend   | React 18 · TypeScript · Vite · Tailwind CSS |
| Database   | SQLite · GORM (pure Go, no CGO)      |
| Deployment | Single static binary · systemd       |

## Make Commands

| Command       | Description                                |
|---------------|--------------------------------------------|
| `make dev`    | Start backend + frontend dev servers       |
| `make build`  | Build production binary                    |
| `make test`   | Run Go tests + TypeScript type-check       |
| `make lint`   | Run golangci-lint + ESLint                 |
| `make clean`  | Remove build artifacts                     |

## Configuration

All options are set via environment variables:

| Variable      | Default           | Description                 |
|---------------|-------------------|-----------------------------|
| `PORT`        | `8080`            | HTTP server port            |
| `DB_PATH`     | `./homestack.db`  | SQLite database file path   |
| `CORS_ORIGIN` | `*`               | Allowed CORS origin         |
| `APP_ENV`     | `production`      | Application environment     |

## API Endpoints

| Method   | Path              | Description          |
|----------|-------------------|----------------------|
| `GET`    | `/api/health`     | Health check         |
| `GET`    | `/api/users`      | List all users       |
| `POST`   | `/api/users`      | Create a user        |
| `DELETE` | `/api/users/{id}` | Delete a user        |

## License

[MIT](./LICENSE)
