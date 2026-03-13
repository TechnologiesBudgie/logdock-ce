package ingest

import (
	"context"
	"io"
	"log"
	"net"
	"net/http"
)

// HandleOTLPHTTP accepts OTLP/HTTP log payloads.
// NOTE: Full OTLP protobuf decoding requires go.opentelemetry.io/proto/otlp.
// Until that dependency is added, we drain the body so the sender doesn't
// stall, log the byte count for observability, and return 202.
func (s *Service) HandleOTLPHTTP(w http.ResponseWriter, r *http.Request) {
	n, _ := io.Copy(io.Discard, r.Body)
	log.Printf("[OTLP-HTTP] received %d bytes (stub: protobuf decode not yet implemented)", n)
	w.WriteHeader(http.StatusAccepted)
}

// StartOTLPGRPC listens for OTLP/gRPC connections.
// BUG-002 fix: the goroutine now exits when ctx is cancelled.
// NOTE: Full gRPC service implementation requires google.golang.org/grpc.
// Until that dependency is added, we accept and immediately close connections
// so clients receive a clean EOF rather than a silent hang.
func (s *Service) StartOTLPGRPC(ctx context.Context, addr string) error {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	go func() {
		<-ctx.Done()
		_ = ln.Close()
	}()
	go func() {
		defer ln.Close()
		for {
			c, err := ln.Accept()
			if err != nil {
				return
			}
			log.Printf("[OTLP-gRPC] connection from %s (stub: gRPC framing not yet implemented)", c.RemoteAddr())
			_ = c.Close()
		}
	}()
	return nil
}
