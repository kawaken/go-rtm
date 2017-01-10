package rtm

import (
	"io/ioutil"
	"net/url"
	"os"
	"testing"

	"github.com/kawaken/go-rtm/methods/auth"
)

func TestSign(t *testing.T) {

	// cf: https://www.rememberthemilk.com/services/api/authentication.rtm
	params := url.Values{}
	params.Add("yxz", "foo")
	params.Add("feg", "bar")
	params.Add("abc", "baz")

	s := sign(params, "BANANAS")

	if s != "82044aae4dd676094f23f1ec152159ba" {
		t.Fatalf("api sign is invalid result: %s", s)
	}
}

func TestConnect(t *testing.T) {
	client, err := NewClient(os.Getenv("RTM_API_KEY"), os.Getenv("RTM_API_SECRET"))
	if err != nil {
		t.Fatal(err)
	}

	res, err := client.Do(auth.GetFrob())

	if err != nil {
		t.Fatal(err)
	}

	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(b))

}
