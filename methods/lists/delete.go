package lists

// This file is generated by methodGenerator.
// DO NOT MOTIFY THIS FILE.

import (
	"net/url"

	"github.com/kawaken/go-rtm/methods"
)

// Delete returns "rtm.lists.delete" method instance.
func Delete(timeline string, listID string) *methods.Method {
	name := "rtm.lists.delete"

	p := url.Values{}
	p.Add("method", name)
	p.Add("timeline", timeline)
	p.Add("list_id", listID)
	
	return &methods.Method{Name: name, Params: p}
}