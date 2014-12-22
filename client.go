package swapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// DefaultClient is the default SWAPI client
var DefaultClient = NewClient(nil)

// A Client communicates with SWAPI
type Client struct {
	// BaseURL is the base url for SWAPI
	BaseURL *url.URL

	// BasePath is the base path for the endpoints
	BasePath string

	// User agent used for HTTP requests to SWAPI
	UserAgent string

	// HTTP client used to communicate with the SWAPI
	httpClient *http.Client
}

// NewClient returns a new SWAPI client.
// If httpClient is nil, http.DefaultClient is used.
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		cloned := *http.DefaultClient
		httpClient = &cloned
	}

	c := &Client{
		BaseURL: &url.URL{
			Scheme: Env("SWAPI_BASE_URL_SCHEME", "http"),
			Host:   Env("SWAPI_BASE_URL_HOST", "swapi.co"),
		},
		BasePath:   Env("SWAPI_BASE_PATH", "/api/"),
		UserAgent:  Env("SWAPI_USER_AGENT", "swapi.go"),
		httpClient: httpClient,
	}

	return c
}

// NewRequest creates an API request.
func (c *Client) NewRequest(s string) (*http.Request, error) {
	rel, err := url.Parse(c.BasePath + s)
	if err != nil {
		return nil, err
	}

	q := rel.Query()
	q.Set("format", "json")

	rel.RawQuery = q.Encode()

	u := c.BaseURL.ResolveReference(rel)

	if EnvBool("SWAPI_VERBOSE", false) {
		fmt.Println("swapi: GET", u.String())
	}

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", c.UserAgent)
	return req, nil
}

// Do sends an API request and returns the API response. The API response is
// decoded and stored in the value pointed to by v, or returned as an error if
// an API error has occurred.
func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
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
