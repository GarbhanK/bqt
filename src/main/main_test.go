package main

import (
	"fmt"
	"testing"
	"time"
)

func TestReadSQL(t *testing.T) {

	result := ReadSQL("test.sql")

	sql := "select * from `{{ params.project }}.transactions.coffee` c\n"
	sql += "where date(insertionTimestamp) >= '{{ ds_nodash }}'\n"
	sql += "left join `{{ params.web_project }}.unified_segment.tracks` t\n"
	sql += "on c.userId = t.userId\n"
	sql += "group by insertionTimestamp desc"
	expected := sql

	t.Logf("result: %s\n", result)
	t.Logf("expected: %s\n", expected)

	if len(result) > 0 {
		t.Logf("ReadSQL('test.sql') PASSED. Is not an empty string\n")
	} else {
		t.Errorf("ReadSQL('test.sql') FAILED. Got an empty string\n")
	}

	if result != expected {
		t.Errorf("ReadSQL('test.sql') FAILED. Expected %s, got %s\n", expected, result)
	} else {
		t.Logf("ReadSQL('test.sql') PASSED. Expected %s, got %s\n", expected, result)
	}
}

func TestCreateMapping(t *testing.T) {

	dt := time.Now()
	ds_nodash := fmt.Sprintf("%d-%d-%d", dt.Year(), dt.Month(), dt.Day())
	ds := fmt.Sprintf("%d%d%d", dt.Year(), dt.Month(), dt.Day())

	expected := map[string]string{
		"ds":                 ds,
		"ds_nodash":          ds_nodash,
		"params.web_project": "livescore-web",
	}

	result_live := CreateMapping("live", true)
	expected_live := make(map[string]string)
	for k, v := range expected {
		expected_live[k] = v
	}
	expected_live["params.project"] = "ls-africa-data-eu-live"

	result_dev := CreateMapping("dev", true)
	expected_dev := make(map[string]string)
	for k, v := range expected {
		expected_dev[k] = v
	}
	expected_dev["params.project"] = "ls-africa-data-eu-dev"

	eq_live := fmt.Sprint(result_live) == fmt.Sprint(expected_live)
	if eq_live {
		t.Logf("CreateMapping('live') PASSED. Expected %s, got %s\n", expected_live, result_live)
	} else {
		t.Errorf("CreateMapping('live') FAILED. Expected %s, got %s\n", expected_live, result_live)
	}

	eq_dev := fmt.Sprint(result_dev) == fmt.Sprint(expected_dev)
	if eq_dev {
		t.Logf("CreateMapping('dev') PASSED. Expected %s, got %s\n", expected_dev, result_dev)
	} else {
		t.Errorf("CreateMapping('dev') FAILED. Expected %s, got %s\n", expected_dev, result_dev)
	}
}
