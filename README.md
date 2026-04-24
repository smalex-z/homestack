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

**Download a pre-built binary from [Releases](https://github.com/smalex-z/homestack/releases):**

```bash
# Stable
wget https://github.com/smalex-z/homestack/releases/latest/download/homestack-linux-amd64
chmod +x homestack-linux-amd64

# Specific version (including pre-releases)
wget https://github.com/smalex-z/homestack/releases/download/v0.1.0-alpha.1/homestack-linux-amd64
chmod +x homestack-linux-amd64
```

Verify your download against checksums on the [releases page](https://github.com/smalex-z/homestack/releases).

**Or build from source:**

```bash
./scripts/build.sh
```

**Run:**

```bash
./homestack                  # start on :8080 (or PORT env var)
./homestack --port 9090      # custom port
./homestack --version        # print version
```

**Install as a systemd service (runs on boot, survives reboots):**

```bash
sudo ./scripts/install.sh

# Service management
sudo systemctl status homestack
sudo systemctl restart homestack
sudo journalctl -u homestack -f
```

**During development — hot-swap without re-installing:**

```bash
./scripts/reinstall.sh    # rebuild + swap binary in the running service
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
├── scripts/             # build.sh, dev.sh, install.sh, reinstall.sh
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

## Releases

The release workflow triggers automatically on any `v*` tag and publishes binaries for `linux/amd64` and `linux/arm64`.

**Versioning scheme:**

| Tag format          | Type        | Notes                          |
|---------------------|-------------|--------------------------------|
| `v1.0.0`            | Stable      | Published as latest release    |
| `v1.0.0-rc.1`       | Release candidate | Pre-release flag set    |
| `v0.1.0-beta.1`     | Beta        | Pre-release flag set           |
| `v0.1.0-alpha.1`    | Alpha       | Pre-release flag set           |

**To cut a release:**

```bash
git tag v0.1.0-alpha.1
git push --tags
```

The CI pipeline builds both architectures, generates a `SHA256SUMS.txt`, and publishes a GitHub release. Tags containing a hyphen (`-alpha`, `-beta`, `-rc`) are automatically marked as pre-releases.

You can also trigger a release manually from the Actions tab using the `workflow_dispatch` input if the tag already exists.

## License

[MIT](./LICENSE)
