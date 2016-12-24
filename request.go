package rtm

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"net/http"
	"net/url"
	"os"
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

var (
	apiKey    = os.Getenv("RTM_API_KEY")
	apiSecret = os.Getenv("RTM_API_SECRET")
)

func appendRequiredParams(v url.Values, apiKey string, apiSecret string) url.Values {

	v.Add("api_key", apiKey)
	v.Add("format", "json")

	// add sign
	v.Add("api_sig", sign(v, apiSecret))

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

	v := appendRequiredParams(m.Params, c.apiKey, c.apiSecret)

	req, err := http.NewRequest(
		http.MethodPost, c.endpoint, strings.NewReader(v.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "go-rtm-client")
	return c.httpClient.Do(req)
}
