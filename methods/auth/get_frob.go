package auth

// This file is generated by methodGenerator.
// DO NOT MOTIFY THIS FILE.

import (
	"net/url"

	"github.com/kawaken/go-rtm/methods"
)

// GetFrob returns "rtm.auth.getFrob" method instance.
func GetFrob() *methods.Method {
	name := "rtm.auth.getFrob"

	p := url.Values{}
	p.Add("method", name)
	
	return &methods.Method{Name: name, Params: p}
}
