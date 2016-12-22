package rtm

import (
	"net/url"
	"testing"
)

func TestSign(t *testing.T) {

	// cf: https://www.rememberthemilk.com/services/api/authentication.rtm
	params := url.Values{}
	params.Add("yxz", "foo")
	params.Add("feg", "bar")
	params.Add("abc", "baz")

	apiSecret = "BANANAS"

	s := sign(params, apiSecret)

	if s != "82044aae4dd676094f23f1ec152159ba" {
		t.Fatalf("api sign is invalid result: %s", s)
	}
}
