package tasks

// This file is generated by methodGenerator.
// DO NOT MOTIFY THIS FILE.

import (
	"net/url"

	"github.com/kawaken/go-rtm/methods"
)

// SetTags returns "rtm.tasks.setTags" method instance.
func SetTags(timeline string, listID string, taskseriesID string, taskID string, tags string) *methods.Method {
	name := "rtm.tasks.setTags"

	p := url.Values{}
	p.Add("method", name)
	p.Add("timeline", timeline)
	p.Add("list_id", listID)
	p.Add("taskseries_id", taskseriesID)
	p.Add("task_id", taskID)
	p.Add("tags", tags)
	
	return &methods.Method{Name: name, Params: p}
}
