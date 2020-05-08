package swapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const (
	defaultBaseURLScheme = "https"
	defaultBaseURLHost   = "swapi.dev"
	defaultBasePath      = "/api/"
	defaultUserAgent     = "swapi.go"
)

// DefaultClient is the default SWAPI client
var DefaultClient = NewClient()

// A Client communicates with SWAPI
type Client struct {
	// baseURL is the base url for SWAPI
	baseURL *url.URL

	// basePath is the base path for the endpoints
	basePath string

	// User agent used for HTTP requests to SWAPI
	userAgent string

	// HTTP client used to communicate with the SWAPI
	httpClient *http.Client
}

// NewClient returns a new SWAPI client.
func NewClient(options ...Option) *Client {
	c := &Client{
		baseURL: &url.URL{
			Scheme: defaultBaseURLScheme,
			Host:   defaultBaseURLHost,
		},
		basePath:   defaultBasePath,
		userAgent:  defaultUserAgent,
		httpClient: http.DefaultClient,
	}

	for _, option := range options {
		option(c)
	}

	return c
}

// newRequest creates an API request.
func (c *Client) newRequest(s string) (*http.Request, error) {
	rel, err := url.Parse(c.basePath + s)
	if err != nil {
		return nil, err
	}

	q := rel.Query()
	q.Set("format", "json")

	rel.RawQuery = q.Encode()

	u := c.baseURL.ResolveReference(rel)

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", c.userAgent)

	return req, nil
}

// do sends an API request and returns the API response. The API response is
// decoded and stored in the value pointed to by v, or returned as an error if
// an API error has occurred.
func (c *Client) do(req *http.Request, v interface{}) (*http.Response, error) {
	// Make sure to close the connection after replying to this request
	req.Close = true

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if v != nil {
		err = json.NewDecoder(resp.Body).Decode(v)
	}

	if err != nil {
		return nil, fmt.Errorf("error reading response from %s %s: %s", req.Method, req.URL.RequestURI(), err)
	}

	return resp, nil
}
