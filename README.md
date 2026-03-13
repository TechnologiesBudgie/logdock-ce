# LogDock Community Edition

LogDock is a self-hosted log observability platform shipped as a **single Go binary** with embedded admin UI and full CLI parity. No external dependencies.

## Features

### Security
- **TOTP MFA** — RFC 6238 authenticator app support (otpauth:// QR provisioning)
- **JWT revocation** — Logout invalidates tokens server-side via JTI blocklist
- **Account lockout** — Configurable max failed attempts + lockout duration per user
- **AES-256-GCM credential encryption** — Webhook URLs never stored plaintext
- **Brute-force guard** — Sliding-window IP rate limiting (240/min) with stale-entry cleanup

### UI Features
- **Real SSE live tail** — True server-sent events; falls back to polling on disconnect
- **Audit Log page** — Filterable event table: logins, exports, user mgmt
- **Keyboard shortcuts** — `/` search, `L` live, `R` refresh, `Esc` close
- **CSV + NDJSON export** — Direct download from explorer toolbar
- **User management** — Create, delete, unlock from Settings > Users

### Observability
- **Log volume histogram** — Click bar to jump to time window
- **Top sources chart** — Ingestion breakdown by source

## Build

```bash
go mod tidy && go build -o logdock ./cmd/logdock
```

## Run

```bash
cp resources/logdock.env.example .env
set -a; source .env; set +a
./logdock serve
```

Access at `http://localhost:2514`. Default: `admin / admin`.

## Docker

```bash
docker build -t logdock:latest .
docker run --rm -p 2514:2514 -p 4317:4317 -p 5140:5140/tcp -p 5140:5140/udp \
  -v $PWD/data:/data -e LOGDOCK_DATA_DIR=/data -e LOGDOCK_MASTER_KEY=change-this \
  logdock:latest
```

## Security Matrix

| Control | Status |
|---|---|
| JWT HS256 + JTI revocation | ✅ |
| AES-256-GCM credential store | ✅ |
| TOTP MFA (RFC 6238) | ✅ |
| Account lockout | ✅ |
| IP rate limiting (sliding window) | ✅ |
| HTTP security headers (CSP/DENY) | ✅ |
| Audit log (all admin actions) | ✅ |
| RBAC (admin / operator / viewer) | ✅ |
| 1MB request body limit | ✅ |
| Login-gated app HTML | ✅ |

## Keyboard Shortcuts

| Key | Action |
|---|---|
| `/` | Focus search |
| `L` | Toggle live tail |
| `R` | Refresh page |
| `Esc` | Close drawer/modal |
| `Ctrl E` | Export CSV |
| `?` | Show shortcuts |

## Environment Variables

| Variable | Description |
|---|---|
| `LOGDOCK_MASTER_KEY` | AES-256 key for credential encryption (required in prod) |
| `LOGDOCK_JWT_SECRET` | JWT signing secret |
| `LOGDOCK_DATA_DIR` | Storage root (default: `./data`) |
| `LOGDOCK_PORT` | HTTP listen port (default: `2514`) |
| `LOGDOCK_SYSLOG_PORT` | Syslog TCP/UDP (default: `5140`) |
