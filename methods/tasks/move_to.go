package tasks

// This file is generated by methodGenerator.
// DO NOT MOTIFY THIS FILE.

import (
	"net/url"

	"github.com/kawaken/go-rtm/methods"
)

// MoveTo returns "rtm.tasks.moveTo" method instance.
func MoveTo(timeline string, fromListID string, toListID string, taskseriesID string, taskID string) *methods.Method {
	name := "rtm.tasks.moveTo"

	p := url.Values{}
	p.Add("method", name)
	p.Add("timeline", timeline)
	p.Add("from_list_id", fromListID)
	p.Add("to_list_id", toListID)
	p.Add("taskseries_id", taskseriesID)
	p.Add("task_id", taskID)
	
	return &methods.Method{Name: name, Params: p}
}