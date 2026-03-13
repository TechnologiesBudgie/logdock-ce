package storage_test

import (
	"testing"
	"time"

	"logdock/internal/storage"
	"logdock/internal/storage/fs"
)

func TestAppendTailSearch(t *testing.T) {
	s, err := fs.New(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	r := storage.LogRecord{
		ID: "1", Timestamp: time.Now().UTC(), Source: "test", Level: "info", Message: "hello world",
	}
	if err := s.Append(r); err != nil {
		t.Fatal(err)
	}
	day := r.Timestamp.Format("2006-01-02")
	rows, err := s.Search(day, "", 10)
	if err != nil || len(rows) != 1 {
		t.Fatalf("search err=%v rows=%d", err, len(rows))
	}
	qres, err := s.Query(storage.QueryFilter{Term: "world", Limit: 10})
	if err != nil || len(qres.Records) != 1 {
		t.Fatalf("query err=%v rows=%d", err, len(qres.Records))
	}
}
