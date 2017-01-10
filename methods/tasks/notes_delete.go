package tasks

// This file is generated by methodGenerator.
// DO NOT MOTIFY THIS FILE.

import (
	"net/url"

	"github.com/kawaken/go-rtm/methods"
)

// NotesDelete returns "rtm.tasks.notes.delete" method instance.
func NotesDelete(timeline string, noteID string) *methods.Method {
	name := "rtm.tasks.notes.delete"

	p := url.Values{}
	p.Add("method", name)
	p.Add("timeline", timeline)
	p.Add("note_id", noteID)
	
	return &methods.Method{Name: name, Params: p}
}
