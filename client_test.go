package ghost

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// setupTest configures a test HTTP server and a Cloud Deploy client
// to mock Cloud Deploy API responses
func setupTest() (*http.ServeMux, *httptest.Server, *Client) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	c := NewClient(server.URL, "", "")

	return mux, server, c
}

func teardownTest(server *httptest.Server) {
	server.Close()
}

func testMethod(t *testing.T, r *http.Request, want string) {
	if got := r.Method; got != want {
		t.Errorf("Request method: %s, want %s", got, want)
	}
}
