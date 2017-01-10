package tasks

// This file is generated by methodGenerator.
// DO NOT MOTIFY THIS FILE.

import (
	"net/url"

	"github.com/kawaken/go-rtm/methods"
)

// SetURL returns "rtm.tasks.setURL" method instance.
func SetURL(timeline string, listID string, taskseriesID string, taskID string, uRL string) *methods.Method {
	name := "rtm.tasks.setURL"

	p := url.Values{}
	p.Add("method", name)
	p.Add("timeline", timeline)
	p.Add("list_id", listID)
	p.Add("taskseries_id", taskseriesID)
	p.Add("task_id", taskID)
	p.Add("url", uRL)
	
	return &methods.Method{Name: name, Params: p}
}