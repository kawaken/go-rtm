package contacts

// This file is generated by methodGenerator.
// DO NOT MOTIFY THIS FILE.

import (
	"net/url"

	"github.com/kawaken/go-rtm/methods"
)

// Delete returns "rtm.contacts.delete" method instance.
func Delete(timeline string, contactID string) *methods.Method {
	name := "rtm.contacts.delete"

	p := url.Values{}
	p.Add("method", name)
	p.Add("timeline", timeline)
	p.Add("contact_id", contactID)
	
	return &methods.Method{Name: name, Params: p}
}
