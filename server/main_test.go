package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dchest/uniuri"
)

func TestPingRoute(t *testing.T) {
	ts := httptest.NewServer(setupServer())
	defer ts.Close()

	resp, err := http.Get(fmt.Sprintf("%s/ping", ts.URL))

	assertNoErrors(resp, err, t)
	assertContentTypeJSON(resp.Header, t)
}

func TestRegisterRoute(t *testing.T) {
	ts := httptest.NewServer(setupServer())
	defer ts.Close()

	token := uniuri.New()
	resp, err := http.Get(fmt.Sprintf("%s/register/%s", ts.URL, token))

	assertNoErrors(resp, err, t)
	assertContentTypeJSON(resp.Header, t)

	// We can't add the same token twice
	resp, err = http.Get(fmt.Sprintf("%s/register/%s", ts.URL, token))
	if resp.StatusCode == http.StatusOK {
		t.Fatalf("Expected error, got none")
	}
}

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
