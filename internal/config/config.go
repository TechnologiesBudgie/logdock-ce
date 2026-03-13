package config

import (
	"log"
	"os"
)

type Config struct {
	DataDir      string
	HTTPAddr     string
	OTLPGRPCAddr string
	OTLPHTTPAddr string
	SyslogTCP    string
	SyslogUDP    string
	TLSCertFile  string
	TLSKeyFile   string
	JWTSecret    string
}

func Load() Config {
	c := Config{
		DataDir:      getEnv("LOGDOCK_DATA_DIR", "./data"),
		HTTPAddr:     getEnv("LOGDOCK_HTTP_ADDR", ":2514"),
		OTLPGRPCAddr: getEnv("LOGDOCK_OTLP_GRPC_ADDR", ":4317"),
		OTLPHTTPAddr: getEnv("LOGDOCK_OTLP_HTTP_ADDR", ":4318"),
		SyslogTCP:    getEnv("LOGDOCK_SYSLOG_TCP_ADDR", ":5140"),
		SyslogUDP:    getEnv("LOGDOCK_SYSLOG_UDP_ADDR", ":5140"),
		TLSCertFile:  os.Getenv("LOGDOCK_TLS_CERT_FILE"),
		TLSKeyFile:   os.Getenv("LOGDOCK_TLS_KEY_FILE"),
		JWTSecret:    getEnv("LOGDOCK_JWT_SECRET", "change-me-now"),
	}

	if c.JWTSecret == "change-me-now" && os.Getenv("LOGDOCK_ENV") == "production" {
		log.Fatal("Security Error: LOGDOCK_JWT_SECRET must be set in production mode.")
	}
	return c
}

func getEnv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
