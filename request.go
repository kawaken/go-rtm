package rtm

import (
	"net/url"
	"os"
)

func newParams() url.Values {
	apiKey := os.Getenv("RTM_API_KEY")

	v := url.Values{}
	v.Add("api_key", apiKey)
	v.Add("format", "json")

	return v
}
