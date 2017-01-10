package tasks

// This file is generated by methodGenerator.
// DO NOT MOTIFY THIS FILE.

import (
	"net/url"

	"github.com/kawaken/go-rtm/methods"
)

// SetEstimate returns "rtm.tasks.setEstimate" method instance.
func SetEstimate(timeline string, listID string, taskseriesID string, taskID string, estimate string) *methods.Method {
	name := "rtm.tasks.setEstimate"

	p := url.Values{}
	p.Add("method", name)
	p.Add("timeline", timeline)
	p.Add("list_id", listID)
	p.Add("taskseries_id", taskseriesID)
	p.Add("task_id", taskID)
	p.Add("estimate", estimate)
	
	return &methods.Method{Name: name, Params: p}
}