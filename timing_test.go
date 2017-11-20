package main

import "testing"

func TestParseYaml(t *testing.T) {
	r := &Rule{}
	parseYaml(r, "./config.yaml")
	if r.Method != "PUT"{
		t.Errorf("test parse yaml file faild expect method is PUT but got %s",r.Method)
	}

	if r.Range["hour"]["gte"] != 7{
		t.Errorf("test parse yaml file faild expect range hour is gte 7 but go %d",r.Range["hour"]["gte"])
	}
}
