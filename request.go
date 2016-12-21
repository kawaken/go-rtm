package rtm

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"net/url"
	"os"
	"sort"
)

var (
	apiKey    = os.Getenv("RTM_API_KEY")
	apiSecret = os.Getenv("RTM_API_SECRET")
)

func newParams() url.Values {

	v := url.Values{}
	v.Add("api_key", apiKey)
	v.Add("format", "json")

	return v
}

func sign(values url.Values) string {
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
