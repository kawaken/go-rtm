package time

// This file is generated by methodGenerator.
// DO NOT MOTIFY THIS FILE.

import (
	"net/url"

	"github.com/kawaken/go-rtm/methods"
)

// Convert returns "rtm.time.convert" method instance.
func Convert(toTimezone string, fromTimezone string, time string) *methods.Method {
	name := "rtm.time.convert"

	p := url.Values{}
	p.Add("method", name)
	p.Add("to_timezone", toTimezone)
	p.Add("from_timezone", fromTimezone)
	p.Add("time", time)
	
	return &methods.Method{Name: name, Params: p}
}
