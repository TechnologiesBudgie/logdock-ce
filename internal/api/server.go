package api

import (
	"embed"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"path"
	"strings"
	"time"

	securejoin "github.com/cyphar/filepath-securejoin"
	"logdock/internal/alerts"
	"logdock/internal/audit"
	"logdock/internal/auth"
	"logdock/internal/ingest"
	pipelinePkg "logdock/internal/pipeline"
	"logdock/internal/security"
	"logdock/internal/settings"
	ssePkg "logdock/internal/sse"
	"logdock/internal/storage"
	"logdock/internal/users"
)

//go:embed webdist/*
var webFS embed.FS

const sessionCookieName = "logdock_session"

type Server struct {
	Mux      *http.ServeMux
	Ingest   *ingest.Service
	Store    storage.Storage
	Auth     *auth.Service
	Users    *users.Service
	Alerts   *alerts.Service
	Creds    *security.CredentialStore
	Settings *settings.Settings
	Audit    *audit.Service
	SSE      *ssePkg.Hub
	Pipeline *pipelinePkg.Engine
}

func New(ing *ingest.Service, st storage.Storage, authSvc *auth.Service,
	userSvc *users.Service, alertSvc *alerts.Service,
	creds *security.CredentialStore, setts *settings.Settings,
	auditSvc *audit.Service) *Server {

	s := &Server{
		Mux:      http.NewServeMux(),
		Ingest:   ing,
		Store:    st,
		Auth:     authSvc,
		Users:    userSvc,
		Alerts:   alertSvc,
		Creds:    creds,
		Settings: setts,
		Audit:    auditSvc,
		SSE:      ssePkg.New(),
		Pipeline: pipelinePkg.New(),
	}
	s.routes()
	return s
}

func (s *Server) routes() {
	// Ingest
	s.Mux.HandleFunc("/api/v1/ingest", s.handleIngestAuthed)
	s.Mux.HandleFunc("/v1/logs", s.Ingest.HandleOTLPHTTP)
	// Auth
	s.Mux.HandleFunc("/api/v1/login", s.handleLogin)
	s.Mux.HandleFunc("/api/v1/logout", s.handleLogout)
	s.Mux.HandleFunc("/api/v1/auth/login", s.handleAuthLogin)
	s.Mux.HandleFunc("/api/v1/auth/me", s.handleAuthMe)
	// Logs
	s.Mux.HandleFunc("/api/v1/logs/tail", s.handleTail)
	s.Mux.HandleFunc("/api/v1/logs/search", s.handleSearch)
	s.Mux.HandleFunc("/api/v1/logs", s.handleLogsQuery)
	s.Mux.HandleFunc("/api/v1/logs/stream", s.handleSSE)
	s.Mux.HandleFunc("/api/v1/logs/export", s.handleExport)
	// Queries
	s.Mux.HandleFunc("/api/v1/query/aggregate", s.handleAggregate)
	// Metrics & health
	s.Mux.HandleFunc("/api/v1/metrics", s.handleMetrics)
	// Users
	s.Mux.HandleFunc("/api/v1/users", s.handleUsers)
	s.Mux.HandleFunc("/api/v1/users/create", s.handleUserCreate)
	s.Mux.HandleFunc("/api/v1/users/delete", s.handleUserDelete)
	s.Mux.HandleFunc("/api/v1/users/unlock", s.handleUserUnlock)
	// Alerts
	s.Mux.HandleFunc("/api/v1/alerts", s.handleAlerts)
	s.Mux.HandleFunc("/api/v1/alerts/history", s.handleAlertsHistory)
	// Settings
	s.Mux.HandleFunc("/api/v1/settings", s.handleSettings)
	// Audit
	s.Mux.HandleFunc("/api/v1/audit", s.handleAudit)
	s.Mux.HandleFunc("/api/v1/audit/stats", s.handleAuditStats)
	// Pipeline
	s.Mux.HandleFunc("/api/v1/pipeline", s.handlePipeline)
	// MFA
	s.Mux.HandleFunc("/api/v1/mfa/enable", s.handleMFAEnable)
	s.Mux.HandleFunc("/api/v1/mfa/disable", s.handleMFADisable)
	// UI
	s.Mux.HandleFunc("/", s.handleUI)
}

// usernameFromRequest extracts the username from the session token (for audit middleware).
func (s *Server) usernameFromRequest(r *http.Request) string {
	token := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
	if token == "" {
		if c, err := r.Cookie(sessionCookieName); err == nil {
			token = c.Value
		}
	}
	if token == "" {
		return ""
	}
	claims, err := s.Auth.Parse(token)
	if err != nil {
		return ""
	}
	u, _ := claims["sub"].(string)
	return u
}

// AuditedHandler wraps the mux with audit middleware. Call after routes() to create the final handler.
func (s *Server) AuditedHandler() http.Handler {
	return s.Audit.Middleware(s.usernameFromRequest, s.Mux)
}

// handleIngestAuthed wraps the ingest handler with authentication + pipeline + anomaly.
func (s *Server) handleIngestAuthed(w http.ResponseWriter, r *http.Request) {
	if s.authorized(w, r, "admin", "operator") == nil {
		return
	}
	s.Ingest.HandleJSON(w, r)
}

func (s *Server) handleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
		OTP      string `json:"otp"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	token, err := s.Auth.Login(body.Username, body.Password, body.OTP, s.Settings.Get().Security.RequireMFA)
	if err != nil {
		// Audit failed login.
		s.Audit.Log(audit.Entry{
			Type: audit.EventLoginFail, Username: body.Username,
			RemoteAddr: r.RemoteAddr, Detail: err.Error(), Success: false,
		})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		_ = json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	// Audit successful login.
	s.Audit.Log(audit.Entry{
		Type: audit.EventLogin, Username: body.Username,
		RemoteAddr: r.RemoteAddr, Success: true,
	})
	http.SetCookie(w, &http.Cookie{
		Name: sessionCookieName, Value: token, Path: "/",
		HttpOnly: true, SameSite: http.SameSiteStrictMode,
	})
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]string{
		"status": "ok",
		"token":  token,
		"cli":    "logdock users login --username " + body.Username,
	})
}

func (s *Server) handleLogout(w http.ResponseWriter, r *http.Request) {
	// Revoke the token from the JTI revocation list.
	token := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
	if token == "" {
		if c, err := r.Cookie(sessionCookieName); err == nil {
			token = c.Value
		}
	}
	if token != "" {
		username := s.usernameFromRequest(r)
		s.Auth.Logout(token)
		s.Audit.Log(audit.Entry{
			Type: audit.EventLogout, Username: username,
			RemoteAddr: r.RemoteAddr, Success: true,
		})
	}
	http.SetCookie(w, &http.Cookie{
		Name: sessionCookieName, Value: "", Path: "/", HttpOnly: true, MaxAge: -1,
	})
	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) handleSettings(w http.ResponseWriter, r *http.Request) {
	if s.authorized(w, r, "admin") == nil {
		return
	}
	if r.Method == http.MethodPost || r.Method == http.MethodPut {
		var body struct {
			InstanceName string                   `json:"instance_name"`
			Timezone     string                   `json:"timezone"`
			Retention    settings.RetentionConfig `json:"retention"`
			Storage      settings.StorageConfig   `json:"storage"`
			Security     settings.SecurityConfig  `json:"security"`
			Webhooks     struct {
				SlackUrl string `json:"slackUrl"`
			} `json:"webhooks"`
		}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		if body.Webhooks.SlackUrl != "" {
			_ = s.Creds.Set("webhook_slack_url", body.Webhooks.SlackUrl)
		}
		snap := s.Settings.Get()
		snap.InstanceName = body.InstanceName
		snap.Timezone = body.Timezone
		snap.Retention = body.Retention
		snap.Storage = body.Storage
		snap.Security = body.Security
		snap.Webhooks.HasSlackURL = s.Creds.Has("webhook_slack_url")
		s.Settings.Update(snap)
		s.Audit.Log(audit.Entry{
			Type: audit.EventConfigSet, Username: s.usernameFromRequest(r),
			RemoteAddr: r.RemoteAddr, Detail: "settings updated", Success: true,
		})
		w.WriteHeader(http.StatusNoContent)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(s.Settings.Get())
}

func (s *Server) handleAudit(w http.ResponseWriter, r *http.Request) {
	if s.authorized(w, r, "admin") == nil {
		return
	}
	limit := 200
	if lStr := r.URL.Query().Get("limit"); lStr != "" {
		fmt.Sscanf(lStr, "%d", &limit)
	}
	eventType := r.URL.Query().Get("type")
	entries := s.Audit.List(limit)
	if eventType != "" {
		filtered := entries[:0]
		for _, e := range entries {
			if string(e.Type) == eventType {
				filtered = append(filtered, e)
			}
		}
		entries = filtered
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]any{
		"history": entries,
		"cli":     "logdock audit --json",
	})
}

func (s *Server) handleAuditStats(w http.ResponseWriter, r *http.Request) {
	if s.authorized(w, r, "admin") == nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(s.Audit.Stats())
}

func (s *Server) handleTail(w http.ResponseWriter, r *http.Request) {
	if s.authorized(w, r, "admin", "operator", "viewer") == nil {
		return
	}
	day := r.URL.Query().Get("day")
	if day == "" {
		day = time.Now().UTC().Format("2006-01-02")
	}
	res, err := s.Store.Query(storage.QueryFilter{From: time.Now().Add(-1 * time.Hour), Limit: 200})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]any{"rows": res.Records, "cli": "logdock logs tail --day " + day})
}

func (s *Server) handleSearch(w http.ResponseWriter, r *http.Request) {
	if s.authorized(w, r, "admin", "operator", "viewer") == nil {
		return
	}
	q := r.URL.Query().Get("q")
	fromStr := r.URL.Query().Get("from")
	limitStr := r.URL.Query().Get("limit")
	limit := 200
	if limitStr != "" {
		fmt.Sscanf(limitStr, "%d", &limit)
	}
	if limit > 10000 {
		limit = 10000
	}
	filter := storage.QueryFilter{Term: q, Limit: limit}
	if fromStr != "" {
		t, _ := time.Parse(time.RFC3339, fromStr)
		filter.From = t
	}
	res, err := s.Store.Query(filter)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]any{"rows": res.Records, "aggregations": res.Aggregations})
}

func (s *Server) handleAggregate(w http.ResponseWriter, r *http.Request) {
	if s.authorized(w, r, "admin", "operator", "viewer") == nil {
		return
	}
	q := r.URL.Query().Get("q")
	fromStr := r.URL.Query().Get("from")
	toStr := r.URL.Query().Get("to")
	filter := storage.QueryFilter{Term: q, Limit: 1000}
	if fromStr != "" {
		filter.From, _ = time.Parse(time.RFC3339, fromStr)
	}
	if toStr != "" {
		filter.To, _ = time.Parse(time.RFC3339, toStr)
	}
	res, err := s.Store.Query(filter)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]any{"aggregations": res.Aggregations})
}

func (s *Server) handleMetrics(w http.ResponseWriter, r *http.Request) {
	if s.authorized(w, r, "admin", "operator", "viewer") == nil {
		return
	}
	metrics, err := s.Store.GetMetrics()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]any{
		"metrics":    metrics,
		"sseClients": s.SSE.ClientCount(),
		"cli":        "logdock storage health --json",
	})
}

func (s *Server) handleUsers(w http.ResponseWriter, r *http.Request) {
	if s.authorized(w, r, "admin") == nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]any{"users": s.Auth.UserList(), "cli": "logdock users list --json"})
}

func (s *Server) handleUserCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", 405)
		return
	}
	if s.authorized(w, r, "admin") == nil {
		return
	}
	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Role     string `json:"role"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	if body.Role == "" {
		body.Role = "viewer"
	}
	if err := s.Auth.CreateUser(body.Username, body.Password, body.Role); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	s.Audit.Log(audit.Entry{
		Type: audit.EventUserCreate, Username: s.usernameFromRequest(r),
		RemoteAddr: r.RemoteAddr, Detail: "created user: " + body.Username, Success: true,
	})
	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) handleUserDelete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost && r.Method != http.MethodDelete {
		http.Error(w, "method not allowed", 405)
		return
	}
	if s.authorized(w, r, "admin") == nil {
		return
	}
	var body struct {
		Username string `json:"username"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	if err := s.Auth.DeleteUser(body.Username); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	s.Audit.Log(audit.Entry{
		Type: audit.EventUserDelete, Username: s.usernameFromRequest(r),
		RemoteAddr: r.RemoteAddr, Detail: "deleted user: " + body.Username, Success: true,
	})
	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) handleUserUnlock(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", 405)
		return
	}
	if s.authorized(w, r, "admin") == nil {
		return
	}
	var body struct {
		Username string `json:"username"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	s.Auth.UnlockUser(body.Username)
	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) handleAlerts(w http.ResponseWriter, r *http.Request) {
	if s.authorized(w, r, "admin", "operator") == nil {
		return
	}
	if r.Method == http.MethodPost {
		var rule alerts.Rule
		if err := json.NewDecoder(r.Body).Decode(&rule); err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		s.Alerts.AddRule(rule)
		s.Audit.Log(audit.Entry{
			Type: audit.EventAlertCreate, Username: s.usernameFromRequest(r),
			RemoteAddr: r.RemoteAddr, Detail: "alert rule: " + rule.Name, Success: true,
		})
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]any{
		"rules": s.Alerts.Rules(), "history": s.Alerts.History(), "cli": "logdock alerts list --json",
	})
}

// handleSSE streams live logs via Server-Sent Events.
func (s *Server) handleSSE(w http.ResponseWriter, r *http.Request) {
	if s.authorized(w, r, "admin", "operator", "viewer") == nil {
		return
	}
	s.SSE.ServeHTTP(w, r)
}

// handleExport exports logs as JSON, CSV, or NDJSON.
func (s *Server) handleExport(w http.ResponseWriter, r *http.Request) {
	if s.authorized(w, r, "admin", "operator") == nil {
		return
	}
	format := r.URL.Query().Get("format") // json | csv | ndjson
	if format == "" {
		format = "json"
	}
	q := r.URL.Query().Get("q")
	limit := 50000
	filter := storage.QueryFilter{Term: q, Limit: limit,
		From: time.Now().Add(-24 * time.Hour)}
	if fromStr := r.URL.Query().Get("from"); fromStr != "" {
		filter.From, _ = time.Parse(time.RFC3339, fromStr)
	}
	if toStr := r.URL.Query().Get("to"); toStr != "" {
		filter.To, _ = time.Parse(time.RFC3339, toStr)
	}
	res, err := s.Store.Query(filter)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	s.Audit.Log(audit.Entry{
		Type: audit.EventDataExport, Username: s.usernameFromRequest(r),
		RemoteAddr: r.RemoteAddr,
		Detail:     fmt.Sprintf("format=%s count=%d", format, len(res.Records)), Success: true,
	})
	switch format {
	case "csv":
		w.Header().Set("Content-Type", "text/csv")
		w.Header().Set("Content-Disposition", `attachment; filename="logdock-export.csv"`)
		cw := csv.NewWriter(w)
		_ = cw.Write([]string{"timestamp", "level", "source", "message", "stream_id"})
		for _, rec := range res.Records {
			_ = cw.Write([]string{
				rec.Timestamp.UTC().Format(time.RFC3339),
				rec.Level, rec.Source, rec.Message, rec.StreamID,
			})
		}
		cw.Flush()
	case "ndjson":
		w.Header().Set("Content-Type", "application/x-ndjson")
		w.Header().Set("Content-Disposition", `attachment; filename="logdock-export.ndjson"`)
		enc := json.NewEncoder(w)
		for _, rec := range res.Records {
			_ = enc.Encode(rec)
		}
	default: // json
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Content-Disposition", `attachment; filename="logdock-export.json"`)
		_ = json.NewEncoder(w).Encode(res.Records)
	}
}

func (s *Server) handlePipeline(w http.ResponseWriter, r *http.Request) {
	if s.authorized(w, r, "admin") == nil {
		return
	}
	switch r.Method {
	case http.MethodGet:
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{"rules": s.Pipeline.Rules()})

	case http.MethodPost:
		var rule pipelinePkg.Rule
		if err := json.NewDecoder(r.Body).Decode(&rule); err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		if err := s.Pipeline.AddRule(rule); err != nil {
			http.Error(w, "invalid rule: "+err.Error(), 400)
			return
		}
		s.Audit.Log(audit.Entry{
			Type: audit.EventRuleCreate, Username: s.usernameFromRequest(r),
			RemoteAddr: r.RemoteAddr, Detail: "pipeline rule: " + rule.Name, Success: true,
		})
		w.WriteHeader(http.StatusNoContent)

	case http.MethodDelete:
		id := r.URL.Query().Get("id")
		if id == "" {
			http.Error(w, "id required", 400)
			return
		}
		s.Pipeline.DeleteRule(id)
		s.Audit.Log(audit.Entry{
			Type: audit.EventRuleDelete, Username: s.usernameFromRequest(r),
			RemoteAddr: r.RemoteAddr, Detail: "rule_id: " + id, Success: true,
		})
		w.WriteHeader(http.StatusNoContent)

	case http.MethodPut:
		var rules []pipelinePkg.Rule
		if err := json.NewDecoder(r.Body).Decode(&rules); err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		if err := s.Pipeline.SetRules(rules); err != nil {
			http.Error(w, "invalid rules: "+err.Error(), 400)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}

func (s *Server) handleMFAEnable(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", 405)
		return
	}
	claims := s.authorized(w, r, "admin", "operator", "viewer")
	if claims == nil {
		return
	}
	username, _ := claims["sub"].(string)
	// Admins can enable MFA for other users.
	var body struct{ Username string `json:"username"` }
	_ = json.NewDecoder(r.Body).Decode(&body)
	role, _ := claims["role"].(string)
	if body.Username != "" && body.Username != username && role != "admin" {
		http.Error(w, "forbidden", 403)
		return
	}
	target := username
	if body.Username != "" && role == "admin" {
		target = body.Username
	}
	uri, err := s.Auth.EnableMFA(target)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	s.Audit.Log(audit.Entry{
		Type: audit.EventMFAChange, Username: username,
		RemoteAddr: r.RemoteAddr, Detail: "mfa enabled for " + target, Success: true,
	})
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]any{
		"provisioning_uri": uri,
		"note":             "Scan with any TOTP authenticator app (Google Authenticator, Authy, etc.)",
	})
}

func (s *Server) handleMFADisable(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", 405)
		return
	}
	claims := s.authorized(w, r, "admin", "operator", "viewer")
	if claims == nil {
		return
	}
	username, _ := claims["sub"].(string)
	var body struct{ Username string `json:"username"` }
	_ = json.NewDecoder(r.Body).Decode(&body)
	role, _ := claims["role"].(string)
	target := username
	if body.Username != "" && role == "admin" {
		target = body.Username
	}
	if err := s.Auth.DisableMFA(target); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	s.Audit.Log(audit.Entry{
		Type: audit.EventMFAChange, Username: username,
		RemoteAddr: r.RemoteAddr, Detail: "mfa disabled for " + target, Success: true,
	})
	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) authorized(w http.ResponseWriter, r *http.Request, roles ...string) map[string]any {
	token := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
	if token == "" {
		if c, err := r.Cookie(sessionCookieName); err == nil {
			token = c.Value
		}
	}
	if token == "" {
		http.Error(w, "missing bearer token or session cookie", http.StatusUnauthorized)
		return nil
	}
	claims, err := s.Auth.Parse(token)
	if err != nil {
		http.Error(w, "invalid or expired token", http.StatusUnauthorized)
		return nil
	}
	if len(roles) > 0 {
		role, _ := claims["role"].(string)
		authOk := false
		for _, r := range roles {
			if role == r {
				authOk = true
				break
			}
		}
		if !authOk {
			http.Error(w, "forbidden", http.StatusForbidden)
			return nil
		}
	}
	return claims
}

const loginPageHTML = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<title>LogDock — Sign In</title>
<link rel="preconnect" href="https://fonts.googleapis.com">
<link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
<link href="https://fonts.googleapis.com/css2?family=Inter:wght@400;500;600;700&display=swap" rel="stylesheet">
<style>
*{margin:0;padding:0;box-sizing:border-box}
body{background:#0a0e1a;color:#f1f5f9;font-family:'Inter',system-ui,sans-serif;font-size:13px;height:100vh;display:flex;align-items:center;justify-content:center}
.box{background:#1a2235;border:1px solid #2d3748;border-radius:12px;padding:36px;width:360px;box-shadow:0 8px 40px rgba(0,0,0,.5)}
.logo{display:flex;align-items:center;gap:12px;margin-bottom:28px}
.logo svg{width:32px;height:32px;color:#7c3aed;flex-shrink:0}
.logo h1{font-size:22px;font-weight:700;letter-spacing:-.5px}
.field{display:flex;flex-direction:column;gap:6px;margin-bottom:14px}
label{font-size:11px;font-weight:500;color:#94a3b8;text-transform:uppercase;letter-spacing:.5px}
input{background:#1f2937;border:1px solid #2d3748;border-radius:6px;color:#f1f5f9;padding:9px 11px;font-size:13px;font-family:inherit;outline:none;transition:border-color .15s}
input:focus{border-color:#7c3aed}
.err{color:#ef4444;font-size:12px;margin-bottom:12px;display:none;padding:8px 10px;background:rgba(239,68,68,.08);border-radius:6px;border:1px solid rgba(239,68,68,.2)}
button{width:100%;padding:10px;background:#7c3aed;border:none;border-radius:6px;color:#fff;font-size:13px;font-weight:600;font-family:inherit;cursor:pointer;transition:background .15s;margin-top:4px}
button:hover{background:#6d28d9}
button:disabled{opacity:.6;cursor:not-allowed}
.hint{text-align:center;font-size:11px;color:#64748b;margin-top:12px}
</style>
</head>
<body>
<div class="box">
  <div class="logo">
    <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M12 2L2 7l10 5 10-5-10-5zM2 17l10 5 10-5M2 12l10 5 10-5"/></svg>
    <h1>LogDock</h1>
  </div>
  <div class="field"><label>Username</label><input type="text" id="u" value="admin" autocomplete="username"></div>
  <div class="field"><label>Password</label><input type="password" id="p" placeholder="Password" autocomplete="current-password"></div>
  <div class="field"><label>MFA Code <span style="color:#64748b;text-transform:none;font-size:10px">(if required)</span></label><input type="text" id="otp" placeholder="6-digit code" autocomplete="one-time-code" inputmode="numeric" maxlength="6"></div>
  <div class="err" id="err"></div>
  <button id="btn" onclick="login()">Sign In</button>
  <p class="hint">Default credentials: admin / admin</p>
</div>
<script>
document.querySelectorAll('input').forEach(el => el.addEventListener('keydown', e => { if(e.key==='Enter') login(); }));
async function login() {
  const btn = document.getElementById('btn');
  const errEl = document.getElementById('err');
  errEl.style.display = 'none';
  btn.disabled = true;
  btn.textContent = 'Signing in…';
  try {
    const res = await fetch('/api/v1/auth/login', {
      method: 'POST',
      headers: {'Content-Type': 'application/json'},
      body: JSON.stringify({
        username: document.getElementById('u').value.trim(),
        password: document.getElementById('p').value,
        otp:      document.getElementById('otp').value.trim()
      })
    });
    if (!res.ok) {
      let msg = 'Login failed';
      try {
        const d = await res.json();
        msg = d.error || msg;
      } catch(e) {
        msg = (await res.text()).trim() || msg;
      }
      throw new Error(msg);
    }
    const data = await res.json();
    localStorage.setItem('logdock_token', data.token);
    window.location.href = '/';
  } catch(e) {
    errEl.textContent = e.message;
    errEl.style.display = 'block';
  } finally {
    btn.disabled = false;
    btn.textContent = 'Sign In';
  }
}
</script>
</body>
</html>`

func (s *Server) handleUI(w http.ResponseWriter, r *http.Request) {
	userPath := strings.TrimPrefix(path.Clean(r.URL.Path), "/")
	p, err := securejoin.SecureJoin("webdist", userPath)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	p = strings.TrimPrefix(p, "webdist/")

	isHTML := p == "" || p == "." || p == "index.html" || (!strings.Contains(p, ".") && !strings.HasPrefix(p, "api/"))

	headers := w.Header()
	headers.Set("X-Content-Type-Options", "nosniff")
	headers.Set("X-Frame-Options", "DENY")
	headers.Set("X-XSS-Protection", "1; mode=block")
	headers.Set("Referrer-Policy", "strict-origin-when-cross-origin")
	headers.Set("Permissions-Policy", "camera=(), microphone=(), geolocation=()")
	headers.Set("Content-Security-Policy",
		"default-src 'self'; script-src 'self' 'unsafe-inline' https://cdn.jsdelivr.net; style-src 'self' 'unsafe-inline' https://fonts.googleapis.com; img-src 'self' data:; font-src 'self' https://fonts.gstatic.com; connect-src 'self'")
	headers.Set("Cache-Control", "no-store")

	if isHTML {
		token := ""
		if c, err := r.Cookie(sessionCookieName); err == nil {
			token = c.Value
		}
		if token == "" {
			token = strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
		}

		authenticated := false
		if token != "" {
			if _, err := s.Auth.Parse(token); err == nil {
				authenticated = true
			}
		}

		headers.Set("Content-Type", "text/html; charset=utf-8")
		if !authenticated {
			tpl, _ := template.New("login").Parse(loginPageHTML)
			_ = tpl.Execute(w, nil)
			return
		}
		b, err := webFS.ReadFile("webdist/index.html")
		if err != nil {
			http.Error(w, "app not found", http.StatusInternalServerError)
			return
		}
		tpl, _ := template.New("index").Parse(string(b))
		_ = tpl.Execute(w, nil)
		return
	}

	b, err := webFS.ReadFile("webdist/" + p)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	switch {
	case strings.HasSuffix(p, ".html"):
		headers.Set("Content-Type", "text/html; charset=utf-8")
	case strings.HasSuffix(p, ".js"):
		headers.Set("Content-Type", "application/javascript")
	case strings.HasSuffix(p, ".css"):
		headers.Set("Content-Type", "text/css")
	case strings.HasSuffix(p, ".svg"):
		headers.Set("Content-Type", "image/svg+xml")
	}

	if strings.HasSuffix(p, ".html") {
		tpl, _ := template.New("page").Parse(string(b))
		_ = tpl.Execute(w, nil)
	} else {
		_, _ = w.Write(b)
	}
}

func (s *Server) handleAuthLogin(w http.ResponseWriter, r *http.Request) { s.handleLogin(w, r) }

func (s *Server) handleAuthMe(w http.ResponseWriter, r *http.Request) {
	claims := s.authorized(w, r, "admin", "operator", "viewer")
	if claims == nil {
		// authorized() already sets http.Error
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]any{
		"username": claims["sub"],
		"role":     claims["role"],
	})
}

func (s *Server) handleLogsQuery(w http.ResponseWriter, r *http.Request) { s.handleSearch(w, r) }

func (s *Server) handleAlertsHistory(w http.ResponseWriter, r *http.Request) {
	if s.authorized(w, r, "admin", "operator") == nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(s.Alerts.History())
}
