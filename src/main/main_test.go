package main

import (
	"fmt"
	"testing"
	"time"
)

// go test -v tests/main_test.go

func TestReadSQL(t *testing.T) {

	result := ReadSQL("test.sql")
	var expected string

	expected = "select * from from `{{ params.project }}.transactions.coffee`" +
		"c where date(insertionTimestamp) >= '{{ ds_nodash }}'" +
		"left join `{{ params.web_project }}.unified_segment.tracks` t" +
		"on c.userId = t.userId group by insertionTimestamp desc"

	t.Log(result)
	t.Log(expected)

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

	result_live := CreateMapping("live")
	expected_live := make(map[string]string)
	for k, v := range expected {
		expected_live[k] = v
	}
	expected_live["params.project"] = "ls-africa-data-eu-live"

	result_dev := CreateMapping("dev")
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
