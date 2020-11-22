package handler

import (
	"net/http"
	"net/http/httptest"
	"notify-me/handler"
	"testing"
)

func assertNoErrors(resp *http.Response, err error, t *testing.T) {
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, got %v", err)
	}
}

func assertContentTypeJSON(header http.Header, t *testing.T) {
	val, ok := header["Content-Type"]

	if !ok {
		t.Fatalf("Expected Content-Type header to be set")
	}

	if val[0] != "application/json; charset=utf-8" {
		t.Fatalf("Expected \"application/json; charset=utf-8, got %s", val[0])
	}
}

func createTestServer() *httptest.Server {
	println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>create test server")
	ts := httptest.NewServer(handler.SetupServer("test-db.db"))

	defer ts.Close()

	println("\n\n---\nserver started\n---\n\n")

	return ts
}
