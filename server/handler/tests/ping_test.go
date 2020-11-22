package handler

import (
	"fmt"
	"net/http"
	"testing"
)

func TestPingRoute(t *testing.T) {
	ts := createTestServer()

	resp, err := http.Get(fmt.Sprintf("%s/api/ping", ts.URL))

	assertNoErrors(resp, err, t)
	assertContentTypeJSON(resp.Header, t)
}
