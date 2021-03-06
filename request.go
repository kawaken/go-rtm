package rtm

import (
	"bytes"
	"context"
	"crypto/md5"
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"

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

// ClientOption is option
type ClientOption func(*Client) error

// NewClient returns a new RTM Client.
func NewClient(apiKey string, apiSecret string, options ...ClientOption) (*Client, error) {

	if apiKey == "" {
		return nil, fmt.Errorf("missing apiKey")
	}

	if apiSecret == "" {
		return nil, fmt.Errorf("missing apiSecret")
	}

	c := &Client{
		apiKey:     apiKey,
		apiSecret:  apiSecret,
		endpoint:   APIEndpoint,
		httpClient: http.DefaultClient,
	}

	for _, o := range options {
		err := o(c)
		if err != nil {
			return nil, err
		}
	}

	return c, nil
}

// WithHTTPClient is used to set HTTP Client
func WithHTTPClient(c *http.Client) ClientOption {
	return func(client *Client) error {
		client.httpClient = c
		return nil
	}
}

// WithEndpoint is used to set endpoint
func WithEndpoint(endpoint string) ClientOption {
	return func(client *Client) error {
		client.endpoint = endpoint
		return nil
	}
}

func (c *Client) do(req *http.Request) (*http.Response, error) {
	req.Header.Set("User-Agent", "go-rtm-client")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	req.WithContext(ctx)

	return c.httpClient.Do(req)
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

	return c.do(req)
}
