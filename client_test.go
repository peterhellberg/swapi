package swapi

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewClient(t *testing.T) {
	userAgent := "test-agent"

	c := NewClient(UserAgent(userAgent))

	if got, want := c.userAgent, userAgent; got != want {
		t.Fatalf("c.userAgent = %q, want %q", got, want)
	}
}

func TestClientNewRequest(t *testing.T) {
	c := NewClient()

	r, err := c.newRequest("test")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if got, want := r.UserAgent(), defaultUserAgent; got != want {
		t.Fatalf("r.UserAgent() = %q, want %q", got, want)
	}

	if got, want := r.URL.String(), "https://swapi.dev/api/test?format=json"; got != want {
		t.Fatalf("r.URL.String() = %q, want %q", got, want)
	}
}

func TestClientDo(t *testing.T) {
	hf := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(Film{Title: "Example"})
	})

	ts := httptest.NewServer(hf)
	defer ts.Close()

	c := NewClient(BaseURL(ts.URL))

	r, err := c.newRequest("test")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var v Film

	if _, err := c.do(r, &v); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if got, want := v.Title, "Example"; got != want {
		t.Fatalf("v.Title = %q, want %q", got, want)
	}
}
