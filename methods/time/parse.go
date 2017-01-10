package time

// This file is generated by methodGenerator.
// DO NOT MOTIFY THIS FILE.

import (
	"net/url"

	"github.com/kawaken/go-rtm/methods"
)

// Parse returns "rtm.time.parse" method instance.
func Parse(text string, timezone string, dateformat string) *methods.Method {
	name := "rtm.time.parse"

	p := url.Values{}
	p.Add("method", name)
	p.Add("text", text)
	p.Add("timezone", timezone)
	p.Add("dateformat", dateformat)
	
	return &methods.Method{Name: name, Params: p}
}