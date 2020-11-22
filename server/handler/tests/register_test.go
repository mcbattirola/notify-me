package handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dchest/uniuri"
)

func TestRegisterRoute(t *testing.T) {
	ts := createTestServer()

	token := uniuri.New()

	resp, err := http.Get(fmt.Sprintf("%s/api/v1/register/%s", ts.URL, token))

	assertNoErrors(resp, err, t)
	assertContentTypeJSON(resp.Header, t)

	assertCantRegisterTokenMoreThanOnce(token, t, ts)

}

func assertCantRegisterTokenMoreThanOnce(token string, t *testing.T, testServer *httptest.Server) {
	_, err := http.Get(fmt.Sprintf("%s/api/v1/register/%s", testServer.URL, token))
	if err == nil {
		t.Fatalf("Expected error, got none")
	}

}
