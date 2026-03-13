package main

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"logdock/internal/alerts"
	"logdock/internal/api"
	"logdock/internal/audit"
	"logdock/internal/auth"
	"logdock/internal/config"
	"logdock/internal/ingest"
	"logdock/internal/pipeline"
	"logdock/internal/security"
	"logdock/internal/settings"
	"logdock/internal/storage/fs"
	"logdock/internal/streams"
	"logdock/internal/users"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] != "serve" {
		runCLI()
		return
	}
	cfg := config.Load()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	store, err := fs.New(cfg.DataDir)
	if err != nil {
		log.Fatal(err)
	}

	// Credential store (AES-256-GCM encrypted secrets)
	creds, err := security.NewCredentialStore()
	if err != nil {
		log.Fatalf("credential store init: %v", err)
	}

	// Settings (runtime config, non-secret)
	setts := settings.New()

	authSvc := auth.New(cfg.JWTSecret)
	_ = authSvc.SeedAdmin(getEnv("LOGDOCK_ADMIN_PASSWORD", "admin"))

	pipe := &pipeline.Pipeline{
		Processors: []pipeline.Processor{&pipeline.JSONExtractor{}},
	}
	streamMgr := &streams.Manager{
		Streams: []streams.Stream{
			{ID: "default", Title: "All Logs"},
		},
	}

	ing := &ingest.Service{Store: store, Pipeline: pipe, Streams: streamMgr}
	userSvc := users.New()
	alertSvc := alerts.New(store)
	go alertSvc.Start(ctx)

	// Background maintenance loop for retention
	go func() {
		ticker := time.NewTicker(1 * time.Hour)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				snap := setts.Get()
				store.RunMaintenance(snap.Retention.MaxDays)
			}
		}
	}()

	must(ing.StartSyslogUDP(ctx, cfg.SyslogUDP))
	must(ing.StartSyslogTCP(ctx, cfg.SyslogTCP))
	must(ing.StartOTLPGRPC(ctx, cfg.OTLPGRPCAddr))

	auditSvc := audit.New()
	srv := api.New(ing, store, authSvc, userSvc, alertSvc, creds, setts, auditSvc)

	guard := security.NewBruteForceGuard()
	guard.Start(ctx)

	httpSrv := &http.Server{
		Addr:              cfg.HTTPAddr,
		Handler:           guard.Middleware(srv.AuditedHandler()),
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      60 * time.Second,
		IdleTimeout:       120 * time.Second,
		TLSConfig:         &tls.Config{MinVersion: tls.VersionTLS12},
	}

	go func() {
		log.Printf("logdock listening on %s", cfg.HTTPAddr)
		var serveErr error
		if cfg.TLSCertFile != "" && cfg.TLSKeyFile != "" {
			serveErr = httpSrv.ListenAndServeTLS(cfg.TLSCertFile, cfg.TLSKeyFile)
		} else {
			serveErr = httpSrv.ListenAndServe()
		}
		if serveErr != nil && serveErr != http.ErrServerClosed {
			log.Printf("http server error: %v", serveErr)
			stop()
		}
	}()

	<-ctx.Done()
	log.Println("logdock: shutdown signal received, draining connections…")

	shutCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := httpSrv.Shutdown(shutCtx); err != nil {
		log.Printf("http shutdown error: %v", err)
	}
	if err := store.Close(); err != nil {
		log.Printf("store close error: %v", err)
	}
	log.Println("logdock: shutdown complete")
}

// ─── CLI Commands ──────────────────────────────────────────────────────────────

func runCLI() {
	if len(os.Args) < 2 {
		fmt.Println(help())
		return
	}
	switch os.Args[1] {
	case "ingest":
		cliIngest()
	case "tui":
		cliTui()
	case "export":
		cliExport()
	case "logs":
		cliLogs()
	case "storage":
		cliStorage()
	case "users":
		cliUsers()
	case "alerts":
		cliAlerts()
	case "version":
		fmt.Println("logdock v2.0.0")
	case "health":
		cliHealth()
	default:
		fmt.Println(help())
	}
}

func cliIngest() {
	fset := flag.NewFlagSet("ingest", flag.ExitOnError)
	endpoint := fset.String("endpoint", "http://127.0.0.1:2514/api/v1/ingest", "endpoint")
	msg := fset.String("message", "", "message")
	source := fset.String("source", "cli", "source")
	level := fset.String("level", "info", "level")
	jsonOut := fset.Bool("json", false, "json output")
	_ = fset.Parse(os.Args[2:])
	body, _ := json.Marshal(map[string]any{
		"timestamp": time.Now().UTC(), "source": *source, "level": *level, "message": *msg,
	})
	resp, err := http.Post(*endpoint, "application/json", bytes.NewReader(body))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	out("ingest", map[string]any{"status": resp.Status}, *jsonOut)
}

func cliExport() {
	fset := flag.NewFlagSet("export", flag.ExitOnError)
	day := fset.String("day", time.Now().UTC().Format("2006-01-02"), "day")
	format := fset.String("format", "json", "json|csv")
	output := fset.String("out", "export.out", "file")
	dataDir := fset.String("data-dir", "./data", "dir")
	jsonOut := fset.Bool("json", false, "json")
	_ = fset.Parse(os.Args[2:])
	st, err := fs.New(*dataDir)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := st.Search(*day, "", 100000)
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create(*output)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if *format == "csv" {
		w := csv.NewWriter(f)
		_ = w.Write([]string{"timestamp", "source", "level", "message"})
		for _, r := range rows {
			_ = w.Write([]string{r.Timestamp.Format(time.RFC3339), r.Source, r.Level, r.Message})
		}
		w.Flush()
	} else {
		enc := json.NewEncoder(f)
		for _, r := range rows {
			_ = enc.Encode(r)
		}
	}
	out("export", map[string]any{"written": *output, "count": len(rows)}, *jsonOut)
}

func cliLogs() {
	if len(os.Args) < 3 {
		fmt.Println("logs tail|search [--day DATE] [--q QUERY] [--data-dir DIR] [--json]")
		return
	}
	fset := flag.NewFlagSet("logs", flag.ExitOnError)
	day := fset.String("day", time.Now().UTC().Format("2006-01-02"), "day")
	q := fset.String("q", "", "query")
	dataDir := fset.String("data-dir", "./data", "dir")
	jsonOut := fset.Bool("json", false, "json")
	_ = fset.Parse(os.Args[3:])
	st, err := fs.New(*dataDir)
	if err != nil {
		log.Fatal(err)
	}
	if os.Args[2] == "tail" {
		rows, _ := st.Search(*day, "", 200)
		out("tail", rows, *jsonOut)
		return
	}
	rows, _ := st.Search(*day, *q, 200)
	out("search", rows, *jsonOut)
}

func cliStorage() {
	if len(os.Args) < 3 || os.Args[2] != "health" {
		fmt.Println("storage health [--json] [--data-dir DIR]")
		return
	}
	fset := flag.NewFlagSet("storage health", flag.ExitOnError)
	dataDir := fset.String("data-dir", "./data", "dir")
	jsonOut := fset.Bool("json", false, "json")
	_ = fset.Parse(os.Args[3:])
	st, err := fs.New(*dataDir)
	if err != nil {
		log.Fatal(err)
	}
	metrics, _ := st.GetMetrics()
	out("storage-health", metrics, *jsonOut)
}

func cliUsers() {
	fset := flag.NewFlagSet("users", flag.ExitOnError)
	jsonOut := fset.Bool("json", false, "json")
	_ = fset.Parse(os.Args[2:])
	out("users", map[string]any{"supported": []string{"list", "create", "mfa"}, "note": "use API for management"}, *jsonOut)
}

func cliAlerts() {
	fset := flag.NewFlagSet("alerts", flag.ExitOnError)
	jsonOut := fset.Bool("json", false, "json")
	_ = fset.Parse(os.Args[2:])
	out("alerts", map[string]any{"supported": []string{"list", "add"}, "channels": []string{"slack", "email", "webhook"}}, *jsonOut)
}

func cliHealth() {
	fset := flag.NewFlagSet("health", flag.ExitOnError)
	endpoint := fset.String("endpoint", "http://127.0.0.1:2514", "server endpoint")
	jsonOut := fset.Bool("json", false, "json")
	_ = fset.Parse(os.Args[2:])
	resp, err := http.Get(*endpoint + "/api/v1/metrics")
	if err != nil {
		out("health", map[string]any{"status": "unreachable", "error": err.Error()}, *jsonOut)
		return
	}
	defer resp.Body.Close()
	out("health", map[string]any{"status": "ok", "code": resp.StatusCode}, *jsonOut)
}

func cliTui() {
	fset := flag.NewFlagSet("tui", flag.ExitOnError)
	dataDir := fset.String("data-dir", "./data", "dir")
	_ = fset.Parse(os.Args[2:])
	st, err := fs.New(*dataDir)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("\033[2J")
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()
	for {
		metrics, _ := st.GetMetrics()
		fmt.Print("\033[H")
		fmt.Println("╔══════════════════════════════════════════╗")
		fmt.Println("║          LogDock TUI — Live View         ║")
		fmt.Println("╠══════════════════════════════════════════╣")
		fmt.Printf("║  Time:            %-24s║\n", time.Now().Format("2006-01-02 15:04:05 UTC"))
		fmt.Println("╠══════════════════════════════════════════╣")
		fmt.Printf("║  Total Ingested:  %-24d║\n", metrics.TotalIngested)
		fmt.Printf("║  Error Count:     %-24d║\n", metrics.ErrorCount)
		fmt.Printf("║  Last Minute:     %-24d║\n", metrics.LastMinuteIngest)
		fmt.Printf("║  Disk Usage:      %-24s║\n", fmtBytesSimple(metrics.DiskBytes))
		fmt.Printf("║  Partitions:      %-24d║\n", metrics.PartitionDays)
		fmt.Println("╚══════════════════════════════════════════╝")
		fmt.Println("  Press Ctrl+C to exit  |  logdock tui v2")
		<-ticker.C
	}
}

func fmtBytesSimple(b int64) string {
	if b < 1024 {
		return fmt.Sprintf("%d B", b)
	} else if b < 1048576 {
		return fmt.Sprintf("%.1f KB", float64(b)/1024)
	} else if b < 1073741824 {
		return fmt.Sprintf("%.1f MB", float64(b)/1048576)
	}
	return fmt.Sprintf("%.2f GB", float64(b)/1073741824)
}

func out(label string, payload any, jsonOut bool) {
	if jsonOut {
		b, _ := json.MarshalIndent(map[string]any{"command": label, "data": payload}, "", "  ")
		fmt.Println(string(b))
		return
	}
	fmt.Printf("%s: %v\n", label, payload)
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getEnv(k, d string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return d
}

func help() string {
	return `LogDock Community Edition — Log Management

Usage: logdock <command> [options]

Server:
  serve                    Start the LogDock server (default if no command)

Log Commands:
  logs tail                Tail recent logs from local store
  logs search --q TERM     Search logs by term
  ingest --message MSG     Send a log event to the server
  export --format json|csv Export logs to file

Administration:
  users                    User account management
  alerts                   Alert rule management
  health                   Check if server is reachable

Interface:
  tui                      Terminal UI (headless server monitoring)
  version                  Show version

Global Flags:
  --json                   Output as JSON
  --data-dir DIR           Override data directory (default: ./data)
  --endpoint URL           Override server endpoint

Environment Variables:
  LOGDOCK_ADMIN_PASSWORD   Admin password
  LOGDOCK_JWT_SECRET       JWT signing secret
  LOGDOCK_MASTER_KEY       Master key for credential encryption
  LOGDOCK_DATA_DIR         Data directory
  LOGDOCK_HTTP_ADDR        HTTP listen address (default: :2514)

Full documentation: MANUAL.md
`
}
