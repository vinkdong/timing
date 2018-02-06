package types

import (
	"testing"
	"github.com/bitly/go-simplejson"
)

func TestChecker_CheckJson(t *testing.T) {
	c := Checker{
		Rule: `{{ $.status eq "green" }}`,
		Name: "test_check",
		Type: "json",
	}
	js, _ := simplejson.NewJson([]byte(`{
  "cluster_name" : "logging",
  "status" : "green",
  "timed_out" : false,
  "number_of_nodes" : 6,
  "number_of_data_nodes" : 2,
  "active_primary_shards" : 3973,
  "active_shards" : 7946,
  "relocating_shards" : 0,
  "initializing_shards" : 0,
  "unassigned_shards" : 0,
  "delayed_unassigned_shards" : 0,
  "number_of_pending_tasks" : 0,
  "number_of_in_flight_fetch" : 0,
  "task_max_waiting_in_queue_millis" : 0,
  "active_shards_percent_as_number" : 100.0
}`))
	c.CheckJson(js)
}
