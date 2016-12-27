package rtm

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"strings"

	"github.com/kawaken/go-rtm/methods"
)

const (
	// APIEndpoint is the REST endpoint URL
	APIEndpoint = "https://api.rememberthemilk.com/services/rest/"
	// AuthenticationURL is authentication URL
	AuthenticationURL = "https://www.rememberthemilk.com/services/auth/"
)

func appendRequiredParams(v url.Values, apiKey string, apiSecret string) url.Values {

	v.Set("api_key", apiKey)
	v.Set("format", "json")

	// add sign
	v.Set("api_sig", sign(v, apiSecret))

	return v
}

func sign(values url.Values, apiSecret string) string {
	var keys []string
	for k := range values {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	buf := bytes.NewBufferString(apiSecret)

	for _, k := range keys {
		buf.WriteString(k)
		for _, v := range values[k] {
			buf.WriteString(v)
		}
	}

	return fmt.Sprintf("%x", md5.Sum(buf.Bytes()))
}

// Client is api client.
type Client struct {
	apiKey     string
	apiSecret  string
	endpoint   string
	httpClient *http.Client
}

// NewClient returns a new RTM Client.
func NewClient(apiKey string, apiSecret string) (*Client, error) {

	if apiKey == "" {
		return nil, fmt.Errorf("no apiKey")
	}

	if apiSecret == "" {
		return nil, fmt.Errorf("no apiSecret")
	}

	return &Client{
		apiKey:     apiKey,
		apiSecret:  apiSecret,
		endpoint:   APIEndpoint,
		httpClient: http.DefaultClient,
	}, nil
}

// Do requests a RTM API method.
func (c *Client) Do(m *methods.Method) (*http.Response, error) {

	if m.Params.Get("method") == "" {
		return nil, fmt.Errorf("API method name is not defined")
	}

	v := appendRequiredParams(m.Params, c.apiKey, c.apiSecret)

	req, err := http.NewRequest(
		http.MethodPost, c.endpoint, strings.NewReader(v.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "go-rtm-client")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return c.httpClient.Do(req)
}
