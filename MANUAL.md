# LogDock Community Edition — Manual

LogDock is a self-hosted log observability platform. It is designed for speed, simplicity, and low operational overhead.

## Quick Start

1. **Configure**: Copy `resources/logdock.env.example` to `.env` and set `LOGDOCK_MASTER_KEY` and `LOGDOCK_JWT_SECRET`.
2. **Start**: `./logdock serve`
3. **Login**: Access `http://localhost:2514`. Default: `admin` / `admin`.

## Core Concepts

### Ingestion
LogDock accepts logs via:
- **Syslog**: TCP/UDP on port 5140 (default).
- **OTLP**: OpenTelemetry gRPC (4317) and HTTP (4318).
- **API**: POST to `/api/v1/ingest`.

### Storage & Retention
Logs are stored in daily partitions. Retention is configured in **Settings > Retention**.
Partitions older than the configured days are automatically deleted by the background maintenance loop.

### Compression Tiers
To balance speed and disk usage, LogDock uses tiered compression:
- **Hot**: Recent logs (uncompressed).
- **Warm**: Compressed with LZ4.
- **Cold**: Max compression with Zstd.

## Administration

### User Management
Manage users via **Settings > Users** or the `logdock users` CLI command.
Supported roles: `admin`, `operator`, `viewer`.

### Alerts
Create rules based on log patterns or error rates. Notifications can be sent to Slack or custom webhooks.

### Security
- **MFA**: TOTP-based two-factor authentication.
- **Credential Store**: Sensitive values are encrypted with AES-256-GCM.
- **Audit Log**: All administrative actions are recorded.

## CLI Usage

```bash
# Search logs
logdock logs search --q "error"

# View status
logdock health

# Create a data snapshot
logdock snapshots --path /backup
```

---
*LogDock Community Edition · [Documentation](https://logdock.io/docs)*
