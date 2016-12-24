package methods

import "net/url"

// Method is a RTM API Method
type Method struct {
	Name   string
	Params url.Values
}
