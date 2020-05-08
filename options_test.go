package swapi

import (
	"net/http"
	"testing"
)

func TestOptions(t *testing.T) {
	t.Run("HTTPClient", func(t *testing.T) {
		c := &Client{}

		HTTPClient(http.DefaultClient)(c)

		if c.httpClient == nil {
			t.Fatalf("c.httpClient not set")
		}
	})

	t.Run("FromEnv", func(t *testing.T) {
		c := &Client{}

		t.Run("fallback", func(t *testing.T) {
			FromEnv(func(string) string { return "" })(c)

			if got, want := c.baseURL.Scheme, defaultBaseURLScheme; got != want {
				t.Fatalf("c.baseURL.Scheme = %q, want %q", got, want)
			}

			if got, want := c.baseURL.Host, defaultBaseURLHost; got != want {
				t.Fatalf("c.baseURL.Host = %q, want %q", got, want)
			}

			if got, want := c.basePath, defaultBasePath; got != want {
				t.Fatalf("c.basePath = %q, want %q", got, want)
			}

			if got, want := c.userAgent, defaultUserAgent; got != want {
				t.Fatalf("c.userAgent = %q, want %q", got, want)
			}
		})

		t.Run("with values", func(t *testing.T) {
			scheme, host, path, userAgent := "http", "example.com", "/path/", "user-agent"

			FromEnv(func(key string) string {
				switch key {
				case "SWAPI_BASE_URL_SCHEME":
					return scheme
				case "SWAPI_BASE_URL_HOST":
					return host
				case "SWAPI_BASE_PATH":
					return path
				case "SWAPI_USER_AGENT":
					return userAgent
				default:
					return ""
				}
			})(c)

			if got, want := c.baseURL.Scheme, scheme; got != want {
				t.Fatalf("c.baseURL.Scheme = %q, want %q", got, want)
			}

			if got, want := c.baseURL.Host, host; got != want {
				t.Fatalf("c.baseURL.Host = %q, want %q", got, want)
			}

			if got, want := c.basePath, path; got != want {
				t.Fatalf("c.basePath = %q, want %q", got, want)
			}

			if got, want := c.userAgent, userAgent; got != want {
				t.Fatalf("c.userAgent = %q, want %q", got, want)
			}
		})
	})
}
