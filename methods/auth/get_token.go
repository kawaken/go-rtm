package auth

// This file is generated by methodGenerator.
// DO NOT MOTIFY THIS FILE.

import (
	"net/url"

	"github.com/kawaken/go-rtm/methods"
)

// GetToken returns "rtm.auth.getToken" method instance.
func GetToken(frob string) *methods.Method {
	name := "rtm.auth.getToken"

	p := url.Values{}
	p.Add("method", name)
	p.Add("frob", frob)
	
	return &methods.Method{Name: name, Params: p}
}
