package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mcbattirola/notify-me/server/models"

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

func TestListSendersRoute(t *testing.T) {
	ts := httptest.NewServer(setupServer())
	defer ts.Close()

	resp, err := http.Get(fmt.Sprintf("%s/sender", ts.URL))

	assertNoErrors(resp, err, t)
	assertContentTypeJSON(resp.Header, t)
}

func TestListSenders(t *testing.T) {
	ts := httptest.NewServer(setupServer())
	defer ts.Close()

	resp, err := http.Get(fmt.Sprintf("%s/sender", ts.URL))

	assertNoErrors(resp, err, t)
	assertContentTypeJSON(resp.Header, t)
}

func TestCreateSendersRoute(t *testing.T) {
	ts := httptest.NewServer(setupServer())
	defer ts.Close()

	requestBody, err := json.Marshal(map[string]string{
		"IP": "127.0.0.1",
	})

	resp, err := http.Post(fmt.Sprintf("%s/sender", ts.URL), "application/json", bytes.NewBuffer(requestBody))

	assertNoErrors(resp, err, t)
	assertContentTypeJSON(resp.Header, t)
}

func TestSubscribeRoute(t *testing.T) {
	ts := httptest.NewServer(setupServer())
	defer ts.Close()

	requestBody, err := json.Marshal(models.Subscription{
		SenderID: 5,
		DeviceID: "1",
	})

	resp, err := http.Post(fmt.Sprintf("%s/subscribe", ts.URL), "application/json", bytes.NewBuffer(requestBody))

	assertNoErrors(resp, err, t)
	assertContentTypeJSON(resp.Header, t)
}

func TestCreateNotificationRoute(t *testing.T) {
	ts := httptest.NewServer(setupServer())
	defer ts.Close()

	requestBody, err := json.Marshal(map[string]string{
		"title": "message title",
		"body":  "body",
	})

	resp, err := http.Post(fmt.Sprintf("%s/notification", ts.URL), "application/json", bytes.NewBuffer(requestBody))

	assertNoErrors(resp, err, t)
	assertContentTypeJSON(resp.Header, t)
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
