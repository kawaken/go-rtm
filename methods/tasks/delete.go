package tasks

// This file is generated by methodGenerator.
// DO NOT MOTIFY THIS FILE.

import (
	"net/url"

	"github.com/kawaken/go-rtm/methods"
)

// Delete returns "rtm.tasks.delete" method instance.
func Delete(timeline string, listID string, taskseriesID string, taskID string) *methods.Method {
	name := "rtm.tasks.delete"

	p := url.Values{}
	p.Add("method", name)
	p.Add("timeline", timeline)
	p.Add("list_id", listID)
	p.Add("taskseries_id", taskseriesID)
	p.Add("task_id", taskID)
	
	return &methods.Method{Name: name, Params: p}
}