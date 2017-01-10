package timezones

// This file is generated by methodGenerator.
// DO NOT MOTIFY THIS FILE.

import (
	"net/url"

	"github.com/kawaken/go-rtm/methods"
)

// GetList returns "rtm.timezones.getList" method instance.
func GetList() *methods.Method {
	name := "rtm.timezones.getList"

	p := url.Values{}
	p.Add("method", name)
	
	return &methods.Method{Name: name, Params: p}
}
