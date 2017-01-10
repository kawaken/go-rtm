package reflection

// This file is generated by methodGenerator.
// DO NOT MOTIFY THIS FILE.

import (
	"net/url"

	"github.com/kawaken/go-rtm/methods"
)

// GetMethodInfo returns "rtm.reflection.getMethodInfo" method instance.
func GetMethodInfo(methodName string) *methods.Method {
	name := "rtm.reflection.getMethodInfo"

	p := url.Values{}
	p.Add("method", name)
	p.Add("method_name", methodName)
	
	return &methods.Method{Name: name, Params: p}
}
