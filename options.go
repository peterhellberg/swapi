package swapi

import (
	"net/http"
	"net/url"
)

type Option func(*Client)

// HTTPClient to use by the client
func HTTPClient(hc *http.Client) Option {
	return func(c *Client) {
		c.httpClient = hc
	}
}

// BaseURL for the client parsed from provided rawurl
func BaseURL(rawurl string) Option {
	return func(c *Client) {
		if baseURL, err := url.Parse(rawurl); err == nil {
			c.baseURL = baseURL
		}
	}
}

// UserAgent to use by the client
func UserAgent(ua string) Option {
	return func(c *Client) {
		c.userAgent = ua
	}
}

// FromEnv client configuration
func FromEnv(getenv func(string) string) Option {
	return func(c *Client) {
		env := func(key, fallback string) string {
			if v := getenv(key); v != "" {
				return v
			}

			return fallback
		}

		c.baseURL = &url.URL{
			Scheme: env("SWAPI_BASE_URL_SCHEME", defaultBaseURLScheme),
			Host:   env("SWAPI_BASE_URL_HOST", defaultBaseURLHost),
		}

		c.basePath = env("SWAPI_BASE_PATH", defaultBasePath)

		c.userAgent = env("SWAPI_USER_AGENT", defaultUserAgent)
	}
}
