package groups

// This file is generated by methodGenerator.
// DO NOT MOTIFY THIS FILE.

import (
	"net/url"

	"github.com/kawaken/go-rtm/methods"
)

// Delete returns "rtm.groups.delete" method instance.
func Delete(timeline string, groupID string) *methods.Method {
	name := "rtm.groups.delete"

	p := url.Values{}
	p.Add("method", name)
	p.Add("timeline", timeline)
	p.Add("group_id", groupID)
	
	return &methods.Method{Name: name, Params: p}
}